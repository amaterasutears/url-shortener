package main

import "github.com/amaterasutears/url-shortener/internal/application"

func main() {
	err := application.Run()
	if err != nil {
		panic(err)
	}
}
