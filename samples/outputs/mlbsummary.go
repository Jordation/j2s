package main

type Root struct {
	Game    *Game  `json:"game"`
	Comment string `json:"_comment"`
}

type Assists struct {
	Outfield int `json:"outfield"`
	Total    int `json:"total"`
}

type Away struct {
	Scoring         []*Scoring       `json:"scoring"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher"`
	Statistics      *Statistics      `json:"statistics"`
	Players         []*Players       `json:"players"`
	Id              string           `json:"id"`
	Lineup          []*Lineup        `json:"lineup"`
	Runs            int              `json:"runs"`
	Name            string           `json:"name"`
	Abbr            string           `json:"abbr"`
	Loss            int              `json:"loss"`
	Used            int              `json:"used"`
	Hits            int              `json:"hits"`
	ProbablePitcher *ProbablePitcher `json:"probable_pitcher"`
	Roster          []*Roster        `json:"roster"`
	Market          string           `json:"market"`
	Errors          int              `json:"errors"`
	Remaining       int              `json:"remaining"`
	Win             int              `json:"win"`
}

type Broadcasts struct {
	Type    string `json:"type"`
	Locale  string `json:"locale"`
	Network string `json:"network"`
}

type Bullpen struct {
	Oab               int      `json:"oab"`
	Slg               float64  `json:"slg"`
	TimesThroughOrder int      `json:"times_through_order"`
	Obp               float64  `json:"obp"`
	Bk                int      `json:"bk"`
	Ip2               float64  `json:"ip_2"`
	Outcome           *Outcome `json:"outcome"`
	Steal             *Steal   `json:"steal"`
	Wp                int      `json:"wp"`
	Gofo              float64  `json:"gofo"`
	Outs              *Outs    `json:"outs"`
	Babip             float64  `json:"babip"`
	Oba               float64  `json:"oba"`
	Lob               int      `json:"lob"`
	PitchCount        int      `json:"pitch_count"`
	Onbase            *Onbase  `json:"onbase"`
	Pitches           *Pitches `json:"pitches"`
	InPlay            *InPlay  `json:"in_play"`
	Whip              float64  `json:"whip"`
	Kbb               int      `json:"kbb"`
	Runs              *Runs    `json:"runs"`
	Games             *Games   `json:"games"`
	Gbfb              float64  `json:"gbfb"`
	K9                int      `json:"k9"`
	Ip1               int      `json:"ip_1"`
	Bf                int      `json:"bf"`
	BfIp              float64  `json:"bf_ip"`
	Era               int      `json:"era"`
}

type CurrentConditions struct {
	Condition  string `json:"condition"`
	Humidity   int    `json:"humidity"`
	DewPointF  int    `json:"dew_point_f"`
	CloudCover int    `json:"cloud_cover"`
	ObsTime    string `json:"obs_time"`
	Wind       *Wind  `json:"wind"`
	TempF      int    `json:"temp_f"`
}

type Errors struct {
	Throwing     int `json:"throwing"`
	Fielding     int `json:"fielding"`
	Interference int `json:"interference"`
	Total        int `json:"total"`
}

type Fielding struct {
	Positions []*Positions `json:"positions"`
	Overall   *Overall     `json:"overall"`
}

type Final struct {
	Inning     int    `json:"inning"`
	InningHalf string `json:"inning_half"`
}

type Forecast struct {
	DewPointF  int    `json:"dew_point_f"`
	CloudCover int    `json:"cloud_cover"`
	ObsTime    string `json:"obs_time"`
	Wind       *Wind  `json:"wind"`
	TempF      int    `json:"temp_f"`
	Condition  string `json:"condition"`
	Humidity   int    `json:"humidity"`
}

