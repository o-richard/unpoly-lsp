Marks this element as a [navigational container](https://unpoly.com/navigation-bars), such as a menu or navigation bar.

When a link within an `[up-nav]` element points to [its layer's location](https://unpoly.com/up.layer.location), it is assigned the `.up-current` class. When the browser navigates to another location, the class is removed automatically.

See [Navigation bars](https://unpoly.com/navigation-bars) for details and examples.

### Example

Let's look at a simple menu with two links:

```html
<div up-nav> <!-- mark-phrase "up-nav" -->
<a href="/foo">Foo</a>
<a href="/bar">Bar</a>
</div>
```

When the browser location changes to `/foo`, the first link is marked as `.up-current`:

```html
<div up-nav>
<a href="/foo" class="up-current">Foo</a> <!-- mark-phrase "up-current" -->
<a href="/bar">Bar</a>
</div>
```

When the browser location changes to `/bar`, the first link loses its `.up-current` class.
Now the second link is marked as `.up-current`:

```html
<div up-nav>
<a href="/foo">Foo</a>
<a href="/bar" class="up-current">Bar</a> <!-- mark-phrase "up-current" -->
</div>
```
