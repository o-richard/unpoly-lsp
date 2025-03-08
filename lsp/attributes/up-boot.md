❗[EXPERIMENTAL] Prevent Unpoly from booting automatically.

By default Unpoly [automatically boots](/install#initialization)con [`DOMContentLoaded`](https://developer.mozilla.org/en-US/docs/Web/API/Window/DOMContentLoaded_event).
To prevent this, add an `[up-boot="manual"]` attribute to the `<script>` element that loads Unpoly:

```html
<script src="unpoly.js" up-boot="manual"></script>
```
You may then call `up.boot()` to manually boot Unpoly at a later time.