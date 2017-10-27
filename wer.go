package wer

import "sort"

// CalculatePercent calculates the WER based on (substitutions + deletions + insertions) / length of reference
func CalculatePercent(reference, hypothesis []string) float64 {
	if len(reference) == 0 {
		return 0
	}

	changes := Changes(reference, hypothesis)

	return float64(changes) / float64(len(reference))
}

// Changes calculates the number of subsitutions + deletions + insertions
func Changes(reference, hypothesis []string) int {
	d := make([][]int, len(reference)+1)
	for i := range d {
		d[i] = make([]int, len(hypothesis)+1)
	}

	// might need to be <= instead of <
	for i := 0; i < len(reference)+1; i++ {
		for j := 0; j < len(hypothesis)+1; j++ {
			if i == 0 {
				d[0][j] = j
			} else if j == 0 {
				d[i][0] = i
			}
		}
	}

	for i := 1; i < len(reference)+1; i++ {
		for j := 1; j < len(hypothesis)+1; j++ {
			if reference[i-1] == hypothesis[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				substitution := d[i-1][j-1] + 1
				insertion := d[i][j-1] + 1
				deletion := d[i-1][j] + 1
				d[i][j] = min(substitution, insertion, deletion)
			}
		}
	}

	return d[len(reference)][len(hypothesis)]
}

func min(vals ...int) int {
	sort.Ints(vals)
	return vals[0]
}
