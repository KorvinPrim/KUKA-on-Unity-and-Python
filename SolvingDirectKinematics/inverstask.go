package main

func (m *Manipulator) invers_kinematics() {
	for _, elem := range m.robotStruct {
		elem["speed"] = 0.0
		elem["acceler"] = 0.0
	}

}
