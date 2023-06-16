package card

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	card                string
	card_int            int
	card_number_in_deck int
}

func NewCard() *Card {
	return &Card{card: "empty"}
}

func NewCardWithParams(card string, card_num_in_deck int) *Card {
	c := &Card{card: card, card_number_in_deck: card_num_in_deck}
	c.card_int = strCard2Int(c.card)
	return c
}

func NewCardFromString(card string) *Card {
	c := &Card{card: card}
	c.card_int = strCard2Int(c.card)
	c.card_number_in_deck = -1
	return c
}

func (c *Card) Empty() bool {
	if c.card == "empty" {
		return true
	}
	return false
}

func (c *Card) GetCard() string {
	return c.card
}

func (c *Card) GetCardInt() int {
	return c.card_int
}

func (c *Card) GetNumberInDeckInt() (int, error) {
	if c.card_number_in_deck == -1 {
		return 0, errors.New("card number in deck cannot be -1")
	}
	return c.card_number_in_deck, nil
}

func Card2Int(card Card) int {
	return strCard2Int(card.GetCard())
}

func strCard2Int(card string) int {
	rank := card[0]
	suit := card[1]
	if len(card) != 2 {
		panic(fmt.Sprintf("card %s not found", card))
	}
	return (RankToInt(string(rank))-2)*4 + SuitToInt(string(suit))
}

func IntCard2Str(card int) string {
	rank := card/4 + 2
	suit := card - (rank-2)*4
	return RankToString(rank) + SuitToString(suit)
}

func BoardCards2Long(cards []string) uint64 {
	cardsObjs := make([]Card, len(cards))
	for i, c := range cards {
		cardsObjs[i] = *NewCardFromString(c)
	}
	return boardCards2Long(cardsObjs)
}

func BoardCard2Long(card *Card) uint64 {
	return boardInt2Long(card.GetCardInt())
}

func boardCards2Long(cards []Card) uint64 {
	boardInts := make([]int, len(cards))
	for i, c := range cards {
		boardInts[i] = Card2Int(c)
	}
	return BoardInts2Long(boardInts)
}

//func BoardCards2Html(cards []Card) string {
//	retHtml := ""
//	for _, oneCard := range cards {
//		if oneCard.Empty() {
//			continue
//		}
//		retHtml += oneCard.ToFormattedHtml()
//	}
//	return retHtml
//}

func BoardsHasIntercept(board1 uint64, board2 uint64) bool {
	return (board1 & board2) != 0
}

func boardInt2Long(board int) uint64 {
	// 这里hard code了一副扑克牌是52张
	if board < 0 || board >= 52 {
		panic(fmt.Sprintf("Card with id %d not found", board))
	}
	// uint64_7 的range 在0 ~ + 2^ 64之间,所以不用太担心溢出问题
	return (uint64(1) << board)
}

func BoardInts2Long(board []int) uint64 {
	if len(board) < 1 || len(board) > 7 {
		panic(fmt.Sprintf("Card length incorrect: %d", len(board)))
	}
	var boardLong uint64 = 0
	for _, oneCard := range board {
		// 这里hard code了一副扑克牌是52张
		if oneCard < 0 || oneCard >= 52 {
			panic(fmt.Sprintf("Card with id %s not found", strconv.Itoa(oneCard)))
		}
		// uint64_7 的range 在0 ~ + 2^ 64之间,所以不用太担心溢出问题
		boardLong += (uint64(1) << oneCard)
	}
	return boardLong
}

var suitMap = map[int]string{
	0: "c",
	1: "d",
	2: "h",
	3: "s",
}

var rankMap = map[int]string{
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "T",
	11: "J",
	12: "Q",
	13: "K",
	14: "A",
}

func SuitToString(suit int) string {
	if s, ok := suitMap[suit]; ok {
		return s
	}
	return "c"
}

func RankToString(rank int) string {
	if r, ok := rankMap[rank]; ok {
		return r
	}
	return "2"
}

func RankToInt(rank string) int {
	switch rank {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		return 2
	}
}

func SuitToInt(suit string) int {
	switch suit {
	case "c":
		return 0
	case "d":
		return 1
	case "h":
		return 2
	case "s":
		return 3
	default:
		return 0
	}
}

func (c *Card) GetSuits() []string {
	return []string{"c", "d", "h", "s"}
}

func (c *Card) ToString() string {
	return c.card
}

func (c *Card) ToFormattedString() string {
	qString := strings.ReplaceAll(c.card, "c", "♣️")
	qString = strings.ReplaceAll(qString, "d", "♦️")
	qString = strings.ReplaceAll(qString, "h", "♥️")
	qString = strings.ReplaceAll(qString, "s", "♠️")
	return qString
}

//func (c *Card) ToFormattedHtml() string {
//	qString := c.card
//	if strings.Contains(qString, "c") {
//		qString = strings.ReplaceAll(qString, "c", "<span style=\"color:black;\">&#9827;<\/span>")
//	} else if strings.Contains(qString, "d") {
//		qString = strings.ReplaceAll(qString, "d", "<span style=\"color:red;\">&#9830;<\/span>")
//	} else if strings.Contains(qString, "h") {
//		qString = strings.ReplaceAll(qString, "h", "<span style=\"color:red;\">&#9829;<\/span>")
//	} else if strings.Contains(qString, "s") {
//		qString = strings.ReplaceAll(qString, "s", "<span style=\"color:black;\">&#9824;<\/span>")
//	}
//	return qString
//}

func Long2board(boardLong uint64) ([]int, error) {
	board := make([]int, 0, 7)
	for i := 0; i < 52; i++ {
		if boardLong&1 == 1 {
			board = append(board, i)
		}
		boardLong >>= 1
	}
	if len(board) < 1 || len(board) > 7 {
		return nil, errors.New(fmt.Sprintf("board length not correct, board length %d", len(board)))
	}
	return board, nil
}

func long2boardCards(boardLong uint64) []Card {
	board, _ := Long2board(boardLong)
	boardCards := make([]Card, len(board))
	for i := 0; i < len(board); i++ {
		oneBoard := board[i]
		boardCards[i] = Card{card: IntCard2Str(oneBoard)}
	}
	if len(boardCards) < 1 || len(boardCards) > 7 {
		panic(fmt.Sprintf("board length not correct, board length %d", len(boardCards)))
	}
	retval := make([]Card, len(boardCards))
	for i := 0; i < len(boardCards); i++ {
		retval[i] = boardCards[i]
	}
	return retval
}
