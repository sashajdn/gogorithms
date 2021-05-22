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
		maxTeam   string
		maxPoints int
	)
	for team, points := range table {
		if points > maxPoints {
			maxPoints, maxTeam = points, team
		}
	}
	return maxTeam
}
