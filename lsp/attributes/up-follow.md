[Follows](https://unpoly.com/up.follow) this link with JavaScript and updates a fragment with the server response.

Following a link is considered [navigation](https://unpoly.com/navigation) by default.

## Example

This link will update the fragment `<div class="content">` with the same element
fetched from `/posts/5`:

```html
<a href="/posts/5" up-follow up-target=".content">Read post</a>
```

If no `[up-target]` attribute is set, the [main target](https://unpoly.com/up-main) is updated.

## Following all links automatically

You can configure Unpoly to follow *all* links on a page without requiring an `[up-follow]` attribute.

See [Handling all links and forms](https://unpoly.com/handling-everything).

## Preventing Unpoly from following links

You can tell Unpoly to ignore clicks on an `[up-follow]` link, causing the link to be non-interactive.
Use one of the following methods:

- Prevent the `up:link:follow` event on the link element
- Prevent the `up:click` event on the link element

To force a [full page load](https://unpoly.com/up.network.loadPage) when a followable link is clicked:

- Set an [`[up-follow=false]`](https://unpoly.com/attributes-and-options#boolean-attributes) attribute on the link element
- Prevent the `up:link:follow` event and call `up.network.loadPage(event.renderOptions)`.

## Making non-interactive elements act as hyperlinks

You can set an `[up-follow]` attribute on any non-interactive element to make it behave like a hyperlink:

```html
<span up-follow up-href="/details">Read more</span>
```

See [Acting like a hyperlink](https://unpoly.com/faux-interactive-elements) for details.

## Advanced fragment changes

Links can update multiple fragments or append content to an existing element.

See [Fragment placement](https://unpoly.com/targeting-fragments) for details.

## Short notation

You may omit the `[up-follow]` attribute if the link has one of the following attributes:

- `[up-target]`
- `[up-layer]`
- `[up-transition]`
- `[up-content]`
- `[up-fragment]`
- `[up-document]`

Such a link will still be followed through Unpoly.
