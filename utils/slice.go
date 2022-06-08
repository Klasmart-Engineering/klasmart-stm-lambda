package utils

import (
	"reflect"
	"strings"
)

func SliceDeduplication(s []string) []string {
	temp := make(map[string]bool)
	for i := range s {
		temp[s[i]] = true
	}

	result := make([]string, 0, len(temp))
	for k, v := range temp {
		if v {
			result = append(result, k)
		}
	}

	return result
}

func IntSliceDeduplication(s []int) []int {
	temp := make(map[int]bool)
	var result []int
	for i := range s {
		if !temp[s[i]] {
			temp[s[i]] = true
			result = append(result, s[i])
		}
	}
	return result
}

func StableSliceDeduplication(s []string) []string {
	temp := make(map[string]bool)
	var result []string
	for i := range s {
		if !temp[s[i]] {
			temp[s[i]] = true
			result = append(result, s[i])
		}
	}
	return result
}

// SliceDeduplicationMap deduplicate slice, then return slice and mapping
// if input slice ["1", "2", "3", "4", "1", "2", "4", "3", "3", "3", "2", "1", "5", "3"]
// you will got slice ["1", "2", "3", "4", "5"]
// and mapping {0:0, 1:1, 2:2, 3:3, 4:0, 5:1, 6:3, 7:2, 8:2, 9:2, 10:1, 11:0, 12:4, 13:2}
func SliceDeduplicationMap(origin []string) ([]string, map[int]int) {
	valueMap := make(map[string]int, len(origin))
	indexMap := make(map[int]int, len(origin))
	result := make([]string, 0, len(origin))

	for index, value := range origin {
		resultIndex, found := valueMap[value]
		if !found {
			result = append(result, value)
			resultIndex = len(result) - 1
			valueMap[value] = resultIndex

		}

		indexMap[index] = resultIndex
	}

	return result, indexMap
}

func ExcludeStrings(source []string, targets []string) []string {
	if len(targets) == 0 {
		return source
	}
	var result []string
	for _, item := range source {
		find := false
		for _, target := range targets {
			if item == target {
				find = true
			}
		}
		if !find {
			result = append(result, item)
		}
	}
	return result
}

func FilterStrings(source []string, whitelist, blacklist []string) []string {
	var result []string
	for _, item := range source {
		pass := false
		for _, target := range whitelist {
			if item == target {
				pass = true
			}
		}
		if !pass {
			find := false
			for _, target := range blacklist {
				if item == target {
					find = true
				}
			}
			if !find {
				pass = true
			}
		}
		if pass {
			result = append(result, item)
		}
	}
	return result
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsInt64(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IntersectAndDeduplicateStrSlice(slice1 []string, slice2 []string) []string {
	m := make(map[string]int)
	m2 := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times > 0 {
			m2[v]++
		}
	}
	result := make([]string, 0, len(m2))
	for key, _ := range m2 {
		result = append(result, key)
	}
	return result
}

func SliceDeduplicationExcludeEmpty(s []string) []string {
	temp := make(map[string]bool)
	for i := range s {
		temp[s[i]] = true
	}

	result := make([]string, 0, len(temp))
	for k, v := range temp {
		if v && k != "" {
			result = append(result, k)
		}
	}

	return result
}

func ParseURLQueryArray(s string) []string {
	return SliceDeduplicationExcludeEmpty(strings.Split(s, ","))
}

// Returns whether the given string slice contains a given string.
func ContainsString(list []string, elem string) bool {
	for _, t := range list {
		if t == elem {
			return true
		}
	}

	return false
}

// Returns whether the given string slice contains any of the following strings.
func ContainsAnyString(list []string, elems ...string) bool {
	for _, e := range elems {
		if ContainsString(list, e) {
			return true
		}
	}

	return false
}

func ReverseSliceInPlace(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
