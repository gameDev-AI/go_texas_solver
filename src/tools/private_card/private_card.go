package private_card

import "go_texas_solver/src/tools/card"

type PrivateCards struct {
	card1        int
	card2        int
	Weight       float32
	RelativeProb float32
	cardVec      []int
	hashCode     int
	boardLong    uint64
}

func NewPrivateCards() *PrivateCards {
	return &PrivateCards{}
}

func NewPrivateCardsWithParams(card1, card2 int, weight float32) *PrivateCards {
	privateCards := &PrivateCards{
		card1:  card1,
		card2:  card2,
		Weight: weight,
	}
	privateCards.RelativeProb = 0
	if card1 > card2 {
		privateCards.hashCode = card1*52 + card2
		privateCards.cardVec = []int{card1, card2}
	} else {
		privateCards.hashCode = card2*52 + card1
		privateCards.cardVec = []int{card2, card1}
	}
	privateCards.boardLong = card.BoardInts2Long(privateCards.cardVec)
	return privateCards
}

func (pc *PrivateCards) toBoardLong() uint64 {
	return pc.boardLong
}

func (pc *PrivateCards) HashCode() int {
	return pc.hashCode
}

func (pc *PrivateCards) ToString() string {
	if pc.card1 > pc.card2 {
		return card.IntCard2Str(pc.card1) + card.IntCard2Str(pc.card2)
	} else {
		return card.IntCard2Str(pc.card2) + card.IntCard2Str(pc.card1)
	}
}

func (pc *PrivateCards) GetHands() []int {
	return pc.cardVec
}
