package main

import (
	"regexp"
	"strings"
)

// title
// next title
// <space>
// desc dasdfasdf
//  <space>
// Log:
// Task:
// Back:
// Change-Id:

const (
	stepTitle int = iota + 1
	stepDesc
	stepKeyValuePairs
)

var lineBreak = "\n"

func parse(str string) (*info, error) {
	var info info
	var step int
	var title string
	lines := strings.Split(str, lineBreak)
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			if step == stepTitle {
				step++ // to desc
			} else if step == stepDesc {
				step++ // to keyValuePairs
			}

		} else {
			if step == 0 {
				step = stepTitle
				typ, scope, line := parseTitle(line)
				info.typ = typ
				info.scope = scope
				title = line + lineBreak
			} else if step == stepTitle {
				title += line + lineBreak
			} else if step == stepDesc {
				info.desc += line + lineBreak
			} else if step == stepKeyValuePairs {
				kv, ok := parseKV(line)
				if ok {
					info.lines = append(info.lines, kv)
				}
			}
		}
	}
	info.title = title
	return &info, nil
}

func parseTitle(str string) (typ string, scope string, title string) {
	fields := strings.SplitN(str, ":", 2)
	if len(fields) == 1 {
		title = strings.TrimSpace(fields[0])
		return
	} else if len(fields) == 0 {
		return
	}
	// now len(fields) == 2
	prefix := strings.TrimSpace(fields[0])
	reg := regexp.MustCompile(`^(\w+)\s*(\(([^()]*)\))?$`)
	// fix
	// fix(scope)
	// fix ()
	// fix  (scope)
	// fix  ( scope )
	match := reg.FindStringSubmatch(prefix)
	if match != nil {
		typ = match[1]
		scope = strings.TrimSpace(match[3])
	}
	title = strings.TrimSpace(fields[1])
	return
}

func parseKV(str string) (line, bool) {
	fields := strings.SplitN(str, ":", 2)
	if len(fields) != 2 {
		return line{}, false
	}
	key := strings.TrimSpace(fields[0])
	value := strings.TrimSpace(fields[1])
	if key == "" || value == "" {
		return line{}, false
	}
	return line{type0: key, content: value}, true
}

type info struct {
	typ   string
	scope string
	title string
	desc  string
	lines []line
}
