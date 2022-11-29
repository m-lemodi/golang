package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Hello")
	resp, err := http.Get("https://www.googdle.cwom")
	if err != nil {
		errors.New("Faild to get http response")
		return
	}
	fmt.Println("Oui")
	fmt.Println(resp)
}
