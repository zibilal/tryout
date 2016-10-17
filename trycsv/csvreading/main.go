package main

import (
	"os"
	"log"
	"encoding/csv"
	"io"
	"fmt"
)

func main() {
	f, err := os.Open("data/datausers.txt")
	if err != nil {
		log.Panic("Failed creating the files", err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
