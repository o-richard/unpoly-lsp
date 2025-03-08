A JavaScript snippet that is run when any error is thrown during the rendering process.

| Expression | Value                                         |
|------------|-----------------------------------------------|
| `this`     | The link being followed                       |
| `error`    | An `Error` object                             |

The callback is also called when the render pass fails due to [network issues](https://unpoly.com/network-issues), or [aborts](https://unpoly.com/aborting-requests).

Also see [Handling errors](https://unpoly.com/render-lifecycle#handling-errors).
