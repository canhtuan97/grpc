package structs
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Types          []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}   `json:"type"`
	}   `json:"types"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	}   `json:"sprites"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}   `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		}   `json:"stat"`
	}   `json:"stats"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}   `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
			}   `json:"move_learn_method"`
		}   `json:"version_group_details"`
	}   `json:"moves"`
}