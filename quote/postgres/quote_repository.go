package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
)

type QuoteRepository struct {
	DB *pgx.Conn
}

func NewQuoteRepository(db *pgx.Conn) *QuoteRepository {
	return &QuoteRepository{DB: db}
}

func (qr *QuoteRepository) Save(ctx context.Context, quote *entity.Quote) error {
	_, err := qr.DB.Exec(ctx, "INSERT INTO quotes (carrier_id, name, service, final_price, deadline) VALUES ($1, $2, $3, $4, $5)", quote.Carrier.CarrierID, quote.Carrier.Name, quote.Service, quote.FinalPrice, quote.Deadline)
	if err != nil {
		return err
	}
	return nil
}

func (qr *QuoteRepository) FindAll(ctx context.Context) ([]*entity.Quote, error) {
	rows, err := qr.DB.Query(ctx, "SELECT * FROM quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*entity.Quote
	for rows.Next() {
		var q entity.Quote
		err := rows.Scan(&q.Carrier.CarrierID, &q.Carrier.Name, &q.Service, &q.FinalPrice, &q.Deadline)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, &q)
	}
	return quotes, nil
}
