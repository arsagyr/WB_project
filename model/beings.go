package beings

type nation struct {
	id   int
	name string
}

type actor struct {
	id         int
	familyname string //
	givenname  string //
	nationid   int    //
	number     int16  //
	honorar    int64  //
}
