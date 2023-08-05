package main

/*
The main package contains code for performing conversions
between degrees and radians, as well as a function for
direct kinematics calculations for a manipulator.

The code defines two types, Radian and Degree, which
represent angles in radians and degrees respectively.
Each type has methods for converting to the other type
and returning the angle as a float64 value.

The package also includes two functions, degreesToRadians
and radiansToDegrees, which perform simple conversions
between degrees and radians using mathematical formulas.

The direct_kinematics function takes in the width and
height of a window and calculates the position of a
manipulator's hand using the given robot structure.
It iterates through the robot structure, updating the
hand position based on the length and angle of each segment.
The final position of the hand is returned as a map with keys
"xHand", "yHand", and "QHand" representing the x-coordinate,
y-coordinate, and angle in degrees respectively.

The commented out main function demonstrates the usage
of the conversion functions and the direct_kinematics function.
*/

import (
	"math"
)

type Radian float64

func (rad Radian) ToDegrees() Degree {
	return Degree(float64(rad) * (180.0 / math.Pi))
}

func (rad Radian) Float64() float64 {
	return float64(rad)
}

type Degree float64

func (deg Degree) ToRadians() Radian {
	return Radian(float64(deg) * (math.Pi / 180.0))
}

func (deg Degree) Float64() float64 {
	return float64(deg)
}

func degreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func radiansToDegrees(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

func (m *Manipulator) direct_kinematics(wind_width float64, wind_height float64) map[string]float64 {
	var xHand float64 = wind_width / 2
	var yHand float64 = wind_height / 2
	var QHand float64

	for i := 0; i < len(m.robotStruct); i++ {
		if i == 0 {
			QHand += 2*math.Pi - m.robotStruct[i]["angle"]
		} else {
			QHand += m.robotStruct[i-1]["angle"] - m.robotStruct[i]["angle"]

		}
		xHand += m.robotStruct[i]["length"] * math.Cos(QHand)
		yHand += m.robotStruct[i]["length"] * math.Sin(QHand)

	}

	return map[string]float64{"xHand": xHand, "yHand": yHand, "QHand": radiansToDegrees(QHand)}

}

//func main() {
//
//	val := radiansToDegrees(1)
//	fmt.Printf("Один радиан равен : %.4f градусов\n", val)
//
//	val2 := degreesToRadians(val)
//	fmt.Printf("%.4f градусов это %.4f радиан\n", val, val2)
//
//	// Конвертация как часть
//	// типа методов
//	val = Radian(1).ToDegrees().Float64()
//	fmt.Printf("Градусы: %.4f градусов\n", val)
//
//	val = Degree(val).ToRadians().Float64()
//	fmt.Printf("Радианы: %.4f радиан\n", val)
//}
