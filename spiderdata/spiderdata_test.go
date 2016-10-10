package spiderdata

import (
	"encoding/json"
	"os"
	"testing"
)

func TestJSON(t *testing.T) {
	file, err := os.Open("sample_requests/simple_request.json")
	if err != nil {
		t.Fatalf("unable to open JSON: %s", err)
	}
	defer file.Close()

	data := &Request{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		t.Fatalf("unable to parse JSON: %s", err)
	}
}
