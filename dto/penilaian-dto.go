package dto

type CreatePenilaian struct {
	Seller				uint32		`json:"seller" form:"seller"`
	Customer			uint32		`json:"customer" form:"customer"`
	Judul				string		`json:"judul" form:"judul"`
	Pesan				string		`json:"pesan" form:"pesan"`
	Star				uint16		`json:"star" form:"star"`
	Trans				uint32		`json:"trans" form:"trans"`
}
