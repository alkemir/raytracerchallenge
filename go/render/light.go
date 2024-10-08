package render

type Light struct {
	position  Tuple
	intensity Tuple
}

func NewPointLight(p Tuple, i Tuple) *Light {
	return &Light{
		position:  p,
		intensity: i,
	}
}
