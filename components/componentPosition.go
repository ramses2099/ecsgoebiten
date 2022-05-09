package components

import cons "github.com/ramses2099/ecsgoebiten/constants"

type componentPosition struct {
	X    int32
	Y    int32
	Name string
}

func (p *componentPosition) GetName() string {
	return p.Name
}

func NewComponentPosition(x, y int32) Component {
	return &componentPosition{
		Name: cons.ComponentPosition,
		X:    x,
		Y:    y,
	}
}
