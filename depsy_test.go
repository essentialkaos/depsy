package depsy

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
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
	data, err := os.ReadFile("testdata/test1.mod")
	c.Assert(err, IsNil)
	deps := Extract(data, false)
	c.Assert(len(deps), Equals, 20)
	deps = Extract(data, true)
	c.Assert(len(deps), Equals, 74)
	c.Assert(deps[5], DeepEquals, Dependency{"go.etcd.io/etcd/api/v3", "3.6.0", "./api"})
	c.Assert(deps[16], DeepEquals, Dependency{"golang.org/x/time", "0.0.0", "20210220033141-f8bda1e9f3ba"})
	c.Assert(deps[27], DeepEquals, Dependency{"github.com/golang-jwt/jwt", "3.2.2", ""})

	data, err = os.ReadFile("go.mod")
	c.Assert(err, IsNil)
	deps = Extract(data, false)
	c.Assert(len(deps), Equals, 1)
	deps = Extract(data, true)
	c.Assert(len(deps), Equals, 4)
	c.Assert(deps[0], DeepEquals, Dependency{"github.com/essentialkaos/check", "1.4.0", ""})

	data, err = os.ReadFile("testdata/test2.mod")
	c.Assert(err, IsNil)
	deps = Extract(data, true)

	c.Assert(deps[5], DeepEquals, Dependency{"github.com/milvus-io/pulsar-client-go", "0.6.10", ""})
	c.Assert(deps[14], DeepEquals, Dependency{"github.com/golang/protobuf", "1.5.3", "/sources/golang/protobuf"})
	c.Assert(deps[18], DeepEquals, Dependency{"github.com/milvus-io/milvus/pkg", "0.0.0", "./pkg"})
}

func (s *DepsySuite) TestAux(c *C) {
	c.Assert(parseDependencyLine(""), DeepEquals, Dependency{})
	c.Assert(parseReplacementLine(""), DeepEquals, replacement{})
	c.Assert(getMajorVersion("12"), Equals, "12")

	var deps Dependencies
	c.Assert(addDep(deps, Dependency{}), IsNil)

	var repls []replacement
	c.Assert(addRepl(repls, replacement{}), IsNil)
}
