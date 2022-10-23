package main

import (
	"log"
	db "world-cup/framework/database/postgres"
	"world-cup/framework/web"
	"world-cup/framework/web/handler"
)

func main() {
	// ctx := context.Background()

	log.Println("Setup Application...")

	// Setup Database: Postgres
	db.Setup()

	// Setup Web Api
	srv := web.SetupHTTPServer(
		handler.NewHealthApiHandler(),
		handler.NewUserApiHandler(),
	)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error something went wrong starting http server: %s", err)
	}

}
