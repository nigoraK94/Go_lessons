package main

import "fmt"

func main() {

type Student struct{
	Name string
    Grades [] int
}

func (s Student) String() string {
	return s.Name
}

func (s Student) Average () float64 {
	if len(s.Grades) == 0 {
		return 0
	}

total := 0

for_, nums := range s.Grades {
	total += nums
}

avg := float64(total) / 


}


}