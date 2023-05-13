package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func URI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error URI mongo")
	}
	return os.Getenv("MONGODB_URI")

}
func PORT() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error PORT")
	}
	return os.Getenv("PORT")
}

func CLOUDINARY() (string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error CLOUDINARY")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET")

}

func CLOUDINARY_URL() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error CLOUDINARY_URL")
	}
	return os.Getenv("CLOUDINARY_URL")
}
func TEST_ACCESS_TOKEN() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error TEST_ACCESS_TOKEN")
	}
	return os.Getenv("TEST_ACCESS_TOKEN")
}
func PUBLICKEY() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error PUBLICKEY")
	}
	return os.Getenv("PUBLICKEY")
}
