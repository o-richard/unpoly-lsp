Add an `[up-expand]` attribute to any element to enlarge the click area of a descendant link.

`[up-expand]` honors all the Unpoly attributes in expanded links, like [`[up-target]`](https://unpoly.com/up-follow#up-target), `[up-instant]` or `[up-preload]`.

### Example

```html
<div class="notification" up-expand>
Record was saved!
<a href="/records">Close</a>
</div>
```

In the example above, clicking anywhere within `.notification` element would [follow](https://unpoly.com/up.follow) the *Close* link.

### Elements with multiple contained links

If a container contains more than one link, you can set the value of the `[up-expand]` attribute to a CSS selector to define which link should be expanded:

```html
<div class="notification" up-expand=".close">
Record was saved!
<a class="details" href="/records/5">Details</a>
<a class="close" href="/records">Close</a>
</div>
```

### Limitations

`[up-expand]` has some limitations for advanced browser users:

- Users won't be able to right-click the expanded area to open a context menu
- Users won't be able to `CTRL`+click the expanded area to open a new tab

To overcome these limitations, consider nesting the entire clickable area in an actual `<a>` tag.
[It's OK to put block elements inside an anchor tag](https://makandracards.com/makandra/43549-it-s-ok-to-put-block-elements-inside-an-a-tag).
