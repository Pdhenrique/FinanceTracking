package app

import (
	"log"
	"os"

	"github.com/Pdhenrique/FinanceTracking/internal/db"
	"github.com/Pdhenrique/FinanceTracking/internal/http"
	"github.com/Pdhenrique/FinanceTracking/pkg/user"
)

func main() {

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("env DATABASE_URL não definido")
	}

	conn, err := db.Connect(dbURL)
	if err != nil {
		log.Fatal("ERROR conectando ao db:", err)
	}
	defer conn.Close()

	userStorage := db.NewUserStorage(conn)
	userService := user.NewService(userStorage)
	handler := http.NewHandler(userService)

	server := http.NewServer(handler, "8080")
	server.Start()
	defer server.Stop()

	select {}
}
