❗[EXPERIMENTAL] Enables keyboard interaction and other accessibility behaviors for non-interactive elements that represent clickable buttons.

It's up to you make the element appear interactive visually, e.g. by assigning a `.button` class from your design system.

See [Clicking on non-interactive elements](https://unpoly.com/faux-interactive-elements) for an overview of similiar techniques.

### Example

Add the `[up-clickable]` attribute to a non-interactive element, like a `<span>`:

```html
<span id="faux-button" up-clickable>Click me</span> <!-- mark-phrase "up-clickable" -->
```

To react the element's effect when activated, handle the `up:click` event:

```js
let button = document.querySelector('#faux-button')

button.addEventListener('up:click', function(event) {
console.log('Click on faux button!')
})
```

### Act on press

To activate the element on `mousedown` instead of `click`, also set an `[up-instant]` attribute:

```html
<span id="faux-button" up-clickable up-instant>Click me</span> <!-- mark-phrase "up-instant" -->
```

### Unobtrusive use

To make elements clickable without an explicit `[up-clickable]` attribute, configure `up.link.config.clickableSelectors`:

```js
up.link.config.clickableSelectors.push('.button')
```

Any matching element will now gain [keyboard interaction and other accessibility behaviors](https://unpoly.com/faux-interactive-elements#accessibility):

```html
<span class="button">I can be used with the keyboard</span>
```