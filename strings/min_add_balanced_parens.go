package strings

// MinAddBalancedParenthesis ...
func MinAddBalancedParenthesis(s string) int {
	var (
		minimumAdd, balance int
	)
	for _, r := range s {
		switch r {
		case ')':
			if balance == 0 {
				minimumAdd++
				continue
			}
			balance--
		case '(':
			balance++
		}
	}
	return minimumAdd + balance
}
