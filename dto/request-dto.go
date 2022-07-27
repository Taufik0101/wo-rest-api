package dto

type CreateRequest struct {
	Name  		string 	`json:"nama" form:"nama" binding:"required"`
	Role      	string 	`json:"role" form:"role"`
	Email     	string 	`json:"email" form:"email" binding:"required"`
	NoHP		uint64	`json:"no_hp" form:"no_hp"`
	NIK			uint64	`json:"nik" form:"nik"`
	Category  	uint32 	`json:"id_kategori" form:"id_kategori" binding:"required"`
	KTP			string 	`json:"ktp" form:"ktp"`
	Selfie    	string 	`json:"selfie_ktp" form:"selfie_ktp"`
	Status      string 	`json:"status" form:"status"`
}

type UpdateRequest struct {
	Name  		string `json:"nama" form:"nama"`
	Role      	string `json:"role" form:"role"`
	Email     	string `json:"email" form:"email"`
	NoHP		uint64	`json:"no_hp" form:"no_hp"`
	NIK			uint64	`json:"nik" form:"nik"`
	Category  	uint32 `json:"id_kategori" form:"id_kategori"`
	KTP			string `json:"ktp" form:"ktp"`
	Selfie    	string `json:"selfie_ktp" form:"selfie_ktp"`
	Status      string `json:"status" form:"status"`
}

type Email struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
