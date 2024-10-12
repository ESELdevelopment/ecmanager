package start_test

import (
	"github.com/ESELDevelopment/ecmanager/internal/pages/start"
	"github.com/ESELDevelopment/ecmanager/testdata"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPageWithNavigation(t *testing.T) {
	m := start.New(testdata.MockRouter{})
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(100, 100))

	testdata.WaitForString(t, tm, " Loading...")

	tm.Type("s")
	testdata.WaitForString(t, tm, "Mock Navigated")
	assert.NoError(t, tm.Quit())
}

func TestPageWithQuit(t *testing.T) {
	m := start.New(testdata.MockRouter{})
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(100, 100))

	testdata.WaitForString(t, tm, " Loading...")

	tm.Type("q")
	tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))
}
