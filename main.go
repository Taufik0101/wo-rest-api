package main

import (
	"fmt"
	"github.com/Taufik0101/wo-rest-api/config"
	"github.com/Taufik0101/wo-rest-api/controller/data"
	"github.com/Taufik0101/wo-rest-api/middleware"
	"github.com/Taufik0101/wo-rest-api/migration"
	"github.com/Taufik0101/wo-rest-api/repository"
	"github.com/Taufik0101/wo-rest-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/xendit/xendit-go/client"
	"gorm.io/gorm"
)

var (
	db							*gorm.DB							= config.SetupDatabaseConnection()
	cache						*redis.Client						= config.SetupRedisConnection()
	xen							*client.API							= config.SetupXenditConnection()
	Migrate						migration.Migrator					= migration.NewMigrator(db)
	Seed						migration.Seeder					= migration.NewSeeder(db)
	Seeder						service.DatabaseService				= service.NewDatabaseService(Seed)
	UserRepository				repository.UserRepository			= repository.NewUserRepository(db)
	AuthRepository				repository.AuthRepository			= repository.NewAuthRepository(db,UserRepository)
	KategoriRepository			repository.KategoriRepository		= repository.NewKategoriRepository(db)
	KotaRepository				repository.KotaRepository			= repository.NewKotaRepository(db)
	RequestRepository			repository.RequestRepository		= repository.NewRequestRepository(db)
	ProdukRepository			repository.ProdukRepository			= repository.NewProdukRepository(db)
	AlbumRepository				repository.AlbumRepository			= repository.NewAlbumRepository(db)
	DetailAlbumRepository		repository.DetailAlbumRepository	= repository.NewDetailAlbumRepository(db)
	TransaksiRepository			repository.TransaksiRepository		= repository.NewTransaksiRepository(db)
	DetailTransaksiRepository	repository.DetailTransaksiRepository= repository.NewDetailTransaksiRepository(db)
	PenilaianRepository			repository.PenilaianRepository		= repository.NewPenilaianRepository(db)
	JwtService					service.JWTService					= service.NewJWTService()
	RedisService				service.RedisService				= service.NewRedisService(cache)
	AuthService					service.AuthService					= service.NewAuthService(AuthRepository)
	UserService					service.UserService					= service.NewUserService(UserRepository)
	KategoriService				service.CategoriService				= service.NewKategoriService(KategoriRepository)
	KotaService					service.KotaService					= service.NewKotaService(KotaRepository)
	RequestService				service.RequestService				= service.NewRequestService(RequestRepository)
	ProdukService				service.ProdukService				= service.NewProdukService(ProdukRepository)
	AlbumService				service.AlbumService				= service.NewAlbumService(AlbumRepository)
	DetailAlbumService			service.DetailAlbumService			= service.NewDetailAlbumService(DetailAlbumRepository)
	TransaksiService			service.TransaksiService			= service.NewTransaksiService(TransaksiRepository)
	DetailTransaksiService		service.DetailTransaksiService		= service.NewDetailTransaksiService(DetailTransaksiRepository)
	PenilaianService			service.PenilaianService			= service.NewPenilaianService(PenilaianRepository)
	AuthControl					data.AuthController					= data.NewAuthController(AuthService,JwtService,RedisService, UserService)
	KategoriControl				data.KategoriController				= data.NewKategoriController(JwtService, KategoriService)
	KotaControl					data.KotaController					= data.NewKotaController(JwtService, KotaService)
	VendorControl				data.VendorController				= data.NewVendorController(UserService, JwtService)
	UserControl					data.UserController					= data.NewUserController(UserService, JwtService, KotaService)
	RequestControl				data.RequestController				= data.NewRequestController(JwtService, RequestService, AuthControl)
	ProdukControl				data.ProdukController				= data.NewProdukController(JwtService, ProdukService)
	AlbumControl				data.AlbumController				= data.NewAlbumController(JwtService, AlbumService)
	DetailAlbumControl			data.DetailAlbumController			= data.NewDetailAlbumController(JwtService, DetailAlbumService)
	DetailTransaksiControl		data.DetailTransaksiController		= data.NewDetailTransaksiController(JwtService, DetailTransaksiService)
	PenilaianControl			data.PenilaianController			= data.NewPenilaianController(JwtService, PenilaianService, DetailTransaksiService, TransaksiService)
	VAControl					data.VAController					= data.NewVAController(xen, TransaksiService, DetailTransaksiService)
)

