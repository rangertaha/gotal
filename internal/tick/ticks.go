package tick

import (
	"sort"

	"github.com/rangertaha/gotal/internal"
)

// Ticks is a named, time-ordered collection of Tick values.
type Ticks struct {
	name  string
	ticks []internal.Tick
}

func NewTicks(name string, tickInputs ...internal.Tick) *Ticks {
	sortTicks(tickInputs)
	return &Ticks{
		name:  name,
		ticks: tickInputs,
	}
}

func (t *Ticks) Name() string           { return t.name }
func (t *Ticks) Ticks() []internal.Tick { return t.ticks }

func (t *Ticks) Add(ticks ...internal.Tick) {
	t.ticks = append(t.ticks, ticks...)
}

func (t *Ticks) Delete(index int) {
	t.ticks = append(t.ticks[:index], t.ticks[index+1:]...)
}

func (t *Ticks) Update(ticks ...internal.Tick) {
	t.ticks = append(t.ticks, ticks...)
	sortTicks(t.ticks)
}

func sortTicks(ticks []internal.Tick) {
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i].Time().Before(ticks[j].Time())
	})
}
