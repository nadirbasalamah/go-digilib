package recommendation

import (
	"context"
	"go-digilib/pkg/ai"
)

type Service interface {
	GetBookRecommendation(ctx context.Context, req *BookRecommendationRequest) (ai.PromptResponse, error)
}

type service struct {
	get
}

var _ Service = (*service)(nil)

func New() Service {
	return service{
		get: get{},
	}
}
