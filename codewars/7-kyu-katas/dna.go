// https://www.codewars.com/kata/554e4a2f232cdd87d9000038/train/go
package kata

func DNAStrand(dna string) string {
	var conv map[string]string = map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}

	output := ""
	for _, char := range dna {
		output += conv[string(char)]
	}

	return output
}
