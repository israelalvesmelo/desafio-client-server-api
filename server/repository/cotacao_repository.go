package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/israelalvesmelo/desafio-client-server-api/server/dto"
	_ "github.com/mattn/go-sqlite3"
)

func SaveCotacao(db *sql.DB, c *dto.CotacaoDto) error {
	log.Println("Salvando cotacao")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("Timeout expired to save cotacao: ", ctx.Err())
			}
		}
	}()

	q := `
		INSERT INTO
			cotacoes(
				code,
				codein,
				name,
				high,
				low,
				varbid,
				pctchange,
				bid,
				ask,
				timestamp,
				create_date
			)
			VALUES (
				?, 
				?, 
				?, 
				?, 
				?, 
				?, 
				?, 
				?, 
				?, 
				?, 
				?
			);
	`

	id, err := db.ExecContext(ctx, q,
		c.Code,
		c.Codein,
		c.Name,
		c.High,
		c.Low,
		c.VarBid,
		c.PctChange,
		c.Bid,
		c.Ask,
		c.Timestamp,
		c.CreateDate,
	)

	if err != nil {
		return fmt.Errorf("failed to save cotacao: %v", err)
	}
	log.Printf("ID inserido: %d", &id)

	return nil
}
