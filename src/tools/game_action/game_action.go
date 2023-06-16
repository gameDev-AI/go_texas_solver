package game_action

import (
	"fmt"
)

type PokerActions int

const (
	BEGIN PokerActions = iota
	ROUNDBEGIN
	BET
	RAISE
	CHECK
	FOLD
	CALL
)

type GameActions struct {
	action PokerActions
	amount float64
}

func NewGameActions(action PokerActions, amount float64) GameActions {
	if (action == RAISE || action == BET) && amount == -1 {
		panic(fmt.Sprintf("raise/bet amount should not be -1, found %f", amount))
	} else if (action == CHECK || action == FOLD || action == CALL) && amount != -1 {
		panic(fmt.Sprintf("check/fold/call amount should be -1, found %f", amount))
	}

	return GameActions{action, amount}
}

func (a GameActions) GetAction() PokerActions {
	return a.action
}

func (a GameActions) GetAmount() float64 {
	return a.amount
}

func pokerActionToString(pokerActions PokerActions) string {
	switch pokerActions {
	case BEGIN:
		return "BEGIN"
	case ROUNDBEGIN:
		return "ROUNDBEGIN"
	case BET:
		return "BET"
	case RAISE:
		return "RAISE"
	case CHECK:
		return "CHECK"
	case FOLD:
		return "FOLD"
	case CALL:
		return "CALL"
	default:
		panic("PokerActions not found")
	}
}

func (a GameActions) ToString() string {
	if a.amount == -1 {
		return pokerActionToString(a.action)
	} else {
		return fmt.Sprintf("%s %.2f", pokerActionToString(a.action), a.amount)
	}
}
