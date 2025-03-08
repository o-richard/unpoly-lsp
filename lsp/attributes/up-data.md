Attaches structured data to an element, to be consumed by a compiler or event handler.

If an element with an `[up-data]` attribute enters the DOM, Unpoly will parse the JSON and pass the resulting object to any matching `up.compiler()` functions and `up.on()` callbacks.

To programmatically parse an `[up-data]` attribute into an object, use `up.data(element)`.

### Example

A container for a [Google Map](https://developers.google.com/maps/documentation/javascript/tutorial)
might attach the location and names of its marker pins:

```html
<div class="google-map" up-data="[
{ lat: 48.36, lng: 10.99, title: 'Friedberg' },
{ lat: 48.75, lng: 11.45, title: 'Ingolstadt' }
]"></div>
```

The JSON will be parsed and handed to your compiler as a second argument:

```js
up.compiler('.google-map', function(element, pins) {
var map = new google.maps.Map(element)
for (let pin of pins) {
    var position = new google.maps.LatLng(pin.lat, pin.lng)
    new google.maps.Marker({ position, map, title: pin.title })
}
})
```

Similarly, when an event is triggered on an element annotated with
[`up-data`], the parsed object will be passed to any matching
[`up.on()`](https://unpoly.com/up.on) handlers:

```js
up.on('click', '.google-map', function(event, element, data) {
console.log("There are %d pins on the clicked map", data.pins.length)
})
```

You may also parse the data object programmatically using the `up.data()` function:

```
let data = up.data('.google-map')
data[0].lat // => 48.36
data[0].lng // => 10.99
data[0].title // => 'Friedberg'
```

### Alternatives

See [attaching data to elements](https://unpoly.com/data).