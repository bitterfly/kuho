package backend

import (
	"encoding/json"
	"os"
	"testing"
)

func unmarshalJSON(t *testing.T, data interface{}, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Could not open file %s - %s", filename, err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		t.Fatalf("Could not decode file %s - %s", filename, err)
	}
}
