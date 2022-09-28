package stack

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	stack = append(stack, asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			stack = append(stack, asteroids[i])
		} else {
			add := true
			for len(stack) > 0 {
				if stack[len(stack)-1] > 0 && asteroids[i]+stack[len(stack)-1] < 0 {
					stack = stack[:len(stack)-1]
					add = true
				} else if stack[len(stack)-1] > 0 && asteroids[i]+stack[len(stack)-1] == 0 {
					stack = stack[:len(stack)-1]
					add = false
					break
				} else if stack[len(stack)-1] < 0 {
					add = true
					break
				} else {
					add = false
					break
				}
			}

			if add {
				stack = append(stack, asteroids[i])
			}
		}
	}

	return stack
}
