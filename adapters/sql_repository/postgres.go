package sql_repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"merchant-account-api/core/domain"
	"os"
)

type postgreSQLRepository struct {
	db *gorm.DB
}

func Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	fmt.Println(DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("Connection error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}
	db.Debug().AutoMigrate(&domain.Merchant{}, &domain.TeamMember{}) //database migration
	return db
}

func NewPostgreSQLRepository() (*postgreSQLRepository, error) {
	postgreSQLDB := Initialize("postgres", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "5432", os.Getenv("DATABASE_HOST"),
		os.Getenv("POSTGRES_DB"))
	db := &postgreSQLRepository{
		db: postgreSQLDB,
	}
	return db, nil
}
