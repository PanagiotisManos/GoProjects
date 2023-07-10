package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Items struct {
	Diamond  string
	Emerald  string
	Sapphire string
}

func main() {
	items := Items{
		Diamond:  "Diamond",
		Emerald:  "Emerald",
		Sapphire: "Sapphire",
	}

	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator

	selectedItem := getRandomItem(items)
	fmt.Println(selectedItem)
}

func getRandomItem(items Items) string {
	itemList := make([]string, 0, 3)
	itemList = append(itemList, items.Diamond)
	itemList = append(itemList, items.Emerald)
	itemList = append(itemList, items.Sapphire)

	randomIndex := rand.Intn(len(itemList))

	return itemList[randomIndex]
}