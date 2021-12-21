package main

import "fmt"

type Player struct {
	position int64
	score    int64
}

func (p *Player) move(roll int64) {
	p.position += roll
	p.position = p.position % 10
	if p.position == 0 {
		p.position = 10
	}
}

func (p *Player) updateScore() {
	p.score += p.position
}

type DeterministicDice int64

func (d *DeterministicDice) roll() (value int64) {
	*d++
	if *d == 101 {
		*d = 1
	}
	value = int64(*d)
	return
}

func playGame(p1, p2 Player, d DeterministicDice) (result int64) {
	rolls := int64(0)
	for i := 0; p1.score < 1000 && p2.score < 1000; i++ {
		var cur *Player
		if i%2 == 0 {
			cur = &p1
		} else {
			cur = &p2
		}

		for j := 0; j < 3; j++ {
			rolls++
			cur.move(d.roll())
		}
		cur.updateScore()
	}

	if p1.score > p2.score {
		result = rolls * p2.score
	} else {
		result = rolls * p1.score
	}
	return
}

type Game struct {
	p1, p2 Player
	turn   int
}

func generateRolls() (rolls map[int64]int64) {
	rolls = make(map[int64]int64)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				rolls[int64(i+j+k)] += 1
			}
		}
	}
	return
}

func partTwo(p1, p2 Player) (result int64) {
	rolls := generateRolls()

	games := make(map[Game]int64)
	p1Wins := int64(0)
	p2Wins := int64(0)

	games[Game{p1: p1, p2: p2, turn: 1}] = 1

	for len(games) > 0 {
		newGames := make(map[Game]int64)
		for game, count := range games {
			for roll, rollCount := range rolls {
				p1 := Player{position: game.p1.position, score: game.p1.score}
				p2 := Player{position: game.p2.position, score: game.p2.score}
				numNewGames := count * rollCount

				if game.turn == 1 {
					p1.move(roll)
					p1.updateScore()
					if p1.score >= 21 {
						p1Wins += numNewGames
					} else {
						newGames[Game{p1: p1, p2: p2, turn: 2}] += numNewGames
					}
				} else {
					p2.move(roll)
					p2.updateScore()
					if p2.score >= 21 {
						p2Wins += numNewGames
					} else {
						newGames[Game{p1: p1, p2: p2, turn: 1}] += numNewGames
					}
				}
			}
		}
		games = newGames
	}

	if p1Wins > p2Wins {
		result = p1Wins
	} else {
		result = p2Wins
	}
	return
}

func main() {
	fmt.Println("===Test Data===")
	testP1 := Player{position: 4, score: 0}
	testP2 := Player{position: 8, score: 0}
	fmt.Println(playGame(testP1, testP2, DeterministicDice(0)) == 739785)
	fmt.Println(partTwo(testP1, testP2) == 444356092776315)

	fmt.Printf("\n===Real Data===\n")
	p1 := Player{position: 4, score: 0}
	p2 := Player{position: 1, score: 0}
	fmt.Println(playGame(p1, p2, DeterministicDice(0)) == 913560)
	fmt.Println(partTwo(p1, p2))

}
