package matchmaking

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const minScore = 0.5

type player struct {
	SteamID   string
	Regions   []int
	MMR       int
	TimeAdded time.Time
}

type group struct {
	players map[*player]struct{}
	regions map[int]int // how many players per region
}

func (g *group) addPlayer(player *player) {
	g.players[player] = struct{}{}
	for _, region := range player.Regions {
		g.regions[region]++
	}
}

func (g *group) isComplete() bool {
	return len(g.players) == 10
}

func (g *group) removePlayer(player *player) {
	delete(g.players, player)

	for _, region := range player.Regions {
		g.regions[region]--
	}
}

// playerFitsByRegion returns true if the players region matches
func (g *group) playerFitsByRegion(player *player) bool {
	fitsRegion := false
	for _, r := range player.Regions {
		if g.regions[r] == len(g.players) {
			fitsRegion = true
			break
		}
	}

	return fitsRegion
}

func (g *group) calculateScore(player *player) float32 {
	if !g.playerFitsByRegion(player) {
		return 0
	}

	totalMMR := 0

	for p := range g.players {
		totalMMR += p.MMR
	}

	averageMMR := float32(totalMMR) / float32(len(g.players))

	diff := math.Abs(float64(averageMMR - float32(player.MMR)))

	return 0
}

type Matchmaker struct {
	sync.Mutex
	groups  map[*group]struct{}
	players map[*player]*group
}

func New() *Matchmaker {
	return &Matchmaker{}
}

func (m *Matchmaker) createGroup(player *player) {
	regions := make(map[int]int, len(player.Regions))
	for _, r := range player.Regions {
		regions[r] = 1
	}

	g := &group{
		players: map[*player]struct{}{
			player: {},
		},
		regions: regions,
	}
	m.groups[g] = struct{}{}
	m.players[player] = g
}

func (m *Matchmaker) RemovePlayer(player *player) {
	g := m.players[player]
	g.removePlayer(player)
	if len(g.players) == 0 {
		delete(m.groups, g)
	}
	delete(m.players, player)
}

func (m *Matchmaker) AddPlayer(steamID string, regions []int, mmr int) {
	p := &player{
		MMR:       mmr,
		Regions:   regions,
		SteamID:   steamID,
		TimeAdded: time.Now(),
	}

	m.addPlayer(p)
}

func (m *Matchmaker) addPlayer(p *player) {
	m.Lock()
	defer m.Unlock()

	if len(m.groups) == 0 {
		m.createGroup(p)
		return
	}

	score, group := m.findBestGroup(p)

	if score <= 0 {
		m.createGroup(p)
		return
	}

	m.addPlayerToGroup(p, group)
	if group.isComplete() {
		m.startMatch(group)
	}
}

func (m *Matchmaker) addPlayerToGroup(player *player, group *group) {
	group.addPlayer(player)
	m.players[player] = group
}

func (m *Matchmaker) startMatch(group *group) {
	fmt.Println("Found match")
	for p := range group.players {
		delete(m.players, p)
	}
	delete(m.groups, group)
}

func (m *Matchmaker) findBestGroup(player *player) (float32, *group) {
	maxScore := float32(-1)
	var maxScoreGroup *group
	for group := range m.groups {
		score := group.calculateScore(player)
		if score > maxScore {
			maxScore = score
			maxScoreGroup = group
		}
	}

	return maxScore, maxScoreGroup
}
