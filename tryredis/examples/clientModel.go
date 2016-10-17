package examples

import (
	"time"
	"fmt"
	"encoding/json"
	"log"
)

type ClientModel struct {
	FirstName string
	LastName string
	Email string
	DateOfBirth time.Time
}

func (c *ClientModel) WriteToRedis() {
	client := GetRedisClient()
	jsonString, err := json.Marshal(*c)
	if err != nil {
		log.Println("Failed marshalling json", err)
	}
	tmp := string(jsonString[:])
	log.Println("The json:", tmp)
	err = client.Set(c.Email, tmp, 0).Err()
	if err != nil {
		log.Println("Failed put the data into redis due to", err)
	}
}

func (c *ClientModel) ReadFromRedis(email string) *ClientModel {
	client := GetRedisClient()
	val, err := client.Get(email).Result()
	log.Println("The VAL", val)
	log.Println("The ERR", err)
	if err != nil {
		log.Printf("Failed reading %s, due to %v", email, err)
		return nil
	}
	clientModel := &ClientModel{}
	err = json.Unmarshal([]byte(val), clientModel)
	if err != nil {
		log.Print("Failed unmarshalling", err)
	}
	return clientModel
}

func NewClientModel(firstName, lastName, email, dob string) *ClientModel {
	c := &ClientModel{}
	c.FirstName = firstName
	c.LastName = lastName
	c.Email = email
	var err error
	c.DateOfBirth, err = time.Parse("02-01-2006", dob)

	fmt.Println("Error", err)
	return c
}


