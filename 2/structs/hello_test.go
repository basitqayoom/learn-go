package hello

import "testing"

func TestSayHello(t *testing.T) {

	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world",
		},
		{
			items:  []string{"Matt"},
			result: "Hello, Matt",
		},
		{
			items:  []string{"Matt", "John"},
			result: "Hello, Matt, John",
		},
	}

	for _, subtest := range subtests {
		if s := Say(subtest.items); s != subtest.result {
			t.Errorf("Expected %q, got %q", subtest.result, s)
		}
	}
}
