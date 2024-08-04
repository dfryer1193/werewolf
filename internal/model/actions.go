package model

import (
	"cmp"
	"time"
)

type ActionSet struct {
}

type actionMeta struct {
	priority   int
	submitTime time.Time
}

type Action interface {
}

func fromActionList(actions []actionMeta) ActionSet {

}

func (a ActionSet) Add(action actionMeta) {

}

func (a actionMeta) Compare(b actionMeta) int {
	priorityCompare := cmp.Compare(a.priority, b.GetPriority())

	if priorityCompare != 0 {
		return priorityCompare
	}

	return cmp.Compare(a.GetSubmittedTime().UnixMilli(), b.GetSubmittedTime().UnixMilli())
}
