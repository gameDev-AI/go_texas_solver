package game_tree_node

import (
	"fmt"
	"go_texas_solver/src/tools/game_action"
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

type GameTreeNodeType int

const (
	ACTION GameTreeNodeType = iota
	SHOWDOWN
	TERMINAL
	CHANCE
)

type GameRound int

const (
	PREFLOP GameRound = iota
	FLOP
	TURN
	RIVER
)

type GameTreeNode struct {
	parent        *GameTreeNode
	Actions       game_action.GameActions `json:"actions"`
	Children      []*GameTreeNode         `json:"children"`
	CurrentPlayer int                     `json:"currentPlayer"`
	Depth         int                     `json:"depth"`
	Pot           float64                 `json:"pot"`
	Round         GameRound               `json:"round"`
}

func NewGameTreeNode() *GameTreeNode {
	return &GameTreeNode{}
}

func NewGameTreeNodeWithParent(round GameRound, pot float64, parent *GameTreeNode) *GameTreeNode {
	return &GameTreeNode{
		Round:  round,
		Pot:    pot,
		parent: parent,
	}
}

func (n *GameTreeNode) GetParent() *GameTreeNode {
	return n.parent
}

func (n *GameTreeNode) SetParent(parent *GameTreeNode) {
	n.parent = parent
}

func (n *GameTreeNode) GetRound() GameRound {
	return n.Round
}

func (n *GameTreeNode) GetPot() float64 {
	return n.Pot
}

func (n *GameTreeNode) PrintHistory() {
	node := n
	for node != nil {
		parent_node := node.parent
		if parent_node == nil {
			break
		}
		if parent_node.GetType() == ACTION {
			action_node := parent_node.(*ActionNode)
			for i, a := range action_node.GetActions() {
				if action_node.GetChildrens()[i] == node {
					fmt.Printf("<- (player %d %s)", action_node.GetPlayer(), ActionToString(a))
				}
			}
		} else if parent_node.GetType() == CHANCE {
			chance_node := parent_node.(*ChanceNode)
			for i, c := range chance_node.GetCards() {
				if chance_node.GetChildrens()[i] == node {
					fmt.Printf("<- (deal card %s)", c)
				}
			}
		} else {
			fmt.Printf("<- (%v)", node)
		}
		node = parent_node
	}
	fmt.Println()
}

func (n *GameTreeNode) ToString() string {
	parent_node := n.parent
	round := ""
	if parent_node == nil {
		return fmt.Sprintf("%s begin", round)
	}
	if parent_node.GetType() == ACTION {
		action_node := parent_node.(*ActionNode)
		for i, a := range action_node.GetActions() {
			if action_node.GetChildrens()[i] == n {
				return fmt.Sprintf("<- (player %d %s)", action_node.GetPlayer(), ActionToString(a))
			}
		}
	} else if parent_node.GetType() == CHANCE {
		chance_node := parent_node.(*ChanceNode)
		for i, c := range chance_node.GetCards() {
			if chance_node.GetChildrens()[i] == n {
				return fmt.Sprintf("<- (deal card %s)", c)
			}
		}
	}
	return ""
}

func PrintNodeHistory(node *GameTreeNode) {
	for node != nil {
		parent_node := node.parent
		if parent_node == nil {
			break
		}
		if parent_node.GetType() == ACTION {
			action_node := parent_node.(*ActionNode)
			for i, a := range action_node.GetActions() {
				if action_node.GetChildrens()[i] == node {
					fmt.Printf("<- (player %d %s)", action_node.GetPlayer(), ActionToString(a))
				}
			}
		} else if parent_node.GetType() == CHANCE {
			chance_node := parent_node.(*ChanceNode)
			for i, c := range chance_node.GetCards() {
				if chance_node.GetChildrens()[i] == node {
					fmt.Printf("<- (deal card %s)", c)
				}
			}
		} else {
			fmt.Printf("<- (%v)", node.ToString())
		}
		node = parent_node
	}
	fmt.Println()
}

func GameRound2Int(gameRound GameRound) int {
	switch gameRound {
	case PREFLOP:
		return 0
	case FLOP:
		return 1
	case TURN:
		return 2
	case RIVER:
		return 3
	default:
		panic(fmt.Sprintf("round %v not found", gameRound))
	}
}

func IntToGameRound(round int) GameRound {
	switch round {
	case 0:
		return PREFLOP
	case 1:
		return FLOP
	case 2:
		return TURN
	case 3:
		return RIVER
	default:
		panic(fmt.Sprintf("round %d not found", round))
	}
}
