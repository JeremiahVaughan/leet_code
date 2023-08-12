package main

import (
	"sort"
)

func main() {

}

func groupAnagrams(strs []string) [][]string {
	groupsMap := make(map[string][]string)
	for _, str := range strs {
		stringBytes := []byte(str)
		sort.Slice(stringBytes, func(i, j int) bool { return stringBytes[i] < stringBytes[j] })
		sortedString := string(stringBytes)
		groupsMap[sortedString] = append(groupsMap[sortedString], str)
	}

	var result [][]string
	for _, g := range groupsMap {
		result = append(result, g)
	}

	return result
}
