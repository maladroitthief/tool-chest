package abstract_factory

// BossRoomFactory is a concrete factory that implements all of RoomFactories methods
type SecretRoomFactory struct {
	RoomFactory
}

func (br *SecretRoomFactory) AddEnemy() string {
	return ""
}

func (br *SecretRoomFactory) AddObstacle() string {
	return "Hidden entrance"
}

func (br *SecretRoomFactory) AddItem() string {
	return "Coins"
}

func (br *SecretRoomFactory) AddReward() string {
	return "Skeleton key"
}
