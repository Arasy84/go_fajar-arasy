package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	if len(cards) == 0 {
		return "Tutup kartu"
	}

	card := cards[0]
	if card[0] == deck[0] || card[0] == deck[1] || card[1] == deck[0] || card[1] == deck[1] {
		return card
	}

	return playingDomino(cards[1:], deck)
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
