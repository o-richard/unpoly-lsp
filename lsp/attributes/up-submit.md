Submits this form via JavaScript and updates a fragment with the server response.

The server must render an element matching the [target selector](https://unpoly.com/targeting-fragments) from the `[up-target]` attribute.
A matching element in the current page is then swapped with the new element from the server response.
The response may include other HTML (even an entire HTML document), but only the matching element will be updated.

The programmatic variant of this is the [`up.submit()`](https://unpoly.com/up.submit) function.

### Example

```html
<form method="post" action="/users" up-submit up-target=".content">
...
</form>
```

### Handling validation errors

When the form could not be submitted due to invalid user input,
Unpoly can re-render the form with validation errors.

See [validating forms](https://unpoly.com/validation) for details and examples.


### Showing that the form is processing

See [Loading state](https://unpoly.com/loading-state) and [Disabling form controls while working](https://unpoly.com/disabling-forms).


### Short notation

You may omit the `[up-submit]` attribute if the form has one of the following attributes:

- `[up-target]`
- `[up-layer]`
- `[up-transition]`

Such a form will still be submitted through Unpoly.

### Handling all forms automatically

You can configure Unpoly to handle *all* forms on a page without requiring an `[up-submit]` attribute.

See [Handling all links and forms](https://unpoly.com/handling-everything).