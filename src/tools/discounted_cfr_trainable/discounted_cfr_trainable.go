package discounted_cfr_trainable

import (
	"go_texas_solver/src/tools/action_node"
	"go_texas_solver/src/tools/private_card"
	"math"
)

type DiscountedCfrTrainable struct {
	action_node   *action_node.ActionNode
	private_cards []*private_card.PrivateCards
	action_number int
	card_number   int
	evs           []float64
	r_plus        []float64
	r_plus_sum    []float64
	cum_r_plus    []float64
}

const alpha float64 = 1.5
const beta float64 = 0.5
const gamma float64 = 2
const theta float64 = 0.9

func NewDiscountedCfrTrainable(private_cards []*private_card.PrivateCards, action_node *action_node.ActionNode) *DiscountedCfrTrainable {
	d := &DiscountedCfrTrainable{}
	d.private_cards = private_cards
	d.action_node = action_node
	//todo
	//d.action_number = len(action_node.Childrens)
	d.card_number = len(private_cards)
	d.evs = make([]float64, d.action_number*d.card_number)
	d.r_plus = make([]float64, d.action_number*d.card_number)
	d.r_plus_sum = make([]float64, d.card_number)
	d.cum_r_plus = make([]float64, d.action_number*d.card_number)
	return d
}

func (d *DiscountedCfrTrainable) isAllZeros(input_array []float64) bool {
	for _, i := range input_array {
		if i != 0 {
			return false
		}
	}
	return true
}

func (d *DiscountedCfrTrainable) getAverageStrategy() []float64 {
	average_strategy := make([]float64, d.action_number*d.card_number)
	for private_id := 0; private_id < d.card_number; private_id++ {
		r_plus_sum := 0.0
		for action_id := 0; action_id < d.action_number; action_id++ {
			index := action_id*d.card_number + private_id
			r_plus_sum += d.cum_r_plus[index]
		}

		for action_id := 0; action_id < d.action_number; action_id++ {
			index := action_id*d.card_number + private_id
			if r_plus_sum != 0 {
				average_strategy[index] = d.cum_r_plus[index] / r_plus_sum
			} else {
				average_strategy[index] = 1.0 / float64(d.action_number)
			}
		}
	}
	return average_strategy
}

func (dct *DiscountedCfrTrainable) getcurrentStrategy() []float64 {
	return dct.GetCurrentStrategyNoCache()
}

func (d *DiscountedCfrTrainable) copyStrategy(otherTrainable DiscountedCfrTrainable) {
	trainableVar := otherTrainable
	copy(d.r_plus, trainableVar.r_plus)
	copy(d.cum_r_plus, trainableVar.cum_r_plus)
}

func (dct *DiscountedCfrTrainable) GetCurrentStrategyNoCache() []float64 {
	currentStrategy := make([]float64, dct.action_number*dct.card_number)
	if len(dct.r_plus_sum) == 0 {
		for i := range currentStrategy {
			currentStrategy[i] = float64(1.0 / float32(dct.action_number))
		}
	} else {
		for actionID := 0; actionID < dct.action_number; actionID++ {
			for privateID := 0; privateID < dct.card_number; privateID++ {
				index := actionID*dct.card_number + privateID
				if dct.r_plus_sum[privateID] != 0 {
					currentStrategy[index] = float64(float32(math.Max(0.0, float64(dct.r_plus[index])) / float64(dct.r_plus_sum[privateID])))
				} else {
					currentStrategy[index] = float64(1.0 / float32(dct.action_number))
				}
				// Debugging
				if math.IsNaN(float64(dct.r_plus[index])) {
					panic("nan found")
				}
			}
		}
	}
	return currentStrategy
}

func (dct *DiscountedCfrTrainable) SetEv(evs []float32) {
	if len(evs) != len(dct.evs) {
		panic("size mismatch in DiscountedCfrTrainable setEV")
	}
	for i := range evs {
		dct.evs[i] = float64(evs[i])
	}
}

func (d *DiscountedCfrTrainable) updateRegrets(regrets []float64, iterationNumber int, reachProbs []float64) {
	if len(regrets) != d.action_number*d.card_number {
		panic("length not match")
	}

	alphaCoef := math.Pow(float64(iterationNumber), float64(alpha))
	alphaCoef = alphaCoef / (1 + alphaCoef)

	for actionID := 0; actionID < d.action_number; actionID++ {
		for privateID := 0; privateID < d.card_number; privateID++ {
			index := actionID*d.card_number + privateID
			oneReg := regrets[index]

			// 更新 R+
			d.r_plus[index] = oneReg + d.r_plus[index]
			if d.r_plus[index] > 0 {
				d.r_plus[index] *= alphaCoef
			} else {
				d.r_plus[index] *= beta
			}

			d.r_plus_sum[privateID] += math.Max(0.0, d.r_plus[index])

			// 更新累计策略
			// d.cumRPlus[index] += d.rPlus[index] * iterationNumber;
			// d.cumRPlusSum[privateID] += d.cumRPlus[index];
		}
	}
	currentStrategy := d.GetCurrentStrategyNoCache()
	strategyCoef := math.Pow((float64(iterationNumber) / (float64(iterationNumber) + 1)), gamma)
	for actionID := 0; actionID < d.action_number; actionID++ {
		for privateID := 0; privateID < d.card_number; privateID++ {
			index := actionID*d.card_number + privateID
			d.cum_r_plus[index] *= theta
			d.cum_r_plus[index] += currentStrategy[index] * strategyCoef // * reachProbs[privateID];
			// d.cumRPlusSum[privateID] += d.cumRPlus[index] ;
		}
	}
}

func (d *DiscountedCfrTrainable) dumpStrategy(withState bool) map[string]interface{} {
	if withState {
		panic("state storage not implemented")
	}

	strategy := make(map[string][]float64)
	averageStrategy := d.getAverageStrategy()
	gameActions := d.action_node.GetActions()
	actionsStr := make([]string, 0)

	for _, oneAction := range gameActions {
		actionsStr = append(actionsStr, oneAction.ToString())
	}

	for i, onePrivateCard := range d.private_cards {
		oneStrategy := make([]float64, d.action_number)

		for j := 0; j < d.action_number; j++ {
			strategyIndex := j*d.card_number + i
			oneStrategy[j] = averageStrategy[strategyIndex]
		}
		strategy[onePrivateCard.ToString()] = oneStrategy
	}

	retJSON := make(map[string]interface{})
	retJSON["actions"] = actionsStr
	retJSON["strategy"] = strategy
	return retJSON
}

func (d *DiscountedCfrTrainable) dumpEvs() map[string]interface{} {
	evs := make(map[string][]float64)
	averageEvs := d.evs
	gameActions := d.action_node.GetActions()
	actionsStr := make([]string, 0)

	for _, oneAction := range gameActions {
		actionsStr = append(actionsStr, oneAction.ToString())
	}

	for i, onePrivateCard := range d.private_cards {
		oneEvs := make([]float64, d.action_number)

		for j := 0; j < d.action_number; j++ {
			evsIndex := j*d.card_number + i
			oneEvs[j] = averageEvs[evsIndex]
		}
		evs[onePrivateCard.ToString()] = oneEvs
	}

	retJSON := make(map[string]interface{})
	retJSON["actions"] = actionsStr
	retJSON["evs"] = evs
	return retJSON
}
