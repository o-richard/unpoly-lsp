Emits a custom event when this element is clicked.

The event is emitted on this element and bubbles up the `document`.
To listen to the event, use [`addEventListener()`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener) or `up.on()` on the element, or on its ancestors.

While the `[up-emit]` attribute is often used with an `<a>` or `<button>` element,
you can also apply to to non-interactive elements, like a `<span>`.
See [clicking on non-interactive elements](https://unpoly.com/faux-interactive-elements) for details and
accessibility considerations.

### Example

This button will emit a `user:select` event when pressed:

```html
<button type="button" up-emit='user:select'>Alice</button>
```

The event can be handled by a listener:

```js
document.addEventListener('user:select', function(event) {
up.reload('#user-details')
})
```

### Event properties

By default `[up-emit]` will emit an event with only basic properties like [`{ target }`](https://developer.mozilla.org/en-US/docs/Web/API/Event/target).

To set custom properties on the event object, encode them as [relaxed JSON](https://unpoly.com/relaxed-json) in an `[up-emit-props]` attribute:

```html
<button type="button"
up-emit="user:select"
up-emit-props="{ id: 5, firstName: 'Alice' }">
Alice
</button>

<script>
up.on('user:select', function(event) {
    console.log(event.id)        // logs 5
    console.log(event.firstName) // logs "Alice"
})
</script>
```

### Fallback URLs {#fallback}

Use `[up-emit]` on a link to define a fallback URL that is rendered in case no listener handles the event:

```html
<a href="/menu" up-emit='menu:open'>Menu</a>
```

When a listener has handled the `menu:open` event, it should call `event.preventDefault()`.
This also prevents the original `click` event, causing the link to no longer be followed:

```js
document.addEventListener('menu:open', function(event) {
event.preventDefault() // prevent the link from being followed
})
```

If no listener prevents the `menu:open` event, the browser will [navigate](https://unpoly.com/up-follow)
to the `/menu` path.

> [tip]
> When an [event closes an overlay](https://unpoly.com/closing-overlays#event-condition) via `[up-accept-event]` or `[up-dismiss-event]`, its default is prevented.
> You can use fallback URLs to make a link that emits a closing event in an overlay, but navigates to a different page on the [root layer](https://unpoly.com/up.layer.root).



