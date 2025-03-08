Show or hide elements when a form field is set to a given value.

The observed elements can use [`[up-show-for]`](https://unpoly.com/up-show-for) and [`[up-hide-for]`](https://unpoly.com/up-hide-for) attributes to indicate for which values they should be shown or hidden.

The `[up-switch]` element and its observed elements must be inside the same `<form>`.

### Example: Select options

The controlling form field gets an `[up-switch]` attribute with a selector for the elements to show or hide:

```html
<select name="advancedness" up-switch=".target">
<option value="basic">Basic parts</option>
<option value="advanced">Advanced parts</option>
<option value="very-advanced">Very advanced parts</option>
</select>
```

The target elements can use [`[up-show-for]`](https://unpoly.com/up-show-for) and [`[up-hide-for]`](https://unpoly.com/up-hide-for) attributes to indicate for which values they should be shown or hidden.

```html
<div class="target" up-show-for="basic">
only shown for advancedness = basic
</div>

<div class="target" up-hide-for="basic">
hidden for advancedness = basic
</div>

<div class="target" up-show-for="advanced, very-advanced">
shown for advancedness = advanced or very-advanced
</div>
```

### Example: Text field

The controlling `<input>` gets an `[up-switch]` attribute with a selector for the elements to show or hide:

```html
<input type="text" name="user" up-switch=".target">

<div class="target" up-show-for="alice">
only shown for user alice
</div>
```

You may also use the pseudo-values `:blank` to match an empty input value,
or `:present` to match a non-empty input value:

```html
<input type="text" name="user" up-switch=".target">

<div class="target" up-show-for=":blank">
please enter a username
</div>
```

### Example: Checkbox

For checkboxes you may match against the pseudo-values `:checked` or `:unchecked`:

```html
<input type="checkbox" name="flag" up-switch=".target">

<div class="target" up-show-for=":checked">
only shown when checkbox is checked
</div>

<div class="target" up-show-for=":unchecked">
only shown when checkbox is unchecked
</div>
```

You may also match against the `[value]` attribute of the checkbox element:

```html
<input type="checkbox" name="flag" value="active" up-switch=".target">

<div class="target" up-show-for="active">
only shown when checkbox is checked
</div>
```

### Example: Radio button

```html
<input type="radio" name="advancedness" value="basic" up-switch=".target">
<input type="radio" name="advancedness" value="advanced" up-switch=".target">
<input type="radio" name="advancedness" value="very-advanced" up-switch=".target">

<div class="target" up-show-for="basic">
only shown for advancedness = basic
</div>

<div class="target" up-hide-for="basic">
hidden for advancedness = basic
</div>

<div class="target" up-show-for="advanced, very-advanced">
shown for advancedness = advanced or very-advanced
</div>
```

### Example: Values containing spaces

If your values might contain spaces, you may also serialize them as a [relaxed JSON](https://unpoly.com/relaxed-json) array:

```html
<select name='advancedness' up-switch='.target'>
<option value='John Doe'>John Doe</option>
<option value='Jane Doe'>Jane Doe</option>
<option value='Max Mustermann'>Max Mustermann</option>
</select>

<div class='target' up-show-for='["John Doe", "Jane Doe"]'>
You selected John or Jane Doe
</div>

<div class='target' up-hide-for='["Max Mustermann"]'>
You selected Max Mustermann
</div>
```