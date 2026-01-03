package backend

import (
	"fmt"
)

type GameSchema struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	FileName string    `json:"fileName"`
	Section  string    `json:"section"`
	Keys     []KeyInfo `json:"keys"`
}

type KeyInfo struct {
	Key         string `json:"key"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Group       string `json:"group"`
}

func GetGames() []GameSchema {
	games := []GameSchema{}

	addGame := func(name, fileName, section string, keys []KeyInfo) {
		games = append(games, GameSchema{
			ID:       fileName,
			Name:     name,
			FileName: fileName,
			Section:  section,
			Keys:     keys,
		})
	}

	addGame("Five Nights at Freddy's", "freddy", "freddy", []KeyInfo{
		{"level", "Current Night", "int", ""},
		{"beatgame", "Beat Night 5 (Star 1)", "bool", ""},
		{"beat6", "Beat Night 6 (Star 2)", "bool", ""},
		{"beat7", "Beat 4/20 (Star 3)", "bool", ""},
		{"lives", "Lives (Unused)", "int", ""},
	})

	fnaf2Keys := []KeyInfo{
		{"level", "Current Night", "int", ""},
		{"cine", "Cutscenes", "int", ""},
		{"beatgame", "Beat Night 5 (Star 1)", "bool", ""},
		{"beat6", "Beat Night 6 (Star 2)", "bool", ""},
		{"beat7", "Beat 10/20 (Star 3)", "bool", ""},
		{"turn", "Death Minigames Seen", "int", ""},
	}
	for i := 1; i <= 10; i++ {
		fnaf2Keys = append(fnaf2Keys, KeyInfo{fmt.Sprintf("c%d", i), fmt.Sprintf("Challenge %d Completed", i), "bool", ""})
	}
	addGame("Five Nights at Freddy's 2", "freddy2", "freddy2", fnaf2Keys)

	fnaf3Keys := []KeyInfo{
		{"level", "Current Night", "int", ""},
		{"cine", "Seen End Minigame", "bool", ""},
		{"beatgame", "Beat Night 5 (Star 1)", "bool", ""},
		{"beat6", "Beat Nightmare (Star 2)", "bool", ""},
		{"goodend", "Happiest Day (Star 3)", "bool", ""},
		{"4thstar", "Beat Aggressive Nightmare (Star 4)", "bool", ""},
		{"fast", "Cheat: Fast Nights", "bool", ""},
		{"nocams", "Cheat: No Errors", "bool", ""},
		{"hyper", "Cheat: Aggressive", "bool", ""},
		{"vents", "Cheat: Radar", "bool", ""},
		{"bb", "Got Balloon (BBs Air Adventure)", "bool", ""},
		{"cake", "Got Cake (Mangles Quest)", "bool", ""},
	}
	for i := 1; i <= 5; i++ {
		fnaf3Keys = append(fnaf3Keys, KeyInfo{fmt.Sprintf("k%d", i), fmt.Sprintf("Gave Cake Child %d", i), "bool", ""})
	}
	addGame("Five Nights at Freddy's 3", "freddy3", "freddy3", fnaf3Keys)

	fnaf4Keys := []KeyInfo{
		{"night", "Current Night", "int", ""},
		{"test", "Seen Sound Test", "bool", ""},
		{"beat5", "Beat Night 5 (Star 1)", "bool", ""},
		{"beat6", "Beat Night 6 (Star 2)", "bool", ""},
		{"beat7", "Beat Nightmare (Star 3)", "bool", ""},
		{"beat8", "Beat 4/20 (Star 4)", "bool", ""},
		{"scene", "Last Seen Minigame", "int", ""},
	}
	for i := 1; i <= 6; i++ {
		fnaf4Keys = append(fnaf4Keys, KeyInfo{fmt.Sprintf("s%d", i), fmt.Sprintf("Challenge %d (Star)", i), "bool", ""})
	}
	addGame("Five Nights at Freddy's 4", "fn4", "fn4", fnaf4Keys)

	for slot := 1; slot <= 3; slot++ {
		fnafwKeys := []KeyInfo{
			{"newgame", "New Game Flag", "bool", ""},
			{"mode", "Game Mode (1-Adv, 2-Fixed)", "int", ""},
			{"diff", "Difficulty (1-Norm, 2-Hard)", "int", ""},
			{"started", "Seen First Fredbear Dialogue", "bool", ""},
			{"locked", "Lock Party Selection", "bool", ""},
			{"cine", "Fredbear Position/Next Dialogue", "int", ""},
			{"tokens", "Faz-Tokens", "int", ""},
			{"x", "Position X", "int", ""},
			{"y", "Position Y", "int", ""},
			{"resetpos", "Set Position to Spawn", "bool", ""},
			{"area", "Current Zone", "int", ""},
			{"sw5", "Opened Key Room", "bool", ""},
			{"w3", "Entered Dusting Fields", "bool", ""},
			{"key", "Have Key", "bool", ""},
			{"find", "Current Clock", "int", ""},
			{"armor", "Current Armor Strength", "int", ""},
			{"fish", "DeeDee Gone / Fish Caught", "int", ""},
			{"pearl", "Times Pearl Caught", "int", ""},
			{"beatgame1", "Defeated Security", "bool", ""},
			{"last", "Encountered Animdude", "bool", ""},
			{"beatgame2", "Defeated Animdude", "bool", ""},
			{"beatgame3", "Defeated Chipper's Revenge", "bool", ""},
			{"portal", "Backstage Portal", "bool", ""},
			{"beatgame7", "Defeated Chica's Magic Rainbow", "bool", ""},
			{"showend", "In End Cutscene", "bool", ""},
			{"seconds", "Playtime (Seconds)", "int", ""},
			{"min", "Playtime (Minutes)", "int", ""},
			{"hour", "Playtime (Hours)", "int", ""},
		}

		for i := 1; i <= 4; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("sw%d", i), fmt.Sprintf("Opened Area %d", i), "bool", ""})
		}
		for i := 6; i <= 9; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("sw%d", i), fmt.Sprintf("Opened Gate %d", i), "bool", ""})
		}
		for i := 1; i <= 5; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("g%d", i), fmt.Sprintf("Completed Clock Minigame %d", i), "bool", ""})
		}
		for i := 1; i <= 8; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("s%d", i), fmt.Sprintf("Character in Party Slot %d", i), "int", ""})
		}
		for i := 1; i <= 48; i++ {
			group := fmt.Sprintf("Character %d", i)
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("%dhave", i), "Have Character", "bool", group})
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("%dnext", i), "EXP", "int", group})
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("%dlv", i), "Level - 1", "int", group})
		}
		for i := 1; i <= 21; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("p%d", i), fmt.Sprintf("Purchased Byte %d", i), "bool", ""})
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("c%d", i), fmt.Sprintf("Have Chip %d", i), "bool", ""})
		}
		for i := 1; i <= 4; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("active%db", i), fmt.Sprintf("Active Byte in Slot %d", i), "int", ""})
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("active%d", i), fmt.Sprintf("Active Chip in Slot %d", i), "int", ""})
		}
		for i := 1; i <= 3; i++ {
			fnafwKeys = append(fnafwKeys, KeyInfo{fmt.Sprintf("ar%d", i), fmt.Sprintf("Have Armor %d", i), "bool", ""})
		}

		addGame(fmt.Sprintf("FNaF World (Slot %d)", slot), fmt.Sprintf("fnafw%d", slot), "fnafw", fnafwKeys)
	}

	infoKeys := []KeyInfo{
		{"first", "Seen Intro", "bool", ""},
		{"beatgame1", "Security Trophy", "bool", ""},
		{"beatgame2", "Animdude Trophy", "bool", ""},
		{"beatgame3", "Chipper's Revenge Trophy", "bool", ""},
		{"beatgame6", "Universe End Trophy", "bool", ""},
		{"beatgame5", "Clock Ending Trophy", "bool", ""},
		{"beatgame4", "4th Layer Trophy", "bool", ""},
		{"gotpearl", "Pearl Trophy", "bool", ""},
		{"all", "Fan Trophy", "bool", ""},
		{"beatgame7", "Chica's Magic Rainbow Trophy", "bool", ""},
	}
	for i := 1; i <= 3; i++ {
		infoKeys = append(infoKeys, KeyInfo{fmt.Sprintf("mode%d", i), fmt.Sprintf("Slot %d: Game Mode", i), "int", ""})
		infoKeys = append(infoKeys, KeyInfo{fmt.Sprintf("diff%d", i), fmt.Sprintf("Slot %d: Difficulty", i), "int", ""})
		infoKeys = append(infoKeys, KeyInfo{fmt.Sprintf("hour%d", i), fmt.Sprintf("Slot %d: Playtime (Hours)", i), "int", ""})
		infoKeys = append(infoKeys, KeyInfo{fmt.Sprintf("min%d", i), fmt.Sprintf("Slot %d: Playtime (Minutes)", i), "int", ""})
	}
	addGame("FNaF World (Global Data)", "info", "info", infoKeys)

	slKeys := []KeyInfo{
		{"current", "Current Night", "int", ""},
		{"intro", "Seen Intro", "bool", ""},
		{"beat1", "Beat Night 5 (Star 1)", "bool", ""},
		{"keycard", "Beat Baby Minigame (Star 2)", "bool", ""},
		{"beat3", "Beat Private Room (Star 3)", "bool", ""},
	}
	addGame("Sister Location", "sl", "sl", slKeys)

	ffpsKeys := []KeyInfo{
		{"night", "Current Day", "int", ""},
		{"money", "Money", "int", ""},
		{"phase", "Current Phase", "int", ""},
		{"newgame", "Blueprint Mode Unlocked", "bool", ""},
	}
	addGame("Freddy Fazbear's Pizzeria Simulator", "FNAF6", "FNAF6", ffpsKeys)

	ucnKeys := []KeyInfo{
		{"hs", "High Score", "int", ""},
		{"bestminute", "50/20 Best Time (Min)", "int", ""},
		{"besttens", "50/20 Best Time (10s)", "int", ""},
		{"bestseconds", "50/20 Best Time (1s)", "int", ""},
		{"besttenths", "50/20 Best Time (0.1s)", "int", ""},
	}
	addGame("Ultimate Custom Night", "CN", "CN", ucnKeys)

	fis2Keys := []KeyInfo{
		{"timer", "Total Seconds", "int", ""},
		{"gems", "Blue Gems", "int", ""},
	}
	addGame("Freddy in Space 2", "FIS2", "FIS2", fis2Keys)

	sbKeys := []KeyInfo{
		{"intro", "Seen Intro", "bool", ""},
		{"c1", "Beat Normal Mode (Star 1)", "bool", ""},
		{"c2", "Beat Hard Mode (Star 2)", "bool", ""},
		{"c3", "Beat Hard (No Death/Animdude) (Star 3)", "bool", ""},
	}
	addGame("Security Breach: Fury's Rage", "SB", "SB", sbKeys)

	fis3Keys := []KeyInfo{
		{"level", "Level - 1", "int", ""},
		{"exp", "XP until next level", "int", ""},
		{"Jason", "Unlocked Freddy", "bool", ""},
		{"Ian", "Unlocked Bonnie", "bool", ""},
		{"Ethan", "Unlocked Foxy", "bool", ""},
		{"Braden", "Unlocked Mangle", "bool", ""},
	}
	addGame("Freddy in Space 3", "c3save", "c3save", fis3Keys)

	return games
}
