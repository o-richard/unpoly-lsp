An explicit language code to set as the [`html[lang]`](https://www.tpgi.com/using-the-html-lang-attribute/) attribute.

By default Unpoly will extract the language from the response and update the `html[lang]`
attribute in the current page.
To prevent the attrribute from being changed, set `[up-lang=false]`.

This attribute is only used when [updating history](https://unpoly.com/updating-history).