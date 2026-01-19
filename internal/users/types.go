package users

import (
	"fmt"
	"github.com/spudmashmedia/gouser/pkg/randomuser"
	"slices"
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
	Name     Name     `json:"name"`
	Contacts Contacts `json:"contacts"`
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

type Contacts []Contact

func (cArray Contacts) FindContactByType(queryContactType string) (string, error) {
	contactId := slices.IndexFunc(cArray, func(item Contact) bool {
		return item.Type == queryContactType
	})

	if contactId == -1 {
		return "", fmt.Errorf("No Items found")
	}
	return cArray[contactId].Value, nil
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
