package main

import (
	"fmt"
	"github.com/a1s0sa/parse-me-the-braces/internal/tokenizer"
)

func main() {
	isSingleDigit, isWhiteSpace, isQuote := tokenizer.IsDigit('1'), tokenizer.IsWhiteSpace('\r'), tokenizer.IsQuote('"')
	isBoolean, tokensConsumed := tokenizer.IsBoolean(0, "tru")
	fmt.Printf("%v\n", isSingleDigit)
	fmt.Printf("%v\n", isWhiteSpace)
	fmt.Printf("%v\n", isQuote)
	fmt.Printf("%v\n", isBoolean)
	fmt.Printf("%v\n", tokensConsumed)
}
