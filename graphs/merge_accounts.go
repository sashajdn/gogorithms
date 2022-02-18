package graphs

import "sort"

// MergeAccounts ...
//
// T -> O()
// S -> O()
func MergeAccounts(accounts [][]string) [][]string {
	dsu := BuildDSU(len(accounts))
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

	components := make(map[int][]string)
	for email, group := range emailGroup {
		representative := dsu.findRepresentative(group)

		_, ok := components[representative]
		if !ok {
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
