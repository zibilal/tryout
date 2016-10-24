package main

import (
	"fmt"
	"sort"
)

type Country struct {
	Name string
}

type Countries []Country

func (sc Countries) Less(i, j int) bool {
	return sc[i].Name < sc[j].Name
}

func (sc Countries) Swap(i, j int) {
	sc[i], sc[j] = sc[j], sc[i]
}

func (sc Countries) Len() int {
	return len(sc)
}

func main() {
	countries := Countries{
		{Name: "United States"},
		{Name: "Bahamas"},
		{Name: "Japan"},
	}

	fmt.Println("Before\n")
	fmt.Println("-->>", countries)

	sort.Sort(countries)

	fmt.Println("\nSorted")
	fmt.Println("-->>", countries)
}
