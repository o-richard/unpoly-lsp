The HTTP method to use for the request.

Common values are `get`, `post`, `put`, `patch` and `delete`. The value is case insensitive.

The HTTP method may also be passed as an `[data-method]` attribute.

By default, methods other than `get` or `post` will be converted into a `post` request, and carry
their original method as a configurable [`_method` parameter](https://unpoly.com/up.protocol.config#config.methodParam).
