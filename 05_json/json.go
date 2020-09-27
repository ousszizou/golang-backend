// Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.

// ===== Encoding =====

// To encode JSON data we use the Marshal function.
// func Marshal(v interface{}) ([]byte, error)

// * Only the exported (public) fields of a struct will be present in the JSON output. Other fields are ignored.
// * A field with a json: struct tag is stored with its tag name instead of its variable name.
// * Pointers will be encoded as the values they point to, or null if the pointer is nil.
// * Channel, complex, and function types cannot be encoded.

// ===== Decoding =====

// To decode JSON data we use the Unmarshal function.
// func Unmarshal(data []byte, v interface{}) error

// ===== arbitrary objects and arrays =====
// The encoding/json package uses

// map[string]interface{} to store arbitrary JSON objects, and
// []interface{} to store arbitrary JSON arrays.


package main

import (
	"encoding/json"
	"fmt"
	"time"
	// "log"
)

// FruitBasket type
type FruitBasket struct {
	Name    string
	Fruit   []string
	ID      int64  `json:"Ref"` // Field appears in JSON as key "Ref".
	private string // An unexported field is not encoded.
	Created time.Time
}

func main() {

	// basket := FruitBasket{
	// 	Name:    "FruitBasket 1",
	// 	Fruit:   []string{"Apple", "Banana", "Orange"},
	// 	ID:      999,
	// 	private: "Second-rate",
	// 	Created: time.Now(),
	// }
	// b, err := json.Marshal(basket)

	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(b) // []byte
	// fmt.Println(string(b))

	// ===== Pretty print: func MarshalIndent
	// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	// b1, err := json.MarshalIndent(basket, "", "    ")
	// fmt.Println(string(b1))

	// ===== Decoding
	// var dbasket FruitBasket

	// err1 := json.Unmarshal(b1, &dbasket)
	// if err1 != nil {
	// 	log.Println(err1)
	// }
	// fmt.Println(dbasket)
	// fmt.Println(dbasket.Name, dbasket.Fruit, dbasket.ID)



	// ===== arbitrary objects and arrays

	child := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var x interface{}

	json.Unmarshal(child, &x)

	fmt.Println(x)

	data := x.(map[string]interface{})

	// var z interface{}

	// z = 7

	// fmt.Println(z.(int)) x.(T) is called a type assertion.

	// v, ok := z.(int)	
	// fmt.Println(v, ok)

	for k, v := range data {
		fmt.Println(k, v)
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")

		case float64:
			fmt.Println(k, v, "(int)")

		case []interface{}:
			for i, u := range v {
        fmt.Println("    ", i, u)
			}
			
		default:
			fmt.Println(k, v, "(unknown)")	
		}
	}

}
