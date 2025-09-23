package config

import "github.com/joho/godotenv"

func init() {

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
