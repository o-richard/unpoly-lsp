⚠️ [DEPRECIATED - use `up-watch`] Watches form fields and runs a callback when a value changes.

Only fields with a `[name]` attribute can be watched.

The programmatic variant of this is the [`up.watch()`](https://unpoly.com/up.watch) function.

### Example

The following would run a log whenever the `<input>` changes:

```html
<input name="query" up-observe="console.log('New value', value)">
```

### Callback context

The script given to `[up-watch]` runs with the following context:

| Name     | Type      | Description                           |
| -------- | --------- | ------------------------------------- |
| `this`   | `Element` | The changed form field                |
| `name`   | `Element` | The `[name]` of the changed field     |
| `value`  | `string`  | The new value of the changed field    |

### Watching multiple fields

You can set `[up-watch]` on any element to observe all contained fields.
The `name` argument contains the name of the field that was changed.
