package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println(s3Bucket)
	fmt.Println(secretKey)
	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() // wait for Enter Key
}

/* ref = https://blog.wu-boy.com/2019/04/how-to-load-env-file-in-go/ */
