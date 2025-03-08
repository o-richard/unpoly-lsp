❗[EXPERIMENTAL] Sets the time when the fragment's underlying data was last changed.

When the fragment is reloaded, its known modification time is sent as an `If-Modified-Since` request header.
The server may check the header and decide to [skip rendering](https://unpoly.com/skipping-rendering).
See [Conditional requests](https://unpoly.com/conditional-requests) for a full example.

The value can either be a Unix timestamp (e.g. `"1445412480"`) or an [RFC 1123](https://www.rfc-editor.org/rfc/rfc1123) time (e.g. `Wed, 21 Oct 2015 07:28:00 GMT`). You can also set the value to `"false"` to prevent a `If-Modified-Since` request header when reloading this fragment.

### How `[up-time]` attributes are set

Unpoly will automatically set an `[up-time]` attribute when a fragment was rendered from a response with a `Last-Modified` header. 
When a fragment was rendered without such a header, Unpoly will set `[up-time=false]` to indicate that its modification time is unknown.

A large response may contain multiple fragments that are later reloaded individually
and should each have their own modification time. In this case the server may also also render multiple
fragments with each their own `[up-time]` attribute.
See [Individual versions per fragment](https://unpoly.com/conditional-requests#fragment-versions) for an example.
