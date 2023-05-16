package DB

type Player struct {
	Name         string `json:"name"`
	Nationality  string `json:"nationality"`
	Club         string `json:"club"`
	JerseyNumber int    `json:"jerseyNumber"`
	Stats        Stats  `json:"stats"`
}

type Stats struct {
	NationalGoal        int `json:"nationalGoal"`        //national team goal
	LeagueGoal          int `json:"leagueGoal"`          //league goal
	ChampionsLeagueGoal int `json:"championsLeagueGoal"` //champions league goal
}

type Players []Player

var PlayerDB = Players{
	Player{
		Name:         "Cristiano",
		Nationality:  "Portugal",
		Club:         "Al Nassr",
		JerseyNumber: 7,
		Stats: Stats{
			NationalGoal:        122,
			LeagueGoal:          12,
			ChampionsLeagueGoal: 140}},
	Player{
		Name:         "Messi",
		Nationality:  "Argentina",
		Club:         "PSG",
		JerseyNumber: 30,
		Stats: Stats{
			NationalGoal:        102,
			LeagueGoal:          31,
			ChampionsLeagueGoal: 129}},
	Player{
		Name:         "Neymar",
		Nationality:  "Brazil",
		Club:         "PSG",
		JerseyNumber: 10,
		Stats: Stats{
			NationalGoal:        70,
			LeagueGoal:          35,
			ChampionsLeagueGoal: 75}},
	Player{
		Name:         "Benzema",
		Nationality:  "France",
		Club:         "Real Madrid",
		JerseyNumber: 9,
		Stats: Stats{
			NationalGoal:        38,
			LeagueGoal:          29,
			ChampionsLeagueGoal: 120}},
	Player{
		Name:         "Casemiro",
		Nationality:  "Brazil",
		Club:         "Man U",
		JerseyNumber: 6,
		Stats: Stats{
			NationalGoal:        21,
			LeagueGoal:          9,
			ChampionsLeagueGoal: 12}},
	Player{
		Name:         "Ramos",
		Nationality:  "Spain",
		Club:         "PSG",
		JerseyNumber: 4,
		Stats: Stats{
			NationalGoal:        28,
			LeagueGoal:          3,
			ChampionsLeagueGoal: 34}},
	Player{
		Name:         "Nuer",
		Nationality:  "Germany",
		Club:         "Bayern",
		JerseyNumber: 1,
		Stats: Stats{
			NationalGoal:        0,
			LeagueGoal:          0,
			ChampionsLeagueGoal: 0}},
	Player{
		Name:         "Valverde",
		Nationality:  "Uruguay",
		Club:         "Real Madrid",
		JerseyNumber: 12,
		Stats: Stats{
			NationalGoal:        3,
			LeagueGoal:          11,
			ChampionsLeagueGoal: 8}},
	Player{
		Name:         "Haaland",
		Nationality:  "Norway",
		Club:         "MAn City",
		JerseyNumber: 10,
		Stats: Stats{
			NationalGoal:        5,
			LeagueGoal:          40,
			ChampionsLeagueGoal: 37}},
}
