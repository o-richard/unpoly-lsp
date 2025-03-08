Whether to [reload the targeted fragment](https://unpoly.com/caching#revalidation) after it was rendered from a cached response.

With `[up-revalidate='auto']` Unpoly will revalidate if the `up.fragment.config.autoRevalidate(response)`
returns `true`. By default this configuration will return true for
[expired](https://unpoly.com/up.fragment.config#config.autoRevalidate) responses.

With `[up-revalidate='true']` Unpoly will always revalidate cached content, regardless
of its age.

With `[up-revalidate='false']` Unpoly will never revalidate cached content.