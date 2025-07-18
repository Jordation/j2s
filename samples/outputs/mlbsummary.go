package main

type Ship struct {
	Combined []*Combined `json:"combined"`
}

type Assists struct {
	Outfield int `json:"outfield"`
	Total int `json:"total"`
}

type Away struct {
	Id string `json:"id"`
	ProbablePitcher *ProbablePitcher `json:"probable_pitcher"`
	Hits int `json:"hits"`
	Loss int `json:"loss"`
	Market string `json:"market"`
	Roster []*Roster `json:"roster"`
	Scoring []*Scoring `json:"scoring"`
	Used int `json:"used"`
	Runs int `json:"runs"`
	Abbr string `json:"abbr"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher"`
	Remaining int `json:"remaining"`
	Players []*Players `json:"players"`
	Errors int `json:"errors"`
	Name string `json:"name"`
	Win int `json:"win"`
	Statistics *Statistics `json:"statistics"`
	Lineup []*Lineup `json:"lineup"`
}

type Broadcasts struct {
	Channel string `json:"channel"`
	Network string `json:"network"`
	Type string `json:"type"`
	Locale string `json:"locale"`
}

type Bullpen struct {
	Kbb int `json:"kbb"`
	Oab int `json:"oab"`
	InPlay *InPlay `json:"in_play"`
	Lob int `json:"lob"`
	Era int `json:"era"`
	Outcome *Outcome `json:"outcome"`
	Outs *Outs `json:"outs"`
	K9 int `json:"k9"`
	Onbase *Onbase `json:"onbase"`
	Whip float64 `json:"whip"`
	TimesThroughOrder int `json:"times_through_order"`
	Games *Games `json:"games"`
	Oba float64 `json:"oba"`
	Gofo int `json:"gofo"`
	Gbfb int `json:"gbfb"`
	Bk int `json:"bk"`
	BfIp float64 `json:"bf_ip"`
	Slg float64 `json:"slg"`
	Obp float64 `json:"obp"`
	Ip2 float64 `json:"ip_2"`
	Bf int `json:"bf"`
	Pitches *Pitches `json:"pitches"`
	Wp int `json:"wp"`
	PitchCount int `json:"pitch_count"`
	Steal *Steal `json:"steal"`
	Ip1 int `json:"ip_1"`
	Runs *Runs `json:"runs"`
	Babip float64 `json:"babip"`
}

type Combined struct {
	Game *Game `json:"game"`
	Comment string `json:"_comment"`
}

type CurrentConditions struct {
	TempF int `json:"temp_f"`
	Condition string `json:"condition"`
	Humidity int `json:"humidity"`
	DewPointF int `json:"dew_point_f"`
	CloudCover int `json:"cloud_cover"`
	ObsTime string `json:"obs_time"`
	Wind *Wind `json:"wind"`
}

type Errors struct {
	Throwing int `json:"throwing"`
	Fielding int `json:"fielding"`
	Interference int `json:"interference"`
	Total int `json:"total"`
}

type Fielding struct {
	Overall *Overall `json:"overall"`
	Positions []*Positions `json:"positions"`
}

type Final struct {
	Inning int `json:"inning"`
	InningHalf string `json:"inning_half"`
}

type Forecast struct {
	Condition string `json:"condition"`
	Humidity int `json:"humidity"`
	DewPointF int `json:"dew_point_f"`
	CloudCover int `json:"cloud_cover"`
	ObsTime string `json:"obs_time"`
	Wind *Wind `json:"wind"`
	TempF int `json:"temp_f"`
}

type Game struct {
	GameNumber int `json:"game_number"`
	DayNight string `json:"day_night"`
	Attendance int `json:"attendance"`
	Weather *Weather `json:"weather"`
	Scheduled string `json:"scheduled"`
	MoundVisits *MoundVisits `json:"mound_visits"`
	DoubleHeader bool `json:"double_header"`
	Broadcasts []*Broadcasts `json:"broadcasts"`
	Home *Home `json:"home"`
	Status string `json:"status"`
	Coverage string `json:"coverage"`
	Venue *Venue `json:"venue"`
	Officials []*Officials `json:"officials"`
	Reference string `json:"reference"`
	Final *Final `json:"final"`
	Away *Away `json:"away"`
	SeasonType string `json:"season_type"`
	AwayTeam string `json:"away_team"`
	Reviews *Reviews `json:"reviews"`
	Id string `json:"id"`
	HomeTeam string `json:"home_team"`
	Duration string `json:"duration"`
	EntryMode string `json:"entry_mode"`
	SeasonId string `json:"season_id"`
	Pitching *Pitching `json:"pitching"`
	SeasonYear int `json:"season_year"`
	TimeZones *TimeZones `json:"time_zones"`
}

