package main

import (
	"fmt"
	"github.com/danieledangeli/concurrency/philosopher/philosopher"
	"github.com/danieledangeli/concurrency/philosopher/table"
	"math/rand"
	"time"
	"bytes"
)

func thinkAndEat(p *philosopher.Philosopher, announce chan *philosopher.Philosopher,) {
	fmt.Printf("%v is thinking.\n", p.Name)
	p.Think(randomDuration(500, 1000))

	fmt.Printf("%v wanna to eat.\n", p.Name)
	p.Eat(randomDuration(500, 1000))

	fmt.Printf("%v has eaten.\n", p.Name)
	announce <- p
}

func randomDuration(minMillicSec, maxMilliSec int) time.Duration {
	rand.Seed(time.Now().Unix())
	amt := time.Duration(minMillicSec + rand.Intn(maxMilliSec - minMillicSec))

	return time.Millisecond * amt
}

func printStatus(philosopher []*philosopher.Philosopher) {
	for {
		var buffer bytes.Buffer

		buffer.WriteString("{\n")
		for _, phil := range philosopher {
			buffer.WriteString(fmt.Sprintf("[%v] %v %v\n", phil.Index, phil.Name, phil.Status))
		}

		buffer.WriteString("}\n");
		fmt.Println(buffer.String())
		time.Sleep(time.Duration(400) * time.Millisecond)
	}
}

func main() {
	names := []string{"Zeno", "Plato", "Epicurus", "Locke", "Aristotle"}

	philosophers := make([]*philosopher.Philosopher, len(names))
	table := table.NewTable(len(names));

	for i, name := range names {
		philosophers[i] = philosopher.NewPhilosopher(name, i, table)
	}

	announce := make(chan *philosopher.Philosopher)

	//run all philosopher to eat and think
	go printStatus(philosophers)

	//run all philosopher to eat and think
	for _, phil := range philosophers {
		go thinkAndEat(phil, announce)
	}

	//run forever: when a philosopher has finhed to think and eat, then starts again to dine
	for {
		phil := <-announce
		go thinkAndEat(phil, announce)
	}
}