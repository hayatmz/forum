package myFuncs

import (
	"strings"
	"slices"
)

func SliceByPrefix(s, prefix string) []string {
	var sliceByPrefix []string
	var sliceFields []string = strings.Fields(s)
	
	for _, str := range sliceFields {
		if strings.HasPrefix(str, prefix) {
			if !(len(str) <= 1) {
				sliceByPrefix = append(sliceByPrefix, str[1:])
			}
		}
	}
	return removeDoublonSlice(sliceByPrefix)
}

func removeDoublonSlice(sliceByPrefix []string) []string {
	var sliceWithoutDoublon []string

	for _, category := range sliceByPrefix {
		if !slices.Contains(sliceWithoutDoublon, category) {
			sliceWithoutDoublon = append(sliceWithoutDoublon, category)
		}
	}
	return sliceWithoutDoublon
}