Controls which fragment to update when the [`[up-target]`](#up-target) selector yields multiple results.

When set to `'region'` Unpoly will prefer to update fragments in the
[region](https://unpoly.com/targeting-fragments#ambiguous-selectors) of the [origin element](https://unpoly.com/up.render#options.origin).

If set to `'first'` Unpoly will always update the first matching fragment.

Defaults to `up.fragment.config.match`, which defaults to `'region'`.