package tennis

import "math"

type Player struct {
	Name  string
	Score int64
}

func (p *Player) SetName(name string) {
	p.Name = name
}
func (p *Player) SetScore(score int64) {
	p.Score = score
}

func InitPlayer(name1, name2 string) (player1, player2 Player) {
	player1.SetName(name1)
	player2.SetName(name2)
	return
}

var ScoreMap = map[int64]string{
	0: "love",
	1: "fifteen",
	2: "thirty",
	3: "forty",
}

func Score(player1 Player, player2 Player) string {
	if player1.Score > 0 || player2.Score > 0 {
		if player1.Score >= 3 || player2.Score >= 3 {
			if player1.Score == player2.Score {
				return "deuce"
			}
			if math.Abs(float64(player1.Score)-float64(player2.Score)) == 1 {
				if player1.Score > player2.Score {
					return player1.Name + " adv"
				}
				return player2.Name + " adv"
			}
			if player1.Score > 3 || player2.Score > 3 {
				if player1.Score > player2.Score {
					return player1.Name + " win"
				}
				return player2.Name + " win"
			}
		}
		if player1.Score == player2.Score {
			return ScoreMap[player1.Score] + " all"
		}
		return ScoreMap[player1.Score] + " " + ScoreMap[player2.Score]
	}

	return "love all"
}
