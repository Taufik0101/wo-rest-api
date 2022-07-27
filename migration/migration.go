package migration

import (
	"github.com/Taufik0101/wo-rest-api/entity"
	"gorm.io/gorm"
)

type Migrator interface {
	DropTable()
	Migration()
}

type migratorConnection struct {
	connection *gorm.DB
}

func (db migratorConnection) DropTable() {
	if db.connection.Migrator().HasTable(&entity.Penilaian{}) {
		db.connection.Migrator().DropTable(&entity.Penilaian{})
	}
	if db.connection.Migrator().HasTable(&entity.Detail_Transaksi{}) {
		db.connection.Migrator().DropTable(&entity.Detail_Transaksi{})
	}
	if db.connection.Migrator().HasTable(&entity.Transaksi{}) {
		db.connection.Migrator().DropTable(&entity.Transaksi{})
	}
	if db.connection.Migrator().HasTable(&entity.Detail_Album{}) {
		db.connection.Migrator().DropTable(&entity.Detail_Album{})
	}
	if db.connection.Migrator().HasTable(&entity.Album{}) {
		db.connection.Migrator().DropTable(&entity.Album{})
	}
	if db.connection.Migrator().HasTable(&entity.Product{}) {
		db.connection.Migrator().DropTable(&entity.Product{})
	}
	if db.connection.Migrator().HasTable(&entity.Request{}) {
		db.connection.Migrator().DropTable(&entity.Request{})
	}
	if db.connection.Migrator().HasTable(&entity.User{}) {
		db.connection.Migrator().DropTable(&entity.User{})
	}
	if db.connection.Migrator().HasTable(&entity.Categories{}) {
		db.connection.Migrator().DropTable(&entity.Categories{})
	}
	if db.connection.Migrator().HasTable(&entity.Kota{}) {
		db.connection.Migrator().DropTable(&entity.Kota{})
	}
}

func (db migratorConnection) Migration() {
	db.connection.AutoMigrate(entity.Kota{}, entity.Categories{}, entity.User{}, entity.Request{}, entity.Product{}, entity.Album{}, entity.Detail_Album{}, entity.Transaksi{}, entity.Detail_Transaksi{}, entity.Penilaian{})
}

func NewMigrator(conn *gorm.DB) Migrator {
	return &migratorConnection{
		connection: conn,
	}
}