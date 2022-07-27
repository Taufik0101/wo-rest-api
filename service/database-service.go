package service

import "github.com/Taufik0101/wo-rest-api/migration"

type DatabaseService interface {
	SeedAll()
}

type databaseService struct {
	databaseSeeder migration.Seeder
}

func (d databaseService) SeedAll() {
	d.databaseSeeder.SeedKota()
	d.databaseSeeder.SeedKategori()
	d.databaseSeeder.SeedUser()
	d.databaseSeeder.SeedProduk()
	d.databaseSeeder.SeedAlbum()
	d.databaseSeeder.SeedDetailAlbum()
	d.databaseSeeder.SeedTransaksi()
	d.databaseSeeder.SeedDetailTransaksi()
	d.databaseSeeder.SeedPenilaian()
}

func NewDatabaseService(seeder migration.Seeder) DatabaseService {
	return &databaseService{
		databaseSeeder: seeder,
	}
}