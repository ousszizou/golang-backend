// Resources

// https://www.guru99.com/difference-between-cookie-session.html
// https://www.aqweeb.com/2019/06/sessions-cookies-php.html

// https://blog.golang.org/gob
// https://golang.org/pkg/encoding/gob/

// https://github.com/gorilla/sessions
// https://github.com/gorilla/securecookie

package main

import (
	"encoding/gob"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"

	"github.com/gorilla/sessions"
)

// User holds a users account information
type User struct {
	Username      string
	Authenticated bool
}

// store will hold all session data
var store *sessions.CookieStore

// tpl holds all parsed templates
var tpl *template.Template

func init() {
	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		hashKey,
		blockKey,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15, // 15 minutes
		HttpOnly: true,    // the session cannot be altered by javascript
	}

	// We register the custom User type with gob encoding package so it can be written as a session value
	gob.Register(User{})

	// Parse all templates in the templates folder
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	router := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/forbidden", forbidden)
	router.HandleFunc("/secret", secret)
	http.ListenAndServe(":8080", loggedRouter)
}

// index serves the index html file
func index(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := getUser(session)
	tpl.ExecuteTemplate(w, "index.html", user)
}

// login authenticates the user
func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.FormValue("code") != "code" {
		session.AddFlash("The code was incorrect")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	}

	username := r.FormValue("username")

	user := &User{
		Username:      username,
		Authenticated: true,
	}

	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/secret", http.StatusFound)
}

// logout revokes authentication for a user
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// secret displays the secret message for authorized users
func secret(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUser(session)

	if auth := user.Authenticated; !auth {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	}

	tpl.ExecuteTemplate(w, "secret.html", user.Username)
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "forbidden.html", flashMessages)
}

// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}