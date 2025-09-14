Follows this link on `mousedown` instead of `click` ("Act on press").

This will save precious milliseconds that otherwise spent on waiting for the user to release the mouse button. Since an AJAX request will be triggered right way, the interaction will appear faster.

Links with the `[up-instant]` attribute are always [followed by Unpoly](https://unpoly.com/up-follow) and will not make a full page load.

To [follow all links on `mousedown`](https://unpoly.com/handling-everything#following-all-links-on-mousedown), configure `up.link.config.instantSelectors`.

### Example

```html
<a href="/users" up-follow up-instant>User list</a> <!-- mark-phrase "up-instant" -->
```

### Accessibility

Links with `[up-instant]` can still be activated with the keyboard.

With `[up-instant]` users can no longer cancel a click by dragging the pressed mouse away from the link.
However, for navigation actions this isn't required. E.g. many operation systems switch tabs on `mousedown` instead of `click`.