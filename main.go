package main

import (
	"fmt"
)

func main() {

	fmt.Println(styledStartApp("APP STARTED"))
	GetReloadStatus()
	// data := GetNotionData()
	GetNotionData()

	Logger()
}
