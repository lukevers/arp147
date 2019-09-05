package entity

import (
	"github.com/Pallinder/go-randomdata"
)

type data struct {
	IPAddr    string
	Abandoned bool
}

func datagen() *data {
	return &data{
		IPAddr:    randomdata.IpV4Address(),
		Abandoned: randomdata.Boolean(),
	}
}
