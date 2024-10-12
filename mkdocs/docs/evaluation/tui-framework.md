# TUI-Framework

We want to use a framework to build a TUI for our application. We will
evaluate the most popular TUI frameworks for Go and choose the one
that best fits our needs.

## Frameworks

There are many frameworks for building TUI in GO but the main ones
are **Tview** and **Bubbletea**. They have the most documentation
and biggest user bases so we will evaluate these two.

## Tview

[Tview](https://github.com/rivo/tview) is a fast, flexible UI framework
for building rich terminal applications in Golang. It offers widgets like tables,
forms, and modals, making interactive console apps easy to create.
Built on `tcell`, it ensures cross-platform compatibility while handling
terminal operations. It works best for UIs with less dynamic updates.

The greates project using tview is [k9s](https://github.com/derailed/k9s)

## Bubbletea

[Bubble Tea](https://github.com/charmbracelet/bubbletea) is a flexible,
reactive framework for building terminal applications in Golang using the
Model-Update-View (MVU) architecture. It allows developers to create
highly interactive and dynamic UIs by handling asynchronous events and
updates seamlessly. Unlike traditional imperative UI frameworks,
Bubble Tea promotes a declarative style, where the UI is a function of
the application's state. This makes it more flexible for building applications
with complex interactions and real-time updates.

As part of the charmbracelet suite, Bubble Tea is well-documented and
actively maintained. It is used in popular projects like
[Superfile](https://github.com/yorukot/superfile). Additionally,
Bubble Tea has also great plugins for
styling like [lipgloss](https://github.com/charmbracelet/lipgloss)
and [glamour](https://github.com/charmbracelet/glamour).

## Comparison

<!-- markdownlint-disable MD013 -->
| Feature       | Tview                      | Bubble Tea                                                                         |
|---------------|----------------------------|------------------------------------------------------------------------------------|
| Architecture  | :orange_circle: Imperative | :green_circle: Declarative                                                         |
| Dynamic UI    | :orange_circle: Limited    | :green_circle: Flexible                                                            |
| Documentation | :green_circle:             | :green_circle:                                                                     |
| Community     | :green_circle:             | :green_circle:                                                                     |
| Templates     | :green_circle: built-in    | :green_circle: extensions like [bubbles](https://github.com/charmbracelet/bubbles) |
<!-- markdownlint-enable MD013 -->

## Conclusion

{{ decision("We use bubble tea with lipgloss and glamour for styling") }}

With the declarative architecture, simplicity and flexibility of Bubble Tea,
we can build a highly interactive and dynamic TUI for our application.
Additionally, the availability of plugins like lipgloss and glamour allows
us to style our UI easily.
