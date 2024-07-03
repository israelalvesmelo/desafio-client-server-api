package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/israelalvesmelo/desafio-client-server-api/server/dto"
	"github.com/israelalvesmelo/desafio-client-server-api/server/model"
	"github.com/israelalvesmelo/desafio-client-server-api/server/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	setupDataBase()
	http.HandleFunc("/cotacao", CotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {

	cotacao, error := GetCotacao()
	if error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	error = SaveCotacao(cotacao)
	if error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cotacao.Bid)
}

func GetCotacao() (*dto.CotacaoDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("Timeout expired to get cotacao: ", ctx.Err())
			}
		}
	}()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var c map[string]dto.CotacaoDto
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	cotacao := c["USDBRL"]
	return &cotacao, nil
}

func SaveCotacao(c *dto.CotacaoDto) error {
	if DB == nil {
		return fmt.Errorf("Database connection is not initialized")
	}
	return repository.SaveCotacao(DB, c)
}

func setupDataBase() {
	db, err := gorm.Open(sqlite.Open("cotacao.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Cotacao{})
	DB = db
}
