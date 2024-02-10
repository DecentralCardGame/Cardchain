package cli

import "encoding/json"

func getJsonArg[T any](from string, into *T) error {
	return json.Unmarshal([]byte(from), into)
}
