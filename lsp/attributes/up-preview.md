The name of a [preview](https://unpoly.com/previews) that temporarily changes the page while new content is loading.

The preview changes will be reverted automatically when the request ends for [any reason](https://unpoly.com/previews#ending).

Links or forms can name a [preview function](https://unpoly.com/previews) that is called while loading content from the server.

When the user interacts with a link or form, its preview function is invoked immediately.
The function will usually [mutate the DOM](https://unpoly.com/previews#basic-mutations) to signal that the app is working, or to provide clues for how the page will ultimately look.
For example, if the user is deleting an item from a list, the preview function could hide that item visually.

See [Previews](https://unpoly.com/previews) for details and examples.

## Usage

To refer to a preview function, set its name as an `[up-preview]` attribute:

```html
<a href="/edit" up-follow up-preview="spinner">Edit page</a> <!-- mark-phrase "spinner" -->
```

To [call multiple previews](https://unpoly.com/previews#multiple), separate their names with a comma:

```html
<a href="/edit" up-follow up-preview="spinner, dim-page">Edit page</a> <!-- mark-phrase "spinner, dim-page" -->
```

[Preview options](https://unpoly.com/up-preview#parameters) can be appended after each preview name, encoded as [Relaxed JSON](https://unpoly.com/relaxed-json):

```html
<a href="/edit"
    up-follow
    up-preview="spinner { size: 20 }, dim-page { animation: 'pulse' }"> <!-- mark-phrase "spinner { size: 20 }, dim-page { animation: 'pulse' }" -->
Edit page
</a>
```