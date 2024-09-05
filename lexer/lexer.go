package lexer

type Lexer struct {
	text        string
	index       int
	punctuators map[rune]TokenType
}

func New(text string) *Lexer {
	result := &Lexer{
		text:        text,
		index:       0,
		punctuators: make(map[rune]TokenType),
	}

	for _, tokenType := range TokenTypes() {
		if tokenType >= 0 {
			result.punctuators[rune(tokenType)] = tokenType
		}
	}

	return result
}

func (l *Lexer) HasNext() bool {
	return l.index < len(l.text)
}

func (l *Lexer) Next() Token {
	for l.index < len(l.text) {
		c := rune(l.text[l.index])

		if tokenType, ok := l.punctuators[c]; ok {
			l.index++

			return NewToken(tokenType, string(c))
		}

		// Skip whitespace
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			l.index++

			continue
		}

		// Parse name
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' {
			start := l.index
			l.index++

			for l.index < len(l.text) {
				c = rune(l.text[l.index])

				if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_' {
					l.index++
				} else {
					break
				}
			}

			return NewToken(TypeName, l.text[start:l.index])
		}

		// Parse number
		if c >= '0' && c <= '9' {
			start := l.index
			l.index++

			for l.index < len(l.text) {
				c = rune(l.text[l.index])

				if (c < '0' || c > '9') && c != '.' {
					break
				}

				l.index++
			}

			return NewToken(TypeNumber, l.text[start:l.index])
		}
	}

	l.index++

	return NewToken(TypeEOF)
}
