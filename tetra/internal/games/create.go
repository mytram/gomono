package games

import (
	"math/rand"
	"time"

	"github.com/mytram/gomono/pkg/model"
)

func Create(team *model.Team) *model.Game {
	players := make([]*model.Player, len(team.Players))

	copy(players, team.Players)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })

	return &model.Game{
		Players: players,
	}
}
