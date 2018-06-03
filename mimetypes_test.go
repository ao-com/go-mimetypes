package mimetypes

import (
	"testing"
)

var getByExtensionTests = []struct {
	ext string
	out string
}{
	{ext: ".323", out: mappings[".323"]},
	{ext: ".zip", out: mappings[".zip"]},
	{ext: ".unknown", out: ""},
}

func TestGetByExtension(t *testing.T) {
	for _, test := range getByExtensionTests {
		result := GetByExtension(test.ext)
		if result != test.out {
			t.Errorf("unexpected out: %s", result)
		}
	}
}
