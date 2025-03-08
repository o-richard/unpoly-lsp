Whether the server response should be considered failed.

For failed responses Unpoly will use attributes prefixed with `up-fail`, e.g. [`[up-fail-target]`](#up-fail-target).
See [handling server errors](https://unpoly.com/failed-responses) for details.

By [default](https://unpoly.com/up.network.config#config.fail) any HTTP status code other than 2xx or [304](https://unpoly.com/skipping-rendering#rendering-nothing) is considered an error code.
Set `[up-fail=false]` to handle *any* response as successful, even with a 4xx or 5xx status code.
