package main

/*
The beginning of each link is the end of the previous one, except
for the first one, its beginning is the middle of the screen
*/

import (
	"math"
)

func (m *Manipulator) calculate_movement(){
	for index, elem := range m.robotStruct{
		if index == 0{
			elem["beginX"] = m.centerX
			elem["beginY"] = m.centerY
			elem["endX"] = elem["beginX"] + math.Cos(elem["angle"])*elem["length"]
			elem["endY"] = elem["beginY"] + math.Sin(elem["angle"])*elem["length"]
			} else {
			elem["beginX"] = m.robotStruct[index-1]["endX"]
			elem["beginY"] = m.robotStruct[index-1]["endY"]
			elem["endX"] = elem["beginX"] + math.Cos(elem["angle"])*elem["length"]
			elem["endY"] = elem["beginY"] + math.Sin(elem["angle"])*elem["length"]
		}
		elem["angle"]+=elem["speed"]
		elem["speed"]+=elem["acceler"]
	}
	//fmt.Println("calc")
}
