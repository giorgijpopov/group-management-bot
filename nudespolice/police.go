package nudespolice

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/koyachi/go-nude"
)

type policeman struct {
}

var _ Policeman = &policeman{}

func NewPoliceman() *policeman {
	return &policeman{}
}

func (p *policeman) CheckNudesInImage(img image.Image) (bool, error) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "nude-check-*.jpg")
	if err != nil {
		return false, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Encode image to temp file
	err = jpeg.Encode(tmpFile, img, nil)
	if err != nil {
		return false, err
	}

	// Check for nudes
	hasNudes, err := nude.IsNude(tmpFile.Name())
	if err != nil {
		return false, err
	}

	return hasNudes, nil
}
