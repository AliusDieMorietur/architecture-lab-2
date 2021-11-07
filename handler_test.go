package lab2

import (
	"bytes"
	"strings"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.

// Test Examples:

// func (s *MySuite) TestPass(c *C) {
// 	c.Assert(42, Equals, 42)
// }

// func (s *MySuite) TestFail(c *C) {
// 	c.Assert(42, Equals, 43)
// }

func (s *MySuite) TestComputeHandler(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("+ 2 2"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(b.String(), Equals, "2 + 2")
}

func (s *MySuite) TestComputeHandlerError(c *C) {
	b := bytes.NewBuffer(make([]byte, 0))

	handler := ComputeHandler{
		Input:  strings.NewReader("23 22 broke another test"),
		Output: b,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}
