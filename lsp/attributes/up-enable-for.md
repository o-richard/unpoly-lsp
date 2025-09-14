A list of input values for which this element should be enabled.

Multiple values can be separated by either a space (foo bar) or a comma (foo, bar).

If your values might contain spaces, you may also serialize them as a [relaxed JSON](https://unpoly.com/relaxed-json) array (["foo", "bar"]).

To react to the [presence or absence](https://unpoly.com/switching-form-state#presence) of a value, use :blank or :present.

For [checkboxes](https://unpoly.com/switching-form-state#checkboxes), you can react to :checked or :unchecked in addition to the checked value.