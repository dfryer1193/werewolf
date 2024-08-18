package model

type Player struct {
	Name   string
	Role   *Role
	IsDead bool
	DiedOn string
}
