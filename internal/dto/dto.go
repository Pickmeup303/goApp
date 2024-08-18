package dto

import (
	"mime/multipart"
)

type RequestEncryptInput struct {
	Alphabet     string                `validate:"required,min=140" label:"K1"`
	KeyShifter   int                   `validate:"required,numeric,gte=1" label:"K2"`
	KeyTranspose int                   `validate:"required,numeric,gte=1" label:"K3"`
	Message      string                `validate:"required" label:"Message"`
	File         *multipart.FileHeader `validate:"required" label:"Video"`
}

type RequestDecryptInput struct {
	Alphabet     string                `validate:"required,min=140" label:"K1"`
	KeyShifter   int                   `validate:"required,numeric,gte=1" label:"K2"`
	KeyTranspose int                   `validate:"required,numeric,gte=1" label:"K3"`
	File         *multipart.FileHeader `validate:"required" label:"Video"`
}

type RequestCapacityVideo struct {
	File *multipart.FileHeader `validate:"required"`
}
