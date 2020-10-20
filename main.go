package main

import (
	"H8/Golang12/Sesi1/models"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	profile := models.Profile{
		Nama: "dicky",
		Umur: 23,
		Alamat: models.Alamat{
			RT: "01",
		},
	}
	profile.Alamat.Kelurahan = "Rawamangun"
	profile.Nama = "Dicky Pratama"
	fmt.Println(profile)
	fmt.Println(uuid)
}
