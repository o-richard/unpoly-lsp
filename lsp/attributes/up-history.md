Whether the browser URL, window title and meta tags will be [updated](https://unpoly.com/updating-history).

If set to `true`, the history will always be updated, using the title and URL from the server response, or from given `[up-title]` and `[up-location]` attributes. In layers, when the overlay is closed, the parent layer's history is restored.

If set to `auto` history will be updated if the `[up-target]` matches a selector in `up.fragment.config.autoHistoryTargets`. By default this contains all [main targets](https://unpoly.com/up-main).

If set to `false`, the history will remain unchanged. In layers, you can still access the overlay's current location using up.layer.location.
