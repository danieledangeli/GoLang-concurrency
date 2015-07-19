package table

import (
	"sync"
)

type Tabler interface {
	GetForks() []*Fork
}

//fork modeled as an array of mutex
type Fork struct {
	sync.Mutex
}

type Table struct {
	forks []*Fork
}

func (t *Table) GetForks() []*Fork {
	return t.forks;
}


//create a new table with a pre-allocated number of forks
func NewTable(forks int) *Table {

	if(forks <= 0) {
		panic("Cannot make table with negative or 0 forks")
	}

	t := new(Table)
	t.forks = make([]*Fork, forks)

	for i := 0; i < forks; i++ {
		t.forks[i] = new(Fork)
	}

	return t
}