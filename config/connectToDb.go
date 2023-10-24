package config

import (
	"fmt"
	"log"
	"tahjib75/restful-crud-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config struct to hold the configuration
type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUserName string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     int    `mapstructure:"DB_PORT"`
}

var DB *gorm.DB
var config Config
var envVariables = []string{"DB_HOST", "DB_USERNAME", "DB_PASSWORD", "DB_NAME", "DB_PORT"}

func LoadEnv() (Config, error) {

	viper.SetConfigType("env") // In this case, it's set to "env" indicating that it expects a key-value pair style configuration.
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // This allows the application to read environment variables directly without any manual binding.

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("error binding environment variables %v", err)
	}

	// Loop through the list of environment variables and bind them to Viper
	for _, env := range envVariables {
		if err := viper.BindEnv(env); err != nil {
			return config, fmt.Errorf("error binding environment variables %v", err)
		}
	}

	// Unmarshal the configuration data from Viper into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("error unmarshalling config %v", err)
	}

	// Validate the populated Config struct using the validator
	validate := validator.New() // Create a new instance of the validator
	if err := validate.Struct(config); err != nil {
		return config, fmt.Errorf("error unmarshalling config %v", err)
	}
	return config, nil
}

func ConnectToDB() *gorm.DB {
	config, err := LoadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment %v", err)
	}

	// Create the DSN (Data Source Name) string for connecting to the database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBName, config.DBPassword, config.DBUserName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}

	fmt.Println("Connected to Database......")

	DB = db // Set the global DB variable to the connected database instance

	DB.AutoMigrate(&models.Admin{})

	return db
}
