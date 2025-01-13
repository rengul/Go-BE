package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"re-home/auth/pkg/auth"
	"re-home/auth/pkg/auth/delivery"
	mysql "re-home/auth/pkg/auth/repository/mongo"
	"re-home/auth/pkg/auth/usecase"
	"time"
)

type App struct {
	httpServer *http.Server

	authUseCase auth.UseCase
}

func NewApp() *App {
	db := initDB()

	repository := mysql.NewUserRepository(db, "users")
	authUseCase := usecase.NewAuthorizer(
		repository,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl")*time.Second,
	)

	return &App{
		authUseCase: authUseCase,
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Endpoints
	api := router.Group("/auth")
	delivery.RegisterHTTPEndpoints(api, a.authUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "gouser:G1nW3bUs3r!@tcp(192.168.178.35:3306)/home?parseTime=true")
	if err != nil {
		log.Fatalf("Error occurred while connecting to MySQL: %v", err)
	}

	// Verifica connessione
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	log.Println("Connected to MySQL successfully")
	return db
}
