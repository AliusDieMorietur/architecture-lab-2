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

// These are tests to test how gocheck works

func (s *MySuite) TestPass(c *C) {
	c.Assert(42, Equals, 42)
}

func (s *MySuite) TestFail(c *C) {
	c.Assert(42, Equals, 43)
}

func (s *MySuite) TestPrefixToInfix(c *C) {
	res, err := PrefixToInfix("+ 5 * - 4 2 3")

	c.Assert(err, Not(Equals), nil)
	c.Assert(res, Equals, "5 + (4 - 2) * 3")
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
