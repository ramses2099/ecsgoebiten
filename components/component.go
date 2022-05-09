package components

import cons "github.com/ramses2099/ecsgoebiten/constants"

type Component interface {
	GetName() string
}

type componentHealth struct {
	Name  string
	Value int32
}

func (h *componentHealth) GetName() string {
	return h.Name
}

func NewComponentHealth(value int32) Component {
	return &componentHealth{
		Name:  cons.ComponentHealth,
		Value: value,
	}
}