type Games struct {
	Save int `json:"save"`
	BlownSave int `json:"blown_save"`
	Play int `json:"play"`
	Finish int `json:"finish"`
	TeamLoss int `json:"team_loss"`
	Svo int `json:"svo"`
	Win int `json:"win"`
	Loss int `json:"loss"`
	Qstart int `json:"qstart"`
	Shutout int `json:"shutout"`
	Complete int `json:"complete"`
	Hold int `json:"hold"`
	TeamWin int `json:"team_win"`
	Start int `json:"start"`
}

type Hitting struct {
	Overall *Overall `json:"overall"`
}

type Hold struct {
	Loss int `json:"loss"`
	Save int `json:"save"`
	BlownSave int `json:"blown_save"`
	LastName string `json:"last_name"`
	JerseyNumber string `json:"jersey_number"`
	Position string `json:"position"`
	PrimaryPosition string `json:"primary_position"`
	Win int `json:"win"`
	Hold int `json:"hold"`
	FullName string `json:"full_name"`
	PreferredName string `json:"preferred_name"`
	FirstName string `json:"first_name"`
	Status string `json:"status"`
	Id string `json:"id"`
}

type Home struct {
	Statistics *Statistics `json:"statistics"`
	Abbr string `json:"abbr"`
	Errors int `json:"errors"`
	Lineup []*Lineup `json:"lineup"`
	Runs int `json:"runs"`
	Loss int `json:"loss"`
	Win int `json:"win"`
	Hits int `json:"hits"`
	StartingPitcher *StartingPitcher `json:"starting_pitcher"`
	Id string `json:"id"`
	Used int `json:"used"`
	Scoring []*Scoring `json:"scoring"`
	Players []*Players `json:"players"`
	Remaining int `json:"remaining"`
	Market string `json:"market"`
	ProbablePitcher *ProbablePitcher `json:"probable_pitcher"`
	Roster []*Roster `json:"roster"`
	Name string `json:"name"`
}

type InPlay struct {
	Linedrive int `json:"linedrive"`
	Groundball int `json:"groundball"`
	Popup int `json:"popup"`
	Flyball int `json:"flyball"`
}

type Lineup struct {
	Id string `json:"id"`
	Inning int `json:"inning"`
	Order int `json:"order"`
	Position int `json:"position"`
	Sequence int `json:"sequence"`
	InningHalf string `json:"inning_half"`
}

type Location struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Loss struct {
	LastName string `json:"last_name"`
	Status string `json:"status"`
	Id string `json:"id"`
	Win int `json:"win"`
	Hold int `json:"hold"`
	PreferredName string `json:"preferred_name"`
	FirstName string `json:"first_name"`
	Position string `json:"position"`
	JerseyNumber string `json:"jersey_number"`
	Save int `json:"save"`
	BlownSave int `json:"blown_save"`
	PrimaryPosition string `json:"primary_position"`
	Loss int `json:"loss"`
	FullName string `json:"full_name"`
}

type MoundVisits struct {
	Home *Home `json:"home"`
	Away *Away `json:"away"`
}

