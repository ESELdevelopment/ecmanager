package testdata

import (
	"github.com/charmbracelet/x/exp/teatest"
	"strings"
	"testing"
	"time"
)

func WaitForString(t *testing.T, tm *teatest.TestModel, s string) {
	teatest.WaitFor(
		t,
		tm.Output(),
		func(b []byte) bool {
			return strings.Contains(string(b), s)
		},
		teatest.WithCheckInterval(time.Millisecond*100),
		teatest.WithDuration(time.Second*1),
	)
}
