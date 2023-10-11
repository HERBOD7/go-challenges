package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dicLines := scanner.Text()
	dicLinesInt, _ := strconv.Atoi(dicLines)
	dictionary := make(map[string]string)
	sentence := ""
	for i := 0; i < dicLinesInt+1; i++ {
		scanner.Scan()
		line := scanner.Text()
		if i < dicLinesInt {
			items := strings.Split(line, "->")
			key := strings.Trim(items[0], " ")
			value := strings.Trim(items[1], " ")
			dictionary[key] = value
		} else {
			sentence = sentence + line
		}
	}

	wordsOfSentence := strings.Split(sentence, " ")
	specialChar := [4]string{
		",",
		"?",
		"!",
		".",
	}

	result := []string{}
	for _, word := range wordsOfSentence {
		punctuationMark := []string{}
		isFirstCharCapital := false
		firstChar := string(word[0])
		lowerCaseWord := ""
		wordWithoutChar := word
		finalWord := ""

		if firstChar == strings.ToUpper(firstChar) {
			isFirstCharCapital = true
		}

		for _, char := range specialChar {
			if strings.Contains(wordWithoutChar, char) {
				punctuationMark = append(punctuationMark, char)
				wordWithoutChar = strings.Trim(wordWithoutChar, char)
				if isFirstCharCapital {
					lowerCaseWord = strings.ToLower(wordWithoutChar)
					wordWithoutChar = strings.ToLower(wordWithoutChar)
				}
			} else {
				lowerCaseWord = strings.ToLower(word)
			}
		}

		lowerCaseTranslate := dictionary[lowerCaseWord]
		translate := dictionary[wordWithoutChar]
		// check zero value
		if translate != "" {
			finalWord = translate
		} else if lowerCaseTranslate != "" && isFirstCharCapital {
			finalWord = lowerCaseTranslate
		} else {
			finalWord = wordWithoutChar
		}

		for i := len(punctuationMark) - 1; i >= 0; i-- {
			finalWord = finalWord + punctuationMark[i]
		}

		if isFirstCharCapital {
			firstCharTranslated := string(finalWord[0])
			trimFirstChar := finalWord[1:]
			capitalFirstChar := strings.ToUpper(firstCharTranslated)
			finalWord = capitalFirstChar + trimFirstChar
		}
		result = append(result, finalWord)
	}

	finalResult := ""
	for _, word := range result {
		finalResult = finalResult + word + " "
	}
	finalResultValue := strings.Trim(finalResult, " ")
	fmt.Println(finalResultValue)
}
