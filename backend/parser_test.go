package backend

import (
	"os"
	"testing"
)

func TestReadWriteINI(t *testing.T) {
	content := `
[freddy]
level=5
beatgame=1
lives=3

[other]
something=else
key_with_equals=value=with=equals
`
	tmpfile, err := os.CreateTemp("", "test_save_*.ini")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	data, err := ReadINI(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadINI failed: %v", err)
	}

	expected := map[string]string{
		"freddy|level":          "5",
		"freddy|beatgame":       "1",
		"freddy|lives":          "3",
		"other|something":       "else",
		"other|key_with_equals": "value=with=equals",
	}

	for k, v := range expected {
		if data[k] != v {
			t.Errorf("Key %s: expected %s, got %s", k, v, data[k])
		}
	}

	data["freddy|level"] = "6"
	data["freddy|newkey"] = "test"
	data["newsection|foo"] = "bar"

	if err := WriteINI(tmpfile.Name(), data); err != nil {
		t.Fatalf("WriteINI failed: %v", err)
	}

	data2, err := ReadINI(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadINI 2 failed: %v", err)
	}

	if data2["freddy|level"] != "6" {
		t.Errorf("Update failed: expected level=6, got %s", data2["freddy|level"])
	}
	if data2["freddy|newkey"] != "test" {
		t.Errorf("New key failed: expected test, got %s", data2["freddy|newkey"])
	}
	if data2["newsection|foo"] != "bar" {
		t.Errorf("New section failed: expected bar, got %s", data2["newsection|foo"])
	}
}
