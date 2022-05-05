package depsy

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
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

// ////////////////////////////////////////////////////////////////////////////////// //

// Extract extracts data from go.mod data
func Extract(data []byte, withIndirect bool) []Dependency {
	var result []Dependency
	var replaceSection bool

	buf := bytes.NewBuffer(data)

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			break
		}

		line = strings.Trim(line, "\n\r\t")

		if !replaceSection {
			if line == "require (" {
				replaceSection = true
				continue
			} else if strings.HasPrefix(line, "require ") {
				line = line[8:]
			} else {
				continue
			}
		} else {
			if line == ")" {
				replaceSection = false
				continue
			}
		}

		if !withIndirect && strings.HasSuffix(line, "// indirect") {
			continue
		}

		result = append(result, parseDependencyLine(line))
	}

	return result
}

// ////////////////////////////////////////////////////////////////////////////////// //

// parseDependencyLine parses line from go.mod with info about module
func parseDependencyLine(data string) Dependency {
	info := strings.Fields(data)

	if len(info) < 2 {
		return Dependency{}
	}

	var path, version, extra string

	path = info[0]
	version = info[1]

	if strings.HasPrefix(version, "v") {
		version = version[1:]
	}

	version = strings.ReplaceAll(version, "+incompatible", "")

	if strings.Contains(version, "-") {
		index := strings.Index(version, "-")
		extra = version[index+1:]
		version = version[:index]
	}

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

// getMajorVersion returns major version for semver string
func getMajorVersion(v string) string {
	if !strings.Contains(v, ".") {
		return v
	}

	return v[:strings.Index(v, ".")]
}
