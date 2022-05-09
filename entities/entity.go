package entities

import (
	"github.com/google/uuid"
	"github.com/ramses2099/ecsgoebiten/components"
)

// Entity struct
type Entity struct {
	Id         string
	Components map[string]components.Component
}

// Factaroy NewEntity
func NewEntity() Entity {
	entity := Entity{Id: uuid.New().String()}
	entity.Components = make(map[string]components.Component)
	return entity
}

// Function Get Components
func (e *Entity) GetComponents() map[string]components.Component {
	return e.Components
}

// Function Get Id
func (e *Entity) GetId() string {
	return e.Id
}

// Function Add Components
func (e *Entity) AddComponents(name string, componet components.Component) {
	e.Components[name] = componet
}

// Function Get Component by name
func (e *Entity) GetComponent(name string) components.Component {
	return e.Components[name]
}

// Function Delete Components
func (e *Entity) DeleteComponent(name string) {
	delete(e.Components, name)
}
