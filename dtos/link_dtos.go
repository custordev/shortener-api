package dtos

type CreateLinkRequest struct {
	ShortCode   string `json:"shortCode" binding:"required"`
	OriginalUrl string `json:"originalUrl" binding:"required"`
	Clicks      uint    ` json:"clicks  omitempty"`
	UserID      uint   `json:"user_id omitempty"`
}