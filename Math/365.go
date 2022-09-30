package Math

/*

You are given two jugs with capacities jug1Capacity and jug2Capacity liters. There is an infinite amount of water supply available. Determine whether it is possible to measure exactly targetCapacity liters using these two jugs.

If targetCapacity liters of water are measurable, you must have targetCapacity liters of water contained within one or both buckets by the end.

Operations allowed:

Fill any of the jugs with water.
Empty any of the jugs.
Pour water from one jug into another till the other jug is completely full, or the first jug itself is empty.

If a or b is negative this means we are emptying a jug of x or y gallons respectively.

Similarly if a or b is positive this means we are filling a jug of x or y gallons respectively.

x = 4, y = 6, z = 8.

GCD(4, 6) = 2

8 is multiple of 2

so this input is valid and we have:

-1 * 4 + 6 * 2 = 8
*/

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == z || y == z {
		return true
	}
	return (z % gcd(x, y)) == 0
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
