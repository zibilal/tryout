package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

func main() {
	var db []byte

	fmt.Println("The bytes", db)
	fmt.Printf("The string %s\n", string(db[:]))

	var b bytes.Buffer
	b.Write([]byte("Hello world"))
	io.Copy(os.Stdout, &b)
}
