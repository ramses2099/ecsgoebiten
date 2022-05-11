package components

import cons "github.com/ramses2099/ecsgoebiten/constants"

type ComponentPosition struct {
	X    int32
	Y    int32
	Name string
}

func (p *ComponentPosition) GetComponentName() string {
	return p.Name
}

func NewComponentPosition(x, y int32) *ComponentPosition {
	return &ComponentPosition{
		Name: cons.ComponentPosition,
		X:    x,
		Y:    y,
	}
}
