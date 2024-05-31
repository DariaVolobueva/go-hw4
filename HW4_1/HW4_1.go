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

	fmt.Print("Введіть рядок для пошуку: ")
	var searchTerm string
	fmt.Scanln(&searchTerm)

	results := search(lines, searchTerm)

	fmt.Println("Знайдені рядки:")
	for _, result := range results {
		fmt.Println(result)
	}
}
