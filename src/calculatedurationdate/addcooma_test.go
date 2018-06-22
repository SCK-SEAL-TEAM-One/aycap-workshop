package calculatedurationdate

import "testing"

func Test_AddComma_Input_1000_should_be_1_comma_000(t *testing.T) {
	number := int64(100000)
	expected := "100,000"

	actual := AddComma(number)

	if actual != expected {
		t.Errorf("expected %s but it is %s", expected, actual)
	}
}

func TestReverse(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := ReverseNumber(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
