Changes the link's destination so it points to the previous URL.

If no previous URL is known, the link will not be changed.

> [NOTE]
> Clicking an `[up-back]` link will *not* call [`history.back()`](https://developer.mozilla.org/en-US/docs/Web/API/History/back).
> Instead the link will [navigate](https://unpoly.com/up.navigate) to the previous URL.

### Example

This link ...

```html
<a href="/default" up-back>
Go back
</a>
```

... will be transformed to:

```html
<a href="/default" up-follow up-href="/previous-page" up-scroll="restore" up-follow>
Go back
</a>
```