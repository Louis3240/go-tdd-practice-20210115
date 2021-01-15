package tennis

import "testing"

func ErrPrintIfNotEqual(t *testing.T, result string, player1, player2 Player) {
	if Score(player1, player2) != result {
		t.Error("not equal")
	}
}

func TestLoveAll(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	ErrPrintIfNotEqual(t, "love all", Player1, Player2)
}
func TestFifteenLove(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(1)
	ErrPrintIfNotEqual(t, "fifteen love", Player1, Player2)
}
func TestThirtyLove(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(2)
	ErrPrintIfNotEqual(t, "thirty love", Player1, Player2)
}
func TestFortyLove(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(3)
	ErrPrintIfNotEqual(t, "forty love", Player1, Player2)
}

func TestLoveFifteen(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player2.SetScore(1)
	ErrPrintIfNotEqual(t, "love fifteen", Player1, Player2)
}
func TestLoveThirty(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player2.SetScore(2)
	ErrPrintIfNotEqual(t, "love thirty", Player1, Player2)
}
func TestLoveForty(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player2.SetScore(3)
	ErrPrintIfNotEqual(t, "love forty", Player1, Player2)
}
func TestFifteenAll(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(1)
	Player2.SetScore(1)
	ErrPrintIfNotEqual(t, "fifteen all", Player1, Player2)
}
func TestThirtyAll(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(2)
	Player2.SetScore(2)
	ErrPrintIfNotEqual(t, "thirty all", Player1, Player2)
}
func TestDeuce(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(3)
	Player2.SetScore(3)
	ErrPrintIfNotEqual(t, "deuce", Player1, Player2)
}
func TestPlayer1Adv(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(4)
	Player2.SetScore(3)
	ErrPrintIfNotEqual(t, "tom adv", Player1, Player2)
}
func TestPlayer1Win(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(5)
	Player2.SetScore(3)
	ErrPrintIfNotEqual(t, "tom win", Player1, Player2)
}
func TestPlayer2Adv(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(5)
	Player2.SetScore(6)
	ErrPrintIfNotEqual(t, "elsa adv", Player1, Player2)
}
func TestPlayer2Win(t *testing.T) {
	Player1, Player2 := InitPlayer("tom", "elsa")
	Player1.SetScore(6)
	Player2.SetScore(8)
	ErrPrintIfNotEqual(t, "elsa win", Player1, Player2)
}
