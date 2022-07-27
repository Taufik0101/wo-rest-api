package dto

type Update struct {
	Name  		string 	`json:"name,omitempty" form:"name"`
	Password  	string 	`json:"password" form:"password"`
	Role      	string 	`json:"role" form:"role"`
	Email     	string 	`json:"email,omitempty" form:"email"`
	Profil		string	`json:"profil" form:"profil"`
	NoHP		uint64	`json:"no_hp" form:"no_hp"`
	Alamat		string	`json:"alamat" form:"alamat"`
	Toko		string	`json:"toko" form:"toko"`				//vendor
	Category  	uint32 	`json:"category" form:"category"`		//Vendor
	Kota      	uint32 	`json:"kota" form:"kota"`				//vendor
	KotaBaru    string 	`json:"kota_baru" form:"kota_baru"`		//vendor
	Bank1		string	`json:"bank_1" form:"bank_1"`			//vendor
	Rekening1	uint64	`json:"rekening_1" form:"rekening_1"`	//vendor
	Bank2		string	`json:"bank_2" form:"bank_2"`			//vendor
	Rekening2	uint64	`json:"rekening_2" form:"rekening_2"`	//vendor
}

type CreateUser struct {
	Name  		string 	`json:"name,omitempty" form:"name"`
	Role      	string 	`json:"role" form:"role"`
	Email     	string 	`json:"email,omitempty" form:"email"`
	Profil		string	`json:"profil" form:"profil"`
	NoHP		uint64	`json:"no_hp" form:"no_hp"`
	Alamat		string	`json:"alamat" form:"alamat"`
	Toko		string	`json:"toko" form:"toko"`				//vendor
	Category  	uint32 	`json:"category" form:"category"`		//Vendor
	Kota      	uint32 	`json:"kota" form:"kota"`				//vendor
	Bank1		string	`json:"bank_1" form:"bank_1"`			//vendor
	Rekening1	uint64	`json:"rekening_1" form:"rekening_1"`	//vendor
	Bank2		string	`json:"bank_2" form:"bank_2"`			//vendor
	Rekening2	uint64	`json:"rekening_2" form:"rekening_2"`	//vendor
}
