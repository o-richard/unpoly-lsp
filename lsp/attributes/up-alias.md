Links within [navigational container](https://unpoly.com/navigation-bars) may use the `[up-alias]` attribute to alternative URLs for which they should also be highlighted as `.up-current`.

See [Highlighting links for multiple URLs](https://unpoly.com/navigation-bars#aliases) for more documentation.

### Example

The link below will be highlighted with `.up-current` at both `/profile` and `/profile/edit` locations:

```html
<nav>
<a href="/profile" up-alias="/profile/edit">Profile</a>
</nav>
```

To configure multiple alternative URLs, use a [URL pattern](https://unpoly.com/url-patterns).