type Game struct {
	Pitching     *Pitching     `json:"pitching"`
	SeasonYear   int           `json:"season_year"`
	Attendance   int           `json:"attendance"`
	Home         *Home         `json:"home"`
	Weather      *Weather      `json:"weather"`
	Scheduled    string        `json:"scheduled"`
	SeasonId     string        `json:"season_id"`
	HomeTeam     string        `json:"home_team"`
	Reference    string        `json:"reference"`
	Duration     string        `json:"duration"`
	SeasonType   string        `json:"season_type"`
	DoubleHeader bool          `json:"double_header"`
	Id           string        `json:"id"`
	Reviews      *Reviews      `json:"reviews"`
	EntryMode    string        `json:"entry_mode"`
	TimeZones    *TimeZones    `json:"time_zones"`
	Broadcasts   []*Broadcasts `json:"broadcasts"`
	GameNumber   int           `json:"game_number"`
	DayNight     string        `json:"day_night"`
	MoundVisits  *MoundVisits  `json:"mound_visits"`
	Away         *Away         `json:"away"`
	Final        *Final        `json:"final"`
	AwayTeam     string        `json:"away_team"`
	Officials    []*Officials  `json:"officials"`
	Coverage     string        `json:"coverage"`
	Venue        *Venue        `json:"venue"`
	Status       string        `json:"status"`
}

type Games struct {
	Save      int `json:"save"`
	BlownSave int `json:"blown_save"`
	Qstart    int `json:"qstart"`
	Start     int `json:"start"`
	TeamWin   int `json:"team_win"`
	TeamLoss  int `json:"team_loss"`
	Complete  int `json:"complete"`
	Loss      int `json:"loss"`
	Hold      int `json:"hold"`
	Svo       int `json:"svo"`
	Shutout   int `json:"shutout"`
	Play      int `json:"play"`
	Finish    int `json:"finish"`
	Win       int `json:"win"`
}

type Hitting struct {
	Overall *Overall `json:"overall"`
}

type Hold struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PrimaryPosition string `json:"primary_position"`
	Loss            int    `json:"loss"`
	Id              string `json:"id"`
	Win             int    `json:"win"`
	Hold            int    `json:"hold"`
	FullName        string `json:"full_name"`
	PreferredName   string `json:"preferred_name"`
	JerseyNumber    string `json:"jersey_number"`
	Position        string `json:"position"`
	BlownSave       int    `json:"blown_save"`
	Status          string `json:"status"`
	Save            int    `json:"save"`
}

type Home struct {
	Win             int              `json:"win"`
	Market          string           `json:"market"`
	Abbr            string           `json:"abbr"`
	Scoring         []*Scoring       `json:"scoring"`
	Used            int              `json:"used"`
	Roster          []*Roster        `json:"roster"`
	Statistics      *Statistics      `json:"statistics"`
	Hits            int              `json:"hits"`
	Remaining       int              `json:"remaining"`
	Lineup          []*Lineup        `json:"lineup"`
	Runs            int              `json:"runs"`
	Errors          int              `json:"errors"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher"`
	Name            string           `json:"name"`
	Id              string           `json:"id"`
	ProbablePitcher *ProbablePitcher `json:"probable_pitcher"`
	Loss            int              `json:"loss"`
	Players         []*Players       `json:"players"`
}

type InPlay struct {
	Linedrive  int `json:"linedrive"`
	Groundball int `json:"groundball"`
	Popup      int `json:"popup"`
	Flyball    int `json:"flyball"`
}

type Lineup struct {
	Position   int    `json:"position"`
	Sequence   int    `json:"sequence"`
	InningHalf string `json:"inning_half"`
	Id         string `json:"id"`
	Inning     int    `json:"inning"`
	Order      int    `json:"order"`
}

type Location struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Loss struct {
	LastName        string `json:"last_name"`
	Status          string `json:"status"`
	Id              string `json:"id"`
	JerseyNumber    string `json:"jersey_number"`
	Win             int    `json:"win"`
	Save            int    `json:"save"`
	BlownSave       int    `json:"blown_save"`
	PreferredName   string `json:"preferred_name"`
	PrimaryPosition string `json:"primary_position"`
	Hold            int    `json:"hold"`
	Position        string `json:"position"`
	Loss            int    `json:"loss"`
	FullName        string `json:"full_name"`
	FirstName       string `json:"first_name"`
}

type MoundVisits struct {
	Home *Home `json:"home"`
	Away *Away `json:"away"`
}

