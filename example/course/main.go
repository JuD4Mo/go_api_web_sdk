package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/JuD4Mo/go_api_web_sdk/course"
)

func main() {
	courseTransport := courseSdk.NewHttpClient("http://localhost:8082", "")

	course, err := courseTransport.Get("8106069c-2253-4e01-a49a-227c3abf8d59")
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found:", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal Server Error:", err.Error())
		os.Exit(1)
	}
	fmt.Println(course)
}
