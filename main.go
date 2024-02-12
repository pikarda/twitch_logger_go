package main

import (
	"fmt"
)

var uselocalConfig bool

func main() {
	fmt.Println(styledStartApp("APP STARTED"))

	GetReloadStatus()

	GetNotionData()

	Logger()
}
