Elements with an `[up-keep]` attribute will be persisted during [fragment updates](https://unpoly.com/up.fragment).

Common use cases for `[up-keep]` include:

- Elements that are expensive to [initialize](https://unpoly.com/up.compiler).
- Media elements (`<video>`, `<audio>`) that should retain their playback state during updates.
- Other elements with client-side state that is difficult to express in a URL or [data object](https://unpoly.com/data).

The element must have a [derivable target selector](https://unpoly.com/target-derivation) so Unpoly can find its position within new content.

Emits the [`up:fragment:keep`](https://unpoly.com/up:fragment:keep) event.

## Example

A common use case is to preserve the playback state of media elements:

```html
<article>
<p>Content</p>
<audio id="player" up-keep src="song.mp3"></audio>
</article>
```

When [targeting](https://unpoly.com/targeting-fragments) the `<article>` fragment, the `<audio>` element and
its playback state will be the same before and after the update. All other elements (like the `<p>`)
will be updated with new content.

## Controlling if an element will be kept

Unpoly will **only** keep an existing element if:

- The existing element has an `[up-keep]` attribute
- The response contains an element matching the [derived target](https://unpoly.com/target-derivation) of the existing element

The element has multiple methods to veto against being kept:

- By setting a `[up-keep=false]` attribute on the new element version.
- By setting a different `[id]` or `[up-id]` attribute so its [derived target](https://unpoly.com/target-derivation) no longer matches the existing element.
- By preventing the [`up:fragment:keep`](https://unpoly.com/up:fragment:keep) event that is [emitted](https://unpoly.com/up.emit) on the existing element.
- By preventing the [`up:fragment:keep`](https://unpoly.com/up:fragment:keep) event that is passed to an [`[up-on-keep]`](#up-on-keep)
callback on the element.

You can also choose to render without keeping elements:

- Link or forms can force a swap of `[up-keep]` elements by setting an [`[up-use-keep=false]`](https://unpoly.com/up-follow#up-use-keep) attribute.
- Rendering functions can force a swap of `[up-keep]` elements by passing an [`{ keep: false }`](https://unpoly.com/up.render#options.keep) option.

### Example for conditional keeping

Let's say we want only keep an `<audio up-keep>` element as long as it plays
the same song (as identified by the tag's `src` attribute).

On the client we can achieve this by listening to an `up:keep:fragment` event
and preventing it if the `src` attribute of the old and new element differ:

```js
up.on('up:fragment:keep', 'audio', function(event) {
if (element.getAttribute('src') !== event.newElement.getAttribute('src')) {
    event.preventDefault()
}
})
```

## Updating data for kept elements

Even when keeping elements, you may reconcile its [data object](https://unpoly.com/data) with the data
from the new element that was discarded.

Let's say you want to display a map within an element. The center of the map
is encoded using an `[up-data]` attribute:

```html
<div id="map" up-keep up-data="{ lat: 50.86, lng: 7.40 }"></div>
```

We can initialize the map using a [compiler](https://unpoly.com/up.compiler) like this:

```js
up.compiler('#map', function(element, data) {
var map = new google.maps.Map(element)
map.setCenter(data)
})
```

While we want to preserve the map during page loads, we *do* want to pick up a new center coordinate when the containing fragment is updated. We can do so by listening to an `up:fragment:keep` event and observing `event.newData`:

```js
up.compiler('#map', function(element, data) {
var map = new google.maps.Map(element)
map.setCenter(data)

map.addEventListener('up:fragment:keep', function(event) { // mark-line
    map.setCenter(event.newData) // mark-line
}) // mark-line
})
```

> [TIP]
> Instead of keeping an element and update its data you may also
> [preserve an element's data through reloads](https://unpoly.com/data#preserving).

## Limitations

- The `[up-keep]` attribute is only supported for elements within the `<body>`.
- If an `<audio up-keep>` or `<video up-keep>` element is a *direct* child of the `<body>`,
it will lose its playback state during a fragment update. To preserve its playback
state, insert a container element between the `<body>` and the media element.
