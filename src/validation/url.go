package validation

type CreateURLInput struct {
	OriginalURL string `validate:"required" json:"originalUrl"`
	Suffix      string `gorm:"unique" json:"suffix"`
}
