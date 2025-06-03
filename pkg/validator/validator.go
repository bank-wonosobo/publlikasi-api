package validator

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(interface{}) error
	RegisterCustomValidor()
	ValidateImageFile(file *multipart.FileHeader) error
}

type validatorImpl struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	return &validatorImpl{
		validator: validator.New(),
	}
}

// Validate implements Validator.
func (v *validatorImpl) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			errors = append(errors, fmt.Sprintf("%s: failed on the '%s' tag", field, tag))
		}
		return fmt.Errorf("%s", strings.Join(errors, ", "))
	}
	return nil
}

func (v *validatorImpl) RegisterCustomValidor() {
	// Register custom enum validation
	v.validator.RegisterValidation("must_category_product", func(fl validator.FieldLevel) bool {
		status := fl.Field().String()
		switch status {
		case string(entities.Deposito), string(entities.Tabungan), string(entities.Kredit), string(entities.Digital):
			return true
		default:
			return false
		}
	})
}

func (v *validatorImpl) ValidateImageFile(file *multipart.FileHeader) error {
	// Validasi ukuran maksimal (contoh: 2 MB)
	if file.Size > 2*1024*1024 {
		return errors.New("image size exceeds 2MB")
	}

	// Validasi ekstensi (bisa juga cek mime secara lebih kompleks)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		return nil
	default:
		return errors.New("invalid image format; only jpg, jpeg, png, gif, webp allowed")
	}
}
