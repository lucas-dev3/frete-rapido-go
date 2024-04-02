package quote

import (
	"context"

	"github.com/lucas-dev3/frete-rapido-go.git/internal/entity"
	"github.com/lucas-dev3/frete-rapido-go.git/internal/http/presenter"
)

type Repository interface {
	Save(ctx context.Context, quote *entity.Quote) error
	FindAll(ctx context.Context) ([]*entity.Quote, error)
}

type UseCase interface {
	ProcessQuote(ctx context.Context, quote *presenter.QuoteRequest) error
	GetMetrics(ctx context.Context) ([]*entity.Quote, error)
}
