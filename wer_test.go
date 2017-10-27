package wer

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/cheekybits/is"
)

func TestMin(t *testing.T) {
	is := is.New(t)
	is.Equal(0, min(5, 9, 1, 0))
}

func TestCalculatePercent(t *testing.T) {
	is := is.New(t)

	is.Equal(0.3333333333333333, CalculatePercent(strings.Split("who is there", " "), strings.Split("is there", " ")))
	is.Equal(0, CalculatePercent([]string{}, strings.Split("is there", " ")))
}

func TestChanges(t *testing.T) {
	is := is.New(t)

	is.Equal(1, Changes(strings.Split("who is there", " "), strings.Split("is there", " ")))
	is.Equal(3, Changes(strings.Split("who is there", " "), strings.Split("", " ")))
	is.Equal(3, Changes(strings.Split("", " "), strings.Split("who is there", " ")))
}

func TestChangesLong(t *testing.T) {
	is := is.New(t)

	r, err := readFile("testdata/lorem_ipsum.txt")
	is.NoErr(err)
	h, err := readFile("testdata/lorem_ipsum2.txt")
	is.NoErr(err)

	is.Equal(3, Changes(r, h))
}

func readFile(file string) ([]string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(bytes))

	return strings.Split(s, " "), nil
}
