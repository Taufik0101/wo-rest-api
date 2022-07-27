package dto

type CreateKategori struct {
	Nama_kategori	string		`json:"nama_kategori" form:"nama_kategori"`
}

type UpdateKategori struct {
	Nama_kategori	string		`json:"nama_kategori" form:"nama_kategori"`
}
