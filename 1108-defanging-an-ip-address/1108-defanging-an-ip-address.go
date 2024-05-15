package leetcode

func defangIPaddr(address string) string {
	// loop backwards through the chars in the string
	for i := len(address) - 1; i > 0; i-- {
		// if the char (rune) is equal to 46 (the byte value of a period)
		if address[i] == 46 {
			// replace it with [.]
			address = address[:i] + "[.]" + address[i+1:]
			// decrement the index counterto avoid hitting the . twice
			i -= 1
		}
	}
	return address
}
