package depsy

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Dependency contains info about used module
type Dependency struct {
	Path    string
	Version string
	Extra   string
}

// Dependencies is slice with dependencies info
type Dependencies []Dependency

// ////////////////////////////////////////////////////////////////////////////////// //

type replacement struct {
	From Dependency
	To   Dependency
}

type replacements []replacement

// ////////////////////////////////////////////////////////////////////////////////// //

// Extract extracts data from go.mod data
func Extract(data []byte, withIndirect bool) Dependencies {
	var deps Dependencies
	var repls replacements

	var reqSection, replSection bool

	buf := bytes.NewBuffer(data)

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		line = strings.Trim(line, "\n\r\t")

		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		if !withIndirect && strings.HasSuffix(line, "// indirect") {
			continue
		}

		switch {
		case line == "require (":
			reqSection = true
		case line == "replace (":
			replSection = true
		case line == ")":
			reqSection, replSection = false, false
		case strings.HasPrefix(line, "require "):
			deps = addDep(deps, parseDependencyLine(line[8:]))
		case strings.HasPrefix(line, "replace "):
			repls = addRepl(repls, parseReplacementLine(line[8:]))
		case reqSection:
			deps = addDep(deps, parseDependencyLine(line))
		case replSection:
			repls = addRepl(repls, parseReplacementLine(line))
		}
	}

REPL:
	for _, repl := range repls {
		for index, dep := range deps {
			if repl.From.Version == "" {
				if dep.Path == repl.From.Path {
					deps[index] = repl.To
					continue REPL
				}
			} else {
				if dep.String() == repl.From.String() {
					deps[index] = repl.To
					continue REPL
				}
			}
		}
	}

	return deps
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string representation of dependency
func (d Dependency) String() string {
	if d.Extra == "" {
		return d.Path + ":" + d.Version
	}

	return d.Path + ":" + d.Version + "+" + d.Extra
}

// ////////////////////////////////////////////////////////////////////////////////// //

// addDep appends dep to slice with dependencies if not empty
func addDep(deps Dependencies, dep Dependency) Dependencies {
	if dep.Path == "" {
		return deps
	}

	return append(deps, dep)
}

// addRepl appends replacement to slice with replacements if not empty
func addRepl(repls replacements, repl replacement) replacements {
	if repl.To.Path == "" || repl.From.Path == "" {
		return repls
	}

	return append(repls, repl)
}

// parseDependencyLine parses line from go.mod with info about module
func parseDependencyLine(data string) Dependency {
	info := strings.Fields(data)

	if len(info) < 2 {
		return Dependency{}
	}

	path := getField(info, 0)
	version, extra := parseVersion(getField(info, 1))
	majorVersion := getMajorVersion(version)

	if strings.HasSuffix(path, "/v"+majorVersion) {
		path = path[:len(path)-(len(majorVersion)+2)]
	}

	return Dependency{
		Path:    path,
		Version: version,
		Extra:   extra,
	}
}

// parseReplacementLine parses line from go.mod with info module replacement
func parseReplacementLine(data string) replacement {
	info := strings.Fields(data)

	if len(info) < 3 {
		return replacement{}
	}

	var fromPath, fromVer, fromExtra string
	var toPath, toVer, toExtra string

	fromPath = getField(info, 0)

	switch getField(info, 1) {
	case "=>":
		toPath = getField(info, 2)
		toVer = getField(info, 3)
	default:
		fromVer = getField(info, 1)
		toPath = getField(info, 3)
		toVer = getField(info, 4)
	}

	if toPath[0] == '.' {
		return replacement{}
	}

	fromVer, fromExtra = parseVersion(fromVer)
	toVer, toExtra = parseVersion(toVer)
	majorVersion := getMajorVersion(toVer)

	if strings.HasSuffix(toPath, "/v"+majorVersion) {
		toPath = toPath[:len(toPath)-(len(majorVersion)+2)]
	}

	return replacement{
		From: Dependency{fromPath, fromVer, fromExtra},
		To:   Dependency{toPath, toVer, toExtra},
	}
}

// parseVersion parses version info
func parseVersion(version string) (string, string) {
	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}

	version = strings.ReplaceAll(version, "+incompatible", "")

	if strings.Contains(version, "-") {
		index := strings.Index(version, "-")
		return version[:index], version[index+1:]
	}

	return version, ""
}

// getMajorVersion returns major version for semver string
func getMajorVersion(v string) string {
	if !strings.Contains(v, ".") {
		return v
	}

	return v[:strings.Index(v, ".")]
}

// getField returns item with given index from slice
func getField(data []string, index int) string {
	if index >= len(data) {
		return ""
	}

	return data[index]
}
