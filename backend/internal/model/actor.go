package model

type Actor struct { //Сущность, которая будет выводиться в конце
	Id         int
	Familyname string // Фамилия актёра
	Givenname  string // Имя актёра
	Nation     string // Национальность (гражданство)
	Number     string // Число фильмов
	Honorar    string // Суммарный гонорар
}
