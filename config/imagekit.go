package config

import (
	"github.com/imagekit-developer/imagekit-go"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func FileUpload() *imagekit.ImageKit {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ik, err := imagekit.New()
	if err != nil {
		panic(err)
	}
	ik = imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  os.Getenv("PRIVATE_KEY"),
		PublicKey:   os.Getenv("PUBLIC_KEY"),
		UrlEndpoint: os.Getenv("URL_ENDPOINT"),
	})

	return ik
}
