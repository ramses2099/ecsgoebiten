package systems

import (
	"github.com/ramses2099/ecsgoebiten/constants"
	"github.com/ramses2099/ecsgoebiten/entities"
)

type SystemMovement struct {
	name     string
	entities []entities.Entity
}

//
func NewSystemMovement(name string, entities []entities.Entity) *System {
	return &SystemMovement{
		name:     name,
		entities: entities,
	}

}

//
func (sm *SystemMovement) GetSystemName() string {
	return sm.name
}

//
func (sm *SystemMovement) ExecuteSystem() {
	for _, e := range sm.entities {
		if e.HasComponent(constants.ComponentPosition) {

		}
	}
}
