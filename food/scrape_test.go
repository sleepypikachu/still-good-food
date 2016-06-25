package food

import "testing"

func TestScrape(t *testing.T) {
	r := "Gooseberry \u0026 custard pies"
	cases := []struct {
		in   string
		want string
	}{
		{"http://www.bbcgoodfood.com/recipes/gooseberry-custard-pies", r},
	}
	for _, c := range cases {
		got, _ := Scrape(c.in)
		if got.Name != c.want {
			t.Errorf("Scrape%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
