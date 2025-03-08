Elements with an `[up-poll]` attribute are [reloaded](https://unpoly.com/up.reload) from the server periodically.

### Example

Assume an application layout with an unread message counter.
You can use `[up-poll]` to refresh the counter every 30 seconds:

```html
<div class="unread-count" up-poll>
2 new messages
</div>
```

### Controlling the reload interval

You may set an optional `[up-interval]` attribute to set the reload interval in milliseconds:

```html
<div class="unread-count" up-poll up-interval="10000">
2 new messages
</div>
```

If the value is omitted, a global default is used. You may configure the default like this:

```js
up.radio.config.pollInterval = 10000
```

### Controlling the source URL

The element will be reloaded from the URL from which it was originally loaded.

To reload from another URL, set an `[up-source]` attribute on the polling element:

```html
<div class="unread-count" up-poll up-source="/unread-count">
2 new messages
</div>
```

### Controlling the target selector

A target selector will be [derived](https://unpoly.com/target-derivation) from the polling element.

### Polling is paused in the background

By default polling will pause while the fragment's [layer](https://unpoly.com/up.layer) is covered by an overlay.
When the layer is uncovered, polling will resume.
To keep polling on background layers, set [`[up-if-layer=any]`](#up-if-layer).

Polling will also pause automatically while the browser tab is hidden.
When the browser tab is re-activated, polling will resume.

When at least one poll interval was spent paused in the background and the user then returns to the layer or tab, Unpoly will immediately reload the fragment.
You can use this to load recent data when the user returns to your app after working on something else for a while. For example, the following would reload your [main](https://unpoly.com/main) element after an absence of 5 minutes or more:

```html
<main up-poll up-interval="300_000">
...
</main>
```

### Skipping updates on the client

Client-side code may skip an update by preventing an `up:fragment:poll` event
on the polling fragment.


### Skipping updates on the server

When polling a fragment periodically we want to avoid rendering unchanged content.
This saves <b>CPU time</b> and reduces the <b>bandwidth cost</b> for a
request/response exchange to about 1 KB (1 packet).

See [Skipping rendering](https://unpoly.com/skipping-rendering) for more details and examples.

When an update is skipped, Unpoly will try to poll again after the configured interval.

### Stopping polling

There are multiple ways to stop the polling interval:

- The fragment from the server response no longer has an `[up-poll]` attribute.
- The fragment from the server response has an `[up-poll="false"]` attribute.
- Client-side code has called `up.radio.stopPolling()` with the polling element.
