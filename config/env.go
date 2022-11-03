package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CheckLoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvCloudName() string {
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	CheckLoadEnvFile()
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	CheckLoadEnvFile()
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	CheckLoadEnvFile()
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}

func EnvDbName() string {
	CheckLoadEnvFile()
	return os.Getenv("DB_NAME")
}
func EnvDbHost() string {
	CheckLoadEnvFile()
	return os.Getenv("DB_HOST")
}
func EnvDbPort() string {
	CheckLoadEnvFile()
	return os.Getenv("DB_PORT")
}
func EnvDbUser() string {
	CheckLoadEnvFile()
	return os.Getenv("DB_USER")
}
func EnvDbPass() string {
	CheckLoadEnvFile()
	return os.Getenv("DB_PASS")
}
func EnvSecretJwt() string {
	CheckLoadEnvFile()
	return os.Getenv("DE6ED21B4E643161949DFCE42DABC")
}
