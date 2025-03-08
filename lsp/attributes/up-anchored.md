Marks this element as being anchored to the right edge of the screen, typically fixed navigation bars.

Since [overlays](https://unpoly.com/up.layer) hide the document scroll bar, elements anchored to the right appear to jump when the dialog opens or closes. 
Applying this attribute to anchored elements will make Unpoly aware of the issue and adjust the `right` property accordingly.
You may customize this behavior by styling the `.up-scrollbar-away` class.

Instead of giving this attribute to any affected element, you can also configure a selector in `up.viewport.config.anchoredRightSelectors`.

> [note]
> Elements with `[up-fixed=top]` or `[up-fixed=bottom]` are also considered to be right-anchored.

### Example

Here is the CSS for a navigation bar that is anchored to the top edge of the screen:

```css
.top-nav {
position: fixed;
top: 0;
left: 0;
right: 0;
}
```

By adding an `up-anchored="right"` attribute to the element, we can prevent the `right` edge from jumping when an [overlay](https://unpoly.com/up.layer) opens or closes:

```html
<div class="top-nav" up-anchored="right">...</div>
```
