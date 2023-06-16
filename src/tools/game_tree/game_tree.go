package game_tree

import (
	"encoding/json"
	"fmt"
	"go_texas_solver/src/tools/deck"
	"go_texas_solver/src/tools/game_tree_build_settings"
	"go_texas_solver/src/tools/game_tree_node"
	"go_texas_solver/src/tools/rule"
	"io/ioutil"
)

type GameRound int

const (
	FLOP GameRound = iota
	TURN
	RIVER
)

type GameActions []string

func (g GameRound) String() string {
	return [...]string{"FLOP", "TURN", "RIVER"}[g]
}

type Rule struct {
	deck            deck.Deck
	oopCommit       float32
	ipCommit        float32
	currentRound    int
	raiseLimit      int
	smallBlind      float32
	bigBlind        float32
	stack           float32
	buildingSetting game_tree_build_settings.GameTreeBuildingSettings
	allinThreshold  float32
}

// Some other code omitted for brevity.

type GameTree struct {
	root         *game_tree_node.GameTreeNode
	treeJsonDir  string
	deck         deck.Deck
	treeSettings game_tree_build_settings.GameTreeBuildingSettings
}

func (g *GameTree) getSettings(round int, player int) interface{} {
	var settings interface{}
	switch GameRound(round) {
	case FLOP:
		if player == 0 {
			settings = g.treeSettings.FlopIP
		} else {
			settings = g.treeSettings.FlopOOP
		}
	case TURN:
		if player == 0 {
			settings = g.treeSettings.TurnIP
		} else {
			settings = g.treeSettings.TurnOOP
		}
	case RIVER:
		if player == 0 {
			settings = g.treeSettings.RiverIP
		} else {
			settings = g.treeSettings.RiverOOP
		}
	}

	return settings
}

func (g *GameTree) Load() error {
	data, err := ioutil.ReadFile(g.treeJsonDir)
	if err != nil {
		return err
	}

	treeData := make(map[string]interface{})
	if err := json.Unmarshal(data, &treeData); err != nil {
		return err
	}

	rootData := treeData["root"].(map[string]interface{})
	g.root = g.recurrentGenerateTreeNode(rootData, nil)

	return nil
}

//func (g *GameTree) recurrentGenerateTreeNode(nodeData map[string]interface{}, parentNode *game_tree_node.GameTreeNode) *game_tree_node.GameTreeNode {
//	actionsData := nodeData["actions"].([]interface{})
//	actions := make(game_action.GameActions, len(actionsData))
//	for i, v := range actionsData {
//		actions[i] = v.(string)
//	}
//
//	childrenData := nodeData["children"].([]interface{})
//	children := make([]*game_tree_node.GameTreeNode, len(childrenData))
//	for i, v := range childrenData {
//		childData := v.(map[string]interface{})
//		child := g.recurrentGenerateTreeNode(childData, parentNode)
//		children[i] = child
//	}
//
//	currentPlayer := int(nodeData["currentPlayer"].(float64))
//	depth := int(nodeData["depth"].(float64))
//	pot := nodeData["pot"].(float64)
//	roundInt := int(nodeData["round"].(float64))
//	round := game_tree_node.GameRound(roundInt)
//
//	return &game_tree_node.GameTreeNode{
//		Actions:       actions,
//		Children:      children,
//		CurrentPlayer: currentPlayer,
//		Depth:         depth,
//		Pot:           pot,
//		Round:         round,
//	}
//}

func (gt *GameTree) _build(node *GameTreeNode, rule rule.Rule, lastAction string, checkTimes int, raiseTimes int) *GameTreeNode {
	switch node.Type() {
	case game_tree_node.ACTION:
		gt.buildAction(node.(*ActionNode), rule, lastAction, checkTimes, raiseTimes)
	case game_tree_node.SHOWDOWN:
	case game_tree_node.TERMINAL:
	case game_tree_node.CHANCE:
		gt.buildChance(node.(*ChanceNode), rule)
	default:
		panic("unknown node type")
	}
	return node
}

func (gt *GameTree) buildChance(root *ChanceNode, rule rule.Rule) {
	pot := float64(rule.get_pot())
	nextrule := rule.Clone()
	if rule.current_round > 3 {
		panic(fmt.Sprintf("current round not valid: %d", rule.current_round))
	}
	var oneNode *GameTreeNode

	if rule.oop_commit == rule.ip_commit && rule.oop_commit == rule.stack {
		if rule.current_round >= 3 {
			p1Commit := rule.ipCommit
			p2Commit := rule.oop_commit
			peaceGetback := (p1Commit + p2Commit) / 2.0

			payoffs := make([][]float64, 2)
			payoffs[0] = []float64{p2Commit, -p2Commit}
			payoffs[1] = []float64{-p1Commit, p1Commit}
			peaceGetbackVec := []float64{peaceGetback - p1Commit, peaceGetback - p2Commit}
			oneNode = NewShowdownNode(peaceGetbackVec, payoffs, intToGameRound(rule.current_round), pot, root)
		} else {
			nextrule.current_round += 1
			if rule.current_round > 3 {
				panic(fmt.Sprintf("current round not valid: %d", rule.current_round))
			}
			oneNode = NewChanceNode(nil, intToGameRound(rule.current_round+1), pot, root, gt.deck.GetCards())
		}
	} else {
		oneNode = NewActionNode(nil, nil, 1, intToGameRound(rule.current_round), pot, root)
	}

	if rule.current_round > 3 {
		panic(fmt.Sprintf("current round not valid: %d", rule.current_round))
	}
	gt._build(oneNode, nextrule, "begin", 0, 0)
	root.SetChildren(oneNode)
}
