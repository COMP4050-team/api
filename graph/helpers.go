package graph

func getOffset(from *int) int {
	if from == nil {
		return 1
	}

	if *from < 1 {
		return 1
	}

	return *from
}
