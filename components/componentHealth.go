package components

import cons "github.com/ramses2099/ecsgoebiten/constants"

type ComponentHealth struct {
	Name  string
	Value int32
}

func (h *ComponentHealth) GetComponentName() string {
	return h.Name
}

func NewComponentHealth(value int32) *ComponentHealth {
	return &ComponentHealth{
		Name:  cons.ComponentHealth,
		Value: value,
	}
}
