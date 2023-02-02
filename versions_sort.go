package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ByVersion []string

func (v ByVersion) Len() int {
	return len(v)
}

func (v ByVersion) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v ByVersion) Less(i, j int) bool {
	var length, lValue, rValue int
	version1 := strings.Split(v[i], ".")
	version2 := strings.Split(v[j], ".")
	len1, len2 := len(version1), len(version2)
	length = len2
	if len1 > len2 {
		length = len1
	}

	for indx := 0; indx < length; indx++ {
		if indx < len1 && indx < len2 {
			if version1[indx] == version2[indx] {
				continue
			}
		}

		lValue = 0
		if indx < len1 {
			lValue, _ = strconv.Atoi(version1[indx])
		}
		rValue = 0
		if indx < len2 {
			rValue, _ = strconv.Atoi(version2[indx])
		}

		if lValue < rValue {
			return true
		} else if lValue > rValue {
			return false
		}
	}

	return false
}

func main() {
	versions := []string{"5.6.9", "2.3.5", "5.6.10", "4.5.8"}
	fmt.Println("Versions before sorting")
	fmt.Println(strings.Join(versions, ", "))
	sort.Sort(ByVersion(versions))
	fmt.Println("Versions after sorting")
	fmt.Println(strings.Join(versions, ", "))
}
