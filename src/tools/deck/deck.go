package deck

type Card struct {
	rank     string
	suit     string
	card_num int
}

type Deck struct {
	ranks    []string
	suits    []string
	cardsStr []string
	cards    []Card
}

func NewDeck() *Deck {
	return &Deck{}
}

func NewDeckWithRanksAndSuits(ranks, suits []string) *Deck {
	deck := &Deck{
		ranks: ranks,
		suits: suits,
	}

	var cardNum int = 0
	for _, rank := range ranks {
		for _, suit := range suits {
			oneCard := rank + suit
			deck.cardsStr = append(deck.cardsStr, oneCard)
			deck.cards = append(deck.cards, Card{rank: rank, suit: suit, card_num: cardNum})
			cardNum += 1
		}
	}
	return deck
}

func (d *Deck) GetCards() []Card {
	return d.cards
}

//func main() {
//	// Create a new deck with ranks and suits
//	deck := NewDeckWithRanksAndSuits([]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}, []string{"H", "D", "C", "S"})
//
//	// Get the cards from the deck
//	cards := deck.GetCards()
//
//	// Print the cards
//	for i, card := range cards {
//		println("Card #" + strconv.Itoa(i+1) + ": " + card.rank + " of " + card.suit)
//	}
//}