type Officials struct {
	Experience string `json:"experience"`
	Id         string `json:"id"`
	FullName   string `json:"full_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Assignment string `json:"assignment"`
}

type Onbase struct {
	H     int     `json:"h"`
	Rov   int     `json:"rov"`
	Hr    int     `json:"hr"`
	H9    float64 `json:"h9"`
	Tb    int     `json:"tb"`
	Ci    int     `json:"ci"`
	S     int     `json:"s"`
	Roe   int     `json:"roe"`
	T     int     `json:"t"`
	Bb    int     `json:"bb"`
	Hr9   int     `json:"hr9"`
	Ibb   int     `json:"ibb"`
	Fc    int     `json:"fc"`
	Cycle int     `json:"cycle"`
	D     int     `json:"d"`
	Hbp   int     `json:"hbp"`
}

type Outcome struct {
	Ball     int `json:"ball"`
	Iball    int `json:"iball"`
	Dirtball int `json:"dirtball"`
	Foul     int `json:"foul"`
	Klook    int `json:"klook"`
	Kswing   int `json:"kswing"`
	Ktotal   int `json:"ktotal"`
}

type Outs struct {
	Go     int `json:"go"`
	Gidp   int `json:"gidp"`
	Sacfly int `json:"sacfly"`
	Klook  int `json:"klook"`
	Kswing int `json:"kswing"`
	Fo     int `json:"fo"`
	Lo     int `json:"lo"`
	Ktotal int `json:"ktotal"`
	Sachit int `json:"sachit"`
	Fidp   int `json:"fidp"`
	Lidp   int `json:"lidp"`
	Po     int `json:"po"`
}

type Overall struct {
	A                 int      `json:"a"`
	Ab                int      `json:"ab"`
	Abhr              int      `json:"abhr"`
	Abk               float64  `json:"abk"`
	AbRisp            int      `json:"ab_risp"`
	Ap                int      `json:"ap"`
	Assists           *Assists `json:"assists"`
	Avg               string   `json:"avg"`
	Babip             int      `json:"babip"`
	Bbk               float64  `json:"bbk"`
	Bbpa              float64  `json:"bbpa"`
	Bf                int      `json:"bf"`
	BfIp              int      `json:"bf_ip"`
	BfStart           int      `json:"bf_start"`
	Bip               int      `json:"bip"`
	Bk                int      `json:"bk"`
	CWp               int      `json:"c_wp"`
	Dp                int      `json:"dp"`
	Era               int      `json:"era"`
	Error             int      `json:"error"`
	Errors            *Errors  `json:"errors"`
	Flyball           int      `json:"flyball"`
	Fpct              int      `json:"fpct"`
	Games             *Games   `json:"games"`
	Gbfb              int      `json:"gbfb"`
	Gofo              float64  `json:"gofo"`
	Groundball        int      `json:"groundball"`
	HitRisp           int      `json:"hit_risp"`
	Inn1              int      `json:"inn_1"`
	Inn2              int      `json:"inn_2"`
	InPlay            *InPlay  `json:"in_play"`
	Ip1               int      `json:"ip_1"`
	Ip2               int      `json:"ip_2"`
	Iso               int      `json:"iso"`
	K9                int      `json:"k9"`
	Kbb               int      `json:"kbb"`
	Linedrive         int      `json:"linedrive"`
	Lob               int      `json:"lob"`
	LobRisp2out       int      `json:"lob_risp_2out"`
	Oab               int      `json:"oab"`
	Oba               int      `json:"oba"`
	Obp               float64  `json:"obp"`
	Onbase            *Onbase  `json:"onbase"`
	Ops               float64  `json:"ops"`
	Outcome           *Outcome `json:"outcome"`
	Outs              *Outs    `json:"outs"`
	Pb                int      `json:"pb"`
	PitchCount        int      `json:"pitch_count"`
	Pitches           *Pitches `json:"pitches"`
	Po                int      `json:"po"`
	Popup             int      `json:"popup"`
	Rbi               int      `json:"rbi"`
	Rbi2out           int      `json:"rbi_2out"`
	Rf                int      `json:"rf"`
	Runs              *Runs    `json:"runs"`
	Seca              float64  `json:"seca"`
	Slg               int      `json:"slg"`
	Steal             *Steal   `json:"steal"`
	Tc                int      `json:"tc"`
	TeamLob           int      `json:"team_lob"`
	TimesThroughOrder int      `json:"times_through_order"`
	Tp                int      `json:"tp"`
	Whip              int      `json:"whip"`
	Wp                int      `json:"wp"`
	Xbh               int      `json:"xbh"`
}

type Pitches struct {
	Ktotal   int     `json:"ktotal"`
	PerIp    int     `json:"per_ip"`
	PerBf    float64 `json:"per_bf"`
	PerStart int     `json:"per_start"`
	Count    int     `json:"count"`
	Btotal   int     `json:"btotal"`
}

type Pitching struct {
	Starters *Starters `json:"starters"`
	Bullpen  *Bullpen  `json:"bullpen"`
	Hold     []*Hold   `json:"hold"`
	Win      *Win      `json:"win"`
	Loss     *Loss     `json:"loss"`
	Save     *Save     `json:"save"`
	Overall  *Overall  `json:"overall"`
}

type Players struct {
	Position        string      `json:"position"`
	PrimaryPosition string      `json:"primary_position"`
	Statistics      *Statistics `json:"statistics"`
	PreferredName   string      `json:"preferred_name"`
	FullName        string      `json:"full_name"`
	LastName        string      `json:"last_name"`
	JerseyNumber    string      `json:"jersey_number"`
	Id              string      `json:"id"`
	FirstName       string      `json:"first_name"`
	Status          string      `json:"status"`
}

type Positions struct {
	Assists  *Assists `json:"assists"`
	CWp      int      `json:"c_wp"`
	Po       int      `json:"po"`
	Pb       int      `json:"pb"`
	Error    int      `json:"error"`
	Inn1     int      `json:"inn_1"`
	Rf       int      `json:"rf"`
	Inn2     int      `json:"inn_2"`
	Tp       int      `json:"tp"`
	Tc       int      `json:"tc"`
	Position string   `json:"position"`
	Games    *Games   `json:"games"`
	Steal    *Steal   `json:"steal"`
	Errors   *Errors  `json:"errors"`
	Fpct     int      `json:"fpct"`
	A        int      `json:"a"`
	Dp       int      `json:"dp"`
}

type ProbablePitcher struct {
	FirstName     string  `json:"first_name"`
	Era           float64 `json:"era"`
	LastName      string  `json:"last_name"`
	JerseyNumber  string  `json:"jersey_number"`
	Id            string  `json:"id"`
	FullName      string  `json:"full_name"`
	Loss          int     `json:"loss"`
	Win           int     `json:"win"`
	PreferredName string  `json:"preferred_name"`
}

type Reviews struct {
	Home *Home `json:"home"`
	Away *Away `json:"away"`
}

type Roster struct {
	Position        string `json:"position"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	FullName        string `json:"full_name"`
	Status          string `json:"status"`
	PrimaryPosition string `json:"primary_position"`
	PreferredName   string `json:"preferred_name"`
	JerseyNumber    string `json:"jersey_number"`
	Id              string `json:"id"`
}

