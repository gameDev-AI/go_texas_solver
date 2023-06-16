package rule

import (
	"errors"
)

type Deck struct {
	// Implementation of Deck struct in Golang is not provided in the given C++ code, so it has to be defined separately.
}

type GameTreeBuildingSettings struct {
	// Implementation of GameTreeBuildingSettings struct in Golang is not provided in the given C++ code, so it has to be defined separately.
}

type Rule struct {
	deck                    Deck
	oop_commit              float32
	ip_commit               float32
	current_round           int
	raise_limit             int
	small_blind             float32
	big_blind               float32
	stack                   float32
	build_settings          GameTreeBuildingSettings
	allin_threshold         float32
	initial_effective_stack float32
}

func NewRule(deck Deck, oop_commit, ip_commit float32, current_round, raise_limit int, small_blind,
	big_blind, stack float32, build_settings GameTreeBuildingSettings, allin_threshold float32) (*Rule, error) {

	if oop_commit != ip_commit {
		return nil, errors.New("ip commit should equal with oop commit")
	}

	return &Rule{
		deck:                    deck,
		oop_commit:              oop_commit,
		ip_commit:               ip_commit,
		current_round:           current_round,
		raise_limit:             raise_limit,
		small_blind:             small_blind,
		big_blind:               big_blind,
		stack:                   stack,
		build_settings:          build_settings,
		allin_threshold:         allin_threshold,
		initial_effective_stack: stack - oop_commit,
	}, nil
}

func (r *Rule) GetPot() float32 {
	return r.oop_commit + r.ip_commit
}

func (r *Rule) GetCommit(player int) float32 {
	if player == 0 {
		return r.ip_commit
	} else if player == 1 {
		return r.oop_commit
	} else {
		panic("unknown player")
	}
}

//func main() {
//	// create a new rule instance
//	deck := Deck{}
//	build_settings := GameTreeBuildingSettings{}
//	r, err := NewRule(deck, 10.0, 10.0, 1, 100, 1.0, 2.0, 1000.0, build_settings, 0.9)
//	if err != nil {
//		panic(err)
//	}
//
//	// get the pot from the rule
//	pot := r.GetPot()
//	println("Pot: ", pot)
//
//	// get the commit amount for a particular player from the rule
//	oopCommit := r.GetCommit(1)
//	println("Out of position commit: ", oopCommit)
//}
