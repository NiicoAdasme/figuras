package figuras

import "math"

type Circulo struct {
	Radio float64
}


func (circ *Circulo) area() float64 {
	return math.Pi * (circ.Radio * circ.Radio)
}

func (circ *Circulo) perimetro() float64 {
	return 2 * math.Pi * circ.Radio
}