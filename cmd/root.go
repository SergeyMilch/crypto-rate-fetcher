package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd представляет корневую команду приложения
var rootCmd = &cobra.Command{
    Use:   "crypto-rate-fetcher",
    Short: "A CLI tool to fetch cryptocurrency rates",
    Long: `Crypto Rate Fetcher is a CLI tool that allows users to fetch 
           cryptocurrency exchange rates from Binance API through a simple 
           and user-friendly interface.`,
}

// Execute добавляет все дочерние команды в rootCmd и устанавливает флаги
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error occurred: %s\n", err)
        os.Exit(1)
    }
}