package table

import (
	"testing"
	gocheck "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	gocheck.TestingT(t)
}

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestGetForks(c *gocheck.C) {
	table := NewTable(4)
	c.Assert(4, gocheck.Equals, len(table.GetForks()))
}

func (s *MySuite) TestPanicWithNegativeForks(c *gocheck.C) {
	c.Assert(func() { NewTable(-2) }, gocheck.PanicMatches, `Cannot make table with negative or 0 forks`)
}

func (s *MySuite) TestPanicWithZeroForks(c *gocheck.C) {
	c.Assert(func() { NewTable(0) }, gocheck.PanicMatches, `Cannot make table with negative or 0 forks`)
}