package kata

import "testing"

func TestNormalizeUsername(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: " Alice ", want: "alice"},
		{in: "Bob Smith", want: "bob_smith"},
		{in: "   ", want: ""},
	}

	for _, tc := range cases {
		got := NormalizeUsername(tc.in)
		if got != tc.want {
			t.Fatalf("NormalizeUsername(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
