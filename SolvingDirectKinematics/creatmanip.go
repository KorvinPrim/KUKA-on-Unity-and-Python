package main

import (
	"fmt"
	"os"
)

type Manipulator struct {
	centerX     int
	centerY     int
	element     int
	robotStruct map[int]map[string]int
	angleList 	map[int]float64
}

func NewManipulator(wScr int,
	hScr int,
	x int,
	y int,
	elem int,
	elemLengths []int,
	angList []float64,
) *Manipulator {
	if (x > wScr) || (y > hScr) || (x < 0) || (y < 0) {
		fmt.Println("The center of the manipulator is out of range of the screen")
		os.Exit(1)
	}

	if len(elemLengths) != elem {
		fmt.Println("The number of elements does not match the number of element lengths")
		os.Exit(1)
	}

	if len(angList) != elem {
		fmt.Println("The number of angles does not match the number of lengths of the elements")
		os.Exit(1)
	}
	rStruct := map[int]map[string]int{}
	rAng := map[int]float64{}
	for i := 0; i < elem; i++ {
		//assign a length value to each elem
		rStruct[i] = map[string]int{}
		rStruct[i]["length"] = elemLengths[i]
		//assign the start and end values to the first element
		if i == 0{
			rStruct[i]["beginX"] = x
			rStruct[i]["beginY"] = y

			rStruct[i]["endX"] = x+elemLengths[i]
			rStruct[i]["endY"] = y+elemLengths[i]
		} else {
			//creating the beginning of the next link from the points of the end of the previous link
			rStruct[i]["beginX"] = rStruct[i-1]["endX"]
			rStruct[i]["beginY"] = rStruct[i-1]["endY"]
			//creating the end of the current link from the points of the end of the previous link
			//plus the length of the current link

			rStruct[i]["endX"] = rStruct[i-1]["endX"]+elemLengths[i]
			rStruct[i]["endY"] = rStruct[i-1]["endY"]+elemLengths[i]
		}

		rAng[i] = angList[i]

	}

	return &Manipulator{x, y, elem, rStruct, rAng}
}
