package model

import (
	"github.com/dfryer1193/werewolf/internal/model/game"
	"math/rand"
)

/*
Roles here: https://gist.github.com/dfryer1193/6d27f4496b63963728c59158c51838b2

Thinking out loud:
	Bot receives "scry A" from B
	The current night adds an entry that looks like B -> Action{Target: "A", meta: {priority: SEER, submitTime: 10:00pm ET}}
	When resolving night actions, call game.scry(actor="B", target="A")
*/

var RoleLUT map[string]*Role

type Team string

const (
	TEAM_WEREWOLF Team = "werewolf"
	TEAM_VILLAGER Team = "villager"
	TEAM_TANNER   Team = "tanner"
)

// TODO: Figure out how the mayor works

type BaseRole string

const (
	SEER            BaseRole = "seer"
	GHOST           BaseRole = "ghost"
	BODYGUARD       BaseRole = "bodyguard"
	CARRIAGE_DRIVER BaseRole = "carriage driver"
	CORONER         BaseRole = "coroner"
	DRUNK           BaseRole = "drunk"
	MASON           BaseRole = "mason"
	HUNTER          BaseRole = "hunter"
	WITCH           BaseRole = "witch"
	VILLAGER        BaseRole = "villager"
	WEREWOLF        BaseRole = "werewolf"
	MINION          BaseRole = "minion"
	TANNER          BaseRole = "tanner"
)

type RoleVariant string

const (
	NONE     RoleVariant = "none"
	CURSED   RoleVariant = "cursed"
	ANCIENT  RoleVariant = "ancient"
	DISEASED RoleVariant = "diseased"
	WW       RoleVariant = "werewolf"
	ALPHA    RoleVariant = "alpha"
	LYCAN    RoleVariant = "lycan"
)

type Role struct {
	Name            BaseRole
	Variant         RoleVariant
	Team            Team
	DisplayTeam     Team
	Action          func(g *game.Game, actor *Player, targets ...string) string
	ActionCondition func(g *game.Game) bool
}

