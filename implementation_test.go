package lab2

import (
	"fmt"
	"testing"
	"strings"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type PrefixToPostfixSuite struct{}

var _ = Suite(&PrefixToPostfixSuite{})

func (s *PrefixToPostfixSuite) TestPrefixToPostfix(c *C) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"+ 3 4", "3 4 +", nil},
		{"* + 3 4 5", "3 4 + 5 *", nil},
		{"- * + 3 4 5 6", "3 4 + 5 * 6 -", nil},
		{"^ 2 3", "2 3 ^", nil},
		{"^ 2 + 3 4", "2 3 4 + ^", nil},
		{"^ + 2 3 4", "2 3 + 4 ^", nil},
		{"", "", fmt.Errorf("invalid expression")},
		{"3", "", fmt.Errorf("invalid expression")},
		{"+ 3", "", fmt.Errorf("invalid expression")},
		{"^ 2", "", fmt.Errorf("invalid expression")},
		{"- * + 3 4 5", "", fmt.Errorf("invalid expression")},
		{"^ 2 + 3 4 5", "", fmt.Errorf("invalid expression")},
	}

	for _, tc := range testCases {
		result, err := PrefixToPostfix(tc.input)
		if tc.err != nil {
			c.Assert(err, ErrorMatches, tc.err.Error())
		} else {
			c.Assert(err, IsNil)
			c.Assert(strings.TrimSpace(result), Equals, strings.TrimSpace(tc.expected))
		}
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
