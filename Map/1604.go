package Map

import (
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	hash := map[string][]int{}
	for i, name := range keyName {
		time := minute(keyTime[i])
		hash[name] = append(hash[name], time)
	}
	names := []string{}
	for k, v := range hash {
		if warning(v) {
			names = append(names, k)
		}
	}
	sort.Strings(names)
	return names
}

func minute(timeStr string) int {
	hour := strings.Split(timeStr, ":")[0]
	minu := strings.Split(timeStr, ":")[1]
	h, _ := strconv.Atoi(hour)
	m, _ := strconv.Atoi(minu)
	return 60*h + m
}

func warning(minutes []int) bool {
	sort.Ints(minutes)
	for i := 2; i < len(minutes); i++ {
		diff := minutes[i] - minutes[i-2]
		if diff <= 60 {
			return true
		}
	}
	return false
}