type Officials struct {
	FullName string `json:"full_name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Assignment string `json:"assignment"`
	Experience string `json:"experience"`
	Id string `json:"id"`
}

type Onbase struct {
	D int `json:"d"`
	Ibb int `json:"ibb"`
	Roe int `json:"roe"`
	Tb int `json:"tb"`
	Fc int `json:"fc"`
	Hr9 int `json:"hr9"`
	Hr int `json:"hr"`
	Bb int `json:"bb"`
	Rov int `json:"rov"`
	Hbp int `json:"hbp"`
	H int `json:"h"`
	T int `json:"t"`
	Cycle int `json:"cycle"`
	Ci int `json:"ci"`
	H9 float64 `json:"h9"`
	S int `json:"s"`
}

type Outcome struct {
	Kswing int `json:"kswing"`
	Ktotal int `json:"ktotal"`
	Ball int `json:"ball"`
	Iball int `json:"iball"`
	Dirtball int `json:"dirtball"`
	Foul int `json:"foul"`
	Klook int `json:"klook"`
}

type Outs struct {
	Go int `json:"go"`
	Kswing int `json:"kswing"`
	Fo int `json:"fo"`
	Lo int `json:"lo"`
	Lidp int `json:"lidp"`
	Ktotal int `json:"ktotal"`
	Gidp int `json:"gidp"`
	Sacfly int `json:"sacfly"`
	Po int `json:"po"`
	Fidp int `json:"fidp"`
	Klook int `json:"klook"`
	Sachit int `json:"sachit"`
}

type Overall struct {
	Seca int `json:"seca"`
	Pitches *Pitches `json:"pitches"`
	Error int `json:"error"`
	PitchCount int `json:"pitch_count"`
	Lob int `json:"lob"`
	Inn1 int `json:"inn_1"`
	BfStart int `json:"bf_start"`
	Oab int `json:"oab"`
	CWp int `json:"c_wp"`
	Bbpa int `json:"bbpa"`
	Bk int `json:"bk"`
	Obp int `json:"obp"`
	Kbb int `json:"kbb"`
	Outcome *Outcome `json:"outcome"`
	Ip2 float64 `json:"ip_2"`
	InPlay *InPlay `json:"in_play"`
	Xbh int `json:"xbh"`
	Ip1 int `json:"ip_1"`
	Games *Games `json:"games"`
	Inn2 int `json:"inn_2"`
	Tc int `json:"tc"`
	Onbase *Onbase `json:"onbase"`
	Tp int `json:"tp"`
	Avg string `json:"avg"`
	Popup int `json:"popup"`
	Ops int `json:"ops"`
	Abhr int `json:"abhr"`
	Linedrive int `json:"linedrive"`
	Assists *Assists `json:"assists"`
	TeamLob int `json:"team_lob"`
	Whip int `json:"whip"`
	Pb int `json:"pb"`
	AbRisp int `json:"ab_risp"`
	HitRisp int `json:"hit_risp"`
	Era int `json:"era"`
	Gbfb int `json:"gbfb"`
	Slg int `json:"slg"`
	Oba int `json:"oba"`
	K9 int `json:"k9"`
	Steal *Steal `json:"steal"`
	Errors *Errors `json:"errors"`
	Bf int `json:"bf"`
	Babip int `json:"babip"`
	Runs *Runs `json:"runs"`
	Iso int `json:"iso"`
	Po int `json:"po"`
	Wp int `json:"wp"`
	BfIp float64 `json:"bf_ip"`
	LobRisp2out int `json:"lob_risp_2out"`
	Rf int `json:"rf"`
	Bbk int `json:"bbk"`
	Gofo int `json:"gofo"`
	Fpct int `json:"fpct"`
	Bip int `json:"bip"`
	Flyball int `json:"flyball"`
	Rbi int `json:"rbi"`
	Ap int `json:"ap"`
	A int `json:"a"`
	Abk float64 `json:"abk"`
	Rbi2out int `json:"rbi_2out"`
	Outs *Outs `json:"outs"`
	Ab int `json:"ab"`
	Dp int `json:"dp"`
	Groundball int `json:"groundball"`
	TimesThroughOrder int `json:"times_through_order"`
}

type Pitches struct {
	PerBf float64 `json:"per_bf"`
	PerStart int `json:"per_start"`
	Count int `json:"count"`
	Btotal int `json:"btotal"`
	Ktotal int `json:"ktotal"`
	PerIp float64 `json:"per_ip"`
}

type Pitching struct {
	Starters *Starters `json:"starters"`
	Bullpen *Bullpen `json:"bullpen"`
	Win *Win `json:"win"`
	Loss *Loss `json:"loss"`
	Save *Save `json:"save"`
	Hold []*Hold `json:"hold"`
	Overall *Overall `json:"overall"`
}

type Players struct {
	FirstName string `json:"first_name"`
	JerseyNumber string `json:"jersey_number"`
	PrimaryPosition string `json:"primary_position"`
	Status string `json:"status"`
	Position string `json:"position"`
	Statistics *Statistics `json:"statistics"`
	LastName string `json:"last_name"`
	Suffix string `json:"suffix"`
	PreferredName string `json:"preferred_name"`
	Id string `json:"id"`
	FullName string `json:"full_name"`
}

type Positions struct {
	Errors *Errors `json:"errors"`
	Games *Games `json:"games"`
	Steal *Steal `json:"steal"`
	Dp int `json:"dp"`
	Position string `json:"position"`
	Po int `json:"po"`
	Inn1 int `json:"inn_1"`
	Assists *Assists `json:"assists"`
	Tc int `json:"tc"`
	Inn2 int `json:"inn_2"`
	Error int `json:"error"`
	Pb int `json:"pb"`
	Rf int `json:"rf"`
	Fpct int `json:"fpct"`
	A int `json:"a"`
	CWp int `json:"c_wp"`
	Tp int `json:"tp"`
}

type ProbablePitcher struct {
	Era float64 `json:"era"`
	LastName string `json:"last_name"`
	FullName string `json:"full_name"`
	Win int `json:"win"`
	Loss int `json:"loss"`
	PreferredName string `json:"preferred_name"`
	FirstName string `json:"first_name"`
	JerseyNumber string `json:"jersey_number"`
	Id string `json:"id"`
}

type Reviews struct {
	Home *Home `json:"home"`
	Away *Away `json:"away"`
}

type Roster struct {
	FullName string `json:"full_name"`
	PrimaryPosition string `json:"primary_position"`
	Id string `json:"id"`
	Suffix string `json:"suffix"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Status string `json:"status"`
	PreferredName string `json:"preferred_name"`
	Position string `json:"position"`
	JerseyNumber string `json:"jersey_number"`
}

