package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("MODE_DB_USERNAME"),
		os.Getenv("MODE_DB_PASSWORD"),
		os.Getenv("MODE_DB_NAME"))
	a.Run(":8080")
}
