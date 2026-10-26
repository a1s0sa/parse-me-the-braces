package tokenizer

type TokenLabel int

const (
	TokenEndOfInput TokenLabel = iota
	TokenDigit
	TokenComma
	TokenLeftCurlyBrace
	TokenRightCurlyBrace
	TokenLeftSquareBrace
	TokenRightSquareBrace
	TokenColon
	TokenQuote
	TokenBoolean
	TokenNull
)

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

func IsBoolean(charIndex int, jsonString string) (bool, int) {
	trueTokens := [4]uint8{'t', 'r', 'u', 'e'}
	falseTokens := [5]uint8{'f', 'a', 'l', 's', 'e'}

	switch jsonString[charIndex] {
	case trueTokens[0]:
		trueTokenCount := 1
		if len(jsonString[charIndex:]) < len(trueTokens) {
			return false, 0
		}
		for i := 1; i < len(trueTokens); i++ {
			if trueTokens[i] == jsonString[charIndex+i] {
				trueTokenCount++
			}
		}
		if trueTokenCount < len(trueTokens) {
			return false, 0
		}
		return true, trueTokenCount
	case falseTokens[0]:
		falseTokenCount := 1
		if len(jsonString[charIndex:]) < len(falseTokens) {
			return false, 0
		}
		for i := 1; i < len(falseTokens); i++ {
			if falseTokens[i] == jsonString[charIndex+i] {
				falseTokenCount++
			}
		}
		if falseTokenCount < len(falseTokens) {
			return false, 0
		}
		return true, falseTokenCount
	default:
		return false, 0
	}
}

func IsNull(charIndex int, jsonString string) (bool, int) {
	nullTokens := [4]byte{'n', 'u', 'l', 'l'}

	switch jsonString[charIndex] {
	case nullTokens[0]:
		nullTokensCount := 1
		if len(jsonString[charIndex:]) < len(nullTokens) {
			return false, 0
		}
		for i := 1; i < len(nullTokens); i++ {
			if nullTokens[i] == jsonString[charIndex+i] {
				nullTokensCount++
			}
		}
		if nullTokensCount < len(nullTokens) {
			return false, 0
		}
		return true, nullTokensCount
	default:
		return false, 0
	}
}

func GetToken(jsonString string) (token string, label TokenLabel, rest string, ok bool) {
	index := 0

	for IsWhiteSpace(jsonString[index]) && index < len(jsonString) {
		index++
	}

	if IsDigit(jsonString[index]) {
		restString := ""
		return string(jsonString[index]), TokenDigit, restString, true
	}

	return
}
