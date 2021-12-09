package search

func MaxValueSequenceAndSequenceInvers(x []int) (max *int) {
	sequence := map[int][]int{}
	sequenceIndex := -1
	boolBefore := false
	sequenceInvers := map[int][]int{}
	sequenceInversIndex := -1
	boolBeforeInvers := false

	// search sequence and sequence invers
	for i, v := range x[1:] {
		// if sequence
		if v-1 == x[i] {
			if !boolBefore {
				sequenceIndex++
				boolBefore = true
			}
			sequence[sequenceIndex] = append(sequence[sequenceIndex], v)
		} else {
			boolBefore = false
		}

		// if sequence invers
		if v == x[i]-1 {
			if !boolBeforeInvers {
				sequenceInversIndex++
				boolBeforeInvers = true
			}
			sequenceInvers[sequenceInversIndex] = append(sequenceInvers[sequenceInversIndex], x[i])
		} else {
			boolBeforeInvers = false
		}
	}

	// check all sequence and sequence invers is palindrome
	for _, seq := range sequence {
		for _, seqInvers := range sequenceInvers {
			// if palindrome and value greater than current max
			// save maximum value to max
			if isPalindrome(seq, seqInvers) {
				if max == nil || *max < seqInvers[0] {
					max = &seqInvers[0]
				}
			}
		}
	}
	return max
}

func isPalindrome(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	yIndex := len(y) - 1
	for i := range x {
		if x[i] != y[yIndex-i] {
			return false
		}
	}
	return true
}
