Configures whether this `<head>` element is updated during [history changes](https://unpoly.com/updating-history).

By [default](https://unpoly.com/up.history.config#config.metaTagSelectors) popular `<meta>` and certain `<link>`
elements in the `<head>` are considered meta tags.
They will be updated when history is changed, in addition to the document's title and URL.

```html
<link rel="canonical" href="https://example.com/dresses/green-dresses"> <!-- mark-line -->
<meta name="description" content="About the AcmeCorp team"> <!-- mark-line -->
<meta prop="og:image" content="https://app.com/og.jpg"> <!-- mark-line -->
<script src="/assets/app.js"></script>
<link rel="stylesheet" href="/assets/app.css">
```

The linked JavaScript and stylesheet are *not* part of history state and will not be updated during history changes.

### Including additional elements {#including-meta-tags}

To update additional `<head>` elements during history changes, mark them with an `[up-meta]` attribute:

```html
<link rel="license" href="https://opensource.org/license/mit/" up-meta>
```

Only elements in the `<head>` can be matched this way.

To include additional elements by default, configure `up.history.config.metaTagSelectors`.

### Excluding elements {#excluding-meta-tags}

To preserve a `<head>` element during history, changes, set an `[up-meta=false]` attribute:

```html
<meta charset="utf-8" up-meta="false">
```

To exclude elements by default, configure `up.history.config.noMetaTagSelectors`.