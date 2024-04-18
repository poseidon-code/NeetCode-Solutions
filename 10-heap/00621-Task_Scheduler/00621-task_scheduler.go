// 621: Task Scheduler
// https://leetcode.com/problems/task-scheduler

package main

import "fmt"

// SOLUTION
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	counts := make([]int, 26)
	for _, task := range tasks {
		counts[task-'A']++
	}

	var max_count, time int
	for _, c := range counts {
		if c > max_count {
			max_count = c
			time = 1
		} else if c == max_count {
			time++
		}
	}

	result := (n+1)*(max_count-1) + time

	if result > len(tasks) {
		return result
	} else {
		return len(tasks)
	}
}

func main() {
	// INPUT
	tasks := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n := 1

	// OUTPUT
	result := leastInterval(tasks, n)
	fmt.Println(result)
}
