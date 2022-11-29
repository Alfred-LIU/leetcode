package Math

/*
Given a characters array tasks, representing the tasks a CPU needs to do, where each letter represents a different task. Tasks could be done in any order. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer n that represents the cooldown period between two same tasks (the same letter in the array), that is that there must be at least n units of time between any two same tasks.

Return the least number of units of times that the CPU will take to finish all the given tasks.
*/

func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	for _, task := range tasks {
		if v, ok := m[task]; !ok {
			m[task] = 1
		} else {
			m[task] = v + 1
		}
	}

	max := 0
	maxGroup := 0

	for _, v := range m {
		if v > max {
			max = v
			maxGroup = 1
		} else if v == max {
			maxGroup++
		}
	}

	return maxValue(len(tasks), (max-1)*(n+1)+maxGroup)
}

func maxValue(i, j int) int {
	if i > j {
		return i
	}
	return j
}
