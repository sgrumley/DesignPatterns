package solid

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

// Primary role is to add journal entries and manipulate some of the data
type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "/n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {

}

var LineSeparator = "\n"

// Separates the concern, rather than writing having a method of the struct
func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// Or creates a new struct to have a new concern
type Persistence struct {
	lineSeparator string
}

// Role is to persist data (file io)
func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

// Breaking pattern: separation of concerns
// Anti pattern: God object
// responsibility of Journal is not to deal with persistence

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func SingleResponsibility() {
	j := Journal{}
	j.AddEntry("First log")
	j.AddEntry("Second log")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "Journal.txt")
}
