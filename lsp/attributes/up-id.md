Sets an unique identifier for this element.

This identifier is used in [target derivation](https://unpoly.com/target-derivation) to create a CSS selector that matches this element precisely.

If the element already has [other attributes that make a good identifier](https://unpoly.com/target-derivation#derivation-patterns), like a good `[id]` or `[class]` attribute, it is not necessary to also set `[up-id]`.

### Example

Take this element:

```html
<a href="/">Homepage</a>
```

Unpoly cannot generate a good CSS selector for this element:

```js
up.fragment.toTarget(element)
// throws error: up.CannotTarget
```

We can improve this by assigning an `[up-id]`:

```html
<a href="/" up-id="link-to-home">Open user 4</a>
```

The attribute value is used to create a better selector:

```js
up.fragment.toTarget(element)
// returns '[up-id="link-to-home"]'
```