Marks this element as the primary content element of your application layout.

Unpoly will update a main element when no more specific render target is given.

A space-separated list of [layer modes](https://unpoly.com/layer-terminology) for which to use this main target.
Omit the attribute value to define a main target for *all* layer modes. 
To use a different main target for all overlays (but not the root layer), set `[up-main=overlay]`.

### Example

Many links simply replace the primary content element in your application layout.

Unpoly lets you mark this elements as a default target using the `[up-main]` attribute:

```html
<body>
<div class="layout">
    <div class="layout--side">
    ...
    </div>
    <div class="layout--content" up-main>
    ...
    </div>
</div>
</body>
```

### Overlays can use different main targets {#overlays}

Overlays often use a different default selector, e.g. to exclude a navigation bar.

To define a different main target for an overlay, set the [layer mode](https://unpoly.com/layer-terminology) as the
value of the `[up-main]` attribute:

```html
<body>
<div class="layout" up-main="root">
    <div class="layout--side">
    ...
    </div>
    <div class="layout--content" up-main="modal">
    ...
    </div>
</div>
</body>
```

### Using existing elements as main targets

Instead of the `[up-main]` attribute you may also use the standard [`<main>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main) element.

You may also configure an existing selector in `up.fragment.config.mainTargets`:

```js
up.fragment.config.mainTargets.push('.layout--content')
```

You may configure layer-specific targets in `up.layer.config`:

```js
up.layer.config.popup.mainTargets.push('.menu')              // for popup overlays
up.layer.config.drawer.mainTargets.push('.menu')             // for drawer overlays
up.layer.config.overlay.mainTargets.push('.layout--content') // for all overlay modes
```
