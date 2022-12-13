package kata

type Dependency interface {
	Value() int
}

type dependencyImpl struct {
	value int
}

func (d dependencyImpl) Value() int {
	return d.value
}

func NewDependencyImpl() Dependency {
	return &dependencyImpl{value: 1}
}
