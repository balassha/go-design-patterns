package main

type maverick struct {
	Gun
}

func Getmaverick() iGun {
	return maverick{
		Gun{
			Name:  "maverick",
			Power: 101,
		},
	}
}
