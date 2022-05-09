package systems

// SystemManager interface
type SystemManager interface {
	GetSystems() map[string]System
}

// SystemManager struct
type systemManager struct {
	systems map[string]System
}

// Function Get Systems
func (s *systemManager) GetSystems() map[string]System {
	return s.systems
}

// Factaroy NewSystemManager
func NewSystemManager() SystemManager {
	return &systemManager{systems: make(map[string]System)}
}
