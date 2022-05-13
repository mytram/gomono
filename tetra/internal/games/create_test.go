package games

import (
	"testing"

	"github.com/mytram/gomono/pkg/model"
)

func TestCreate(t *testing.T) {

	p1 := &model.Player{Name: "player 1", Number: "1"}
	p2 := &model.Player{Name: "player 2", Number: "2"}
	p3 := &model.Player{Name: "player 3", Number: "3"}

	team := &model.Team{
		Name: "a team",
		Players: []*model.Player{
			p1, p2, p3,
		},
	}

	Create(team)
}
