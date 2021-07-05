package main

import "fmt"

func main() {
	ak47 := GetGun("ak47")
	if ak47 == nil {
		fmt.Println("Error while creating AK47")
	}
	PrintDetails(ak47)
	maverick := GetGun("maverick")
	if maverick == nil {
		fmt.Println("Error while creating Maverick")
	}
	PrintDetails(maverick)
}

func PrintDetails(g iGun) {
	fmt.Printf("Name : %s\n", g.GetName())
	fmt.Printf("Power : %d\n", g.GetPower())
}
