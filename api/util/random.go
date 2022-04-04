package util

import (
	"fmt"
	"github.com/marcos-travasso/library-system/models"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Random functions from https://github.com/techschool/simplebank/blob/master/util/random.go

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(8)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomDate() string {
	return fmt.Sprintf("%s-%s-%s", RandomInt(1, 31), RandomInt(1, 12), RandomInt(1000, 3000))
}

func RandomAddress() models.Address {
	var a models.Address

	a.ID = RandomInt(1, 256)
	a.CEP = strconv.Itoa(int(RandomInt(8, 8)))
	a.Complement = RandomString(8)
	a.City = RandomString(8)
	a.Neighborhood = RandomString(8)
	a.Street = RandomString(8)
	a.Number = int(RandomInt(1, 256))

	return a
}

func RandomPerson() models.Person {
	var p models.Person

	p.ID = RandomInt(1, 256)
	p.Name = RandomString(8)
	p.Gender = RandomString(1)
	p.Birthday = RandomDate()

	return p
}

func RandomUser() models.User {
	var u models.User

	u.ID = RandomInt(1, 256)
	u.Person = RandomPerson()
	u.Address = RandomAddress()
	u.Email = RandomEmail()
	u.CellNumber = strconv.Itoa(int(RandomInt(9, 9)))
	u.PhoneNumber = strconv.Itoa(int(RandomInt(8, 8)))
	u.CPF = strconv.Itoa(int(RandomInt(11, 11)))
	u.CreationDate = RandomDate()

	return u
}
