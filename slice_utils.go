package utils

import "sort"

//Reverses the order of elements in the given string slice s.  This does not take the element's value into consideration.
func ReverseStringSlice(s []string) []string {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}

//MapStrings applies the given func to every string in the slice
func MapStrings(slice []string, f func(string) string) []string {
	newslice := make([]string, 0)
	for _, v := range slice {
		newslice = append(newslice, f(v))
	}
	return newslice
}

// Returns the first element in the given string slice
func FirstString(slice []string) string {
	return slice[0]
}

// Returns the last element in the given string slice
func LastString(slice []string) string {
	return slice[len(slice)-1]
}

//FilterStrings returns a new slice containing all strings in the slice that satisfy the predicate f.
func FilterStrings(slice []string, f func(string) bool) []string {
	sliceCopy := make([]string, 0)
	for _, v := range slice {
		if f(v) {
			sliceCopy = append(sliceCopy, v)
		}
	}
	return sliceCopy
}

//StringSliceContains returns true if the provided string is located a given slice of strings.
func StringSliceContains(slice []string, s string) bool {
	sort.Strings(slice)
	i := sort.SearchStrings(slice, s)
	// s is present at slice[i], otherwise s is not present in slice, but i is the index where it would be inserted.
	return i < len(slice) && slice[i] == s
}

// Divides slice s into n subslices such that the elements are distributed
// as evenly as possible. In other words, if there are 10 elements in s,
// and n is 3, there will be one subslice with 4 elements and the others will
// have only 3.
func SubsliceInt(s []int, n int) (ret [][]int) {
	for ; n > 0; n-- {
		s, ret = s[len(s)/n:], append(ret, s[:len(s)/n])
	}
	return ret
}
