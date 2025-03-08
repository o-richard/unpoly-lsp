A string of HTML comprising *only* the new fragment's [outer HTML](https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML).

With an `[up-fragment]` attribute you can omit the `[up-target]` attribute.
The target will be [derived](https://unpoly.com/target-derivation) from the root element in the given HTML.

See [Rendering a string that only contains the fragment](https://unpoly.com/providing-html#fragment).

Instead of passing an HTML string you can also [pass a template selector](https://unpoly.com/templates),
optionally with [variables](https://unpoly.com/placeholders#dynamic-templates).
