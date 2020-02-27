package termui

import (
	"strings"
)

var PRINTABLE_KEYS = append(
	strings.Split(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,./<>?;:'\"[]\\{}|`~!@#$%^&*()-_=+",
		"",
	),
	"<Space>",
	"<Tab>",
	"<Enter>",
)
