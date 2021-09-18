package abstract_factory

// RoomFactory is an abstract interface to define how ConcreteFactories should behave
type RoomFactory interface {
	AddObstacle() string
	AddEnemy() string
	AddReward() string
	AddItem() string
}
