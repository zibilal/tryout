package main

import (
	"encoding/csv"
	"os"
	"log"
)

func main() {
	records := [][] string {
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	f, err := os.Create("data/datausers.txt")
	if err != nil {
		log.Panicln("Failed opening file datausers.txt")
	}

	defer f.Close()

	w := csv.NewWriter(f)
	for _, r := range records {
		if err := w.Write(r); err != nil {
			log.Fatalln("error writing record to csv", err)
		}
	}

	// write any buffered data to the undelying writer
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
