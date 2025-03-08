Automatically submits a form when a field changes.

The programmatic variant of this is the [`up.autosubmit()`](https://unpoly.com/up.autosubmit) function.

### Example

The following would automatically submit the form when the `query` field is changed:

```html
<form method="GET" action="/search">
<input type="search" name="query" up-autosubmit>
<input type="checkbox" name="archive"> Include archived
</form>
```

### Auto-submitting multiple fields

You can set `[up-autosubmit]` on any element to submit the form when a contained field changes.

For instance, to auto-submit a form when any field changes, set the `[up-autosubmit]` on the `<form>` element:

```html
<form method="GET" action="/search" up-autosubmit>
<input type="search" name="query">
<input type="checkbox" name="archive"> Include archived
</form>
```

#### Auto-submitting radio buttons

Multiple radio buttons with the same `[name]` (a radio button group) produce a single value for the form.

To auto-submit group of radio buttons, use the `[up-autosubmit]` attribute on an element containing the entire button group:

```html
<div up-autosubmit>
<input type="radio" name="format" value="html"> HTML format
<input type="radio" name="format" value="pdf"> PDF format
<input type="radio" name="format" value="txt"> Text format
</div>
```
