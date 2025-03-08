Only piggy-back on updates on [layers](https://unpoly.com/up.layer) that match the given [layer reference](https://unpoly.com/layer-option).

Relative references like `'parent'` or `'child'` will be resolved in relation to the hungry element's layer.

When set to `'front'`, polling will pause while the fragment's layer is covered by an overlay.
When the fragment's layer is uncovered, polling will resume.

To match a hungry element when updating one of multiple layers, separate the references using and `or` delimiter.
For example, `'current or child'` will match for updates on either the hungry element's layer, or
its direct child.

To match a hungry element when updating *any* layer, set this attribute to `'any'`.

When set to `'any'`, polling will continue on background layers.


