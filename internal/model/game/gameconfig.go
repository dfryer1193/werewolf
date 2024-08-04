package game

import "time"

type GameConfig struct {
	StartTime time.Time
	PlayerCap int
	Roles     map[string]int
}
