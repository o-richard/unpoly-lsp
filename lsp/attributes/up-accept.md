[Accepts](https://unpoly.com/closing-overlays) the [current layer](https://unpoly.com/up.layer.current) when the link is clicked.

The JSON value of the `[up-accept]` attribute becomes the overlay's [acceptance value](https://unpoly.com/closing-overlays#overlay-result-values).

### Example

```html
<a href="/users/5" up-accept="{ id: 5 }">Choose user #5</a>
```

### Fallback for the root layer

The link's `[href]` will only be followed when this link is clicked in the [root layer](https://unpoly.com/up.layer).
In an overlay the `click` event's default action is prevented.

You can also omit the `[href]` attribute to make a link that only works in overlays.
