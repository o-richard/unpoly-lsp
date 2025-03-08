Whether existing [cache](https://unpoly.com/caching) entries will be [expired](https://unpoly.com/caching#expiration) with this request.

By default a non-GET request will expire the entire cache.
You may also pass a [URL pattern](https://unpoly.com/url-patterns) to only expire matching requests.

Also see [`up.request({ expireCache })`](https://unpoly.com/up.request#options.expireCache) and `up.network.config.expireCache`.
