package main

import "server/src/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}

func main() {

}
