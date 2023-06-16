package solver

import "go_texas_solver/src/tools/game_tree"

type Solver struct {
	tree *game_tree.GameTree
}

func NewSolver() *Solver {
	return &Solver{}
}

func NewSolverWithTree(tree *game_tree.GameTree) *Solver {
	return &Solver{tree: tree}
}

func (s *Solver) GetTree() *game_tree.GameTree {
	return s.tree
}
