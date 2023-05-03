package main

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/goleafcore"
)

func main() {
	dto := goleafcore.Dto{
		"id":   1,
		"name": "Dimas",
	}

	dto.Put("umur", 21)
	fmt.Println("Dto? ", dto)
}
