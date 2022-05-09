package components

import cons "github.com/ramses2099/ecsgoebiten/constants"

type Component interface {
	GetComponentName() string
}

type componentHealth struct {
	Name  string
	Value int32
}

func (h *componentHealth) GetComponentName() string {
	return h.Name
}

func NewComponentHealth(value int32) Component {
	return &componentHealth{
		Name:  cons.ComponentHealth,
		Value: value,
	}
}
