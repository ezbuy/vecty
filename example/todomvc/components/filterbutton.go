package components

import (
	"github.com/ezbuy/vecty"
	"github.com/ezbuy/vecty/elem"
	"github.com/ezbuy/vecty/event"
	"github.com/ezbuy/vecty/example/todomvc/actions"
	"github.com/ezbuy/vecty/example/todomvc/dispatcher"
	"github.com/ezbuy/vecty/example/todomvc/store"
	"github.com/ezbuy/vecty/example/todomvc/store/model"
	"github.com/ezbuy/vecty/prop"
)

// FilterButton is a vecty.Component which allows the user to select a filter
// state.
type FilterButton struct {
	vecty.Core

	Label  string            `vecty:"prop"`
	Filter model.FilterState `vecty:"prop"`
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

// Render implements the vecty.Component interface.
func (b *FilterButton) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.Markup(
				vecty.MarkupIf(store.Filter == b.Filter, vecty.Class("selected")),
				prop.Href("#"),
				event.Click(b.onClick).PreventDefault(),
			),

			vecty.Text(b.Label),
		),
	)
}
