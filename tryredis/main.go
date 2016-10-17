package main

import (
	"github.com/zibilal/tryout/tryredis/examples"
	"fmt"
)

func main() {
	c := examples.NewClientModel("Bilal", "Muhammad", "zibilal@outlook.com", "11-12-1977")
	fmt.Println("Client", c)

	c.WriteToRedis()

	fmt.Println("Now read from redis")

	d := c.ReadFromRedis("zibilal@outlook.com")
	fmt.Println("Read", d)
}
