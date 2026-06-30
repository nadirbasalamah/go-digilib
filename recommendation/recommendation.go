package recommendation

type BookRecommendationRequest struct {
	Quantity uint   `json:"quantity" validate:"required,min=1"`
	Topic    string `json:"topic" validate:"required"`
}
