Whether to [animate the scroll motion](https://unpoly.com/scroll-tuning#animating-the-scroll-motion) when [prepending or appending](https://unpoly.com/targeting-fragments#appending-or-prepending) content.

To animate the scroll motion, pass `{ scrollBehavior: 'smooth' }.`
To instantly jump to the new scroll position, pass `{ scrollBehavior: 'instant' }` (the default).

When `{ scrollBehavior: 'auto' }` is passed, the behavior is determined by the CSS property
[`scroll-behavior`](https://developer.mozilla.org/en-US/docs/Web/CSS/scroll-behavior) of the viewport element.

> [important]
> When [swapping a fragment](https://unpoly.com/targeting-fragments#swapping), the scroll motion cannot be animated.
> You *can* animate the scroll motion when [prepending, appending](https://unpoly.com/targeting-fragments#appending-or-prepending) or [destroying](https://unpoly.com/up.destroy) a fragment.