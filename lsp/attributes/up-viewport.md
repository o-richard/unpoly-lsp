Marks this element as a scrolling container ("viewport").

Apply this attribute if your app uses a custom panel layout with fixed positioning instead of scrolling the `<body>` element. As an alternative you can also push a selector matching your custom viewport to the `up.viewport.config.viewportSelectors` array.

When [scrolling](https://unpoly.com/scrolling) Unpoly will always scroll the viewport closest to the updated element. By default this is the `<body>` element.

Elements with the `[up-viewport]` attribute must also have a [derivable target selector](https://unpoly.com/target-derivation).

### Example

Here is an example for a layout for an e-mail client, showing a list of e-mails
on the left side and the e-mail text on the right side:

```css
.side {
position: fixed;
top: 0;
bottom: 0;
left: 0;
width: 100px;
overflow-y: scroll;
}

.main {
position: fixed;
top: 0;
bottom: 0;
left: 100px;
right: 0;
overflow-y: scroll;
}
```

This would be the HTML (notice the `up-viewport` attribute):

```html
<div class=".side" up-viewport>
<a href="/emails/5001" up-target=".main">Re: Your invoice</a>
<a href="/emails/2023" up-target=".main">Quote for services</a>
<a href="/emails/9002" up-target=".main">Fwd: Room reservation</a>
</div>

<div class="main" up-viewport>
<h1>Re: Your Invoice</h1>
<p>
    Lorem ipsum dolor sit amet, consetetur sadipscing elitr.
    Stet clita kasd gubergren, no sea takimata sanctus est.
</p>
</div>