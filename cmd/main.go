package main

import (
	"example.com/my-medium-clone/internal/common/db_connection"
	articleHandler "example.com/my-medium-clone/internal/ports/http/handler"
	userDAO "example.com/my-medium-clone/internal/repoImpl"
	"example.com/my-medium-clone/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	httpServer()
}

func httpServer() *chi.Mux {
	db, err := db_connection.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	var (
		userRepo    = userDAO.NewUserRepo(db)
		userUseCase = usecase.NewUserUseCase(userRepo)
		userHandler = articleHandler.NewUserHandler(userUseCase)
	)
	var (
		articleRepo    = userDAO.NewArticleRepo(db)
		articleUseCase = usecase.NewArticleUseCase(articleRepo)
		articleHandler = articleHandler.NewArticleHandler(articleUseCase)
	)

	router.Route("/api", func(r chi.Router) {

		r.Route("/user", func(r chi.Router) {

			r.Post("/sign-up-user", userHandler.SignUpUser)

			r.Post("/sign-in-user", userHandler.SignInUser)

			r.Get("/get-id", userHandler.GetById)

			r.Get("/get-email", userHandler.GetByEmail)

			r.Get("/criteria", userHandler.GetList)

			r.Put("/put", userHandler.Update)

			r.Delete("/delete", userHandler.Delete)
		})

		r.Route("/article", func(r chi.Router) {
			r.Post("/create", articleHandler.Create)
		})
	})

	server := &http.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	log.Println("Starting server on port...", os.Getenv("HTTP_PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()

	return router
}
