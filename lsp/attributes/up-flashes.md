Use an `[up-flashes]` element to show confirmations, alerts or warnings.

You application layout should have an empty `[up-flashes]` element to indicate where flash messages should be inserted:

```html
<nav>
Navigation items ...
</nav>
<div up-flashes></div> <!-- mark-line -->
<main>
Main page content ...
</main>
```

To render a flash message, include an `[up-flashes]` element in your response.
The element's content should be the messages you want to render:

```html
<div up-flashes>
<strong>User was updated!</strong>
</div>

<main>
Main response content ...
</main>
```

See [notification flashes](https://unpoly.com/flashes) for more details and examples.
