[Preloads](https://unpoly.com/preloading) this link before the user clicks it.
When the link is clicked, the response will already be [cached](https://unpoly.com/caching), making the interaction feel instant.

Unpoly will only preload [links with safe methods](https://unpoly.com/up.link.isSafe). The `[up-preload]` attribute has no effect on unsafe links.

Links with the `[up-preload]` attribute are always [followed by Unpoly](https://unpoly.com/up-follow) and will not make a full page load.

When set to `'hover'` (the default), preloading will start when the user hovers over this link [for a while](#up-preload-delay). 
On touch devices preloading will begin when the user places her finger on the link. Also see [preloading on hover](https://unpoly.com/preloading#on-hover).

When set to `'insert'`, preloading will start immediatedly when this link is inserted into the DOM. Also see [eagerly preloading on insertion](https://unpoly.com/preloading#on-insert).

When set to `'reveal'`, preloading will start when the link is scrolled into the [viewport](https://unpoly.com/up-viewport). 
If the link is already visible when inserted, preloading will start immediately.  Also see [preloading when a link becomes visible](https://unpoly.com/preloading#on-reveal).

See [preloading links](https://unpoly.com/preloading) for more details and examples.

### Example

```html
<a href="/path" up-preload>Hover over me to preload my content</a>
```