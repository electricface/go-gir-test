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
	//var title string
	lines := strings.Split(str, lineBreak)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}

		if line == "" {
			if step == stepTitle && info.title != "" {
				step++ // to desc
			} else if step == stepDesc && info.desc != "" {
				step++ // to keyValuePairs
			}

		} else {
			if step == 0 {
				step = stepTitle
				typ, scope, line := parseTitle(line)
				info.typ = typ
				info.scope = scope
				info.title = line
			} else {
				kv, ok := parseKV(line)
				if ok {
					info.lines = append(info.lines, kv)
					step = stepKeyValuePairs
				} else {
					switch step {
					case stepTitle:
						if info.title != "" {
							info.title += lineBreak
						}
						info.title += line
					case stepDesc:
						if info.desc != "" {
							info.desc += lineBreak
						}
						info.desc += line
					}
				}
			}
		}
	}
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
