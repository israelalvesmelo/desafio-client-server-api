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
		Code:       dto.USDBRL.Code,
		Codein:     dto.USDBRL.Codein,
		Name:       dto.USDBRL.Name,
		High:       dto.USDBRL.High,
		Low:        dto.USDBRL.Low,
		VarBid:     dto.USDBRL.VarBid,
		PctChange:  dto.USDBRL.PctChange,
		Bid:        dto.USDBRL.Bid,
		Ask:        dto.USDBRL.Ask,
		Timestamp:  dto.USDBRL.Timestamp,
		CreateDate: dto.USDBRL.CreateDate,
	})

	if tx.Error != nil {
		return fmt.Errorf("failed to save cotacao: %v", tx.Error)
	}

	return nil
}
