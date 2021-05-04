package arrays

const (
	homeTeamWon = 1
	awayTeamWon = 0
)

// TournamentWinner T(N) -> O(n), S(N) -> O(k)
// n -  number of competitions
// k - number of teams
func TournamentWinner(competitions [][]string, results []int) string {
	table := map[string]int{}
	for i, game := range competitions {
		h, a := game[0], game[1]
		switch results[i] {
		case homeTeamWon:
			table[h] += 3
		case awayTeamWon:
			table[a] += 3
		}
	}
	var (
		maxTeam  string
		maxValue int
	)
	for k, v := range table {
		if v > maxValue {
			maxTeam = k
		}
	}
	return maxTeam
}
