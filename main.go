package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./wap.txt")
	if err != nil {
		fmt.Println(err)
	}
	countWords := make(map[string]int)
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var lowerText = strings.ToLower(scanner.Text())

		regex := regexp.MustCompile("(!|,|\\.|\\?|'|\\)|\\(|–|[0-9]|;|:|)|«|»") // [a-zA-Z]

		lowerText = regex.ReplaceAllString(lowerText, "")

		var strSplit = strings.Fields(lowerText)
		wordsCountInString(strSplit, countWords)
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

func wordsCountInString(strSplit []string, countWords map[string]int) {

	for _, word := range strSplit {
		countWords[word]++
	}

}
