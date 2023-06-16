package street_setting

type StreetSetting struct {
	betSizes   []float32
	raiseSizes []float32
	donkSizes  []float32
	allin      bool
}

func NewStreetSetting(betSizes []float32, raiseSizes []float32, donkSizes []float32, allin bool) *StreetSetting {
	return &StreetSetting{
		betSizes:   betSizes,
		raiseSizes: raiseSizes,
		donkSizes:  donkSizes,
		allin:      allin,
	}
}
