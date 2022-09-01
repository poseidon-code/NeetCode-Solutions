// 853: Car Fleet
// https://leetcode.com/problems/car-fleet/

package main

import (
	"fmt"
	"sort"
)


type car struct {
	p int
	s int
}

// SOLUTION
func carFleet(target int, position []int, speed []int) int {
	cars := []car{}
	n := len(position)
	for i:=0; i<n; i++ {
		cars = append(cars, car{position[i], speed[i]})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].p < cars[j].p
	})

	s := []float64{}
	for i:=0; i<n; i++ {
		time := float64((target - cars[i].p) / cars[i].s)
		for len(s)!=0 && time >= s[len(s)-1] {s = s[:len(s)-1]}
		s = append(s, time)
	}

	return len(s);
}


func main() {
	// INPUT
	target := 12
    position := []int{10,8,0,5,3}
    speed := []int{2,4,1,1,3}

    // OUTPUT
    result := carFleet(target, position, speed)
    fmt.Println(result)
}