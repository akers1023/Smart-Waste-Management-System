package connections

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Sql struct {
	Db       *gorm.DB
	Host     string
	Port     int
	UserName string
	Password string
	SSLMode  string
	DbName   string
}

func (s *Sql) Connect() (*gorm.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		s.Host, s.Port, s.UserName, s.Password, s.DbName, s.SSLMode)
	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		return db, err
	}
	fmt.Println("😂 Smart-Waste-Management-System Conected Database ✊")
	return db, nil
}

func (s *Sql) Close() {
	if s.Db != nil {
		db, _ := s.Db.DB()
		db.Close()
		fmt.Println("Closed Database Connection")
	}
}
