package systems

type System interface {
	GetSystemName() string
	ExecuteSystem()
}
