package mjsonl

import "errors"

// SplitMultilineJSONL splits input into JSONs. Each token is a json.
// Separated by new lines and whitespace.
// Implemented by counting opening and closing brackets.
// JSONs have to be objects.
// Escaped brackets are not supported, may return wrong result.
func SplitMultilineJSONL(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return 0, nil, nil
	}

	iStart := -1
	iEnd := -1
	openCount := 0

	for i, b := range data {
		if b == '{' {
			if openCount == 0 {
				iStart = i
			}
			openCount++
			continue
		}
		if b == '}' {
			openCount--
			if openCount <= 0 {
				iEnd = i
				break
			}
		}
	}

	if openCount > 0 {
		return 0, nil, errors.New("got more { than }")
	}
	if openCount < 0 {
		return 0, nil, errors.New("got more } than {")
	}

	if iStart == -1 || iEnd == -1 {
		// none found, just advance further
		return len(data), nil, nil
	}

	if iStart >= iEnd {
		return 0, nil, errors.New("end index has to be after start index")
	}

	return iEnd + 1, data[iStart : iEnd+1], nil
}
