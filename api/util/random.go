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
	return fmt.Sprintf("%d-%d-%d", RandomInt(1, 31), RandomInt(1, 12), RandomInt(1000, 3000))
}

func RandomAddress() (a models.Address) {
	a.ID = RandomID()
	a.CEP = strconv.Itoa(int(RandomInt(8, 8)))
	a.Complement = RandomString(8)
	a.City = RandomString(8)
	a.Neighborhood = RandomString(8)
	a.Street = RandomString(8)
	a.Number = int(RandomInt(1, 256))

	return
}

func RandomPerson() (p models.Person) {
	p.ID = RandomID()
	p.Name = RandomString(8)
	p.Gender = RandomString(1)
	p.Birthday = RandomDate()

	return
}

func RandomUser() (u models.User) {
	u.ID = RandomID()
	u.Person = RandomPerson()
	u.Address = RandomAddress()
	u.Email = RandomEmail()
	u.CellNumber = strconv.Itoa(int(RandomInt(9, 9)))
	u.PhoneNumber = strconv.Itoa(int(RandomInt(8, 8)))
	u.CPF = strconv.Itoa(int(RandomInt(11, 11)))
	u.CreationDate = RandomDate()

	return
}

func RandomAuthor() (a models.Author) {
	a.ID = RandomID()
	a.Person = RandomPerson()

	return
}

func RandomGenre() (g models.Genre) {
	g.ID = RandomID()
	g.Name = RandomName()

	return
}

func RandomBook() (b models.Book) {
	b.ID = RandomID()
	b.Title = RandomName()
	b.Pages = int(RandomInt(1, 256))
	b.Year = int(RandomInt(1000, 3000))
	b.Author = RandomAuthor()
	b.Genre = RandomGenre()

	return
}

func RandomLending() (l models.Lending) {
	l.ID = RandomID()
	l.User = RandomUser()
	l.Book = RandomBook()
	l.LendDay = RandomDate()
	l.Returned = 0
	l.Devolution = models.Devolution{
		ID:   RandomID(),
		Date: RandomDate(),
	}

	return
}

func RandomID() int64 {
	return RandomInt(1, 1024)
}
