package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func search(lines []string, searchTerm string) []string {
	var results []string
	for _, line := range lines {
		if strings.Contains(line, searchTerm) {
			results = append(results, line)
		}
	}
	return results
}

func main() {
	filename := "text.txt"
	lines, err := readFile(filename)
	if err != nil {
		fmt.Println("Помилка читання файлу:", err)
		return
	}

	cache := make(map[string][]string)

	for {
		fmt.Print("Введіть рядок для пошуку: ")
		var searchTerm string
		fmt.Scanln(&searchTerm)

		if cachedResults, found := cache[searchTerm]; found {
			fmt.Println("Знайдені рядки (з кешу):")
			for _, result := range cachedResults {
				fmt.Println(result)
			}
		} else {
			results := search(lines, searchTerm)
			cache[searchTerm] = results
			fmt.Println(cache)

			fmt.Println("Знайдені рядки:")
			for _, result := range results {
				fmt.Println(result)
			}
		}

		fmt.Print("Продовжити пошук? (так/ні): ")
		var continueSearch string
		fmt.Scanln(&continueSearch)

		if strings.ToLower(continueSearch) != "так" {
			break
		}
	}
}
