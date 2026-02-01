package main

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

// TODO
