package main

import (
	"github.com/joho/godotenv"
	"github.com/majoo-test/repository"
	"github.com/majoo-test/routes"
	"github.com/majoo-test/util"
)

func main() {
	// Load env
	godotenv.Load(".env")

	// Open connection
	db := util.SetupDB()

	// Note: update password to hash
	userRepository := repository.NewRepositoryUser(db)
	userRepository.UpdatePasswordHash()

	// use router
	server := routes.SetupRouter(db)
	server.Run()
}
