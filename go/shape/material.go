package shape

import "raytracerchallenge/tuple"

type Material struct {
	color     tuple.Tuple
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

var DefaultMaterial *Material = &Material{
	color:     tuple.NewColor(1, 1, 1),
	ambient:   0.1,
	diffuse:   0.9,
	specular:  0.9,
	shininess: 200,
}

func NewMaterial(color tuple.Tuple, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{color, ambient, diffuse, specular, shininess}
}
