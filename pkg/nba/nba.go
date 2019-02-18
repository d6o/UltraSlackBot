package nba

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	urlScoreboard  = "http://data.nba.net/10s/prod/v1/%s/scoreboard.json"
	urlTeams       = "http://data.nba.net/10s/prod/v2/2018/teams.json"
	urlStanding = "http://data.nba.net/10s/prod/v1/current/standings_all_no_sort_keys.json"
)

type (
	NBA struct {
		httpClient HTTP
	}

	HTTP interface {
		Do(req *http.Request) (*http.Response, error)
	}

	ScoreboardResponse struct {
		NumGames int `json:"numGames"`
		Games    []struct {
			SeasonStageID int    `json:"seasonStageId"`
			SeasonYear    string `json:"seasonYear"`
			GameID        string `json:"gameId"`
			Arena         struct {
				Name       string `json:"name"`
				IsDomestic bool   `json:"isDomestic"`
				City       string `json:"city"`
				StateAbbr  string `json:"stateAbbr"`
				Country    string `json:"country"`
			} `json:"arena"`
			IsGameActivated       bool      `json:"isGameActivated"`
			StatusNum             int       `json:"statusNum"`
			ExtendedStatusNum     int       `json:"extendedStatusNum"`
			StartTimeEastern      string    `json:"startTimeEastern"`
			StartTimeUTC          time.Time `json:"startTimeUTC"`
			StartDateEastern      string    `json:"startDateEastern"`
			Clock                 string    `json:"clock"`
			IsBuzzerBeater        bool      `json:"isBuzzerBeater"`
			IsPreviewArticleAvail bool      `json:"isPreviewArticleAvail"`
			IsRecapArticleAvail   bool      `json:"isRecapArticleAvail"`
			Tickets               struct {
				MobileApp    string `json:"mobileApp"`
				DesktopWeb   string `json:"desktopWeb"`
				MobileWeb    string `json:"mobileWeb"`
				LeagGameInfo string `json:"leagGameInfo"`
				LeagTix      string `json:"leagTix"`
			} `json:"tickets"`
			HasGameBookPdf bool `json:"hasGameBookPdf"`
			IsStartTimeTBD bool `json:"isStartTimeTBD"`
			Nugget         struct {
				Text string `json:"text"`
			} `json:"nugget"`
			Attendance   string `json:"attendance"`
			GameDuration struct {
				Hours   string `json:"hours"`
				Minutes string `json:"minutes"`
			} `json:"gameDuration"`
			Period struct {
				Current       int  `json:"current"`
				Type          int  `json:"type"`
				MaxRegular    int  `json:"maxRegular"`
				IsHalftime    bool `json:"isHalftime"`
				IsEndOfPeriod bool `json:"isEndOfPeriod"`
			} `json:"period"`
			VTeam struct {
				TeamID     string        `json:"teamId"`
				TriCode    string        `json:"triCode"`
				Win        string        `json:"win"`
				Loss       string        `json:"loss"`
				SeriesWin  string        `json:"seriesWin"`
				SeriesLoss string        `json:"seriesLoss"`
				Score      string        `json:"score"`
				Linescore  []interface{} `json:"linescore"`
			} `json:"vTeam"`
			HTeam struct {
				TeamID     string        `json:"teamId"`
				TriCode    string        `json:"triCode"`
				Win        string        `json:"win"`
				Loss       string        `json:"loss"`
				SeriesWin  string        `json:"seriesWin"`
				SeriesLoss string        `json:"seriesLoss"`
				Score      string        `json:"score"`
				Linescore  []interface{} `json:"linescore"`
			} `json:"hTeam"`
			Watch struct {
				Broadcast struct {
					Broadcasters struct {
						National []interface{} `json:"national"`
						Canadian []struct {
							ShortName string `json:"shortName"`
							LongName  string `json:"longName"`
						} `json:"canadian"`
						VTeam []struct {
							ShortName string `json:"shortName"`
							LongName  string `json:"longName"`
						} `json:"vTeam"`
						HTeam []struct {
							ShortName string `json:"shortName"`
							LongName  string `json:"longName"`
						} `json:"hTeam"`
						SpanishHTeam    []interface{} `json:"spanish_hTeam"`
						SpanishVTeam    []interface{} `json:"spanish_vTeam"`
						SpanishNational []interface{} `json:"spanish_national"`
					} `json:"broadcasters"`
					Video struct {
						RegionalBlackoutCodes string `json:"regionalBlackoutCodes"`
						CanPurchase           bool   `json:"canPurchase"`
						IsLeaguePass          bool   `json:"isLeaguePass"`
						IsNationalBlackout    bool   `json:"isNationalBlackout"`
						IsTNTOT               bool   `json:"isTNTOT"`
						IsVR                  bool   `json:"isVR"`
						TntotIsOnAir          bool   `json:"tntotIsOnAir"`
						IsNextVR              bool   `json:"isNextVR"`
						IsNBAOnTNTVR          bool   `json:"isNBAOnTNTVR"`
						IsMagicLeap           bool   `json:"isMagicLeap"`
						IsOculusVenues        bool   `json:"isOculusVenues"`
						Streams               []struct {
							StreamType            string `json:"streamType"`
							IsOnAir               bool   `json:"isOnAir"`
							DoesArchiveExist      bool   `json:"doesArchiveExist"`
							IsArchiveAvailToWatch bool   `json:"isArchiveAvailToWatch"`
							StreamID              string `json:"streamId"`
							Duration              int    `json:"duration"`
						} `json:"streams"`
						DeepLink []struct {
							Broadcaster         string `json:"broadcaster"`
							RegionalMarketCodes string `json:"regionalMarketCodes"`
							IosApp              string `json:"iosApp"`
							AndroidApp          string `json:"androidApp"`
							DesktopWeb          string `json:"desktopWeb"`
							MobileWeb           string `json:"mobileWeb"`
						} `json:"deepLink"`
					} `json:"video"`
					Audio struct {
						National struct {
							Streams []struct {
								Language string `json:"language"`
								IsOnAir  bool   `json:"isOnAir"`
								StreamID string `json:"streamId"`
							} `json:"streams"`
							Broadcasters []interface{} `json:"broadcasters"`
						} `json:"national"`
						VTeam struct {
							Streams []struct {
								Language string `json:"language"`
								IsOnAir  bool   `json:"isOnAir"`
								StreamID string `json:"streamId"`
							} `json:"streams"`
							Broadcasters []struct {
								ShortName string `json:"shortName"`
								LongName  string `json:"longName"`
							} `json:"broadcasters"`
						} `json:"vTeam"`
						HTeam struct {
							Streams []struct {
								Language string `json:"language"`
								IsOnAir  bool   `json:"isOnAir"`
								StreamID string `json:"streamId"`
							} `json:"streams"`
							Broadcasters []struct {
								ShortName string `json:"shortName"`
								LongName  string `json:"longName"`
							} `json:"broadcasters"`
						} `json:"hTeam"`
					} `json:"audio"`
				} `json:"broadcast"`
			} `json:"watch"`
		} `json:"games"`
	}

	TeamsResponse struct {
		League struct {
			Standard []Team `json:"standard"`
		} `json:"league"`
	}

	Team struct {
		IsNBAFranchise bool   `json:"isNBAFranchise"`
		IsAllStar      bool   `json:"isAllStar"`
		City           string `json:"city"`
		AltCityName    string `json:"altCityName"`
		FullName       string `json:"fullName"`
		Tricode        string `json:"tricode"`
		TeamID         string `json:"teamId"`
		Nickname       string `json:"nickname"`
		URLName        string `json:"urlName"`
		ConfName       string `json:"confName"`
		DivName        string `json:"divName"`
	}

	StandingResponse struct {
		League struct {
			Standard struct {
				SeasonYear    int `json:"seasonYear"`
				SeasonStageID int `json:"seasonStageId"`
				Teams         []struct {
					TeamID                 string `json:"teamId"`
					Win                    string `json:"win"`
					Loss                   string `json:"loss"`
					WinPct                 string `json:"winPct"`
					WinPctV2               string `json:"winPctV2"`
					LossPct                string `json:"lossPct"`
					LossPctV2              string `json:"lossPctV2"`
					GamesBehind            string `json:"gamesBehind"`
					DivGamesBehind         string `json:"divGamesBehind"`
					ClinchedPlayoffsCode   string `json:"clinchedPlayoffsCode"`
					ClinchedPlayoffsCodeV2 string `json:"clinchedPlayoffsCodeV2"`
					ConfRank               string `json:"confRank"`
					ConfWin                string `json:"confWin"`
					ConfLoss               string `json:"confLoss"`
					DivWin                 string `json:"divWin"`
					DivLoss                string `json:"divLoss"`
					HomeWin                string `json:"homeWin"`
					HomeLoss               string `json:"homeLoss"`
					AwayWin                string `json:"awayWin"`
					AwayLoss               string `json:"awayLoss"`
					LastTenWin             string `json:"lastTenWin"`
					LastTenLoss            string `json:"lastTenLoss"`
					Streak                 string `json:"streak"`
					DivRank                string `json:"divRank"`
					IsWinStreak            bool   `json:"isWinStreak"`
					TieBreakerPts          string `json:"tieBreakerPts"`
				} `json:"teams"`
			} `json:"standard"`
		} `json:"league"`
	}

)


func New() *NBA {
	return &NBA{
		httpClient: http.DefaultClient,
	}
}

func (n *NBA) Today() (*ScoreboardResponse, error) {
	return n.Games(time.Now())
}

func (n *NBA) Games(date time.Time) (*ScoreboardResponse, error) {
	url := fmt.Sprintf(urlScoreboard, date.Format("20060102"))
	response := &ScoreboardResponse{}
	return response, n.Get(url, response)
}

func (n *NBA) Teams() (*TeamsResponse, error) {
	response := &TeamsResponse{}
	return response, n.Get(urlTeams, response)
}

func (n *NBA) Standing() (*StandingResponse, error) {
	response := &StandingResponse{}
	return response, n.Get(urlStanding, response)
}

func (n *NBA) Get(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error making GET request to %s. StatusCode: %d", resp.Request.URL, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
