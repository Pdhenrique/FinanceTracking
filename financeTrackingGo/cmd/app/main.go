package app

import (
	"log"
	"os"

	"github.com/Pdhenrique/FinanceTracking/db"
	"github.com/Pdhenrique/FinanceTracking/internal/http"
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

	handler := http.NewHandler()

	server := http.NewServer(handler, "8080")
	server.Start()
	defer server.Stop()
}
