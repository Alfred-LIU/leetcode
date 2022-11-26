package swipLine

import (
	"sort"
)

/*
We are given a list schedule of employees, which represents the working time for each employee.

Each employee has a list of non-overlapping Intervals, and these intervals are in sorted order.

Return the list of finite intervals representing common, positive-length free time for all employees, also in sorted order.

(Even though we are representing Intervals in the form [x, y], the objects inside are Intervals, not lists or arrays. For example, schedule[0][0].start = 1, schedule[0][0].end = 2, and schedule[0][0][0] is not defined).  Also, we wouldn't include intervals like [5, 5] in our answer, as they have zero length.



Example 1:

Input: schedule = [[[1,2],[5,6]],[[1,3]],[[4,10]]]
Output: [[3,4]]
Explanation: There are a total of three employees, and all common
free time intervals would be [-inf, 1], [3, 4], [10, inf].
We discard any intervals that contain inf as they aren't finite.
Example 2:

Input: schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
Output: [[5,6],[7,9]]
*/

type Interval struct {
	Start int
	End   int
}

type element struct {
	time  int
	value int
}

type elementList []element

func (e elementList) Len() int {
	return len(e)
}

func (e elementList) Less(i, j int) bool {
	if e[i].time == e[j].time {
		return e[i].value < e[j].value
	}
	return e[i].time < e[j].time
}

func (e elementList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	elements := make([]element, 0)
	for _, person := range schedule {
		for _, i := range person {
			eStart := element{
				time:  i.Start,
				value: 1,
			}
			elements = append(elements, eStart)
			eEnd := element{
				time:  i.End,
				value: -1,
			}
			elements = append(elements, eEnd)
		}
	}

	res := make([]*Interval, 0)
	sort.Sort(elementList(elements))
	pre, total, start, end := -1, 0, 0, 0

	for _, e := range elements {
		total += e.value
		if total == 0 {
			start = e.time
		}
		if pre == 0 {
			end = e.time
			if end != start {
				res = append(res, &Interval{
					Start: start,
					End:   end,
				})
			}
		}
		pre = total
	}

	return res
}
