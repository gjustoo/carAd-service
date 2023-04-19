package main

import (
	"fm-scrapper-go/app/controller"
	"fm-scrapper-go/app/repo/db"
)

func main() {
	db.Init()
	controller.Start()
}