type Runs struct {
	Ira int `json:"ira"`
	Bqr int `json:"bqr"`
	Bqra int `json:"bqra"`
	Total int `json:"total"`
	Unearned int `json:"unearned"`
	Earned int `json:"earned"`
	Ir int `json:"ir"`
}

type Save struct {
	PreferredName string `json:"preferred_name"`
	Id string `json:"id"`
	Loss int `json:"loss"`
	JerseyNumber string `json:"jersey_number"`
	Status string `json:"status"`
	PrimaryPosition string `json:"primary_position"`
	Win int `json:"win"`
	Hold int `json:"hold"`
	FirstName string `json:"first_name"`
	Position string `json:"position"`
	FullName string `json:"full_name"`
	LastName string `json:"last_name"`
	Save int `json:"save"`
	BlownSave int `json:"blown_save"`
}

type Scoring struct {
	Errors int `json:"errors"`
	Type string `json:"type"`
	Number int `json:"number"`
	Sequence int `json:"sequence"`
	Runs int `json:"runs"`
	Hits int `json:"hits"`
}

type Starters struct {
	PitchCount int `json:"pitch_count"`
	Steal *Steal `json:"steal"`
	Bk int `json:"bk"`
	K9 float64 `json:"k9"`
	TimesThroughOrder int `json:"times_through_order"`
	Slg float64 `json:"slg"`
	Ip2 int `json:"ip_2"`
	Oba float64 `json:"oba"`
	Outcome *Outcome `json:"outcome"`
	Runs *Runs `json:"runs"`
	Pitches *Pitches `json:"pitches"`
	Gbfb int `json:"gbfb"`
	Oab int `json:"oab"`
	BfStart int `json:"bf_start"`
	Era float64 `json:"era"`
	Lob int `json:"lob"`
	BfIp float64 `json:"bf_ip"`
	Wp int `json:"wp"`
	Games *Games `json:"games"`
	Whip float64 `json:"whip"`
	Ip1 int `json:"ip_1"`
	Gofo float64 `json:"gofo"`
	Babip float64 `json:"babip"`
	Kbb float64 `json:"kbb"`
	InPlay *InPlay `json:"in_play"`
	Onbase *Onbase `json:"onbase"`
	Bf int `json:"bf"`
	Obp float64 `json:"obp"`
	Outs *Outs `json:"outs"`
}

type StartingPitcher struct {
	LastName string `json:"last_name"`
	FirstName string `json:"first_name"`
	JerseyNumber string `json:"jersey_number"`
	Id string `json:"id"`
	Loss int `json:"loss"`
	PreferredName string `json:"preferred_name"`
	FullName string `json:"full_name"`
	Win int `json:"win"`
	Era float64 `json:"era"`
}

type Statistics struct {
	Fielding *Fielding `json:"fielding"`
	Hitting *Hitting `json:"hitting"`
	Pitching *Pitching `json:"pitching"`
}

type Steal struct {
	Caught int `json:"caught"`
	Stolen int `json:"stolen"`
	Pickoff int `json:"pickoff"`
	Pct int `json:"pct"`
}

type TimeZones struct {
	Venue string `json:"venue"`
	Home string `json:"home"`
	Away string `json:"away"`
}

type Venue struct {
	Id string `json:"id"`
	Surface string `json:"surface"`
	Country string `json:"country"`
	Location *Location `json:"location"`
	Address string `json:"address"`
	FieldOrientation string `json:"field_orientation"`
	StadiumType string `json:"stadium_type"`
	Name string `json:"name"`
	City string `json:"city"`
	Zip string `json:"zip"`
	TimeZone string `json:"time_zone"`
	Market string `json:"market"`
	Capacity int `json:"capacity"`
	State string `json:"state"`
}

type Weather struct {
	Forecast *Forecast `json:"forecast"`
	CurrentConditions *CurrentConditions `json:"current_conditions"`
}

type Win struct {
	BlownSave int `json:"blown_save"`
	LastName string `json:"last_name"`
	Status string `json:"status"`
	Win int `json:"win"`
	Loss int `json:"loss"`
	Hold int `json:"hold"`
	PreferredName string `json:"preferred_name"`
	JerseyNumber string `json:"jersey_number"`
	Save int `json:"save"`
	FullName string `json:"full_name"`
	FirstName string `json:"first_name"`
	Position string `json:"position"`
	PrimaryPosition string `json:"primary_position"`
	Id string `json:"id"`
}

type Wind struct {
	SpeedMph int `json:"speed_mph"`
	Direction string `json:"direction"`
}


