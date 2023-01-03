package shared

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// SQL wrapper
	SQL *gorm.DB

	// Database info
)

// Connect to the database
func Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Silent, // Log level
			Colorful: true,          // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		Logger:                                   newLogger,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		AllowGlobalUpdate:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})

	if err != nil {
		log.Print(err)
	}

	err = automigrate(db)
	if err != nil {
		log.Print(err)
	}
	SQL = db
}

func automigrate(db *gorm.DB) (err error) {
	mg := db.AutoMigrate(&Order{}, &ProductType{})
	return mg
}
