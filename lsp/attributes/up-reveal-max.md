How many pixel lines of [high element to reveal](https://unpoly.com/scroll-tuning#revealing-with-padding) when scrolling to an element.

When [revealing](https://unpoly.com/up.reveal) an element, the viewport will scroll as far as necessary to make the element visible.

For an element higher than the viewport, this would mean that the top edge of the element would be moved to the top of the viewport. This large change of scroll positions may disorient the user.

To limit the scroll motion when revealing an element, pass a pixel value as `{ revealMax }`. Unpoly will then pretend that the element is no higher than the given value. You may also pass a function that accepts the element and returns a pixel value.

The default is `{ revealMax: () => 0.5 * innerHeight }`, meaning that Unpoly will reveal high elements until half the screen height is filled. You may change this default in `up.viewport.config.revealMax`.