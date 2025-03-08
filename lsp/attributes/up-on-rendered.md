A JavaScript snippet that is executed when Unpoly has updated fragments.

The snippet runs in the following scope:

| Expression | Value                                                |
|------------|------------------------------------------------------|
| `this`     | The link being followed                              |
| `result`   | The `up.RenderResult` for the respective render pass |

The snippet will be called zero, one or two times:

- When the server rendered an [empty response](https://unpoly.com/skipping-rendering#rendering-nothing), no fragments are updated. `[up-on-rendered]` is not called.
- When the server rendered a matching fragment, it will be updated on the page. `[up-on-rendered]` is called with the [result](https://unpoly.com/up.RenderResult).
- When [revalidation](https://unpoly.com/caching#revalidation) renders a second time, `[up-on-rendered]` is called again with the final result.

Also see [Running code after rendering](https://unpoly.com/render-lifecycle#running-code-after-rendering).
