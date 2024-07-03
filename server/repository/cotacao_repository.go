package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/israelalvesmelo/desafio-client-server-api/server/dto"
	"github.com/israelalvesmelo/desafio-client-server-api/server/model"
	"gorm.io/gorm"
)

func SaveCotacao(db *gorm.DB, dto *dto.CotacaoDto) error {

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

	tx := db.WithContext(ctx).Create(model.Cotacao{
		Code:       dto.Code,
		Codein:     dto.Codein,
		Name:       dto.Name,
		High:       dto.High,
		Low:        dto.Low,
		VarBid:     dto.VarBid,
		PctChange:  dto.PctChange,
		Bid:        dto.Bid,
		Ask:        dto.Ask,
		Timestamp:  dto.Timestamp,
		CreateDate: dto.CreateDate,
	})

	if tx.Error != nil {
		return fmt.Errorf("failed to save cotacao: %v", tx.Error)
	}

	return nil
}
