Renders a new form state when a field changes, to show validation errors or update [dependent fields](https://unpoly.com/dependent-fields).

When a form field with an `[up-validate]` attribute is changed, the form is submitted to the server
which is expected to render a new form state from its current field values.
The [form group](https://unpoly.com/up-form-group) around the changed field is updated with the server response.

### Updating a different fragment

If you don't want to update the field's form group, you can set the `[up-validate]` attribute to any [target selector](https://unpoly.com/targeting-fragments):

```html
<input type="text" name="email" up-validate=".email-errors"> <!-- mark-phrase ".email-errors" -->
<div class="email-errors"></div>
```

You may also [update multiple fragments](https://unpoly.com/targeting-fragments#multiple) by separating their target selectors with a comma:

```html
<input type="text" name="email" up-validate=".email-errors, .base-errors"> <!-- mark-phrase ".email-errors, .base-errors" -->
```

## Validating multiple fields

You can set `[up-validate]` on any element to validate *all contained fields* on change.

In the [example above](#marking-fields-for-validation), instead of setting `[up-validate]` on each individual `<input>`, we can also set it on the `<form>`:

```html
<form action="/users" up-validate> <!-- mark-phrase "up-validate" -->

<fieldset>
    <label for="email" up-validate>E-mail</label>
    <input type="text" id="email" name="email">
</fieldset>

<fieldset>
    <label for="password" up-validate>Password</label>
    <input type="password" id="password" name="password">
</fieldset>

<button type="submit">Register</button>

</form>
```

### Validating radio buttons

Multiple radio buttons with the same `[name]` produce a single value for the form.

To watch radio buttons group, use the `[up-validate]` attribute on an element that contains all radio button elements with a given name:

```html
<fieldset up-validate>
<input type="radio" name="format" value="html"> HTML format
<input type="radio" name="format" value="pdf"> PDF format
<input type="radio" name="format" value="txt"> Text format
</fieldset>
```
