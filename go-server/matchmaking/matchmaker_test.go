package matchmaking

import "testing"

func TestMatchmaker(t *testing.T) {
	m := New()

	regions := []int{
		1,
	}

	m.AddPlayer("1", regions, 3000)
	m.AddPlayer("2", regions, 3000)
	m.AddPlayer("3", regions, 3000)
	m.AddPlayer("4", regions, 3000)
	m.AddPlayer("5", regions, 3000)
	m.AddPlayer("6", regions, 3000)
	m.AddPlayer("7", regions, 3000)
	m.AddPlayer("8", regions, 3000)
	m.AddPlayer("9", regions, 3000)
	m.AddPlayer("10", regions, 3000)
}
