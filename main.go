package main

import (
	"fmt"
)

func main() {
	fmt.Println(styledStartApp("APP STARTED"))

	GetReloadStatus()

	GetNotionData()

	Logger()
}
