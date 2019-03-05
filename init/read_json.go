package main

import (
	"fmt";
	"os";
	"encoding/json";
	"io/ioutil"
)

func Read_json(filename string) {
	content, err := os.Open(filename)
	// fmt.Println(os.Getwd())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Successfully opened %s\n", filename)
	}

	byteVal, _ := ioutil.ReadAll(content)
	var users Users
	json.Unmarshal(byteVal, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
	}

	defer content.Close()
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name 	string 	`json:"name"`
	Type 	string 	`json:"type"`
	Age 	int		`json:"age"`
	Social	Social 	`json:"social"`
}

type Social struct {
	Facebook	string `json:"facebook"`
	Twitter 	string `json:"twitter"`
}