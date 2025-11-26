package underStandingBasics

import (
	"math"
)

func maxElement(arr []int) int {
	var maxElement int = math.MinInt32
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
	}
	return maxElement
}
func minElement(arr []int) int {
	var minElement int = math.MaxInt32
	n := len(arr)
	for i := 0; i < n; i++ {
		if minElement < arr[i] {
			minElement = arr[i]
		}
	}
	return minElement
}
func sumOfArray(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
func avgOfArray(arr []int) int {
	var avgRes int
	var sum int
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		sum += arr[i]
	}
	return sum / avgRes
}
func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
	return arr
}
func secMaxElement(arr []int) int {
	var maxElement int = math.MinInt32
	var secondMaxElement int = math.MinInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxElement {
			secondMaxElement = maxElement
			maxElement = arr[i]
		} else if arr[i] > secondMaxElement && arr[i] < maxElement {
			secondMaxElement = arr[i]
		}
	}
	return secondMaxElement
}
func secondMinElement(arr []int) int {
	var minElement int = math.MaxInt32
	var secondMinElement int = math.MaxInt32
	for i := 0; i < len(arr); i++ {
		if arr[i] < minElement {
			secondMinElement = minElement
			minElement = arr[i]
		} else if arr[i] < secondMinElement && arr[i] != minElement {
			secondMinElement = arr[i]
		}
	}
	return secondMinElement
}
func searchElementInArr(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}
func mergeTwoSortedArray(arr1 []int, arr2 []int) []int {
	i, j := 0, 0
	mergedArr := make([]int, 0, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArr = append(mergedArr, arr1[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		mergedArr = append(mergedArr, arr1[i])
		i++
	}
	for j < len(arr2) {
		mergedArr = append(mergedArr, arr2[j])
		j++
	}
	return mergedArr
}
func FrequencyOfEachElement(arr []int) map[int]int {
	frequencyMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		value, isExists := frequencyMap[arr[i]]
		if isExists {
			frequencyMap[arr[i]] = value + 1
		} else {
			frequencyMap[arr[i]] = 1
		}
	}
	return frequencyMap
}
func RemoveDuplicateElements(arr []int) []int {
	seenMap := make(map[int]bool)
	result := []int{}
	for _, v := range arr {
		if !seenMap[v] {
			seenMap[v] = true
			result = append(result, v)
		}
	}
	return arr
}
