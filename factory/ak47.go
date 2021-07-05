package main

type ak47 struct {
	Gun
}

func Getak47() iGun {
	return ak47{
		Gun{
			Name:  "ak47",
			Power: 100,
		},
	}
}
