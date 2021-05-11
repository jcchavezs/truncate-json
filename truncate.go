package truncate

import (
	"errors"
)

func isTruncatable(payload string) bool {
	if payload[0] == '{' || payload[0] == '[' {
		return true
	}

	return false
}

func truncateJSON(payload string, tl int) (string, error) {
	if len(payload) <= tl {
		return payload, nil
	}

	if !isTruncatable(payload) {
		return payload, errors.New("not truncatable")
	}

	tPayload := payload[:tl]

	nQuotes := 0
	nSemicolon := 0
	symbols := []rune{}
	for _, c := range tPayload {
		switch c {
		case '[':
			symbols = append(symbols, '[')
		case '{':
			nQuotes = 0
			symbols = append(symbols, '{')
		case ':':
			nSemicolon += 1
		case '"':
			nQuotes += 1
		case '}':
			symbols = symbols[:len(symbols)-1]
		case ']':
			symbols = symbols[:len(symbols)-1]
		}
	}

	closingsymbols := ""
	nQuotes = nQuotes % 4
	if symbols[len(symbols)-1] == '{' {
		if nQuotes == 1 {
			closingsymbols = "\": \"\""
		} else if nQuotes == 2 {
			closingsymbols = "\"\""
		} else if nQuotes == 3 {
			closingsymbols = "\""
		}
	} else if symbols[len(symbols)-1] == '[' {
		if nQuotes == 1 || nQuotes == 3 {
			closingsymbols = "\""
		}
	}

	for i := len(symbols) - 1; i >= 0; i-- {
		switch symbols[i] {
		case '{':
			closingsymbols += "}"
		case '[':
			closingsymbols += "]"
		}
	}

	if tPayload[0] == '{' {
		tPayload = "{\"hypertrace\": \"truncated\", " + tPayload[1:] + closingsymbols
	} else if tPayload[0] == '[' {
		tPayload = "[{\"hypertrace\": \"truncated\"}, " + tPayload[1:] + closingsymbols
	}

	return tPayload, nil
}
