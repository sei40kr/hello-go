package main

import (
	"fmt"
	"math/rand"
)

// rps-simulator-extended.go --- Rock-Paper-Scissors Simulator (extended)
// author: Seong Yong-ju <sei40kr@gmail.com>

const (
	Rock     = 0
	Paper    = 1
	Scissors = 2
)

type Player struct{}

func (p *Player) RandomHand() int {
	return rand.Intn(3)
}

func (p *Player) RandomHandExtended() int {
	r := rand.Intn(4)
	if r == 3 {
		return Paper
	} else {
		return r
	}
}

func main() {
	p1, p2 := &Player{}, &Player{}
	total := 1000000
	draws, wins, loses := 0, 0, 0

	for i := 0; i < total; i++ {
		hand1, hand2 := p1.RandomHand(), p2.RandomHandExtended()

		t := (3 + hand1 - hand2) % 3
		switch t {
		case 0:
			// draw
			draws++
		case 1:
			// 1p wins
			wins++
		case 2:
			// 2p wins
			loses++
		}
	}

	fmt.Printf("勝利=%d回\n", wins)
	fmt.Printf("引分=%d回\n", draws)
	fmt.Printf("敗北=%d回\n", loses)
	fmt.Printf("------------------------\n")
	fmt.Printf("合計=%d回\n", total)
}
