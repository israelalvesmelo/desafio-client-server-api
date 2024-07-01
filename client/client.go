package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao string

func main() {
	cotacao, err := GetCotacao()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Dólar: %s", *cotacao)

	err = CreateFileCotacao(cotacao)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetCotacao() (*Cotacao, error) {
	log.Print("Buscando cotação...")

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("Timeout expired to get cotacao: ", ctx.Err())
			}
		}
	}()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
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
	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return nil, err
	}
	return &cotacao, nil
}

func CreateFileCotacao(cotacao *Cotacao) error {
	log.Print("Criando arquivo cotacao.txt...")
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return fmt.Errorf("Erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", *cotacao))
	log.Print("Arquivo criado com sucesso!")
	return nil
}
