package entities

// EntityManager
type EntityManager struct {
	entities []Entity
}

// NewEntityManager
func NewEntityManager(entities ...Entity) EntityManager {
	en := []Entity{}
	en = append(en, entities...)
	return EntityManager{entities: en}
}

// Lenght
func (em *EntityManager) GetLength() int {
	return len(em.entities)
}

// Get Entities
func (em *EntityManager) GetEntities() []Entity {
	return em.entities
}

// remove index slice
func removeIndex(entities []Entity, index int) []Entity {
	return append(entities[:index], entities[index+1:]...)
}

// Remove Entity
func (em *EntityManager) GetRemove(Id string) {
	for i, e := range em.entities {
		if e.Id != Id {
			removeIndex(em.entities, i)
		}
	}
}

// Add Entity
func (em *EntityManager) AddEntity(entity Entity) {
	em.entities = append(em.entities, entity)
}
