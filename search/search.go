package search

func MaxValueSequenceAndSequenceInvers(slice []int) (max *int) {

	//get sequence and sequence invers
	sequence, sequenceInvers := getSequenceAndSequenceInvers(slice)

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

func getSequenceAndSequenceInvers(slice []int) (sequence, sequenceInvers map[int][]int) {
	sequence, sequenceInvers = map[int][]int{}, map[int][]int{}
	sequenceIndex, sequenceInversIndex := -1, -1
	isSequenceBefore, isSequenceInversBefore := false, false

	// search sequence and sequence invers
	for index, currentValue := range slice[1:] {
		// if sequence
		if currentValue-1 == slice[index] {
			if !isSequenceBefore {
				sequenceIndex++
				isSequenceBefore = true
			}
			sequence[sequenceIndex] = append(sequence[sequenceIndex], currentValue)
		} else {
			isSequenceBefore = false
		}

		// if sequence invers
		if currentValue == slice[index]-1 {
			if !isSequenceInversBefore {
				sequenceInversIndex++
				isSequenceInversBefore = true
			}
			sequenceInvers[sequenceInversIndex] = append(sequenceInvers[sequenceInversIndex], slice[index])
		} else {
			isSequenceInversBefore = false
		}
	}
	return sequence, sequenceInvers
}

func isPalindrome(sequence, sequenceInvers []int) bool {
	if len(sequence) != len(sequenceInvers) {
		return false
	}
	sequenceInversLength := len(sequenceInvers) - 1
	for sequenceIndex := range sequence {
		if sequence[sequenceIndex] != sequenceInvers[sequenceInversLength-sequenceIndex] {
			return false
		}
	}
	return true
}
