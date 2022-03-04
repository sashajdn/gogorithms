package graphs

import "sort"

// MergeAccounts ...
//
// T -> O(e) where n is the total number of edges (in this case emails)
// S -> O(max(accounts, emails))
func MergeAccounts_Improved(accounts [][]string) [][]string {
	dsu := BuildDSU(len(accounts))
	groupings := make(map[string]int)

	for group, account := range accounts {
		for _, email := range account[1:] {
			otherGroup, seen := groupings[email]
			if !seen {
				groupings[email] = group
				continue
			}

			dsu.unionBySize(group, otherGroup)
		}
	}

	components := make(map[int][]string)
	for email, group := range groupings {
		accountName := accounts[group][0]
		representative := dsu.findRepresentative(group)

		if _, ok := components[representative]; !ok {
			components[representative] = []string{accountName}
		}

		components[representative] = append(components[representative], email)
	}

	// T -> O(len(accounts) * elog(e)) where e is the total number of distinct emails.
	var mergedAccounts [][]string
	for _, component := range components {
		sort.Strings(component[1:])
		mergedAccounts = append(mergedAccounts, component)
	}

	return mergedAccounts
}

// MergeAccounts ...
//
// T -> O()
// S -> O()
func MergeAccounts(accounts [][]string) [][]string {
	// Build DSU.
	dsu := BuildDSU(len(accounts))

	// Perform Union on Groups.
	// { "email" index(connected component) }
	emailGroup := make(map[string]int)
	for j := 0; j < len(accounts); j++ {
		accountSize := len(accounts[j])

		for i := 1; i < accountSize; i++ {
			email := accounts[j][i]

			group, seen := emailGroup[email]
			if !seen {
				emailGroup[email] = j
				continue
			}

			dsu.unionBySize(j, group)
		}
	}

	// [0, 0, 2] -> r
	// [1, 1, 1] -> s

	// Iterate through group - find unioned representatives.
	// Here we basically gather all connected components.
	// eg. email_4 in group 0
	//
	// email_4, 0
	// rep -> 0
	components := make(map[int][]string)
	for email, group := range emailGroup {
		representative := dsu.findRepresentative(group)

		if _, ok := components[representative]; !ok {
			components[representative] = []string{}
		}

		components[representative] = append(components[representative], email)
	}

	var mergedAccounts [][]string
	for group, component := range components {
		sort.Strings(component)
		component = append([]string{accounts[group][0]}, component...)
		mergedAccounts = append(mergedAccounts, component)
	}

	return mergedAccounts
}

func BuildDSU(size int) *DSU {
	var r = make([]int, 0, size)
	for i := 0; i < size; i++ {
		r = append(r, i)
	}

	var s = make([]int, 0, size)
	for i := 0; i < size; i++ {
		s = append(s, 1)
	}

	return &DSU{
		representative: r,
		size:           s,
	}
}

type DSU struct {
	representative []int
	size           []int
}

func (d *DSU) findRepresentative(x int) int {
	if x == d.representative[x] {
		return x
	}

	// Path compression.
	d.representative[x] = d.findRepresentative(d.representative[x])
	return d.representative[x]
}

func (d *DSU) unionBySize(a, b int) {
	ra, rb := d.findRepresentative(a), d.findRepresentative(b)

	if ra == rb {
		return
	}

	switch {
	case d.size[ra] >= d.size[rb]:
		d.size[ra] += d.size[rb]
		d.representative[rb] = d.representative[ra]
	default:
		d.size[rb] += d.size[ra]
		d.representative[ra] = d.representative[rb]
	}
}
