When Unpoly inserts a fragment, the `[up-source]` attribute is automatically set to the URL from which the fragment's HTML was loaded.

When an element is [reloaded](https://unpoly.com/up.reload), Unpoly will request the URL from the closest `[up-source]` attribute. You may manually set `[up-source]` attribute to indicate a different source URL for a fragment or a fragment's descendant.

To access the source URL from JavaScript, use `up.fragment.source()`.
