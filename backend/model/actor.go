package model

type Actor struct { //Сущность, которая будет выводиться в конце
	Id         int    `json:"id"`         // ID актёра
	Familyname string `json:"familyname"` // Фамилия актёра
	Givenname  string `json:"givenname"`  // Имя актёра
	Nation     string `json:"nation"`     // Национальность (гражданство)
	Number     string `json:"number"`     // Число фильмов
	Honorar    string `json:"honorar"`    // Суммарный гонорар
}
