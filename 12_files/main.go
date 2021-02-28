package main

func main() {

	// Read whole data within file
	// b, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))

	// Read particular data within file
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// b1 := make([]byte, 10)
	// n1, err := f.Read(b1)
	// fmt.Println(string(b1[:n1]))
	// fmt.Println(len(b1))

	// r := bufio.NewReader(f)
	// b, err := r.Peek(10)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))
	// fmt.Println(len(b))

	// f.Close()

	// Writing a file
	// d1 := []byte("hello\ngo\n")
	// err := ioutil.WriteFile("test1.txt", d1, 777)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// f, err := os.Create("test2.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// w := bufio.NewWriter(f)
	// _, err1 := w.WriteString("Learn Golang\n")
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }

	// w.Flush()

}
