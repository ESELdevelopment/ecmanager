package example_test

import (
	"github.com/ESELDevelopment/ecmanager/internal/ui/pages/example"
	"github.com/ESELDevelopment/ecmanager/testdata"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExamplePage(t *testing.T) {
	router := testdata.MockRouter{}
	m := example.New(router)
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(300, 100))

	testdata.WaitForString(t, tm, "This is the example screen. Press any key to switch to the second screen.")

	tm.Type("n")

	testdata.WaitForString(t, tm, "Mock Navigated")
	assert.NoError(t, tm.Quit())
}
