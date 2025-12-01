package protein

import "errors"

var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid rna")
)

var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			break
		}
		protein, err := FromCodon(rna[i : i+3])
		if errors.Is(err, ErrStop) {
			break
		}
		if err != nil {
			return nil, err
		}

		proteins = append(proteins, protein)
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	protein, exists := codonToProtein[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil

}