var (
	ROLE_SEER = Role{
		Name:        SEER,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Scry,
	}
	ROLE_CURSED_SEER = Role{
		Name:        SEER,
		Variant:     CURSED,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Scry,
	}
	ROLE_ANCIENT_SEER = Role{
		Name:        SEER,
		Variant:     ANCIENT,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Scry,
	}
	ROLE_GHOST = Role{
		Name:        GHOST,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_BODYGUARD = Role{
		Name:        BODYGUARD,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Guard,
	}
	ROLE_CARRIAGE_DRIVER = Role{
		Name:        CARRIAGE_DRIVER,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Swap,
	}
	ROLE_CURSED = Role{
		Name:        VILLAGER,
		Variant:     CURSED,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_CORONER = Role{
		Name:        CORONER,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Autopsy,
	}
	ROLE_MASON = Role{
		Name:        MASON,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_LYCAN = Role{
		Name:        VILLAGER,
		Variant:     LYCAN,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_WEREWOLF,
	}
	ROLE_HUNTER = Role{
		Name:        HUNTER,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_WITCH = Role{
		Name:        WITCH,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_DISEASED = Role{
		Name:        VILLAGER,
		Variant:     DISEASED,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_VILLAGER = Role{
		Name:        VILLAGER,
		Variant:     NONE,
		Team:        TEAM_VILLAGER,
		DisplayTeam: TEAM_VILLAGER,
	}
	ROLE_WEREWOLF = Role{
		Name:        WEREWOLF,
		Variant:     NONE,
		Team:        TEAM_WEREWOLF,
		DisplayTeam: TEAM_WEREWOLF,
	}
	ROLE_ALPHA_WEREWOLF = Role{
		Name:        WEREWOLF,
		Variant:     ALPHA,
		Team:        TEAM_WEREWOLF,
		DisplayTeam: TEAM_WEREWOLF,
		Action:      AlphaConversion,
	}
	ROLE_WEREWOLF_SEER = Role{
		Name:        SEER,
		Variant:     WW,
		Team:        TEAM_WEREWOLF,
		DisplayTeam: TEAM_WEREWOLF,
		Action:      Scry,
	}
	ROLE_MINION = Role{
		Name:        MINION,
		Variant:     NONE,
		Team:        TEAM_WEREWOLF,
		DisplayTeam: TEAM_VILLAGER,
		Action:      Sabotage,
	}
	ROLE_TANNER = Role{
		Name:        TANNER,
		Variant:     NONE,
		Team:        TEAM_TANNER,
		DisplayTeam: TEAM_VILLAGER,
	}
)

func init() {
	RoleLUT = make(map[string]*Role)

	RoleLUT["SEER"] = &ROLE_SEER
	RoleLUT["CURSED_SEER"] = &ROLE_CURSED_SEER
	RoleLUT["ANCIENT_SEER"] = &ROLE_ANCIENT_SEER
	RoleLUT["GHOST"] = &ROLE_GHOST
	RoleLUT["BODYGUARD"] = &ROLE_BODYGUARD
	RoleLUT["CARRIAGE_DRIVER"] = &ROLE_CARRIAGE_DRIVER
	RoleLUT["CURSED"] = &ROLE_CURSED
	RoleLUT["CORONER"] = &ROLE_CORONER
	RoleLUT["MASON"] = &ROLE_MASON
	RoleLUT["LYCAN"] = &ROLE_LYCAN
	RoleLUT["HUNTER"] = &ROLE_HUNTER
	RoleLUT["WITCH"] = &ROLE_WITCH
	RoleLUT["DISEASED"] = &ROLE_DISEASED
	RoleLUT["VILLAGER"] = &ROLE_VILLAGER
	RoleLUT["WEREWOLF"] = &ROLE_WEREWOLF
	RoleLUT["ALPHA_WEREWOLF"] = &ROLE_ALPHA_WEREWOLF
	RoleLUT["WEREWOLF_SEER"] = &ROLE_WEREWOLF_SEER
	RoleLUT["MINION"] = &ROLE_MINION
	RoleLUT["TANNER"] = &ROLE_TANNER
}

func RoleLookup(name string) *Role {
	return RoleLUT[name]
}

func Scry(g *game.Game, actor *Player, targets ...string) string {
	resolvedTarget := g.CurrentNight.PlayerTargetMap[targets[0]]
	targetPlayer := g.Players[resolvedTarget]

	if targetPlayer.Role.Name == GHOST {
		return string(GHOST)
	}

	switch actor.Role.Variant {
	case CURSED:
		options := []Team{TEAM_VILLAGER, TEAM_WEREWOLF}
		return string(options[rand.Intn(len(options))])
	case ANCIENT:
		return string(targetPlayer.Role.Variant) + " " + string(targetPlayer.Role.Name)
	case WW:
		return string(targetPlayer.Role.Variant) + " " + string(targetPlayer.Role.Name)
	default:
		return string(targetPlayer.Role.DisplayTeam)
	}
}

type GuardState struct {
	target string
	guard  string
}

func Guard(g *game.Game, actor *Player, targets ...string) string {
	resolvedTarget := g.CurrentNight.PlayerTargetMap[targets[0]]
	guardState := GuardState{
		target: resolvedTarget,
		guard:  actor.Name,
	}

	g.CurrentNight.GuardedPlayers = append(g.CurrentNight.GuardedPlayers, guardState)

	return ""
}

func Autopsy(g *game.Game, actor *Player, targets ...string) string {
	targetPlayer := g.Players[targets[0]]
	return string(targetPlayer.Role.Variant) + " " + string(targetPlayer.Role.Name)
}

func Swap(g *game.Game, actor *Player, targets ...string) string {
	resolvedTarget1 := g.CurrentNight.PlayerTargetMap[targets[0]]
	resolvedTarget2 := g.CurrentNight.PlayerTargetMap[targets[1]]

	g.CurrentNight.PlayerTargetMap[resolvedTarget1] = resolvedTarget2
	g.CurrentNight.PlayerTargetMap[resolvedTarget2] = resolvedTarget1

	return ""
}

func AlphaConversion(g *game.Game, actor *Player, targets ...string) string {
	target := targets[0]
	targetPlayer := g.Players[target]

	targetPlayer.Role = &ROLE_WEREWOLF

	return ""
}

func Sabotage(g *game.Game, actor *Player, targets ...string) string {
	resolvedTarget := g.CurrentNight.PlayerTargetMap[targets[0]]

	g.CurrentNight.Sabotages = append(g.CurrentNight.Sabotages, resolvedTarget)

	return ""
}
