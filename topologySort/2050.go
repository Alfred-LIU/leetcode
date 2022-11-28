package topologySort

/*
You are given an integer n, which indicates that there are n courses labeled from 1 to n. You are also given a 2D integer array relations where relations[j] = [prevCoursej, nextCoursej] denotes that course prevCoursej has to be completed before course nextCoursej (prerequisite relationship). Furthermore, you are given a 0-indexed integer array time where time[i] denotes how many months it takes to complete the (i+1)th course.

You must find the minimum number of months needed to complete all the courses following these rules:

You may start taking a course at any time if the prerequisites are met.
Any number of courses can be taken at the same time.
Return the minimum number of months needed to complete all the courses.

Note: The test cases are generated such that it is possible to complete every course (i.e., the graph is a directed acyclic graph).
*/

func minimumTime(n int, relations [][]int, time []int) int {
	neighbour := make(map[int][]int)
	ins := make([]int, n+1)
	resNodes := make([]int, n+1)
	res := -1

	for _, r := range relations {
		if _, ok := neighbour[r[0]]; !ok {
			ns := make([]int, 0)
			neighbour[r[0]] = ns
		}
		neighbour[r[0]] = append(neighbour[r[0]], r[1])
		ins[r[1]] += 1
	}

	queue := make([]int, 0)
	for idx, in := range ins {
		if idx == 0 {
			continue
		}
		if in == 0 {
			queue = append(queue, idx)
			resNodes[idx] = time[idx-1]
		}
	}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		paths := neighbour[current]

		for _, p := range paths {
			if time[p-1]+resNodes[current] > resNodes[p] {
				resNodes[p] = time[p-1] + resNodes[current]
			}
			ins[p]--
			if ins[p] == 0 {
				queue = append(queue, p)
			}

		}
	}

	for _, v := range resNodes {
		if v > res {
			res = v
		}
	}

	return res
}
