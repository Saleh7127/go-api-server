package api

import (
	"fmt"
	"github.com/Saleh7127/go-api-server/auth"
	"github.com/Saleh7127/go-api-server/handler"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

var SecretKey = []byte("CHAGOL")

func CheckServer(port int) {

	fmt.Printf("Server starting at port %v\n", port)
	fmt.Printf("Server link: localhost:%v\n", port)
	fmt.Println(auth.GenerateJWT(SecretKey))

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.HandleFunc("/", auth.IsAuthorized(handler.HomePage))

	router.Get("/players", auth.IsAuthorized(handler.AllPlayer))
	router.Get("/player/{name}", auth.IsAuthorized(handler.PlayerSearch))
	router.Post("/players/new", auth.IsAuthorized(handler.CreateNewEntry))
	router.Put("/update/{name}", auth.IsAuthorized(handler.UpdatePlayer))
	router.Delete("/delete/{name}", auth.IsAuthorized(handler.DeletePlayer))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
