package action_node

import (
	"errors"
	"fmt"
	"go_texas_solver/src/tools/discounted_cfr_trainable"
	"go_texas_solver/src/tools/game_action"
	"go_texas_solver/src/tools/private_card"
)

type GameTreeNodeType int

const (
	ROOT GameTreeNodeType = iota
	ACTION
	CHANCE
)

type GameRound int

type GameTreeNode struct {
	round      GameRound
	pot        float64
	parent     *GameTreeNode
	trainables []*discounted_cfr_trainable.DiscountedCfrTrainable
}

type ActionNode struct {
	GameTreeNode
	actions        []game_action.GameActions
	childrens      []*GameTreeNode
	player         int
	playerPrivates []*private_card.PrivateCards
}

func NewActionNode(actions []game_action.GameActions, childrens []*GameTreeNode, player int, round GameRound, pot float64, parent *GameTreeNode) *ActionNode {
	return &ActionNode{
		GameTreeNode: GameTreeNode{
			round:      round,
			pot:        pot,
			parent:     parent,
			trainables: make([]*discounted_cfr_trainable.DiscountedCfrTrainable, len(childrens)),
		},
		actions:   actions,
		childrens: childrens,
		player:    player,
	}
}

func (a *ActionNode) GetType() GameTreeNodeType {
	return ACTION
}

func (a *ActionNode) GetPlayer() int {
	return a.player
}

func (a *ActionNode) GetActions() []game_action.GameActions {
	return a.actions
}

func (a *ActionNode) GetChildrens() []*GameTreeNode {
	return a.childrens
}

func (a *ActionNode) GetTrainable(i int, createOnSite bool) (*discounted_cfr_trainable.DiscountedCfrTrainable, error) {
	if i >= len(a.trainables) {
		return nil, errors.New(fmt.Sprintf("size unacceptable %d > %d", i, len(a.trainables)))
	}
	if a.trainables[i] == nil && createOnSite {
		a.trainables[i] = discounted_cfr_trainable.NewDiscountedCfrTrainable(a.playerPrivates, a)
	}
	return a.trainables[i], nil
}

func (a *ActionNode) SetTrainable(trainables []*discounted_cfr_trainable.DiscountedCfrTrainable, playerPrivates []*private_card.PrivateCards) {
	a.trainables = trainables
	a.playerPrivates = playerPrivates
}

func (a *ActionNode) SetActions(actions []game_action.GameActions) {
	a.actions = actions
}

func (a *ActionNode) SetChildrens(childrens []*GameTreeNode) {
	a.childrens = childrens
}
