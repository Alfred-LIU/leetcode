package topologySort

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	path := make(map[int][]int)
	inDegree := make(map[int]int)

	for i := 0; i < numCourses; i++ {
		inDegree[i] = 0
		path[i] = make([]int, 0)
	}

	for _, p := range prerequisites {
		path[p[1]] = append(path[p[1]], p[0])
		inDegree[p[0]] += 1
	}

	queue := make([]int, 0)
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)

		}
	}

	res := make([]int, 0)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		for _, p := range path[cur] {
			inDegree[p] -= 1
			if inDegree[p] == 0 {
				queue = append(queue, p)
			}
		}
	}

	return len(res) == numCourses
}
