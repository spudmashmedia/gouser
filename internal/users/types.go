package users

import (
	"github.com/spudmashmedia/gouser/pkg/randomuser"
)

type QueryData struct {
	Count            int
	IsConcurrent     bool
	IsSimLongProcess bool
}

type UsersResponse struct {
	Results []User `json:"results"`
}

type User struct {
	Name     Name      `json:"name"`
	Contacts []Contact `json:"contacts"`
}

type Name struct {
	Title     string `json:"title"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ContactType int

const (
	Email ContactType = iota + 1
	Phone
	Mobile
)

var contactName = map[ContactType]string{
	Email:  "email",
	Phone:  "phone",
	Mobile: "mobile",
}

func (ct ContactType) String() string {
	return contactName[ct]
}

type Contact struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (u *User) RuToUser(src *randomuser.User) *User {
	if src == nil {
		return u
	}

	u.Name = Name{
		Title:     src.Name.Title,
		FirstName: src.Name.First,
		LastName:  src.Name.Last,
	}

	u.Contacts = []Contact{}
	u.Contacts = append(u.Contacts, Contact{
		Type:  ContactType.String(Email),
		Value: src.Email,
	})
	u.Contacts = append(u.Contacts, Contact{
		Type:  ContactType.String(Phone),
		Value: src.Phone,
	})
	u.Contacts = append(u.Contacts, Contact{
		Type:  ContactType.String(Mobile),
		Value: src.Cell,
	})

	return u
}
