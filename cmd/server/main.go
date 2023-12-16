package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CVWO-Backend/internal/auth"
	"github.com/CVWO-Backend/internal/database"
	"github.com/CVWO-Backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Listening on port 8000 at http://localhost:8000!")

	_, err := database.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	auth.GenerateAuthInfo()
	// util.Migrate()

	log.Fatalln(http.ListenAndServe(":8000", r))
}
