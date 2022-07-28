package migration

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

type Seeder interface {
	SeedUser()
	SeedKategori()
	SeedKota()
	SeedProduk()
	SeedAlbum()
	SeedDetailAlbum()
	SeedPenilaian()
	SeedTransaksi()
	SeedDetailTransaksi()
}

type seederConnection struct {
	connection *gorm.DB
}

func (s seederConnection) SeedUser() {
	var password = "@Passwordku"
	var password1 = "@Bismillah"
	var password2 = "@Vendors"
	var password3 = "@123456"
	hashing, errPwd := repository.Hash(password)
	if errPwd != nil {
		panic(errPwd)
	}
	pwd := string(hashing)
	admin := entity.User{
		Name:      "Admin",
		Role:      "Admin",
		Email:     "admin@gmail.com",
		Password:  pwd,
	}
	hashed, errPwd1 := repository.Hash(password1)
	if errPwd1 != nil {
		panic(errPwd1)
	}
	Pwd1 := string(hashed)
	vendor := entity.User{
		Name:      "John",
		Role:      "Vendor",
		Email:     "altairphotography@gmail.com",
		Profil:    "altair.jpg",
		NoHP:      62812347180,
		Alamat:    "Jalanin Aja Dulu",
		Toko:      "Altair Photography",
		Category:  1,
		Kota:      1,
		Bank1:     "BCA",
		Rekening1: 212312,
		Bank2:     "BRI",
		Rekening2: 123,
		Password:  Pwd1,
	}
	hashes, errPwd2 := repository.Hash(password2)
	if errPwd2 != nil {
		panic(errPwd2)
	}
	Pwd2 := string(hashes)
	user := entity.User{
		Name:      "Alexander",
		Role:      "User",
		Email:     "hid09h@gmail.com",
		NoHP:      6291283123,
		Alamat:    "Jalan Mulu Jadian Kaga",
		Password:  Pwd2,
	}
	hasheds, errPwd3 := repository.Hash(password3)
	if errPwd3 != nil {
		panic(errPwd3)
	}
	Pwd3 := string(hasheds)
	vendor1 := entity.User{
		Name:      "Muhammad Taufik Hidayanto",
		Role:      "Vendor",
		Email:     "hid09h1@gmail.com",
		NoHP:      6291283123,
		Category: 1,
		NIK: 123123,
		Password:  Pwd3,
	}
	s.connection.Create(&admin)
	s.connection.Create(&vendor)
	s.connection.Create(&user)
	s.connection.Create(&vendor1)
}

func (s seederConnection) SeedKategori() {
	category := entity.Categories{
		Nama_kategori: "Venue",
	}
	category1 := entity.Categories{
		Nama_kategori: "Bridal dan Jas Pengantin",
	}
	category2 := entity.Categories{
		Nama_kategori: "Make Up Artist",
	}
	category3 := entity.Categories{
		Nama_kategori: "Catering",
	}
	category4 := entity.Categories{
		Nama_kategori: "Foto dan Videografer",
	}
	category5 := entity.Categories{
		Nama_kategori: "Undangan dan Souvenir",
	}
	category6 := entity.Categories{
		Nama_kategori: "Wedding Organizer",
	}
	s.connection.Create(&category)
	s.connection.Create(&category1)
	s.connection.Create(&category2)
	s.connection.Create(&category3)
	s.connection.Create(&category4)
	s.connection.Create(&category5)
	s.connection.Create(&category6)
}

func (s seederConnection) SeedKota() {
	Kota1 := entity.Kota{
		Nama_Kota: "Yogyakarta",
	}
	Kota2 := entity.Kota{
		Nama_Kota: "Surakarta",
	}
	s.connection.Create(&Kota1)
	s.connection.Create(&Kota2)
}

func (s seederConnection) SeedProduk() {
	Produk := entity.Product{
		Seller:     2,
		Nama_Produk: "Silver Wedding Package",
		Harga:       "4000000",
		DP:          30,
		Image:       "altair.jpg",
		Deskripsi:   "Custom Wedding Package :\n\n- Full Day Documentation 1 Crew for Photo + 1 Crew for Video\n- Full Video 1-2 Hours Duration\n- 25 Edit Photos\n- 1 Minute Cinematic Video for Instagram\n- Free All RAW Files\n- Flashdisk for all raw files photos",
		Is_Active: "1",
	}
	Produk1 := entity.Product{
		Seller:     4,
		Nama_Produk: "Gold Wedding Package",
		Harga:       "4000000",
		DP:          30,
		Image:       "altair.jpg",
		Deskripsi:   "Custom Wedding Package :\n\n- Full Day Documentation 1 Crew for Photo + 1 Crew for Video\n- Full Video 1-2 Hours Duration\n- 25 Edit Photos\n- 1 Minute Cinematic Video for Instagram\n- Free All RAW Files\n- Flashdisk for all raw files photos",
		Is_Active: "1",
	}
	s.connection.Create(&Produk)
	s.connection.Create(&Produk1)
}

func (s seederConnection) SeedAlbum() {
	Album := entity.Album{
		IdVendor:             2,
		Nama:                "Wedding X and Y",
		Tanggal_Pelaksanaan: carbon.Parse(carbon.Now().ToDateTimeString()).Time,
	}
	s.connection.Create(&Album)
}

func (s seederConnection) SeedDetailAlbum() {
	DetailAlbum := entity.Detail_Album{
		IdAlbum:  1,
		Foto:      "altair.jpg",
	}
	s.connection.Create(&DetailAlbum)
}

func (s seederConnection) SeedPenilaian() {
	Penilaian := entity.Penilaian{
		Seller:       2,
		Customer:     3,
		Trans:        1,
		Judul:        "Wedding X dan Y",
		Pesan:        "Pelayanan Cukup Bagus",
		Star:         5,
	}
	s.connection.Create(&Penilaian)
}

func (s seederConnection) SeedTransaksi() {
	Transaksi := entity.Transaksi{
		Customer:       3,
		TotalPrice:     2000000,
		Status:         "2",
	}
	s.connection.Create(&Transaksi)
}

func (s seederConnection) SeedDetailTransaksi() {
	DetailTX := entity.Detail_Transaksi{
		IdTransaksi:   1,
		Customer:       3,
		Seller:         2,
		Produk:         1,
		Pax:            1,
		DownPayment:    2000000,
		TanggalRes:     carbon.Parse(carbon.Now().ToDateTimeString()).Time,
		//PaymentMethods: "bca",
		Status:         "2",
	}
	s.connection.Create(&DetailTX)
}

func NewSeeder(conn *gorm.DB) Seeder {
	return &seederConnection{
		connection: conn,
	}
}