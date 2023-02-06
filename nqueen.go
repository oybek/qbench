package main

func backtrack(qs []queen, boardSize int, par int) [][]queen {
	if len(qs) == boardSize {
		return [][]queen{qs}
	}

	if par == 1 {
		x := len(qs)
		placements := [][]queen{}
		for y := 0; y < boardSize; y++ {
			candidate := queen{x, y}
			if !attack(qs, candidate) {
				possiblePlacement := append(qs, candidate)
				placements = append(placements, backtrack(possiblePlacement, boardSize, 1)...)
			}
		}
		return placements
	}

	ch := make(chan [][]queen)

	candidates := make(map[int][]queen)
	x := len(qs)
	for y := 0; y < boardSize; y++ {
		candidates[y%par] = append(candidates[y%par], queen{x, y})
	}

	for _, cs := range candidates {
		go func(cs []queen) {
			placements := [][]queen{}
			for _, c := range cs {
				if !attack(qs, c) {
					possiblePlacement := append(qs, c)
					placements = append(placements, backtrack(possiblePlacement, boardSize, 1)...)
				}
			}
			ch <- placements
		}(cs)
	}

	placements := [][]queen{}
	for i := 0; i < par; i++ {
		placements = append(placements, <-ch...)
	}

	return placements
}
