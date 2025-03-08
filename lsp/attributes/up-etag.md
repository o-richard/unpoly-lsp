❗[EXPERIMENTAL] Sets an [ETag](https://en.wikipedia.org/wiki/HTTP_ETag) for the fragment's underlying data.

ETags can be used to skip unnecessary rendering of unchanged content. See [Conditional requests](https://unpoly.com/conditional-requests) for a full example.

### How `[up-etag]` attributes are set

Unpoly will automatically set an `[up-etag]` attribute when a fragment was rendered from a response with a `ETag` header. When a fragment was rendered without such a header, Unpoly will set `[up-etag=false]` to indicate that its ETag is unknown.

A large response may contain multiple fragments that are later reloaded individually and should each have their own ETag. In this case the server may also also render multiple fragments with each their own `[up-etag]` attribute.
See [Individual versions per fragment](https://unpoly.com/conditional-requests#fragment-versions) for an example.
