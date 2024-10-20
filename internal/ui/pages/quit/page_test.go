package quit_test

import (
	"github.com/ESELDevelopment/ecmanager/internal/ui/pages/quit"
	"github.com/ESELDevelopment/ecmanager/testdata"
	"github.com/charmbracelet/x/exp/teatest"
	"testing"
	"time"
)

func TestQuitPage(t *testing.T) {
	m := quit.New()
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(300, 100))

	testdata.WaitForString(t, tm, "This is the end screen. Press any key to quit.")

	tm.Type("a")

	tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))
}
