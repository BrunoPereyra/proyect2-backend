package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PORT() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error")
	}
	return os.Getenv("PORT")
}

func CLOUDINARY() (string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET")

}

func CLOUDINARY_URL() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error")
	}
	return os.Getenv("CLOUDINARY_URL")
}
