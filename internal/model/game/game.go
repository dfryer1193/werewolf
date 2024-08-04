package game

import "github.com/dfryer1193/werewolf/internal/model"

type GameEgg struct {
	Id               string
	Signups          map[string]model.DiscordUser
	ConfirmedPlayers []model.DiscordUser
	Config           GameConfig
}

type Game struct {
	Id          string
	Players     map[string]model.Player // Discord username -> Player
	Config      GameConfig
	IsCompleted bool
}

type Day struct {
	GameID string
	Id     string
	Votes  map[string]string
}

type Night struct {
	GameID  string
	Id      string
	Actions ActionSet
}

func NewGame(signups []model.DiscordUser, cfg GameConfig) *GameEgg {
	// TODO: Get game ID back from database once implemented
	id := "tmp"

	signupMap := make(map[string]model.DiscordUser)
	for _, u := range signups {
		signupMap[u.Username] = u
	}

	return &GameEgg{
		Id:      id,
		Signups: signupMap,
		Config:  cfg,
	}
}

func (g GameEgg) StartGame() *Game {
	players := g.AssignRoles()

	return &Game{
		Id:          g.Id,
		Players:     players,
		IsCompleted: false,
	}
}

func (g GameEgg) AssignRoles() map[string]model.Player {
	playerMap := make(map[string]model.Player)
	for _, u := range g.ConfirmedPlayers {

	}

	return playerMap
}
