package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	poke struct {
		httpClient HTTP
	}

	Ability struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	Form struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	Version struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	GameIndice struct {
		GameIndex int     `json:"game_index"`
		Version   Version `json:"version"`
	}

	Specie struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	Stat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	GetPokemonResponse struct {
		Abilities []struct {
			Ability  Ability `json:"ability"`
			IsHidden bool    `json:"is_hidden"`
			Slot     int     `json:"slot"`
		} `json:"abilities"`
		BaseExperience         int           `json:"base_experience"`
		Forms                  []Form        `json:"forms"`
		GameIndices            []GameIndice  `json:"game_indices"`
		Height                 int           `json:"height"`
		HeldItems              []interface{} `json:"held_items"`
		ID                     int           `json:"id"`
		IsDefault              bool          `json:"is_default"`
		LocationAreaEncounters string        `json:"location_area_encounters"`
		Moves                  []struct {
			Move struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move"`
			VersionGroupDetails []struct {
				LevelLearnedAt  int `json:"level_learned_at"`
				MoveLearnMethod struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"move_learn_method"`
				VersionGroup struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"version_group"`
			} `json:"version_group_details"`
		} `json:"moves"`
		Name    string `json:"name"`
		Order   int    `json:"order"`
		Species Specie `json:"species"`
		Sprites struct {
			BackDefault      string      `json:"back_default"`
			BackFemale       interface{} `json:"back_female"`
			BackShiny        string      `json:"back_shiny"`
			BackShinyFemale  interface{} `json:"back_shiny_female"`
			FrontDefault     string      `json:"front_default"`
			FrontFemale      interface{} `json:"front_female"`
			FrontShiny       string      `json:"front_shiny"`
			FrontShinyFemale interface{} `json:"front_shiny_female"`
		} `json:"sprites"`
		Stats []struct {
			BaseStat int  `json:"base_stat"`
			Effort   int  `json:"effort"`
			Stat     Stat `json:"stat"`
		} `json:"stats"`
		Types []struct {
			Slot int  `json:"slot"`
			Type Type `json:"type"`
		} `json:"types"`
		Weight int `json:"weight"`
	}

	GetCharacteristicResponse struct {
		Descriptions []struct {
			Description string `json:"description"`
			Language    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"descriptions"`
		GeneModulo  int `json:"gene_modulo"`
		HighestStat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"highest_stat"`
		ID             int   `json:"id"`
		PossibleValues []int `json:"possible_values"`
	}

	GetSpeciesResponse struct {
		BaseHappiness int `json:"base_happiness"`
		CaptureRate   int `json:"capture_rate"`
		Color         struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"color"`
		EggGroups []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"egg_groups"`
		EvolutionChain struct {
			URL string `json:"url"`
		} `json:"evolution_chain"`
		EvolvesFromSpecies struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"evolves_from_species"`
		FlavorTextEntries []struct {
			FlavorText string `json:"flavor_text"`
			Language   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"flavor_text_entries"`
		FormDescriptions []interface{} `json:"form_descriptions"`
		FormsSwitchable  bool          `json:"forms_switchable"`
		GenderRate       int           `json:"gender_rate"`
		Genera           []struct {
			Genus    string `json:"genus"`
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"genera"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		GrowthRate struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"growth_rate"`
		Habitat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"habitat"`
		HasGenderDifferences bool   `json:"has_gender_differences"`
		HatchCounter         int    `json:"hatch_counter"`
		ID                   int    `json:"id"`
		IsBaby               bool   `json:"is_baby"`
		Name                 string `json:"name"`
		Names                []struct {
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
			Name string `json:"name"`
		} `json:"names"`
		Order             int `json:"order"`
		PalParkEncounters []struct {
			Area struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"area"`
			BaseScore int `json:"base_score"`
			Rate      int `json:"rate"`
		} `json:"pal_park_encounters"`
		PokedexNumbers []struct {
			EntryNumber int `json:"entry_number"`
			Pokedex     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"pokedex"`
		} `json:"pokedex_numbers"`
		Shape struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"shape"`
		Varieties []struct {
			IsDefault bool `json:"is_default"`
			Pokemon   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"pokemon"`
		} `json:"varieties"`
	}

	Poke interface {
		Get(url string, v interface{}) error
		Pokemon(number int) (*GetPokemonResponse, error)
		Characteristic(number int) (*GetCharacteristicResponse, error)
		Species(number int) (*GetSpeciesResponse, error)
		SetHTTPClient(httpClient HTTP)
	}

	HTTP interface {
		Do(req *http.Request) (*http.Response, error)
	}
)

const (
	urlPokemon        = "https://pokeapi.co/api/v2/pokemon/%d"
	urlCharacteristic = "https://pokeapi.co/api/v2/characteristic/%d"
	urlSpecies        = "https://pokeapi.co/api/v2/pokemon-species/%d"
)

func New() Poke {
	return &poke{
		httpClient: &http.Client{},
	}
}

func (p *poke) SetHTTPClient(httpClient HTTP) {
	p.httpClient = httpClient
}

func (p *poke) Pokemon(number int) (*GetPokemonResponse, error) {
	url := fmt.Sprintf(urlPokemon, number)
	response := &GetPokemonResponse{}
	return response, p.Get(url, response)
}

func (p *poke) Characteristic(number int) (*GetCharacteristicResponse, error) {
	url := fmt.Sprintf(urlCharacteristic, number)
	response := &GetCharacteristicResponse{}
	return response, p.Get(url, response)
}

func (p *poke) Species(number int) (*GetSpeciesResponse, error) {
	url := fmt.Sprintf(urlSpecies, number)
	response := &GetSpeciesResponse{}
	return response, p.Get(url, response)
}

func (p *poke) Get(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error making GET request to %s. StatusCode: %d", resp.Request.URL, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (r *GetPokemonResponse) AllTypes() []string {
	var types []string
	for _, t := range r.Types {
		types = append(types, t.Type.Name)
	}
	return types
}

func (r *GetPokemonResponse) AllStats() []string {
	var stats []string
	for _, s := range r.Stats {
		stats = append(stats, fmt.Sprintf("%s (%d)", s.Stat.Name, s.BaseStat))
	}
	return stats
}

func (r *GetCharacteristicResponse) AllDescriptions() []string {
	var desc []string
	for _, s := range r.Descriptions {
		if s.Language.Name != "en" {
			continue
		}
		desc = append(desc, s.Description)
	}
	return desc
}

func (r *GetSpeciesResponse) AllEggGroups() []string {
	var eggs []string
	for _, s := range r.EggGroups {
		eggs = append(eggs, s.Name)
	}
	return eggs
}
