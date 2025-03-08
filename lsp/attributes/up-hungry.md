Elements with an `[up-hungry]` attribute are updated whenever the server sends a matching element, even if the element isn't [targeted](https://unpoly.com/targeting-fragments).

Hungry elements are optional in the same way as `:maybe` targets.
When an `[up-hungry]` element does not match in the server response, the element will not be updated,
but no error is thrown.

### Use cases

Common use cases for `[up-hungry]` are elements that live in the application layout,
outside of the fragment that is typically being targeted. Examples include:

- Unread message counters
- Page-specific subnavigation
- Account-wide notifications (e.g. about an expired credit card)

Instead of explicitly including such elements in every [target selector](https://unpoly.com/targeting-fragments)
(e.g. `.content, .unread-messages:maybe`) we can mark the element as `[up-hungry]`:

```html
<div class="unread-messages" up-hungry>
You have no unread messages
</div>
```

An selector for the hungry element (`.unread-messages`) will be added to target selectors automatically.

### Derivable target required {#derivable-target-required}

When an `[up-hungry]` fragment piggy-backs on another fragment update, Unpoly will [derive a target selector](https://unpoly.com/target-derivation) for the hungry element.

For this to work the hungry element must have an [identifying attribute](https://unpoly.com/target-derivation#derivation-patterns), like an `[id]` or a unique `[class]` attribute.
When no good target can be derived, the hungry element is excluded from the update.

### Behavior with multiple layers

By default only hungry elements on the targeted [layer](https://unpoly.com/up.layer) are updated.

To match a hungry element when updating other layers, set an [`[up-if-layer]`](#up-if-layer) attribute.
For example, a hungry element with `[up-if-layer="subtree"]` will piggy-back on render passes for both
its own layer and any overlay covering it.

### Conflict resolution

When Unpoly renders new content, each element in that content can only be inserted once.
When multiple hungry elements conflict with each other or with the the [primary render target](https://unpoly.com/targeting-fragments), that conflict is resolved using the following rules:

1. When both a [target selector](https://unpoly.com/targeting-fragments) and a hungry elements target the same fragment in the response, only the direct render target will be updated.
2. When hungry elements are nested within each other, the outmost fragment will be updated. Note that we recommend to not over-use the hungry mechanism, and prefer to explicit render targets instead.
3. When hungry elements on different layers target the same fragment in the response, the layer closest to the rendering layer will be chosen.

### Disabling

By default hungry fragments are processed for all updates of the current layer.
You can disable the processing of hungry fragments using one of the following methods:

- Rendering with an [`{ hungry: false }`](https://unpoly.com/up.render#options.hungry) option will not process any hungry fragments.
- Setting an [`[up-use-hungry="false"]`](https://unpoly.com/up-follow#up-use-hungry) attribute on a link or form will not update hungry fragments when the element is activated.
- Preventing an `up:fragment:hungry` event will prevent the hungry fragment from being updated.
- Calling `event.preventDefault()` in an `[up-on-hungry]` attribute handler will prevent the hungry fragment from being updated.

