package tokenizer

// digit chars '0'-'9' equal 48-57 int range
func IsDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}

func IsWhiteSpace(char uint8) bool {
	switch char {
	case '\t', '\n', '\r', ' ':
		return true
	default:
		return false
	}
}

func IsPunctuation(char uint8) bool {
	switch char {
	case '[', ']', '{', '}', ',', ':':
		return true
	default:
		return false
	}
}

func IsQuote(char uint8) bool {
	return char == '"'
}
