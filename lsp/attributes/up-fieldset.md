⚠️ [DEPRECIATED - use `up-form-group`] Marks this element as a from group, which (usually) contains a label, input and error message.

You are not required to use form groups to [submit forms through Unpoly](https://unpoly.com/up-submit).
However, structuring your form into groups will help Unpoly to make smaller changes to the DOM when
working with complex form. For instance, when [validating](https://unpoly.com/up-validate) a field,
Unpoly will re-render the [closest](https://developer.mozilla.org/en-US/docs/Web/API/Element/closest)
form group around that field.

By default Unpoly will also consider a `<fieldset>` or `<label>` around a field to be a form group.
You can configure this in `up.form.config.groupSelectors`.