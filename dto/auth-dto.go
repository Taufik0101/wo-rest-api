package dto

type Register struct {
	Name  		string 	`json:"name" form:"name" binding:"required"`
	Role      	string 	`json:"role" form:"role"`
	NoHP		uint64	`json:"no_hp" form:"no_hp"`
	NIK			uint64	`json:"nik" form:"nik"`
	Email     	string 	`json:"email" form:"email" binding:"required"`
	Password  	string 	`json:"password" form:"password" binding:"required"`
	Category  	uint32 	`json:"category" form:"category"`
}

type Login struct {
	Email     	string 	`json:"email" form:"email" binding:"required"`
	Password  	string 	`json:"password" form:"password" binding:"required"`
}

type UPassword struct {
	OPassword  	string 	`json:"opassword" form:"opassword" binding:"required"`
	NPassword  	string 	`json:"npassword" form:"npassword" binding:"required"`
}

type Forgot struct {
	URL			string	`json:"url" form:"url" binding:"required"`
}
