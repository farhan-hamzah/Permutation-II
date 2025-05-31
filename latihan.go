package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		// Ubah kata menjadi urutan huruf yang telah diurutkan (sorted)
		sortedWord := sortString(word)
		// Gunakan kata yang telah diurutkan sebagai key
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	// Kumpulkan hasilnya ke slice 2 dimensi
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

// Helper: mengurutkan huruf dalam string
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {
	// Contoh penggunaan
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	grouped := groupAnagrams(input)

	// Cetak hasilnya
	for _, group := range grouped {
		fmt.Println(group)
	}
}
