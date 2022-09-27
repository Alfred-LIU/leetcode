package binary_search

type elem struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]elem
}

func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]elem)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = []elem{}
	}

	this.m[key] = append(this.m[key], elem{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.m[key]; !ok {
		return ""
	}

	index := findFirst(this.m[key], timestamp)

	if index < 0 {
		return ""
	}

	return this.m[key][index].value
}

func findFirst(input []elem, target int) int {
	left, right := 0, len(input)-1
	for left <= right {
		mid := (right-left)/2 + left
		if input[mid].timestamp == target {
			return mid
		} else if input[mid].timestamp < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
