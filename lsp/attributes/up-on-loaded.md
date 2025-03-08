A JavaScript snippet that is executed when the server responds with new HTML, but before the HTML is rendered.

The snippet runs in the following scope:

| Expression | Value                                         |
|------------|-----------------------------------------------|
| `this`     | The link being followed                       |
| `event`    | A preventable `up:fragment:loaded` event      |

The snippet will also run for [failed responses](https://unpoly.com/failed-responses).