package main

import (
	"fmt"
	"sort"
)

type IntArray []int

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray []day

func (p dayArray) Len() int           { return len(p) }
func (p dayArray) Less(i, j int) bool { return p[i].num < p[j].num }
func (p dayArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntArray(data) //conversion to type IntArray from package sort
	sort.Sort(a)
	fmt.Println(a)

	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	data1 := []day{Tuesday, Thursday, Wednesday, Sunday, Monday, Friday, Saturday}
	d := dayArray(data1)
	sort.Sort(d)
	fmt.Print(d)

}
