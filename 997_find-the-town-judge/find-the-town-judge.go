package find_the_town_judge

func findJudge(N int, trust [][]int) int {
	p := [1000][]int{}

	for i := 0; i < N; i++ {
		p[i] = []int{0, 0}
	}

	for _, e := range trust {
		v1 := e[0]
		v2 := e[1]
		p[v1-1][1]++
		p[v2-1][0]++
	}

	maxI := -1
	for i := 0; i < N; i++ {
		if p[i][0] == N-1 {
			maxI = i
			break
		}
	}

	if maxI != -1 {
		if p[maxI][1] == 0 {
			return maxI + 1
		}
	}

	return -1
}
