The name of an [animation](https://unpoly.com/up.motion) to reveal a new fragment when
[prepending or appending content](https://unpoly.com/targeting-fragments#appending-or-prepending) or when opening/closing an overlay.

If you are replacing content (the default), use the `[up-transition]` attribute instead.

For layers, it defaults to overlay's [preconfigured close animation](https://unpoly.com/up.layer.config).

Predefined animations
=====================

| Animation          | Visual effect  |
|--------------------|----------------|
| `fade-in`          | Changes the element's opacity from 0% to 100% |
| `fade-out`         | Changes the element's opacity from 100% to 0% |
| `move-to-top`      | Moves the element upwards until it exits the screen at the top edge |
| `move-from-top`    | Moves the element downwards from beyond the top edge of the screen until it reaches its current position |
| `move-to-bottom`   | Moves the element downwards until it exits the screen at the bottom edge |
| `move-from-bottom` | Moves the element upwards from beyond the bottom edge of the screen until it reaches its current position |
| `move-to-left`     | Moves the element leftwards until it exists the screen at the left edge |
| `move-from-left`   | Moves the element rightwards from beyond the left edge of the screen until it reaches its current position |
| `move-to-right`    | Moves the element rightwards until it exists the screen at the right edge |
| `move-from-right`  | Moves the element leftwards from beyond the right  edge of the screen until it reaches its current position |
| `none`             | An animation that has no visible effect. Sounds useless at first, but can save you a lot of `if` statements. |

## Custom animations

To define a custom animation, use `up.animation()`.
