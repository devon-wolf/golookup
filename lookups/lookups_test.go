package lookups

import "testing"

func TestLookups(t *testing.T) {
	cases := []struct {
		word           string
		resource       string
		expectedStatus int
	}{
		{"purple", "d", 200},
		{"healthy", "t", 200},
	}
	for _, c := range cases {
		_, actualStatus := Lookup(c.word, c.resource)
		if actualStatus != c.expectedStatus {
			t.Errorf("Lookup(%q, %q) == %q, expected %q", c.word, c.resource, actualStatus, c.expectedStatus)
		}
	}
}
