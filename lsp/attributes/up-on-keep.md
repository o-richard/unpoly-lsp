Code to run before an existing element is kept during a page update.

Calling `event.preventDefault()` will prevent the element from being kept. It will then be swapped with `newFragment`.

The code may use the variables `event` (of type `up:fragment:keep`), `this` (the old fragment), `newFragment` and `newData`.
