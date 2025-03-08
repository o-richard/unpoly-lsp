Whether to read from and write to the [cache](https://unpoly.com/caching).

With `[up-cache=true]` Unpoly will try to re-use a cached response before connecting
to the network. To prevent display of stale content, cached responses are
[reloaded once rendered](#up-revalidate). If no cached response exists,
Unpoly will make a request and cache the server response.

With `[up-cache=auto]` Unpoly will use the cache only if `up.network.config.autoCache`
returns `true` for the request. By default this only caches `GET` requests.

With `[up-cache=false]` Unpoly will always make a network request.