package services

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

var names = []string{
	"Eddie",
	"Garrett",
	"Nolan",
	"Paul",
	"Sergey",
	"Robert",
	"Hayden",
	"Ryder",
	"Zaiden",
	"Marty",
	"Eduardo",
	"Gabriella",
	"Christy",
	"Finley",
	"Nerissa",
	"Alicia",
	"Anna",
	"Sophia",
	"Lilly",
	"Melissa",
}

var types = []string{
	"system", "normal", "public", "private", "file_info",
}

var g *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandName() string {
	length := len(names)
	return names[g.Intn(length)]
}

func RandType_tag() string {
	return types[g.Intn(5)]
}

func Namehash(name string) string {
	h := sha1.New()
	h.Write([]byte(name))

	bs := fmt.Sprintf("%x", h.Sum(nil))
	return bs

}
