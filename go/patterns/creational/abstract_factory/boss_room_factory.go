package abstract_factory

// BossRoomFactory is a concrete factory that implements all of RoomFactories methods
type BossRoomFactory struct {
	RoomFactory
}

func (br *BossRoomFactory) AddEnemy() string {
	return "Dragon"
}

func (br *BossRoomFactory) AddObstacle() string {
	return "Rough terrain"
}

func (br *BossRoomFactory) AddItem() string {
	return "Heart container"
}

func (br *BossRoomFactory) AddReward() string {
	return "Potions"
}
