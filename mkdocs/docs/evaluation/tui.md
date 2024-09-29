# TUI

This is an Evaluation for a TUI (Terminal User Interface) Library in GO.

There are many Libarys for Buildung TUI in GO but the main ones are Tview and Bubbletea.
They have the most documentation and biggest user bases so we will only evaluati these two.

## Tview

Widget-based Imperative UI: Built around a traditional imperative programming model, tview offers predefined widgets such as tables, forms, lists, and trees, which are easy to use and integrate into an application.
Efficiency: It is particularly useful for building structured and static UIs that do not require a high level of interactivity or reactivity. Its widgets are simple to implement, making it a great choice for dashboards, editors, or tools that focus on data presentation rather than interaction.
Lower Learning Curve: tview’s design is familiar to developers who have worked with widget-based toolkits like GTK or Qt, making it easier to learn and faster to implement.
Simplicity over Flexibility: While efficient, tview is not as flexible as Bubble Tea when it comes to handling asynchronous events or creating custom UI components. It works best for UIs with less dynamic updates.
Best for: Applications that require static layouts with minimal interactivity, such as terminal dashboards, text editors, or simple data browsers.

## Bubbletea

Declarative UI: Ideal for projects that require a declarative approach to building user interfaces, similar to frameworks like React or Elm. The UI is defined based on the state, which ensures that the view is always synchronized with the underlying state.
Elm Architecture: Follows the Model-Update-View (MUV) pattern, offering reactive state management where the state drives the UI. This is excellent for handling asynchronous events, like background processes, network requests, or timed actions.
Customization & Flexibility: Well-suited for applications with highly dynamic UIs that need custom rendering or complex behavior, such as real-time dashboards, interactive CLIs, or text-based games.
Learning Curve: The architecture requires understanding of functional programming principles, which might be unfamiliar to some Go developers, but it offers powerful abstractions once mastered.
Best for: Projects requiring real-time interactivity, asynchronous handling, or highly customized UIs where flexibility is key.
