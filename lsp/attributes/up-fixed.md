## Top

Marks this element as being fixed to the top edge of the screen using `position: fixed`.

When [following a fragment link](https://unpoly.com/up-follow), the viewport is scrolled so the targeted element becomes visible. 
By using this attribute you can make Unpoly aware of fixed elements that are obstructing the viewport contents.
Unpoly will then scroll the viewport far enough that the revealed element is fully visible.

Instead of using this attribute,
you can also configure a selector in `up.viewport.config.fixedTopSelectors`.

### Example

```html
<div class="top-nav" up-fixed="top">...</div>
```

## Bottom

Marks this element as being fixed to the bottom edge of the screen using `position: fixed`.

When [following a fragment link](https://unpoly.com/up-follow), the viewport is scrolled so the targeted element becomes visible. 
By using this attribute you can make Unpoly aware of fixed elements that are obstructing the viewport contents.
Unpoly will then scroll the viewport far enough that the revealed element is fully visible.

Instead of using this attribute, you can also configure a selector in `up.viewport.config.fixedBottomSelectors`.

### Example

```html
<div class="bottom-nav" up-fixed="bottom">...</div>
```