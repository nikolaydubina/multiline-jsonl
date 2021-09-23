package mjsonl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// FormatJSONL will read JSONL from input
// If there are multiline JSONs then they will be either
// expanded, meaning pretty printed, or un-expanded, meaning become single lines.
func FormatJSONL(r io.Reader, o io.Writer, expand bool) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(SplitMultilineJSONL)

	for scanner.Scan() {
		var inJSON map[string]interface{}
		if err := json.Unmarshal(scanner.Bytes(), &inJSON); err != nil {
			return fmt.Errorf("can not decode json: %w", err)
		}

		var outJSON []byte
		var err error
		if expand {
			outJSON, err = json.MarshalIndent(inJSON, "", "    ")
		} else {
			outJSON, err = json.Marshal(inJSON)
		}

		if err != nil {
			return fmt.Errorf("can not encode json: %w", err)
		}

		outJSON = append(outJSON, '\n')

		_, err = o.Write(outJSON)
		if err != nil {
			return fmt.Errorf("can not write json bytes: %w", err)
		}
	}

	return scanner.Err()
}
