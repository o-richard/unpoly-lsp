⚠️ [DEPRECIATED] Follows this link as fast as possible.

This is done by:

- [Following the link through AJAX](https://unpoly.com/up-follow) instead of a full page load
- [Preloading the link's destination URL](https://unpoly.com/preloading)
- [Triggering the link on `mousedown`](https://unpoly.com/up-instant) instead of on `click`

### Example

Use `[up-dash]` like this:

    <a href="/users" up-dash=".main">User list</a>

This is shorthand for:

    <a href="/users" up-target=".main" up-instant up-preload>User list</a>
