A selector for the fragment to render the content in.

By default the element itself will be replaced with the loaded content.
For this the element must have a [derivable target selector](https://unpoly.com/target-derivation).

You may target one or [multiple](https://unpoly.com/targeting-fragments#multiple) fragments.
To target the placeholder itself, you can use `:origin` target instead of spelling out a selector.

By default Unpoly expects all targeted fragments to be present in both the current page and the server response. If a target selector doesn't match in either, an error `up.CannotMatch` is thrown.

You may mark a target as optional by using the  `:maybe` pseudo selector or you can also set an `up-hungry` attribute on the element that should optionally be updated.

Instead of swapping an entire fragment you may append children to an existing fragment by using the `:after` pseudo selector. In the same fashion, you can use `:before` to prepend the loaded content.

Use `:layer` to replace all visible elements of a layer. To only update a layer's main content area while keeping static layout elements around it, see `:main`.

To make a server request without changing a fragment, use the `:none` [target](https://unpoly.com/targeting-fragments).

With `up-follow`, if omitted, a [main target](https://unpoly.com/up-main) will be rendered.