type Runs struct {
	Bqra     int `json:"bqra"`
	Unearned int `json:"unearned"`
	Earned   int `json:"earned"`
	Total    int `json:"total"`
	Ir       int `json:"ir"`
	Ira      int `json:"ira"`
	Bqr      int `json:"bqr"`
}

type Save struct {
	PrimaryPosition string `json:"primary_position"`
	Win             int    `json:"win"`
	Hold            int    `json:"hold"`
	PreferredName   string `json:"preferred_name"`
	Position        string `json:"position"`
	Loss            int    `json:"loss"`
	BlownSave       int    `json:"blown_save"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Status          string `json:"status"`
	Id              string `json:"id"`
	Save            int    `json:"save"`
	JerseyNumber    string `json:"jersey_number"`
	FullName        string `json:"full_name"`
}

type Scoring struct {
	Hits     int    `json:"hits"`
	Errors   int    `json:"errors"`
	Type     string `json:"type"`
	Number   int    `json:"number"`
	Sequence int    `json:"sequence"`
	Runs     int    `json:"runs"`
}

type Starters struct {
	TimesThroughOrder int      `json:"times_through_order"`
	Whip              float64  `json:"whip"`
	Pitches           *Pitches `json:"pitches"`
	Wp                int      `json:"wp"`
	InPlay            *InPlay  `json:"in_play"`
	PitchCount        int      `json:"pitch_count"`
	Bk                int      `json:"bk"`
	Kbb               int      `json:"kbb"`
	Obp               float64  `json:"obp"`
	Slg               float64  `json:"slg"`
	Onbase            *Onbase  `json:"onbase"`
	Era               float64  `json:"era"`
	Ip1               int      `json:"ip_1"`
	BfStart           int      `json:"bf_start"`
	Babip             float64  `json:"babip"`
	Outs              *Outs    `json:"outs"`
	Gofo              int      `json:"gofo"`
	Oba               float64  `json:"oba"`
	Oab               int      `json:"oab"`
	Bf                int      `json:"bf"`
	Ip2               float64  `json:"ip_2"`
	Gbfb              float64  `json:"gbfb"`
	Runs              *Runs    `json:"runs"`
	Outcome           *Outcome `json:"outcome"`
	BfIp              float64  `json:"bf_ip"`
	Steal             *Steal   `json:"steal"`
	Games             *Games   `json:"games"`
	Lob               int      `json:"lob"`
	K9                float64  `json:"k9"`
}

type StartingPitcher struct {
	Era           float64 `json:"era"`
	FirstName     string  `json:"first_name"`
	Id            string  `json:"id"`
	PreferredName string  `json:"preferred_name"`
	FullName      string  `json:"full_name"`
	Win           int     `json:"win"`
	LastName      string  `json:"last_name"`
	JerseyNumber  string  `json:"jersey_number"`
	Loss          int     `json:"loss"`
}

type Statistics struct {
	Fielding *Fielding `json:"fielding"`
	Hitting  *Hitting  `json:"hitting"`
	Pitching *Pitching `json:"pitching"`
}

type Steal struct {
	Caught  int     `json:"caught"`
	Stolen  int     `json:"stolen"`
	Pct     float64 `json:"pct"`
	Pickoff int     `json:"pickoff"`
}

type TimeZones struct {
	Venue string `json:"venue"`
	Home  string `json:"home"`
	Away  string `json:"away"`
}

type Venue struct {
	Surface          string    `json:"surface"`
	Zip              string    `json:"zip"`
	Country          string    `json:"country"`
	State            string    `json:"state"`
	TimeZone         string    `json:"time_zone"`
	Location         *Location `json:"location"`
	Capacity         int       `json:"capacity"`
	Address          string    `json:"address"`
	City             string    `json:"city"`
	Id               string    `json:"id"`
	FieldOrientation string    `json:"field_orientation"`
	StadiumType      string    `json:"stadium_type"`
	Name             string    `json:"name"`
	Market           string    `json:"market"`
}

type Weather struct {
	Forecast          *Forecast          `json:"forecast"`
	CurrentConditions *CurrentConditions `json:"current_conditions"`
}

type Win struct {
	Position        string `json:"position"`
	FullName        string `json:"full_name"`
	Id              string `json:"id"`
	Win             int    `json:"win"`
	Hold            int    `json:"hold"`
	LastName        string `json:"last_name"`
	PrimaryPosition string `json:"primary_position"`
	Loss            int    `json:"loss"`
	PreferredName   string `json:"preferred_name"`
	Save            int    `json:"save"`
	BlownSave       int    `json:"blown_save"`
	FirstName       string `json:"first_name"`
	JerseyNumber    string `json:"jersey_number"`
	Status          string `json:"status"`
}

type Wind struct {
	Direction string `json:"direction"`
	SpeedMph  int    `json:"speed_mph"`
}
