package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"resume/internal/resumeBot"
)

type son struct {
	Message1 string
}

func main() {
	var mir son
	filename, err := os.Open("texts.json")
	if err != nil {
		panic(err)
	}
	defer filename.Close()
	data, err := ioutil.ReadAll(filename)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &mir)
	fmt.Println(mir)
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("Token not found")
	}

	//fmt.Println("token=", token)
	if err := resumeBot.Start(token); err != nil {
		panic(err)
	}
}
