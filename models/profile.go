package models

type Profile struct {
	Nama     string
	Umur     int
	Alamat   Alamat
	Password string
}

type Alamat struct {
	RT            string
	RW            string
	Kelurahan     string
	AlamatLengkap string
}
