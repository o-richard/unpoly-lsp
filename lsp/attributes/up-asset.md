Tracks an element as a [frontend asset](https://unpoly.com/handling-asset-changes), usually JavaScripts and stylesheets.

When [rendering](https://unpoly.com/up.render), Unpoly compares the assets on the current page with the new assets from the server response. If the assets don't match, an `up:assets:changed` event is emitted.

### Default assets

By default all remote scripts and stylesheets in the `<head>` are considered assets:

```html
<html>
<head>
    <link rel="stylesheet" href="/assets/frontend-5f3aa101.css"> <!-- mark-line -->
    <script src="/assets/frontend-81ba23a9.js"></script> <!-- mark-line -->
</head>
<body>
    ...
</body>
</html>
```

Unpoly only tracks assets in the `<head>`. Elements in the `<body>` are never tracked.

[Inline scripts](https://simpledev.io/lesson/inline-script-javascript-1/) and
[internal styles](https://www.tutorialspoint.com/How-to-use-internal-CSS-Style-Sheet-in-HTML)
are not tracked by default, but you can [include them explicitly](#including-assets).

### Excluding assets from tracking {#excluding-assets}

To *exclude* an element in the `<head>` from tracking, mark it with an `[up-asset="false"]` attribute:

```html
<script src="/assets/analytics.js" up-asset="false"></script>
```

To exclude assets by default, configure `up.script.config.noAssetSelectors`.


### Tracking additional assets {#including-assets}

To track additional assets in the `<head>`, mark them with an `[up-asset]` attribute.

For example, [inline scripts](https://simpledev.io/lesson/inline-script-javascript-1/) are not tracked by default, but you can include them explictily:

```html
<script up-asset>
window.SALE_START = new Date('2024-05-01')
</script>
```

Only elements in the `<head>` can be matched this way.

To track additional assets by default, configure `up.script.config.assetSelectors`.

### Tracking the backend version {#tracking-backend-versions}

To detect a new deployment of your *backend* code, consider including the deployed commit hash in a `<meta>` tag.

By marking the `<meta>` tag with `[up-asset]` it will also emit an `up:assets:changed` event when the commit hash changes:

```html
<meta name="backend-version" value="d50c6dd629e9bbc80304e14a6ba99a18c32ba738" up-asset>
```

