When to [snap to the top](https://unpoly.com/scroll-tuning#snapping-to-the-screen-edge) when scrolling to an element near the top edge of the viewport's scroll buffer.

When [revealing](https://unpoly.com/up.reveal) an element near the top edge of a viewport's scroll buffer, you often want to scroll to the very top for aesthetic reasons. 
For example, if you reveal a navigation bar that sits 50px below the top logo, you probably want to scroll to zero (instead of 50) pixels.

In order to snap to the top edge in such cases, pass a `{ revealSnap }` option.
When the the revealed element would be closer to the viewport's top edge than this value, Unpoly will scroll the viewport to the top.

The default is `{ revealSnap: 200 }`.
You may change this default in `up.viewport.config.revealSnap`.

To disable snapping, use `{ revealSnap: 0 }`.