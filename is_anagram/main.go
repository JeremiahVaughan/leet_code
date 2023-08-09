package main

import "strings"

func isAnagram(s string, t string) bool {
	// key is a letter in the word, the value is the number of times that letter occurs in the word
	leftLetterMap := fillMap(s)
	rightLetterMap := fillMap(t)
	if len(leftLetterMap) != len(rightLetterMap) {
		return false
	}

	for k, leftValue := range leftLetterMap {
		rightValue, ok := rightLetterMap[k]
		if !ok {
			return false
		}

		if leftValue != rightValue {
			return false
		}
	}
	return true
}

func fillMap(s string) map[string]int {
	letterMap := make(map[string]int)
	leftSplit := strings.Split(s, "")
	for _, currentLetter := range leftSplit {
		if _, ok := letterMap[currentLetter]; ok {
			letterMap[currentLetter] = letterMap[currentLetter] + 1
		} else {
			letterMap[currentLetter] = 1
		}
	}
	return letterMap
}
