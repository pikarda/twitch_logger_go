package main

import (
	"fmt"
)

func main() {

	fmt.Println("APP STARTED")
	GetReloadStatus()
	// data := GetNotionData()
	GetNotionData()

	Logger()
}
