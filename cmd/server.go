package cmd

import (
	"log"
	"net/http"

	"github.com/SergeyMilch/crypto-rate-fetcher/pkg/server"
	"github.com/spf13/cobra"
)


var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Start the server",
    Run: func(cmd *cobra.Command, args []string) {
        http.HandleFunc("/api/v1/rates", server.HandleRequest)
        log.Println("Listening at port 3001...")
        http.ListenAndServe(":3001", nil)
    },
}

func init() {
    rootCmd.AddCommand(serverCmd)
}