package systems

import (
	"reflect"

	"github.com/ramses2099/ecsgoebiten/constants"
	"github.com/ramses2099/ecsgoebiten/entities"
)

type systemMovement struct {
	name     string
	entities []entities.Entity
}

func NewSystemMovement(name string, entities []entities.Entity) System {
	return &systemMovement{
		name:     name,
		entities: entities,
	}

}

//
func (sm *systemMovement) GetSystemName() string {
	return sm.name
}

//
func (sm *systemMovement) ExecuteSystem() {
	for _, e := range sm.entities {
		if e.HasComponent(constants.ComponentPosition) {
			cp := e.GetComponent(constants.ComponentPosition)
			pos := reflect.ValueOf(cp).Elem()

		}
	}
}
