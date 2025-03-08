When to [move a revealed element to the top](https://unpoly.com/scroll-tuning#moving-revealed-elements-to-the-top) when scrolling to an element.

When [revealing](https://unpoly.com/up.reveal) an element, the viewport will only scroll
*as little as possible* to make the element visible. For instance, if the viewport is
already fully visible, the scroll position will not change.

To always align the top edges of the revealed element and viewport,
pass `{ revealTop: true }`.

The default is `{ revealTop: false }`.
You may change this default in `up.viewport.config.revealTop`.