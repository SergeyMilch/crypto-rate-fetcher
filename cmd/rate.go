package cmd

import (
	"fmt"

	"github.com/SergeyMilch/crypto-rate-fetcher/pkg/client"
	"github.com/spf13/cobra"
)


var rateCmd = &cobra.Command{
    Use:   "rate",
    Short: "Get the rate of a currency pair",
    Run: func(cmd *cobra.Command, args []string) {
        pair, _ := cmd.Flags().GetString("pair")
        rate, err := client.GetRate(pair)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(rate)
    },
}

func init() {
    rateCmd.Flags().String("pair", "", "Pair of currencies to get the rate")
    rootCmd.AddCommand(rateCmd)
}