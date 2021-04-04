package nudespolice

import (
	"image"

	"github.com/koyachi/go-nude"
)

type policeman struct {
}

var _ Policeman = &policeman{}

func NewPoliceman() *policeman {
	return &policeman{}
}

func (p *policeman) CheckNudesInImage(img image.Image) (bool, error) {
	hasNudes, err := nude.IsImageNude(img)
	if err != nil {
		return false, err
	}

	return hasNudes, err
}
