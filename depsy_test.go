package depsy

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io/ioutil"
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type DepsySuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&DepsySuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (s *DepsySuite) TestExtract(c *C) {
	data, err := ioutil.ReadFile("testdata/etcd.mod")
	c.Assert(err, IsNil)
	deps := Extract(data, false)
	c.Assert(len(deps), Equals, 20)
	deps = Extract(data, true)
	c.Assert(len(deps), Equals, 74)
	c.Assert(deps[5], DeepEquals, Dependency{"go.etcd.io/etcd/api", "3.6.0", "alpha.0"})
	c.Assert(deps[27], DeepEquals, Dependency{"github.com/golang-jwt/jwt", "3.2.2", ""})

	data, err = ioutil.ReadFile("go.mod")
	c.Assert(err, IsNil)
	deps = Extract(data, false)
	c.Assert(len(deps), Equals, 1)
	deps = Extract(data, true)
	c.Assert(len(deps), Equals, 4)
	c.Assert(deps[0], DeepEquals, Dependency{"github.com/essentialkaos/check", "1.4.0", ""})
}

func (s *DepsySuite) TestAux(c *C) {
	c.Assert(parseDependencyLine(""), DeepEquals, Dependency{})
	c.Assert(getMajorVersion("12"), Equals, "12")
}
