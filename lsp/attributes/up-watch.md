Watches form fields and runs a callback when a value changes.

Only fields with a `[name]` attribute can be watched.

The programmatic variant of this is the [`up.watch()`](https://unpoly.com/up.watch) function.

## Example

The following would log a message whenever the `<input>` changes:

```html
<input name="query" up-watch="console.log('New value', value)">
```

## Callback context

The script given to `[up-watch]` runs with the following context:

| Name                  | Type       | Description                                                                                                                                   |
|-----------------------|------------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| `value`               | `string`   | The changed field value.                                                                                                                      |
| `name`                | `string`   | The `[name]` of the changed field.                                                                                                            |
| `options.origin`      | `Element`  | The element that caused the change.<br>This is usually the changed field.                                                                     |
| `options.feedback`    | `boolean`  | Whether to set [feedback classes](https://unpoly.com/feedback-classes) while working.<br>Parsed from the field's `[up-watch-feedback]` attribute.               |
| `options.disable`     | `boolean`  | Which [fields to disable](https://unpoly.com/disabling-forms) while working.<br>Parsed from the field's `[up-watch-disable]` attribute.                         |
| `options.preview`     | `string`   | The name of a [preview](https://unpoly.com/previews) to run while working.<br>Parsed from the field's `[up-watch-preview]` attribute.                           |
| `options.placeholder` | `string`   | The HTML or selector for a [placeholder](https://unpoly.com/placeholders) to show while working.<br>Parsed from the field's `[up-watch-placeholder]` attribute. |


## Watching multiple fields

You can set `[up-watch]` on any element to observe all contained fields.
The `name` argument contains the name of the field that was changed:

```html
<form>
<div up-watch="console.log(`New value of ${name} is ${value}`)">
    <input type="email" name="email">
    <input type="password" name="password">
</div>

<!-- This field is outside the [up-watch] container and will not be watched -->
<input type="text" name="screen-name">
</form>
```

You may also set `[up-watch]` on a `<form>` element to watch *all* fields in a form:

```html
<form up-watch="console.log(`New value of ${name} is ${value}`)">
<input type="email" name="email">
<input type="password" name="password">
<input type="text" name="screen-name">
</form>
```

### Watching radio buttons

Multiple radio buttons with the same `[name]` produce a single value for the form.

To watch radio buttons group, use the `[up-watch]` attribute on an
element that contains all radio button elements with a given name:

```html
<div up-watch="console.log('New value is', value)">
<input type="radio" name="format" value="html"> HTML format
<input type="radio" name="format" value="pdf"> PDF format
<input type="radio" name="format" value="txt"> Text format
</div>
```

## Async callbacks

When your callback does async work (like fetching data over the network) it must return a promise
that settles once the work concludes:

```html
<input name="query" up-watch="return asyncWork()"> <!-- mark-phrase "return" -->
```

Unpoly will guarantee that only one async callback is running concurrently.
If the form is changed while an async callback is still processing, Unpoly will wait until the callback concludes and then run it again with the latest field values.