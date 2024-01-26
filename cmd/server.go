package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/SergeyMilch/crypto-rate-fetcher/pkg/server"
	"github.com/spf13/cobra"
)


var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Start the server",
    Run: func(cmd *cobra.Command, args []string) {
       // Получение порта из переменных окружения
       port := os.Getenv("PORT")
       if port == "" {
           port = "3001" // Значение по умолчанию, если переменная не установлена
       }

       // Получение роута из переменных окружения
       url := os.Getenv("API_ROUTE")
       if url == "" {
           url = "/api/v1/rates" // Значение по умолчанию, если переменная не установлена
       }

       http.HandleFunc(url, server.HandleRequest)
       log.Printf("Listening at port %s...", port)
       http.ListenAndServe(":" + port, nil)
    },
}

func init() {
    rootCmd.AddCommand(serverCmd)
}