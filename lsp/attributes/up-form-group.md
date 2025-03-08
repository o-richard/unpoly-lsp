Marks this element as a form group, which (usually) contains a label, input and error message.

You are not required to use form groups to [submit forms through Unpoly](https://unpoly.com/up-submit).
However, structuring your form into groups will help Unpoly to make smaller changes to the DOM when working with complex form. 
For instance, when [validating](https://unpoly.com/validation#validating-after-changing-a-field) a field, Unpoly will re-render the [closest](https://developer.mozilla.org/en-US/docs/Web/API/Element/closest) form group around that field.

By default Unpoly will also consider a `<fieldset>` or `<label>` around a field to be a form group.
You can configure this in `up.form.config.groupSelectors`.

### Example

Many apps use form groups to wrap a label, input field, error message and help text:

```html
<form up-validate>
<div up-form-group>
    <label for="email">E-mail</label>
    <input type="text" name="email" id="email">
</div>
<div up-form-group>
    <label for="password">Password</label>
    <input type="text" name="password" id="password">
    <div class="error">Must be 8 characters or longer</div>
</div>
</form>
```

The form above also uses the `[up-validate]` attribute to [validate](https://unpoly.com/validation#validating-after-changing-a-field) form groups after changing a field:

- After changing the *E-Mail* field, Unpoly will validate the `[up-form-group]:has(#email)` target.
- After changing the *Password* field, Unpoly will validate the `[up-form-group]:has(#password)` target.
