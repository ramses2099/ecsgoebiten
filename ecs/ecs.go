package ecs

import (
	"reflect"
)

// Entity
type Id uint32

type Entity struct {
	components map[Id]interface{}
}

func NewEntity() *Entity {
	return &Entity{
		components: make(map[Id]interface{}),
	}
}

// GetComponent
func (e *Entity) GetComponet(id Id) (interface{}, bool) {
	val, ok := e.components[id]
	return val, ok
}

// AddComponent
func (e *Entity) AddComponet(id Id, val interface{}) {
	e.components[id] = val
}

// Entity

// Manager Entity
type ManagerEntity struct {
	reg       map[string]*Entity
	idCounter Id
}

func NewManagerEntity() *ManagerEntity {
	return &ManagerEntity{
		reg:       make(map[string]*Entity),
		idCounter: 0,
	}
}

func (m *ManagerEntity) NewId() Id {
	id := m.idCounter
	m.idCounter++
	return id
}

func name(t interface{}) string {
	name := reflect.TypeOf(t).String()
	if name[0] == '*' {
		return name[1:]
	}
	return name
}

func GetEntity(m *ManagerEntity, t interface{}) *Entity {
	name := name(t)
	entity, ok := m.reg[name]
	if !ok {
		m.reg[name] = NewEntity()
		entity = m.reg[name]
	}
	return entity
}

func Read(e *ManagerEntity, id Id, val Component) bool {
	// Get entity
	entity := GetEntity(e, val)
	// Get component value
	componets, ok := entity.GetComponet(id)
	//
	if ok {
		val.ComponentSet(componets)
	}

	return ok
}

func Write(e *ManagerEntity, id Id, val interface{}) {
	// Get entity
	entity := GetEntity(e, val)
	// Add component value
	entity.AddComponet(id, val)
}

func Each(en *ManagerEntity, t interface{}, f func(id Id, a interface{})) {
	// Get entity
	entity := GetEntity(en, t)
	// Loof component value
	for id, a := range entity.components {
		f(id, a)
	}
}

// Manager Entity

//components

type Component interface {
	ComponentSet(i interface{})
}

type Transform struct {
	X, Y int32
}

func (t *Transform) ComponentSet(val interface{}) {
	*t = val.(Transform)
}
