package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)

	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// SRPに違反しているため以下のアプローチは取らない
// func (j *Journal) Save(filename string) {
// 	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
// }

// func (j *Journal) Load(filename string) {
// 	//
// }

// func (j *Journal) LoadFromWeb(url *url.URL) {
// 	//
// }

// こちらのアプローチを採用する

// 関数を切り出すパターン
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// Structで定義するパターン
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("牛乳を買った")
	j.AddEntry("ゴルフに行った")
	fmt.Println(j.String())

	// Persistence
	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal_p.txt")
}
