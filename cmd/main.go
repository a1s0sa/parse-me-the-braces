package main

import (
	"fmt"
	"parse-me-the-braces/internal/tokenizer"
)

func main() {
	isSingleDigit, isWhiteSpace, isQuote := tokenizer.IsDigit('1'), tokenizer.IsWhiteSpace('\r'), tokenizer.IsQuote('"')
	fmt.Printf("%v\n", isSingleDigit)
	fmt.Printf("%v\n", isWhiteSpace)
	fmt.Printf("%v\n", isQuote)
}
