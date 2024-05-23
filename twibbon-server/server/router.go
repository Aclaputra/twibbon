package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"twibbon-server/repository"
	"twibbon-server/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	server struct {
		r            *gin.Engine
		repositories repositories
		usecases     usecases
	}

	repositories struct {
		userRepository repository.UserRepository
	}

	usecases struct {
		userUsecase usecase.UserUsecase
	}
)

func (s *server) initRepositores(db *gorm.DB) {
	s.repositories = repositories{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (s *server) initUsecase() {
	s.usecases = usecases{
		userUsecase:        usecaseImpl.NewUserUsecase(s.repositories.userRepository),
		transactionUsecase: usecaseImpl.NewTransactionUsecase(s.repositories.transactionRepository),
	}
}

func (s *server) initHandler() {
	handler.NewUserHandler(s.usecases.userUsecase).Route(s.r)
	handler.NewTransactionHandler(s.usecases.transactionUsecase).Route(s.r)
}

func StartServer(db *gorm.DB) {
	server := server{
		r: gin.Default(),
	}
	server.r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	server.r.ContextWithFallback = true
	server.r.Use(middleware.WithTimeout)
	server.r.Use(middleware.RequestLogger())

	server.initRepositores(db)
	server.initUsecase()
	server.initHandler()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	default:
	}
	log.Println("Server exiting")
}
