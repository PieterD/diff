package diff

func New(iface Interface) []Diff {
	table := lcs(iface)
	diff := walk(iface, table)
	// Reverse
	i := 0
	j := len(diff) - 1
	for i < j {
		diff[i], diff[j] = diff[j], diff[i]
		i++
		j--
	}
	return diff
}

// Constructs a LCSLength table
// http://en.wikipedia.org/wiki/Longest_common_subsequence_problem#Computing_the_length_of_the_LCS
func lcs(iface Interface) [][]int {
	lnum, rnum := iface.Length()
	rows, cols := lnum+1, rnum+1
	table := make([][]int, rows)
	for i := 0; i < rows; i++ {
		table[i] = make([]int, cols)
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if iface.Equal(i-1, j-1) {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				a := table[i-1][j]
				b := table[i][j-1]
				if b > a {
					a = b
				}
				table[i][j] = a
			}
		}
	}
	return table
}

// Walk the lcs table
// http://en.wikipedia.org/wiki/Longest_common_subsequence_problem#Example
func walk(iface Interface, table [][]int) (diff []Diff) {
	i, j := iface.Length()
	for {
		if i == 0 && j == 0 {
			return
		} else if i == 0 {
			j--
			diff = append(diff, Diff{Delta: Right, Index: j})
		} else if j == 0 {
			i--
			diff = append(diff, Diff{Delta: Left, Index: i})
		} else {
			if iface.Equal(i-1, j-1) {
				i--
				j--
				diff = append(diff, Diff{Delta: Both, Index: i})
			} else {
				if table[i-1][j] > table[i][j-1] {
					i--
					diff = append(diff, Diff{Delta: Left, Index: i})
				} else {
					j--
					diff = append(diff, Diff{Delta: Right, Index: j})
				}
			}
		}
	}
}
