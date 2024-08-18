package game

import (
	"fmt"
	"github.com/dfryer1193/werewolf/internal/model"
	"math/rand/v2"
)

type GameEgg struct {
	Id               string
	Signups          map[string]model.DiscordUser
	ConfirmedPlayers []model.DiscordUser
	Config           GameConfig
}

type Game struct {
	Id           string
	VillageName  string
	Players      map[string]model.Player // Discord username -> Player
	Config       GameConfig
	IsCompleted  bool
	CurrentNight *Night
	CurrentDay   *Day
}

type Day struct {
	GameID   string
	Id       int
	Votes    map[string]string
	Complete bool
}

type Night struct {
	GameID          string
	Id              int
	PlayerTargetMap map[string]string
	GuardedPlayers  []model.GuardState
	Sabotages       []string
	Complete        bool
}

const baseNighttimeAnnouncement = `# Good Evening %s Village!

* @%s was executed by the village. They died with **%d** votes.

##NOTIFY_ROLE##
It is now **Night %d**. You have until ##NOTIFY TIME## to send your night actions.`

const wwKillMessage = `* @%s was killed by werewolves.`
const daytimeAnnouncementHeader = `# Good Morning %s Village!
`
const daytimeAnnouncementFooter = `##NOTIFY_ROLE##
It is now **Day %d You have until 9pm ET to send your votes on who you would like to execute. If there is a tie, a tied player will be picked at random. If no votes are received, a living player will be killed at random.`

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
	players := g.ConfirmedPlayers

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	playerIdx := 0
	for role, count := range g.Config.Roles {
		for i := 0; i < count; i++ {
			playerMap[players[playerIdx].Username] = model.Player{
				Name: players[playerIdx].Username,
				Role: model.RoleLookup(role),
			}
			playerIdx++
		}
	}

	return playerMap
}

// StartNight - usage: POST /werewolves.fyi/game/<id>/startNight
// This allows us to use a cron job created at the beginning of a game in order to handle the day/night cycle
func (g Game) StartNight() *Night {
	targetMap := make(map[string]string)
	for player, _ := range g.Players {
		targetMap[player] = player
	}

	announcement := g.CurrentDay.Finalize(g.VillageName)
	g.SendAnnouncement(announcement)
	g.SendRoleMessages()

	// TODO: Write new night to db
	g.CurrentNight = &Night{
		GameID:          g.Id,
		Id:              g.CurrentDay.Id,
		PlayerTargetMap: targetMap,
		GuardedPlayers:  make([]model.GuardState, 0),
		Sabotages:       make([]string, 0),
		Complete:        false,
	}

	return g.CurrentNight
}

// StartDay - usage: POST /werewolves.fyi/game/<id>/startDay
// This allows us to use a cron job created at the beginning of a game in order to handle the day/night cycle
func (g Game) StartDay() *Day {
	g.CurrentNight.Finalize(g.VillageName)
	g.SendRoleResponses()

	// TODO: Write new day to db
	g.CurrentDay = &Day{
		GameID:   g.Id,
		Id:       g.CurrentNight.Id + 1,
		Votes:    make(map[string]string),
		Complete: false,
	}

	return g.CurrentDay
}

func (d Day) Finalize(villageName string) string {
	// TODO: Write to db
	d.Complete = true
	// TODO: Support GM-supplied death messages
	voteCounts := make(map[string]int)
	for _, vote := range d.Votes {
		voteCounts[vote]++
	}

	topVotes := make([]string, 0)
	maxVoteCount := 0
	for player, count := range voteCounts {
		if count > maxVoteCount {
			topVotes = make([]string, 0)
			topVotes = append(topVotes, player)
			maxVoteCount = count
		} else if count == maxVoteCount {
			topVotes = append(topVotes, player)
		}
	}

	var killedPlayer string
	if len(topVotes) == 1 {
		killedPlayer = topVotes[0]
	} else {
		killedPlayer = topVotes[rand.IntN(len(topVotes))]
	}

	deathMessage := fmt.Sprintf(baseNighttimeAnnouncement, villageName, killedPlayer, maxVoteCount, d.Id)

	return deathMessage
}
