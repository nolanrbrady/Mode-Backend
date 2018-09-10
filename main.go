package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("MODE_DB_USERNAME"), // need to declare this environment var
		os.Getenv("MODE_DB_PASSWORD"), // need to declare this environment var
		os.Getenv("MODE_DB_NAME"),     // need to declare this environment var
	)
	a.Run(":8080")
}
