package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	info0, _ := parse(`fix(scope): title
title next line

description desc1
desc2
desc3

Log: log1
Bug: url1
Bug: url2
`)
	assert.Equal(t, &info{
		typ:   "fix",
		scope: "scope",
		title: "title\ntitle next line\n",
		desc:  "description desc1\ndesc2\ndesc3\n",
		lines: []line{
			{
				type0:   "Log",
				content: "log1",
			},
			{
				type0:   "Bug",
				content: "url1",
			},
			{
				type0:   "Bug",
				content: "url2",
			},
		},
	}, info0)
}

func TestParseTitle(t *testing.T) {
	typ, scope, title := parseTitle("fix: title")
	assert.Equal(t, "fix", typ)
	assert.Equal(t, "", scope)
	assert.Equal(t, "title", title)

	typ, scope, title = parseTitle("fix(): title")
	assert.Equal(t, "fix", typ)
	assert.Equal(t, "", scope)
	assert.Equal(t, "title", title)

	typ, scope, title = parseTitle("fix (): title")
	assert.Equal(t, "fix", typ)
	assert.Equal(t, "", scope)
	assert.Equal(t, "title", title)

	typ, scope, title = parseTitle("fix(scope): title")
	assert.Equal(t, "fix", typ)
	assert.Equal(t, "scope", scope)
	assert.Equal(t, "title", title)

	typ, scope, title = parseTitle(" fix : title")
	assert.Equal(t, "fix", typ)
	assert.Equal(t, "", scope)
	assert.Equal(t, "title", title)

	typ, scope, title = parseTitle(": title")
	assert.Equal(t, "", typ)
	assert.Equal(t, "", scope)
	assert.Equal(t, "title", title)

	typ, scope, title = parseTitle("title")
	assert.Equal(t, "", typ)
	assert.Equal(t, "", scope)
	assert.Equal(t, "title", title)
}

func TestParseKV(t *testing.T) {
	kv, ok := parseKV("Bug: url1")
	assert.True(t, ok)
	assert.Equal(t, line{
		type0:   "Bug",
		content: "url1",
	}, kv)

	kv, ok = parseKV(": value")
	assert.False(t, ok)

	kv, ok = parseKV("key:")
	assert.False(t, ok)

	kv, ok = parseKV(":")
	assert.False(t, ok)

	kv, ok = parseKV("")
	assert.False(t, ok)
}
