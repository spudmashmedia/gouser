package randomuser

type RandomUserResponse struct {
	Results []User `json:"results"`
	Info    Info   `json:"info"`
}

type User struct {
	Gender       string     `json:"gender"`
	Name         Name       `json:"name"`
	Location     Location   `json:"location"`
	Email        string     `json:"email"`
	Login        Login      `json:"login"`
	Dob          TimeAlive  `json:"dob"`
	Registration TimeAlive  `json:"registration"`
	Phone        string     `json:"phone"`
	Cell         string     `json:"cell"`
	Id           Identifier `json:"id"`
	Picture      ImageSet   `json:"picture"`
	Nat          string     `json:"nat"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    any         `json:"postcode"`
	Coordinates GpsLocation `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type GpsLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Login struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type TimeAlive struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type Identifier struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ImageSet struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}
