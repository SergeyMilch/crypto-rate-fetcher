package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// Структура для ответа от API Binance
type BinanceResponse struct {
    Symbol string  `json:"symbol"`
    Price  float64 `json:"price,string"`
}

// Структура для ответа сервера
type ServerResponse map[string]float64

// Функция обработки запросов к серверу
func HandleRequest(w http.ResponseWriter, r *http.Request) {

    var symbols []string

    // Обработка GET запроса
    if r.Method == "GET" {
        pairs := r.URL.Query().Get("pairs")
        symbols = strings.Split(pairs, ",")
    }

    // Обработка POST запроса
    if r.Method == "POST" {
        var postData struct {
            Pairs []string `json:"pairs"`
        }
        json.NewDecoder(r.Body).Decode(&postData)
        symbols = postData.Pairs
    }

	response := make(ServerResponse)

    if len(symbols) == 0 {
        http.Error(w, "No currency pairs provided", http.StatusBadRequest)
        return
    }

    for _, symbol := range symbols {

        if symbol == "" {
            log.Println("Empty currency pair detected, skipping")
            continue
        }

        if !IsValidSymbol(symbol) {
            log.Printf("Invalid currency pair format: %s", symbol)
            http.Error(w, fmt.Sprintf("Invalid currency pair format: %s", symbol), http.StatusBadRequest)
            return
        }
        
        binanceResp, err := queryBinanceAPI(symbol)
        if err != nil {
            log.Printf("Error querying Binance API for %s: %v", symbol, err)
            http.Error(w, "Error querying Binance API", http.StatusInternalServerError)
            return
        }

        response[symbol] = binanceResp.Price
    }

    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Printf("Error encoding response: %v", err)
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
    }
}

func queryBinanceAPI(symbol string) (*BinanceResponse, error) {

	formattedSymbol := strings.ReplaceAll(symbol, "-", "")

    url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", formattedSymbol)

    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making request to Binance: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Binance API returned non-OK status: %d", resp.StatusCode)
    }

    var binanceResp BinanceResponse

    err = json.NewDecoder(resp.Body).Decode(&binanceResp)
    if err != nil {
        return nil, fmt.Errorf("error decoding Binance response: %w", err)
    }

    return &binanceResp, nil
}

func IsValidSymbol(symbol string) bool {
   
    re := regexp.MustCompile(`^[A-Z]{2,12}-[A-Z]{2,12}$`)
    
    return re.MatchString(symbol)
}