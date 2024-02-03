package main

import (
	articleDAO "example.com/my-medium-clone/internal/article/adapters"
	"example.com/my-medium-clone/internal/article/app"
	articleHandler "example.com/my-medium-clone/internal/article/ports/http/handler"
	"example.com/my-medium-clone/internal/pkg/db_connection"
	userDAO "example.com/my-medium-clone/internal/users/adapters"
	usecase "example.com/my-medium-clone/internal/users/app"
	"example.com/my-medium-clone/internal/users/jwt"
	"example.com/my-medium-clone/internal/users/ports/http/handler"
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
		userHandler = handler.NewUserHandler(userUseCase)
	)
	var (
		articleRepo    = articleDAO.NewArticleRepo(db)
		articleUseCase = app.NewArticleUseCase(articleRepo)
		articleHandler = articleHandler.NewArticleHandler(articleUseCase)
	)

	router.Route("/api", func(r chi.Router) {

		r.Route("/user", func(r chi.Router) {

			r.Post("/sign-up-user", userHandler.SignUpUser)

			//r.Post("/sign-in-user", userHandler.SignInUser)

			r.Get("/get-id", userHandler.GetById)

			r.Get("/criteria", userHandler.GetList)

			r.Put("/put", userHandler.Update)

			r.Delete("/delete", userHandler.Delete)

			r.With(jwt.AuthMiddleWare).Get("/get-user{email}", userHandler.GetUserByEmail)
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
