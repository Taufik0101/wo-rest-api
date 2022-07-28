package dto

type CreateProduk struct {
	Seller			uint32		`json:"seller" form:"seller"`
	Nama_Produk		string		`json:"nama_produk" form:"nama_produk"`
	Harga			string		`json:"harga" form:"harga"`
	DP				uint16		`json:"dp" form:"dp"`
	Image			string		`json:"image" form:"image"`
	Deskripsi		string		`json:"deskripsi" form:"deskripsi"`
}

type UpdateProduk struct {
	Seller			uint32		`json:"seller" form:"seller"`
	Nama_Produk		string		`json:"nama_produk" form:"nama_produk"`
	Harga			string		`json:"harga" form:"harga"`
	DP				uint16		`json:"dp" form:"dp"`
	Image			string		`json:"image" form:"image"`
	Deskripsi		string		`json:"deskripsi" form:"deskripsi"`
	Is_Active		string		`json:"is_active" form:"is_active"`
}
