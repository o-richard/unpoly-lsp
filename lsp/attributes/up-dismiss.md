[Dismisses](https://unpoly.com/closing-overlays) the [current layer](https://unpoly.com/up.layer.current) when the link is clicked.

The [relaxed JSON](https://unpoly.com/relaxed-json) value of the `[up-dismiss]` attribute becomes the overlay's [dismissal value](https://unpoly.com/closing-overlays#overlay-result-values).

### Example

```html
<a href="/dashboard" up-dismiss="'sidebar-close'">Close</a>
```

### Fallback for the root layer

The link's `[href]` will only be followed when this link is clicked in the [root layer](https://unpoly.com/up.layer).
In an overlay the `click` event's default action is prevented.

You can also [omit the `[href]` attribute](https://unpoly.com/providing-html#omitting-href) to make a link that only works in overlays.
