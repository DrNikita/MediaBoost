package main

import (
	"fmt"
)

type Error int

const (
	Unknown Error = -1
)

func (e Error) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e, e.Title(), e.Description())
}

func (e Error) Title() string {
	switch e {
	case Unknown:
		return "Unknown"
	}

	return ""
}

func (e Error) Description() string {
	switch e {
	case Unknown:
		return "an unexpected server error occurred"
	}

	return ""
}
