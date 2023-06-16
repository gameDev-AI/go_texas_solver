package trainable

import "errors"

type Trainable interface {
	GetAverageStrategy() []float32
	GetCurrentStrategy() []float32
	UpdateRegrets(regrets []float32, iterationNumber int, reachProbs []float32)
	SetEv(evs []float32)
	CopyStrategy(otherTrainable Trainable)
	//DumpStrategy(withState bool) struct{}
	//DumpEvs() struct{}
	GetType() int
}

const (
	CFR_PLUS_TRAINABLE = iota
	DISCOUNTED_CFR_TRAINABLE
)

type trainable struct {
	// Implementation of trainable struct in Golang is not provided in the given C++ code, so it has to be defined separately.
}

func NewTrainable(tType int) (Trainable, error) {
	switch tType {
	case CFR_PLUS_TRAINABLE:
		return &trainable{}, nil
	case DISCOUNTED_CFR_TRAINABLE:
		return &trainable{}, nil
	default:
		return nil, errors.New("invalid trainable type")
	}
}

func (t *trainable) GetAverageStrategy() []float32 {
	// Implementation of GetAverageStrategy method in Golang is not provided in the given C++ code, so it has to be defined separately.
	return []float32{}
}

func (t *trainable) GetCurrentStrategy() []float32 {
	// Implementation of GetCurrentStrategy method in Golang is not provided in the given C++ code, so it has to be defined separately.
	return []float32{}
}

func (t *trainable) UpdateRegrets(regrets []float32, iterationNumber int, reachProbs []float32) {
	// Implementation of UpdateRegrets method in Golang is not provided in the given C++ code, so it has to be defined separately.
}

func (t *trainable) SetEv(evs []float32) {
	// Implementation of SetEv method in Golang is not provided in the given C++ code, so it has to be defined separately.
}

func (t *trainable) CopyStrategy(otherTrainable Trainable) {
	// Implementation of CopyStrategy method in Golang is not provided in the given C++ code, so it has to be defined separately.
}

//func (t *trainable) DumpStrategy(withState bool) struct{} {
//	// Implementation of DumpStrategy method in Golang is not provided in the given C++ code, so it has to be defined separately.
//	return json.JSON{}
//}
//
//func (t *trainable) DumpEvs() json.JSON {
//	// Implementation of DumpEvs method in Golang is not provided in the given C++ code, so it has to be defined separately.
//	return json.JSON{}
//}

func (t *trainable) GetType() int {
	// Implementation of GetType method in Golang is not provided in the given C++ code, so it has to be defined separately.
	return 0
}

//func main() {
//	trainable, err := NewTrainable(CFR_PLUS_TRAINABLE)
//	if err != nil {
//		panic(err)
//	}
//
//	avgStrategy := trainable.GetAverageStrategy()
//	println("Average strategy: ", avgStrategy)
//
//	currentStrategy := trainable.GetCurrentStrategy()
//	println("Current strategy: ", currentStrategy)
//
//	regrets := []float32{1.0, 2.0, 3.0}
//	reachProbs := []float32{0.5, 0.4, 0.1}
//	trainable.UpdateRegrets(regrets, 100, reachProbs)
//
//	evs := []float32{10.0, 20.0, 30.0}
//	trainable.SetEv(evs)
//
//	otherTrainable, err := NewTrainable(DISCOUNTED_CFR_TRAINABLE)
//	if err != nil {
//		panic(err)
//	}
//
//	trainable.CopyStrategy(otherTrainable)
//
//	dumpStrategy := trainable.DumpStrategy(true)
//	println("Dump strategy: ", dumpStrategy)
//
//	dumpEvs := trainable.DumpEvs()
//	println("Dump EVs: ", dumpEvs)
//
//	tType := trainable.GetType()
//	println("Trainable type: ", tType)
//}
