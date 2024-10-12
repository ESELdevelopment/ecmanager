package pages_test

import (
  "github.com/ESELDevelopment/ecmanager/internal/pages"
  "github.com/charmbracelet/x/exp/teatest"
  "github.com/stretchr/testify/assert"
  "strings"
  "testing"
  "time"
)

func TestExamplePage(t *testing.T) {
  m := pages.NewExamplePage(pages.Router())
  tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(300, 100))

  waitForString(t, tm, "This is the example screen. Press any key to switch to the second screen.")

  tm.Type("n")

  waitForString(t, tm, "ECManager")
  assert.NoError(t, tm.Quit())
}

func waitForString(t *testing.T, tm *teatest.TestModel, s string) {
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
