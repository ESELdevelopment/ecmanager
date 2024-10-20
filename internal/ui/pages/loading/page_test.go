package loading_test

import (
	"github.com/ESELDevelopment/ecmanager/internal/ui/pages/loading"
	"github.com/ESELDevelopment/ecmanager/testdata"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPageWithNavigation(t *testing.T) {
	m := loading.New(testdata.MockRouter{})
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(100, 100))

	testdata.WaitForString(t, tm, " Loading...")

	tm.Type("s")
	testdata.WaitForString(t, tm, "Mock Navigated")
	assert.NoError(t, tm.Quit())
}

func TestPageWithQuit(t *testing.T) {
	m := loading.New(testdata.MockRouter{})
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(100, 100))

	testdata.WaitForString(t, tm, " Loading...")

	tm.Type("q")
	tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))
}
