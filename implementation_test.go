package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// Test Examples:

// func (s *MySuite) TestPass(c *C) {
// 	c.Assert(42, Equals, 42)
// }

// func (s *MySuite) TestFail(c *C) {
// 	c.Assert(42, Equals, 43)
// }

func (s *MySuite) TestPrefixToInfixDefault(c *C) {
	res, err := PrefixToInfix("+ 5 * - 4 2 32222")

	c.Assert(err, Equals, nil)
	c.Assert(res, Equals, "5 + (4 - 2) * 3")
}

func (s *MySuite) TestPrefixToInfix222(c *C) {
	res, err := PrefixToInfix("+ 2 * 2 2")

	c.Assert(err, Equals, nil)
	c.Assert(res, Equals, "2 + 2 * 2")
}

func (s *MySuite) TestPrefixToInfixMultiply(c *C) {
	res, err := PrefixToInfix("* 7 - 1 * 13 4")

	c.Assert(err, Equals, nil)
	c.Assert(res, Equals, "7 * (1 - 13 * 4)")
}

func (s *MySuite) TestPrefixToInfixLong(c *C) {
	res, err := PrefixToInfix("+ 11 + 22 + 33 + 4 + 5 + 6 + 7 8")

	c.Assert(err, Equals, nil)
	c.Assert(res, Equals, "11 + 22 + 33 + 4 + 5 + 6 + 7 + 8")
}

func (s *MySuite) TestPrefixToInfixLong2(c *C) {
	res, err := PrefixToInfix("/ + * 14 - 15 7 * 3 17 + 10 1")

	c.Assert(err, Equals, nil)
	c.Assert(res, Equals, "(14 * (15 - 7) + 3 * 17) / (10 + 1)")
}

func (s *MySuite) TestPrefixToInfixEmptyString(c *C) {
	res, err := PrefixToInfix("")

	c.Assert(err, Not(Equals), nil)
	c.Assert(err.Error(), Equals, "Empty string")
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestPrefixToInfixExtraNumber(c *C) {
	res, err := PrefixToInfix("* 7 - 1 * 13 4 7")

	c.Assert(err, Not(Equals), nil)
	c.Assert(res, Equals, "")
}

func (s *MySuite) TestPrefixToInfixFewNumbers(c *C) {
	res, err := PrefixToInfix("- * 2 2")

	c.Assert(err, Not(Equals), nil)
	c.Assert(res, Equals, "")
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("* + 2 - 5 8 - 101 / 36 6")
	fmt.Println(res)

	// Output:
	// (2 + 5 - 8) * (101 - 36 / 6)
}
