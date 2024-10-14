# Pages

This site contains conventions about our pages.

## Where to find and create pages

All pages are located in the `internal/ui/pages` folder. To create a new page,
create a new folder with the name of the page and add a `page.go` file.

## Go Conventions

### File structure

All pages extending the `bubbletea.Model` interface. The `page.go` file
should look like that:

```go
package example

import (
  tea "github.com/charmbracelet/bubbletea"
)

// page-model, implements the bubbletea
type page struct {
}

// Static function to create the Pagemodel
func New() tea.Model {
  return page{}
}

// Implement the tea.Model interface
```

### Use Bubbles over custom components

We want to use [bubbles](https://github.com/charmbracelet/bubbles)
over custom components. This is to ensure a consistent look and feel of our
application.

### Use lipgloss for styling

We use [lipgloss](https://github.com/charmbracelet/lipgloss) for styling.
This reduces the amount of custom CSS in our View() Methods.

### Routing

The routing is done by our `internal/route.gor` service. To enable routing,
simply add the `Router` interface to your page model.

> :warning: Disclaimer: Currently our routing is very basic and only supports
> routing forward.

### Testing

To test your page, create a new test file in the same folder as your page.
We use the [teatest](https://charm.sh/blog/teatest/) for our tests. We **only**
write page-tests, which means we only test the page itself, not the components.

All actions should be triggered by user input (type or click actions).

### Useful Sources

- [Tips for Bubble Tea](https://leg100.github.io/en/posts/building-bubbletea-programs/)
- [Information about Teatest](https://charm.sh/blog/teatest/)
- [Example Pages](https://github.com/charmbracelet/bubbletea/tree/main/examples)
