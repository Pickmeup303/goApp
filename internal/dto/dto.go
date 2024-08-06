package dto

import (
	"mime/multipart"
)

type RequestEncryptInput struct {
	Alphabet string                `validate:"required,min=140"`
	Key      int                   `validate:"required,numeric,gte=1"`
	Message  string                `validate:"required"`
	File     *multipart.FileHeader `validate:"required"`
}

type RequestDecryptInput struct {
	Alphabet string                `validate:"required,min=140"`
	Key      int                   `validate:"required,numeric,gte=1"`
	File     *multipart.FileHeader `validate:"required"`
}

type RequestCapacityVideo struct {
	File *multipart.FileHeader `validate:"required"`
}
