package philosopher

import (
	t "github.com/danieledangeli/concurrency/philosopher/table"
	"time"
	"fmt"
)

type Philosopher struct {
	Name string
	Index int
	Table t.Tabler
	Status string
}

func NewPhilosopher(name string, i int, t t.Tabler) *Philosopher {

	if(i < 0) {
		panic("Cannot make philospher with negative index")
	}

	p := new(Philosopher)

	p.Name = name
	p.Index = i
	p.Table = t
	p.Status = "SITTING_ON_TABLE"

	return p
}

func (p *Philosopher) Think(d time.Duration) {
	p.Status = "THINK";
	time.Sleep(d)
}


func (p *Philosopher) Eat(d time.Duration) {
	rF := p.takeRightFork();
	rF.Lock();

	lF := p.takeLeftFork();
	lF.Lock();

	p.Status = "EAT";
	time.Sleep(d)

	rF.Unlock();
	lF.Unlock();

	p.Status = "FED";
}

func (p *Philosopher) takeLeftFork() *t.Fork {
	var i = p.Index - 1;

	if(p.Index == 0) {
		i = len(p.Table.GetForks()) - 1;
	}

	p.Status = fmt.Sprintf("WAIT_FOR_LEFT_FORK[%v]", i)
	return p.Table.GetForks()[i];
}

func (p *Philosopher) takeRightFork() *t.Fork {
	p.Status = fmt.Sprintf("WAIT_FOR_RIGHT_FORK[%v]", p.Index)
	return p.Table.GetForks()[p.Index];
}
