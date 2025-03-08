A JavaScript snippet that is called when the overlay was [accepted](https://unpoly.com/closing-overlays).

The snippet runs in the following scope:

| Expression | Value                                         |
|------------|-----------------------------------------------|
| `this`     | The link that originally opened the overlay   |
| `layer`    | An `up.Layer` object for the accepted overlay |
| `value`    | The overlay's [acceptance value](https://unpoly.com/closing-overlays#overlay-result-values) |
| `response` | The server response that caused the overlay to close |
| `event`    | An `up:layer:accepted` event                  |
