package main

import (
	"fmt"
	"os"
)

type Manipulator struct {
	centerX     float64
	centerY     float64
	element     int
	robotStruct map[int]map[string]float64
}

func NewManipulator(
	wScr float64,
	hScr float64,
	x float64,
	y float64,
	elem int,
	elemLengths []float64,
	angList []float64,
	speedList []float64,
	accelerList []float64,
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
	rStruct := map[int]map[string]float64{}
	for i := 0; i < elem; i++ {
		//assign a length value to each elem
		rStruct[i] = map[string]float64{}
		rStruct[i]["length"] = float64(elemLengths[i])
		//assign the start and end values to the first element
		if i == 0{
			rStruct[i]["beginX"] = float64(x)
			rStruct[i]["beginY"] = float64(y)

			rStruct[i]["endX"] = float64(x+elemLengths[i])
			rStruct[i]["endY"] = float64(y+elemLengths[i])

		} else {
			//creating the beginning of the next link from the points of the end of the previous link
			rStruct[i]["beginX"] = float64(rStruct[i-1]["endX"])
			rStruct[i]["beginY"] = float64(rStruct[i-1]["endY"])
			//creating the end of the current link from the points of the end of the previous link
			//plus the length of the current link
			rStruct[i]["endX"] = float64(rStruct[i-1]["endX"]+elemLengths[i])
			rStruct[i]["endY"] = float64(rStruct[i-1]["endY"]+elemLengths[i])
		}
		rStruct[i]["speed"] = speedList[i]
		rStruct[i]["acceler"] = accelerList[i]
		rStruct[i]["angle"] = angList[i]

	}

	return &Manipulator{x, y, elem, rStruct}
}
