package main

import (
	"fmt"
	"sort"
)

type Item struct {
	ID int
}

func uniqueAndSorted(items []Item) []Item {
	uniqueMap := make(map[int]struct{})
	var uniqueItems []Item

	for _, item := range items {
		if _, exists := uniqueMap[item.ID]; !exists {
			uniqueMap[item.ID] = struct{}{}
			uniqueItems = append(uniqueItems, item)
		}
	}

	sort.Slice(uniqueItems, func(i, j int) bool {
		return uniqueItems[i].ID < uniqueItems[j].ID
	})

	return uniqueItems
}

func main() {
	items := []Item{{3}, {5}, {2}, {2}, {3}, {5}}
	uniqueSortedItems := uniqueAndSorted(items)

	fmt.Println("Унікальні відсортовані елементи:")
	for _, item := range uniqueSortedItems {
		fmt.Println(item)
	}
}
