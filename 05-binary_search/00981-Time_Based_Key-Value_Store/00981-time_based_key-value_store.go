// 981: Time Based Key-Value Store
// https://leetcode.com/problems/time-based-key-value-store/

package main

import (
	"fmt"
	"sort"
)

type Data struct {
    val string
    time int
}

type TimeMap struct {
    tm map[string][]Data
}


func Constructor() TimeMap {
    return TimeMap{tm: make(map[string][]Data)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.tm[key] = append(
        this.tm[key], 
        Data{time: timestamp,val: value},
    )
}


func (this *TimeMap) Get(key string, timestamp int) string {
    var list []Data
    var ok bool
    
    if list, ok = this.tm[key]; !ok {
        return ""
    }

    index := sort.Search(len(list), func(i int) bool {
        return list[i].time > timestamp
    })
    
    if index == 0 {
        return ""
    } else {
        return list[index - 1].val
    }  
}

func main() {
    // OPERATIONS
    o := Constructor()
    o.Set("foo", "bar", 1)
    fmt.Println(o.Get("foo", 1))
    fmt.Println(o.Get("foo", 3))
    o.Set("foo", "bar2", 4)
    fmt.Println(o.Get("foo", 4))
    fmt.Println(o.Get("foo", 5))
}