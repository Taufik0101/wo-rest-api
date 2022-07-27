package dto

type CreateAlbum struct {
	IdVendor				uint32		`json:"id_vendor" form:"id_vendor"`
	Nama					string		`json:"nama" form:"nama"`
	Tanggal_Pelaksanaan		string	`json:"tanggal_pelaksanaan" form:"tanggal_pelaksanaan"`
}

type UpdateAlbum struct {
	IdVendor				uint32		`json:"id_vendor" form:"id_vendor"`
	Nama					string		`json:"nama" form:"nama"`
	Tanggal_Pelaksanaan		string	`json:"tanggal_pelaksanaan" form:"tanggal_pelaksanaan"`
}
