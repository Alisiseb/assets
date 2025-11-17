package strand

import "strings"

func ToRNA(dna string) string {
	var RNA string
	for i := 0; i < len(dna); i++ {
		upperCase := strings.ToUpper(string(dna[i]))
		RNA += transcriptionMap[upperCase]
	}
	return RNA
}
var transcriptionMap = map[string]string{
	"A": "U",
	"T": "A",
	"C": "G",
	"G": "C",
}
