package abstract_factory

type Room struct {
	Enemy    string
	Item     string
	Obstacle string
	Reward   string
}

// NewRoom accepts the abstract RoomFactory interface for creating rooms
func NewRoom(f RoomFactory) *Room {
	r := Room{}

	r.Enemy = f.AddEnemy()
	r.Item = f.AddItem()
	r.Obstacle = f.AddObstacle()
	r.Reward = f.AddReward()

	return &r
}
