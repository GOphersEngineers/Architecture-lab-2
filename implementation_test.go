package lab2

import (
	"fmt"
	"gopkg.in/check.v1"
	"testing"
)

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func Test(t *testing.T) { check.TestingT(t) }

func (s *MySuite) TestIsOperator(c *check.C) {
	c.Check(isOperator("+"), check.Equals, true)
	c.Check(isOperator("-"), check.Equals, true)
	c.Check(isOperator("*"), check.Equals, true)
	c.Check(isOperator("/"), check.Equals, true)
	c.Check(isOperator("^"), check.Equals, true)
	c.Check(isOperator("1"), check.Equals, false)
}

func (s *MySuite) TestPrefixToPostfix(c *check.C) {
	result, err := PrefixToPostfix("+ 1 2")
	c.Check(err, check.IsNil)
	c.Check(result, check.Equals, "1 2 +")

	result, err = PrefixToPostfix("- * 3 4 5")
	c.Check(err, check.IsNil)
	c.Check(result, check.Equals, "3 4 * 5 -")

	result, err = PrefixToPostfix("+ 1")
	c.Check(err, check.NotNil)
	c.Check(result, check.Equals, "")

	result, err = PrefixToPostfix("+ - * 1 2 3 * / 4 5 6")
	c.Check(err, check.IsNil)
	c.Check(result, check.Equals, "2 1 * 3 - 4 5 / 6 * +")

	result, err = PrefixToPostfix("+ - * 1 2 3 * / 4 5 6 7")
	c.Check(err, check.ErrorMatches, "invalid expression")
	c.Check(result, check.Equals, "")

	result, err = PrefixToPostfix("")
	c.Check(err, check.ErrorMatches, "invalid expression")
	c.Check(result, check.Equals, "")

	result, err = PrefixToPostfix("+ 3")
	c.Check(err, check.ErrorMatches, "invalid expression")
	c.Check(result, check.Equals, "")

	result, err = PrefixToPostfix("^ 2")
	c.Check(err, check.ErrorMatches, "invalid expression")
	c.Check(result, check.Equals, "")
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
