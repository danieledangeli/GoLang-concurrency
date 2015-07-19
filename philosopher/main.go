package main

import (
	"fmt"
	"github.com/danieledangeli/concurrency/philosopher/philosopher"
	"github.com/danieledangeli/concurrency/philosopher/table"
	"math/rand"
	"time"
	"bytes"
	"log"
	"github.com/danieledangeli/concurrency/philosopher/config"
)

func thinkAndEat(p *philosopher.Philosopher, announce chan *philosopher.Philosopher, c config.Config) {
	fmt.Printf("%v is thinking.\n", p.Name)
	p.Think(randomDuration(c.MinThinkTimeMs, c.MaxThinkTimeMs))

	fmt.Printf("%v wanna to eat.\n", p.Name)
	p.Eat(randomDuration(c.MinEatTimeMs, c.MaxEatTimeMs))

	fmt.Printf("%v has eaten.\n", p.Name)
	announce <- p
}

func randomDuration(minMilliSec int, maxMilliSec int) time.Duration {
	rand.Seed(time.Now().Unix())
	amt := time.Duration(minMilliSec + rand.Intn(maxMilliSec - minMilliSec))

	return time.Millisecond * amt
}

func printStatus(philosopher []*philosopher.Philosopher, refreshTimeMs int) {
	for {
		var buffer bytes.Buffer

		buffer.WriteString("{\n")
		for _, phil := range philosopher {
			buffer.WriteString(fmt.Sprintf("[%v] %v %v\n", phil.Index, phil.Name, phil.Status))
		}

		buffer.WriteString("}\n");
		fmt.Println(buffer.String())
		time.Sleep(time.Duration(refreshTimeMs) * time.Millisecond)
	}
}

func main() {
	config, err := config.GetConfig("conf.yml")

	if(err != nil) {
		log.Fatal(err)
	}

	names := config.PhilosopherNames
	philosophers := make([]*philosopher.Philosopher, len(names))
	table := table.NewTable(len(names));

	for i, name := range names {
		philosophers[i] = philosopher.NewPhilosopher(name, i, table)
	}

	announce := make(chan *philosopher.Philosopher)

	//worker to keep all statuses
	go printStatus(philosophers, config.RefreshStateTimeMs)

	//run all philosopher to eat and think
	for _, phil := range philosophers {
		go thinkAndEat(phil, announce, config)
	}

	//run forever: when a philosopher has finished to think and eat, then starts again to dine
	//on the same announce channel
	for {
		phil := <-announce
		go thinkAndEat(phil, announce, config)
	}
}