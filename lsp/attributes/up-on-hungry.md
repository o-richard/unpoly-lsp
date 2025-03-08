Code to run before this element is included in a fragment update.

Calling `event.preventDefault()` will prevent the hungry fragment from being updated.

For instance, you want to auto-update an hungry navigation bar, but only if we're changing history entries:

```html
<nav id="side-nav" up-hungry up-on-hungry="if (!renderOptions.history) event.preventDefault()">
    ...
</nav>
```

The code may use the variables `event` (of type `up:fragment:hungry`), `this` (the hungry element), `newFragment` and `renderOptions`.
