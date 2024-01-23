package Initializers

import (
	"log"
	"os"
	"time"
	"wan-api-spam-report/Const"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	//Get database url from environment variables (defined in .env file)
	var dsn string = os.Getenv("DB_URL")

	//Connect with postgres
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             50 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Warn,           // Log level
				IgnoreRecordNotFoundError: false,                 // Dont ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,                 // Include params in the SQL log
				Colorful:                  true,                  // Disable color
			},
		),
	})

	// !TODO

	// Zone Comment
	// Research connection pool
	// sqlDB.SetMaxIdleConns(viper.GetInt("database.mysql.pool.min"))
	// sqlDB.SetMaxOpenConns(viper.GetInt("database.mysql.pool.max"))

	//Set up connection pool
	sqlDB, err := DB.DB()
	//Minimum number of connections that can use the db in a period of time
	sqlDB.SetMaxIdleConns(Const.MIN_CONNECTIONS)
	//Maximum number of connections that can use the db in a period of time
	sqlDB.SetMaxOpenConns(Const.MAX_CONNECTIONS)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
