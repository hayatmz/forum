package myFuncs

import (
	"fmt"
	"strings"
)

func SliceByPrefix(s, prefix string) []string {
	var sliceByPrefix []string
	var sliceFields []string = strings.Fields(s)
	
	for _, str := range sliceFields {
		if strings.HasPrefix(str, prefix) {
			sliceByPrefix = append(sliceByPrefix, str[1:])
		}
	}
	fmt.Println(sliceByPrefix)
	return sliceByPrefix
}
