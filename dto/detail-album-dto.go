package dto

type CreateDetailAlbum struct {
	IdAlbum		uint32		`json:"id_album" form:"id_album"`
	Foto		string		`json:"foto" form:"foto"`
}

type UpdateDetailAlbum struct {
	IdAlbum		uint32		`json:"id_album" form:"id_album"`
	Foto		string		`json:"foto" form:"foto"`
}
