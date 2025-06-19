package model

type Nation struct {
	Id   int
	Name string
}

type Actor struct {
	Id         int
	Familyname string // Фамилия актёра
	Givenname  string // Имя актёра
	Nation     string // Национальность (гражданство)
	Number     int16  // Число фильмов
	Honorar    int64  // Суммарный гонорар
}

type Name struct {
	Familyname string // Фамилия актёра
	Givenname  string // Имя актёра
}
