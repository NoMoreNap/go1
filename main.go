package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	file, err := os.Open("./wap.txt")
	if err != nil {
		fmt.Println(err)
	}
	countWords := make(map[string]int)
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		countWords[regexp.MustCompile("(!|,|\\.|\\?|'|\\)|\\(|–|[0-9]|;|:|)|«|»").
			ReplaceAllString(scanner.Text(), "")]++
	}

	type key_value struct {
		Key   string
		Value int
	}
	var sorted []key_value

	for key, value := range countWords {
		sorted = append(sorted, key_value{key, value})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	file2, _ := os.Create("counter.txt")

	for _, key_value := range sorted {
		str := fmt.Sprintf("%s: %d\n", key_value.Key, key_value.Value)
		file2.WriteString(str)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
