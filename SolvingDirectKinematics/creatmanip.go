package main

/*
The code defines a struct called "Manipulator" which represents a robot arm manipulator.
The struct has several fields including the center coordinates, the number of elements
in the arm, a map representing the structure of the arm, and a map representing the
working area of the arm.

The code also defines a function called "NewManipulator" which is a constructor
for creating a new instance of the Manipulator struct. The function takes several
parameters including the screen width and height, the center coordinates of the
manipulator, the number of elements in the arm, and lists of element lengths,
angles, speeds, and accelerations.

The function first checks if the center coordinates are within the screen boundaries.
If they are not, it prints an error message and exits the program. Then it
checks if the number of element lengths and angles match the number of elements
in the arm. If they don't match, it prints an error message and exits the program.

Next, the function creates a map called "rStruct" to represent the structure of
the arm. It iterates over the number of elements and assigns the length value
to each element. For the first element, it assigns the start and end coordinates
based on the center coordinates and the length of the element. For the
subsequent elements, it calculates the start and end coordinates based
on the end coordinates of the previous element and the length of the
current element. It also assigns the speed, acceleration, and angle values to each element.

After that, the function calculates the working area of the arm.
It iterates over the elements in the "rStruct" map and accumulates
the total distance by adding up the lengths of all elements. Then
it creates a new map called "working_area" and assigns the distance,
as well as the minimum and maximum values of the X and Y coordinates
based on the center coordinates and the total distance.

Finally, the function returns a new instance of the Manipulator
struct with the center coordinates, number of elements, the
"rStruct" map, and the "working_area" map.

In summary, the code defines a struct for a robot arm manipulator
and provides a constructor function to create a new instance of
the struct with specified parameters. The constructor initializes
the structure and working area of the arm based on the given parameters.
*/

import (
	"fmt"
	"os"
)

type Manipulator struct {
	centerX      float64
	centerY      float64
	element      int
	robotStruct  map[int]map[string]float64
	working_area map[string]float64
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
	//Working with the links of the robot arm
	rStruct := map[int]map[string]float64{}
	for i := 0; i < elem; i++ {
		//assign a length value to each elem
		rStruct[i] = map[string]float64{}
		rStruct[i]["length"] = float64(elemLengths[i])
		//assign the start and end values to the first element
		if i == 0 {
			rStruct[i]["beginX"] = float64(x)
			rStruct[i]["beginY"] = float64(y)

			rStruct[i]["endX"] = float64(x + elemLengths[i])
			rStruct[i]["endY"] = float64(y + elemLengths[i])

		} else {
			//creating the beginning of the next link from the points of the end of the previous link
			rStruct[i]["beginX"] = float64(rStruct[i-1]["endX"])
			rStruct[i]["beginY"] = float64(rStruct[i-1]["endY"])
			//creating the end of the current link from the points of the end of the previous link
			//plus the length of the current link
			rStruct[i]["endX"] = float64(rStruct[i-1]["endX"] + elemLengths[i])
			rStruct[i]["endY"] = float64(rStruct[i-1]["endY"] + elemLengths[i])
		}
		rStruct[i]["speed"] = speedList[i]
		rStruct[i]["acceler"] = accelerList[i]
		rStruct[i]["angle"] = angList[i]

	}
	//Сalculation of the working area of the robot arm
	distance := 0.0
	working_area := make(map[string]float64, 5)
	for _, elem := range rStruct {
		distance += elem["length"]
	}

	working_area["Distance"] = distance

	working_area["Xmin"] = x - distance
	working_area["Xmax"] = x + distance

	working_area["Ymin"] = y - distance
	working_area["Ymax"] = y + distance

	return &Manipulator{x, y, elem, rStruct, working_area}
}
