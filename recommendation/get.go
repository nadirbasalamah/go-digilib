package recommendation

import (
	"context"
	"go-digilib/pkg/ai"
)

type get struct{}

func (g get) GetBookRecommendation(ctx context.Context, req *BookRecommendationRequest) (ai.PromptResponse, error) {
	service := ai.InitService()
	recomReq := ai.BookRecommendationRequest{
		Quantity: req.Quantity,
		Topic:    req.Topic,
	}

	return service.GetBookRecommendation(recomReq)
}
