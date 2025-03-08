A JavaScript snippet that is executed when no further DOM changes will be caused by this render pass.

In particular:

- [Animations](https://unpoly.com/up.motion) have concluded and [transitioned](https://unpoly.com/up-transition) elements were removed from the DOM tree.
- A [cached response](#up-cache) was [revalidated with the server](https://unpoly.com/caching#revalidation). If the server has responded with new content, this content has also been rendered.

| Expression | Value                                                                  |
|------------|------------------------------------------------------------------------|
| `this`     | The link being followed                                                |
| `result`   | The `up.RenderResult` for the last render pass that updated a fragment |

If [revalidation](https://unpoly.com/caching#revalidation) re-rendered the fragment, `result` describes updates from the second render pass. 
If no revalidation was performed, or if revalidation yielded an [empty response](https://unpoly.com/caching#when-nothing-changed), it is the result from the initial render pass.

Also see [Awaiting postprocessing](https://unpoly.com/render-lifecycle#awaiting-postprocessing).
