package game_tree_build_settings

import (
	"errors"
	"go_texas_solver/src/tools/street_setting"
)

type GameTreeBuildingSettings struct {
	FlopIP   street_setting.StreetSetting
	TurnIP   street_setting.StreetSetting
	RiverIP  street_setting.StreetSetting
	FlopOOP  street_setting.StreetSetting
	TurnOOP  street_setting.StreetSetting
	RiverOOP street_setting.StreetSetting
}

func NewGameTreeBuildingSettings(flopIP, turnIP, riverIP, flopOOP, turnOOP, riverOOP street_setting.StreetSetting) *GameTreeBuildingSettings {
	return &GameTreeBuildingSettings{
		FlopIP:   flopIP,
		TurnIP:   turnIP,
		RiverIP:  riverIP,
		FlopOOP:  flopOOP,
		TurnOOP:  turnOOP,
		RiverOOP: riverOOP,
	}
}

func (settings *GameTreeBuildingSettings) GetSetting(player, round string) (*street_setting.StreetSetting, error) {
	switch {
	case player == "ip" && round == "flop":
		return &settings.FlopIP, nil
	case player == "ip" && round == "turn":
		return &settings.TurnIP, nil
	case player == "ip" && round == "river":
		return &settings.RiverIP, nil
	case player == "oop" && round == "flop":
		return &settings.FlopOOP, nil
	case player == "oop" && round == "turn":
		return &settings.TurnOOP, nil
	case player == "oop" && round == "river":
		return &settings.RiverOOP, nil
	default:
		return nil, errors.New("get_setting not valid")
	}
}
