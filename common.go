package tmdb

type CastCredit []struct {
	Character   string
	CreditID    string `json:"credit_id"`
	ID          int
	Name        string
	Gender      int `json:"gender"`
	Order       int
	ProfilePath string `json:"profile_path"`
}

type CrewCredit []struct {
	CreditID    string `json:"credit_id"`
	Department  string
	Gender      int `json:"gender"`
	ID          int
	Job         string
	Name        string
	ProfilePath string `json:"profile_path"`
}

type SpokenLanguages []struct {
	Iso639_1    string "json:\"iso_639_1\""
	Name        string
	EnglishName string "json:\"english_name\""
}

type ProductionCountries []struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string
}

type ProductionCompanies []struct {
	ID        int
	Name      string
	LogoPath  string `json:"logo_path"`
	Iso3166_1 string `json:"origin_country"`
}
