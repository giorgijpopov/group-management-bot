package nudespolice

import "image"

type Policeman interface {
	CheckNudesInImage(img image.Image) (bool, error)
}
