package dto

type CreateKota struct {
	Nama_Kota	string	`json:"nama_kota" form:"nama_kota"`
}

type UpdateKota struct {
	Nama_Kota	string	`json:"nama_kota" form:"nama_kota"`
}
