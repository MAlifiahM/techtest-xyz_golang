package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"time"
	"xyz_golang/internal/domain"
)

var db *gorm.DB

func dbSetup() {
	var err error
	l := gormLogger.Default.LogMode(gormLogger.Silent)
	db, err = gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{
		Logger: l,
	})

	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database")

	// enable uuid extension
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	// Run migration
	if cfg.IsDevelopment {
		if err := db.AutoMigrate(
			&domain.Consumer{},
			&domain.Limit{},
			&domain.Transaction{},
		); err != nil {
			panic(err)
		}
	}

	// Run seed
	var count int64
	db.Model(&domain.Consumer{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	log.Println("Seeding database")
	consumers := []domain.Consumer{
		{
			NIK:          "1234567890123456",
			FullName:     "Budi Santoso",
			LegalName:    "Budi Santoso",
			PlaceOfBirth: "Jakarta",
			DateOfBirth:  time.Date(1990, 1, 10, 0, 0, 0, 0, time.UTC),
			Salary:       5000000,
			PhotoKTP:     "budi_ktp.jpg",
			PhotoSelfie:  "budi_selfie.jpg",
		},
		{
			NIK:          "6543210987654321",
			FullName:     "Annisa Putri",
			LegalName:    "Annisa Putri",
			PlaceOfBirth: "Bandung",
			DateOfBirth:  time.Date(1992, 5, 20, 0, 0, 0, 0, time.UTC),
			Salary:       8000000,
			PhotoKTP:     "annisa_ktp.jpg",
			PhotoSelfie:  "annisa_selfie.jpg",
		},
	}

	if err = db.Create(&consumers).Error; err != nil {
		panic(err)
	}

	limits := []domain.Limit{
		{ConsumerID: consumers[0].ID, Tenor: 1, Amount: 100000},
		{ConsumerID: consumers[0].ID, Tenor: 2, Amount: 200000},
		{ConsumerID: consumers[0].ID, Tenor: 3, Amount: 500000},
		{ConsumerID: consumers[0].ID, Tenor: 6, Amount: 700000},
		{ConsumerID: consumers[1].ID, Tenor: 1, Amount: 1000000},
		{ConsumerID: consumers[1].ID, Tenor: 2, Amount: 1200000},
		{ConsumerID: consumers[1].ID, Tenor: 3, Amount: 1500000},
		{ConsumerID: consumers[1].ID, Tenor: 6, Amount: 2000000},
	}

	if err = db.Create(&limits).Error; err != nil {
		panic(err)
	}

	log.Println("Database seeded")
}