func main()  {
	defer config.CloseDatabaseConnection(db)
	defer cache.Close()
	Migrate.DropTable()
	Migrate.Migration()
	Seeder.SeedAll()
	router := gin.Default()

	router.Use(CORSMiddleware())
	router.Static("/album", "./upload/images/album")
	router.Static("/ktp", "./upload/images/ktp")
	router.Static("/produk", "./upload/images/produk")
	router.Static("/profil", "./upload/images/profil")
	router.Static("/selfiektp", "./upload/images/selfie")
	//router.GET("/", KategoriControl.AllKategori)

	Auth := router.Group("api/auth")
	{
		Auth.POST("/login", AuthControl.Login)
		Auth.POST("/register", AuthControl.Register)
		Auth.POST("/forget", AuthControl.SendTokenForgot)
		Auth.POST("/reset", AuthControl.ResetPassword)
		Auth.POST("/logout", middleware.Auth(JwtService, UserService, RedisService), AuthControl.Logout)
	}

	User := router.Group("api/user")
	{
		User.GET("/:user_id", middleware.Auth(JwtService, UserService, RedisService), UserControl.FindByID)
		User.POST("/update/:user_id", middleware.Auth(JwtService, UserService, RedisService), UserControl.UpdateUser)
		User.POST("/updatePass/:user_id", middleware.Auth(JwtService, UserService, RedisService), UserControl.UpdatePassword)
	}

	Kategori := router.Group("api/kategori")
	{
		Kategori.GET("/list", KategoriControl.AllKategori)
		Kategori.POST("/add", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), KategoriControl.SimpanKategori)
		Kategori.POST("/update/:kategori_id", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), KategoriControl.UpdateKategori)
		Kategori.POST("/delete/:kategori_id", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), KategoriControl.DeleteKategori)
	}

	Kota := router.Group("api/kota")
	{
		Kota.GET("/list", KotaControl.AllKota)
		Kota.POST("/add", middleware.Auth(JwtService, UserService, RedisService), KotaControl.SimpanKota)
		Kota.POST("/update/:kota_id", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), KotaControl.UpdateKota)
		Kota.POST("/delete/:kota_id", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), KotaControl.DeleteKota)
	}

	Vendor := router.Group("api/vendor")
	{
		Vendor.GET("/list", VendorControl.AllVendor)
		Vendor.POST("/register", RequestControl.SimpanRequest)
		Vendor.GET("/filterKC/:KC", VendorControl.VendorByKategoryAndCity)
		Vendor.GET("/filterKota/:kota", VendorControl.VendorByCity)
		Vendor.GET("/filterKategori/:category", VendorControl.VendorByCategory)
	}

	Request := router.Group("api/request")
	{
		Request.GET("/list", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), RequestControl.AllRequest)
		Request.POST("/update/:req_id", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), RequestControl.UpdateRequest)
	}

	Produk := router.Group("api/produk")
	{
		Produk.GET("/listAll", ProdukControl.AllProduk)
		Produk.GET("/listRandom", ProdukControl.RandomProductLimit)
		Produk.GET("/list/:vendor_id", ProdukControl.ProdukByVendorUser)
		Produk.GET("/listforVendor/:vendor_id", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), ProdukControl.ProdukByVendor)
		Produk.POST("/add", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), ProdukControl.SimpanProduk)
		Produk.POST("/update/:produk_id", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), ProdukControl.UpdateProduk)
		Produk.POST("/publish/:produk_id", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), ProdukControl.PublishProduct)
		Produk.POST("/delete/:produk_id", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), ProdukControl.DeleteProduk)
	}

	Album := router.Group("api/album")
	{
		Album.GET("/list/:vendor_id", AlbumControl.AlbumByVendor)
		Album.POST("/add", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), AlbumControl.SimpanAlbum)
		Album.POST("/update/:album_id", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), AlbumControl.Update)
	}

	Koleksi := router.Group("api/koleksi")
	{
		Koleksi.GET("/list/:album_id", DetailAlbumControl.DetailAlbumByAlbum)
		Koleksi.POST("/add", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), DetailAlbumControl.SimpanDetailAlbum)
		Koleksi.POST("/delete/:detail_id", middleware.CheckRole(JwtService, UserService, RedisService, "Vendor"), DetailAlbumControl.HapusDetailAl)
	}

	Transaksi := router.Group("api/transaksi")
	{
		Transaksi.GET("/listAll", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), VAControl.AllTransaksi)
		Transaksi.GET("/list/:vendor_id", DetailTransaksiControl.DetailTransaksiSeller)
		Transaksi.POST("/add", middleware.CheckRole(JwtService, UserService, RedisService, "User"), VAControl.CreateInvoice)
		Transaksi.POST("/callback", VAControl.CallbackTransaksi)
		Transaksi.GET("/riwayat/:customer_id", middleware.CheckRole(JwtService, UserService, RedisService, "User"), VAControl.HistoryTransaksi)
		Transaksi.GET("/riwayatDetil/:transaksi_id", middleware.CheckRole(JwtService, UserService, RedisService, "User"), VAControl.DetailHistoryTransaksi)
		Transaksi.GET("/riwayatDetilAll/:transaksi_id", middleware.CheckRole(JwtService, UserService, RedisService, "Admin"), VAControl.DetailHistoryTransaksi)
	}

	Penilaian := router.Group("api/penilaian")
	{
		Penilaian.GET("/list/:vendor_id", PenilaianControl.PenilaianByVendor)
		Penilaian.POST("/add", middleware.CheckRole(JwtService, UserService, RedisService, "User"), PenilaianControl.SimpanPenilaian)
	}

	VA := router.Group("api/va")
	{
		VA.GET("/list", VAControl.AllVA)
	}

	router.Run(":1000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}