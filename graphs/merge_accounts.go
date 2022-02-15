package graphs

import "sort"

// MergeAccount_DFS ...
// T -> O(nk log(nk))
// S -> O(nk)
func MergeAccount_DFS(accounts [][]string) [][]string {
	// 1. adjacency map.
	var adjacency map[string][]string
	for _, account := range accounts {
		if len(account) < 2 {
			continue
		}

		firstAccountEmail := account[0]
		for i := 2; i < len(account); i++ {
			neighbour := account[i]

			_, faeok := adjacency[firstAccountEmail]
			switch {
			case faeok:
				adjacency[firstAccountEmail] = append(adjacency[firstAccountEmail], neighbour)
			default:
				adjacency[firstAccountEmail] = []string{neighbour}
			}

			_, nok := adjacency[neighbour]
			switch {
			case nok:
				adjacency[neighbour] = append(adjacency[neighbour], firstAccountEmail)
			default:
				adjacency[neighbour] = []string{firstAccountEmail}
			}

		}
	}

	// Build merged arrays using depth first search.
	var (
		mergedAccounts [][]string
		visited        = map[string]struct{}{}
	)

	// T -> O(n) where n is the number of accounts.
	for _, account := range accounts {
		if len(account) == 0 {
			continue
		}

		firstAccountEmail := account[0]
		if _, ok := visited[firstAccountEmail]; ok {
			continue
		}

		// T -> O(nk)
		var mergedAccount = []string{account[0]}
		mergeSortDFS(firstAccountEmail, &mergedAccount, &visited, &adjacency)

		// T -> O(log(nk))
		sort.Strings(mergedAccount[1:])

		mergedAccounts = append(mergedAccounts, mergedAccount)
	}

	return mergedAccounts
}

func mergeSortDFS(email string, mergedAccount *[]string, visited *map[string]struct{}, adjacency *map[string][]string) {
	(*visited)[email] = struct{}{}
	*mergedAccount = append(*mergedAccount, email)

	neighbours := (*adjacency)[email]
	for _, neighbour := range neighbours {
		if _, ok := (*visited)[neighbour]; ok {
			continue
		}

		mergeSortDFS(neighbour, mergedAccount, visited, adjacency)
	}
}
