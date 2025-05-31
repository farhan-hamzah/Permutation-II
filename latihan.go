package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums) // Urutkan agar duplikat bisa dihindari dengan mudah
	used := make([]bool, len(nums))
	var backtrack func(path []int)

	backtrack = func(path []int) {
		if len(path) == len(nums) {
			// Buat salinan sebelum menambahkan
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// Lewati jika angka sudah digunakan
			if used[i] {
				continue
			}
			// Hindari duplikat: jika sama dengan sebelumnya dan sebelumnya belum dipakai di level ini
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]int{})
	return res
}

func main() {
	input := []int{1, 1, 2}
	result := permuteUnique(input)
	for _, perm := range result {
		fmt.Println(perm)
	}
}
