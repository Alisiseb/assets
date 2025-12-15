package house

import "strings"

type Part struct {
	noun string
	verb string
}

var parts = []Part{
	{"house that Jack built", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

// This is the rat
// that ate the malt
// that lay in the house that Jack built.

// 3

func Verse(n int) string {
	if n < 1 || n > len(parts) {
		return ""
	}
	verseWriter := strings.Builder{}
	verseWriter.WriteString("This is the " + parts[n-1].noun)
	for i := n - 1; i > 0; i-- {
		verseWriter.WriteString("\n" + "that " + parts[i].verb + " the " + parts[i-1].noun)

	}
	return verseWriter.String() + "."
}

func Song() string {
	lyricist := strings.Builder{}
	for i := 1; i <= len(parts); i++ {
		lyricist.WriteString(Verse(i))
		if i != len(parts) {
			lyricist.WriteString("\n\n")
		}
	}
	return lyricist.String()

}
