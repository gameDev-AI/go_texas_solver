package private_card_manager

import (
	"errors"
	"go_texas_solver/src/tools/card"
	"go_texas_solver/src/tools/private_card"
)

type PrivateCardsManager struct {
	privateCards    [][]private_card.PrivateCards
	cardPlayerIndex [][]int
	playerNumber    int
	initialBoard    uint64
}

func NewPrivateCardsManager() *PrivateCardsManager {
	return &PrivateCardsManager{}
}

func NewPrivateCardsManagerWithParams(privateCards [][]private_card.PrivateCards, playerNumber int, initialBoard uint64) *PrivateCardsManager {
	pcm := &PrivateCardsManager{
		privateCards:    privateCards,
		playerNumber:    playerNumber,
		initialBoard:    initialBoard,
		cardPlayerIndex: make([][]int, 52*52),
	}

	for i := 0; i < 52*52; i++ {
		pcm.cardPlayerIndex[i] = make([]int, pcm.playerNumber)
	}

	// 用一个二维数组记录每个Private Combo的对应index,方便从一方的手牌找对方的同名卡牌的index
	for playerId := 0; playerId < pcm.playerNumber; playerId++ {
		privateCombos := pcm.privateCards[playerId]
		for i := 0; i < len(privateCombos); i++ {
			onePrivateCombo := privateCombos[i]
			pcm.cardPlayerIndex[onePrivateCombo.HashCode()][playerId] = i
		}
	}

	pcm.setRelativeProbs()

	return pcm
}

func (pcm *PrivateCardsManager) GetPreflopCards(player int) []private_card.PrivateCards {
	return pcm.privateCards[player]
}

func (pcm *PrivateCardsManager) IndPlayer2Player(fromPlayer, toPlayer, index int) (int, error) {
	if index < 0 || index >= len(pcm.GetPreflopCards(fromPlayer)) {
		return -1, errors.New("index out of range")
	}

	//PrivateCards player_combo = this->getPreflopCards(from_player)[index];
	toPlayerIndex := pcm.cardPlayerIndex[pcm.privateCards[fromPlayer][index].HashCode()][toPlayer]
	if toPlayerIndex == -1 {
		return -1, nil
	} else {
		return toPlayerIndex, nil
	}
}

func (pcm *PrivateCardsManager) GetInitialReachProb(player int, initialBoard uint64) []float32 {
	cardsLen := len(pcm.privateCards[player])
	probs := make([]float32, cardsLen)
	for i := 0; i < cardsLen; i++ {
		pc := pcm.privateCards[player][i]
		if card.BoardsHasIntercept(initialBoard, card.BoardInts2Long(pc.GetHands())) {
			probs[i] = 0
		} else {
			probs[i] = pc.Weight
		}
	}
	return probs
}

func (pcm *PrivateCardsManager) setRelativeProbs() {
	players := len(pcm.privateCards)
	for playerId := 0; playerId < players; playerId++ {
		// TODO 这里只考虑了两个玩家的情况
		oppo := 1 - playerId
		playerProbSum := float32(0)

		for i := 0; i < len(pcm.privateCards[playerId]); i++ {
			oppoProbSum := float32(0)
			playerCard := &pcm.privateCards[playerId][i]
			playerLong := card.BoardInts2Long(playerCard.GetHands())

			//
			if card.BoardsHasIntercept(playerLong, pcm.initialBoard) {
				continue
			}

			for _, oppoCard := range pcm.privateCards[oppo] {
				oppoLong := card.BoardInts2Long(oppoCard.GetHands())
				if card.BoardsHasIntercept(oppoLong, pcm.initialBoard) || card.BoardsHasIntercept(oppoLong, playerLong) {
					continue
				}
				oppoProbSum += oppoCard.Weight

			}
			playerCard.RelativeProb = oppoProbSum * playerCard.Weight
			playerProbSum += playerCard.RelativeProb
		}
		for i := 0; i < len(pcm.privateCards[playerId]); i++ {
			pcm.privateCards[playerId][i].RelativeProb = pcm.privateCards[playerId][i].RelativeProb / playerProbSum
			//player_card.relative_prob = player_card.relative_prob / player_prob_sum;
		}

	}
}
