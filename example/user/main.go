package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/JuD4Mo/go_api_web_sdk/user"
)

func main() {
	userTransport := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTransport.Get("fa81b5ab-de0b-4d9e-aabc-cb4fe0b587")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(user)
}
