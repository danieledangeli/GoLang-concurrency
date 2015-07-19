package philosopher

import (
	"testing"
	gocheck "gopkg.in/check.v1"
	"github.com/stretchr/testify/mock"
	"github.com/danieledangeli/concurrency/philosopher/table"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	gocheck.TestingT(t)
}

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})


func (s *MySuite) TestCreate(c *gocheck.C) {

	forks := make([]*table.Fork, 4)

	mockTabler := new(MockTabler)
	mockTabler.On("GetForks").Return(forks)

	phil := NewPhilosopher("Daniele", 1, mockTabler)

	c.Assert("Daniele", gocheck.Equals, phil.Name)
	c.Assert("SITTING_ON_TABLE", gocheck.Equals, phil.Status)
}

func (s *MySuite) TestTakeLeftForkIfIsFirst(c *gocheck.C) {

	forks := make([]*table.Fork, 4)

	mockTabler := new(MockTabler)
	mockTabler.On("GetForks").Return(forks)

	phil := NewPhilosopher("Daniele", 0, mockTabler)

	fork := phil.takeLeftFork()

	c.Assert(forks[3], gocheck.Equals, fork)
	c.Assert("WAIT_FOR_LEFT_FORK[3]", gocheck.Equals, phil.Status) //even if it has been taken
}

func (s *MySuite) TestTakeLeftForkI(c *gocheck.C) {

	forks := make([]*table.Fork, 4)

	mockTabler := new(MockTabler)
	mockTabler.On("GetForks").Return(forks)

	phil := NewPhilosopher("Daniele", 2, mockTabler)

	fork := phil.takeLeftFork()

	c.Assert(forks[1], gocheck.Equals, fork)
	c.Assert("WAIT_FOR_LEFT_FORK[1]", gocheck.Equals, phil.Status) //even if it has been taken
}

func (s *MySuite) TestTakeRightFork(c *gocheck.C) {

	forks := make([]*table.Fork, 4)

	mockTabler := new(MockTabler)
	mockTabler.On("GetForks").Return(forks)

	phil := NewPhilosopher("Daniele", 2, mockTabler)

	fork := phil.takeRightFork()

	c.Assert(forks[2], gocheck.Equals, fork)
	c.Assert("WAIT_FOR_RIGHT_FORK[2]", gocheck.Equals, phil.Status) //even if it has been taken
}

func (s *MySuite) TestPanicWithNegativeIndex(c *gocheck.C) {
	c.Assert(func() { NewPhilosopher("", -1, nil) }, gocheck.PanicMatches, `Cannot make philospher with negative index`)
}


type MockTabler struct {
	mock.Mock
}

func (m MockTabler) GetForks() []*table.Fork {
	args := m.Called()
	return args.Get(0).([]*table.Fork)
}