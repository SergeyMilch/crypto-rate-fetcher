package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SergeyMilch/crypto-rate-fetcher/pkg/server"
)


func GetRate(pair string) (float64, error) {
    
    if pair == "" {
        return 0, fmt.Errorf("no currency pair specified")
    }

    if !server.IsValidSymbol(pair) {
        return 0, fmt.Errorf("invalid currency pair format: %s", pair)
    }

    // Получение URL из переменных окружения
    apiURL := os.Getenv("API_URL")
    if apiURL == "" {
        apiURL = "http://localhost:3001/api/v1/rates" // URL по умолчанию
    }

    // Формирование запроса
    requestURL := fmt.Sprintf("%s?pairs=%s", apiURL, pair)

    resp, err := http.Get(requestURL)
    if err != nil {
        log.Printf("Error making GET request for %s: %v", pair, err)
        return 0, fmt.Errorf("error making GET request: %v", err)
    }
    defer resp.Body.Close()

    var rates map[string]float64
    if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
        log.Printf("Error decoding response body for %s: %v", pair, err)
        return 0, fmt.Errorf("error decoding response body: %v", err)
    }

    if rate, ok := rates[pair]; ok {
        return rate, nil
    } else {
        return 0, fmt.Errorf("rate for %s not found", pair)
    }
}