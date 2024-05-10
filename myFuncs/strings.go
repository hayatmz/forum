package myFuncs

import (
	"strings"
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
	return sliceByPrefix
}
