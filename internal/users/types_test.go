package users

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/spudmashmedia/gouser/pkg/randomuser"
)

func BuildRuUser() randomuser.User {
	return randomuser.User{
		Gender: "male",
		Name: randomuser.Name{
			Title: "Mr",
			First: "Leeroy",
			Last:  "Jenkins",
		},
		Location: randomuser.Location{
			Street: randomuser.Street{
				Number: 1441,
				Name:   "",
			},
			City:     "Louisville",
			State:    "KY",
			Country:  "United States",
			Postcode: "40213",
			Coordinates: randomuser.GpsLocation{
				Latitude:  "38.2045",
				Longitude: "-85.6983",
			},
			Timezone: randomuser.Timezone{},
		},
		Email: "leeroyjenkins@wow.com",
		Login: randomuser.Login{
			Uuid:     "B0B44364-18E6-41FA-8DBC-2D0D73ED6744",
			Username: "leeroy",
			Password: "CD1EDA8E-B0F2-45D7-978E-810FC43ECBDC",
			Salt:     "C7612542-B5B7-47F0-8119-D27886D7DA1B",
			Md5:      "b15835f133ff2e27c7cb28117bfae8f4",
			Sha1:     "2ace62c1befa19e3ea37dd52be9f6d508c5163e6",
			Sha256:   "8a331fdde7032f33a71e1b2e257d80166e348e00fcb17914f48bdb57a1c63007",
		},
		Dob: randomuser.TimeAlive{
			Date: "2005-05-11T00:00:00.000z",
			Age:  20,
		},
		Registration: randomuser.TimeAlive{
			Date: "2005-05-11T00:00:00.000z",
			Age:  20,
		},
		Phone: "0414555555",
		Cell:  "0414555555",
		Id: randomuser.Identifier{
			Name:  "leeroy",
			Value: "9001",
		},
		Picture: randomuser.ImageSet{
			Large:     "https://warcraft.wiki.gg/wiki/File:Leeroy_Jenkins_TCG.jpg",
			Medium:    "https://warcraft.wiki.gg/wiki/File:Leeroy_Jenkins_TCG.jpg",
			Thumbnail: "https://warcraft.wiki.gg/images/thumb/Leeroy_Jenkins_TCG.jpg/200px-Leeroy_Jenkins_TCG.jpg?d9452d",
		},
		Nat: "American",
	}
}

func TestRuToUser(t *testing.T) {
	// Arrange
	var testRuUser randomuser.User
	testRuUser = BuildRuUser()

	var actualUser User

	// Act
	actualUser.RuToUser(&testRuUser)

	// Assert
	require.NotEmpty(t, actualUser, "actualUser should not be empty")

	assert.Equal(t, actualUser.Name.Title, testRuUser.Name.Title, "actualUser.Name.Title should be eqaul")
	assert.Equal(t, actualUser.Name.FirstName, testRuUser.Name.First, "actualUser.Name.FirstName should be equal")
	assert.Equal(t, actualUser.Name.LastName, testRuUser.Name.Last, "actualUser.Name.LastName should be equal")

	assert.Equal(t, len(actualUser.Contacts), 3, "expextedUser.Contacts[] should have 3 records")

	// Find Email
	actualEmail, err := actualUser.Contacts.FindContactByType("email")
	require.NoError(t, err, "FindContactByType email should not throw error")
	assert.Equal(t, actualEmail, testRuUser.Email, "Email Should be equal")

	// Find Phone
	actualPhone, err := actualUser.Contacts.FindContactByType("phone")
	require.NoError(t, err, "FindContactByType phone should not throw error")
	assert.Equal(t, actualPhone, testRuUser.Phone, "Phone Should be equal")

	// Find Mobile - Cell to Mobile conversion
	actualMobile, err := actualUser.Contacts.FindContactByType("mobile")
	require.NoError(t, err, "FindContactByType Mobile should not throw error")
	assert.Equal(t, actualMobile, testRuUser.Cell, "Mobile Should be equal")
}
