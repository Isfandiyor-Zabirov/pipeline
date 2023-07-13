package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions. Dev.to is awesome. Lets test it" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions. Dev.to is awesome", result)
	}
}
