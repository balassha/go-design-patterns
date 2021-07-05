package main

func GetGun(name string) iGun {
	if name == "ak47" {
		return Getak47()
	} else if name == "maverick" {
		return Getmaverick()
	}
	return nil
}
