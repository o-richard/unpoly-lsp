package lsp

import (
	"embed"
	"errors"
	"fmt"
	"os"
)

type attribute struct {
	detail           string
	documentation    string
	insertText       string
	insertTextFormat int
	isDepreciated    bool
	choices          []string
}

var attributes map[string]*attribute = map[string]*attribute{
	"up-clickable": {
		detail:     "Enables keyboard interaction and other accessibility behaviors for non-interactive elements that represent clickable buttons",
		insertText: "up-clickable", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-dash": {
		detail:     "Follows this link as fast as possible",
		insertText: `up-dash="$1"`, insertTextFormat: snippetInsertTextFormat, isDepreciated: true,
	},
	"up-background": {
		detail:           "Whether this request will load in the background",
		insertText:       `up-background="${1|false,true|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"false", "true"},
	},
	"up-headers": {
		detail:     "A relaxed JSON object with additional request headers",
		insertText: `up-headers="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-href": {
		detail:     "The URL from which to load the content",
		insertText: `up-href="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-defer": {
		detail:           "When to load and render the deferred content",
		insertText:       `up-defer="${1|insert,reveal,manual|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"insert", "reveal", "manual"},
	},
	"up-intersect-margin": {
		detail:     "Enlarges the viewport by the given number of pixels before computing the intersection",
		insertText: `up-intersect-margin="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-target": {
		detail:     "A selector for the fragment to render the content in",
		insertText: `up-target="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-expand": {
		detail:     "Enlarge the click area of a descendant link",
		insertText: "up-expand", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-follow": {
		detail:     "Follows this link with JavaScript and updates a fragment with the server response",
		insertText: `up-navigate="${1|false,true|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"false", "true"},
	},
	"up-navigate": {
		detail:           "Whether this fragment update is considered navigation",
		insertText:       `up-navigate="${1|false,true|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"false", "true"},
	},
	"up-fallback": {
		detail:     "Specifies behavior if the target selector is missing from the current page or the server response",
		insertText: `up-fallback="${1:true}"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-match": {
		detail:     "Controls which fragment to update when the `[up-target]` selector yields multiple results",
		insertText: `up-match="${1|region,first|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"region", "first"},
	},
	"up-method": {
		detail:     "The HTTP method to use for the request",
		insertText: `up-method="${1|get,post,put,patch,delete|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"get", "post", "put", "patch", "delete"},
	},
	"up-params": {
		detail:     "A relaxed JSON object with additional parameters that should be sent as the request's query string",
		insertText: `up-params="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-content": {
		detail:     "The new inner HTML for the targeted fragment",
		insertText: `up-content="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-fragment": {
		detail:     "A string of HTML comprising *only* the new fragment's outer HTML",
		insertText: `up-fragment="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-document": {
		detail:     "A string of HTML containing the targeted fragment",
		insertText: `up-document="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-use-data": {
		detail:     "A relaxed JSON object that overrides properties from the new fragment's data",
		insertText: `up-use-data="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-fail": {
		detail:     "Whether the server response should be considered failed",
		insertText: `up-fail="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-fail-target": {
		detail:     "The target selector to update after a failed response",
		insertText: `up-fail-target="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-history": {
		detail:     "Whether the browser URL, window title and meta tags will be updated",
		insertText: `up-history="${1|true,auto,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "auto", "false"},
	},
	"up-fail-history": {
		detail:     "Whether to update history when the server responds with an error code",
		insertText: `up-fail-history="${1|true,auto,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "auto", "false"},
	},
	"up-title": {
		detail:     "An explicit document title to set before rendering",
		insertText: `up-title="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-location": {
		detail:     "An explicit URL to set before rendering",
		insertText: `up-location="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-meta-tags": {
		detail:     "Whether to update meta tags `<head>`",
		insertText: `up-meta-tags="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-lang": {
		detail:     "An explicit language code to set as the `html[lang]` attribute",
		insertText: `up-lang="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-transition": {
		detail:           "The name of an transition to morph between the old and few fragment/notification flashes",
		insertText:       `up-transition="${1:|cross-fade,move-up,move-down,move-left,move-right,none|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"cross-fade", "move-up", "move-down", "move-left", "move-right", "none"},
	},
	"up-fail-transition": {
		detail:           "The transition to use when the server responds with an error code",
		insertText:       `up-fail-transition="${1:|cross-fade,move-up,move-down,move-left,move-right,none|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"cross-fade", "move-up", "move-down", "move-left", "move-right", "none"},
	},
	"up-animation": {
		detail:           "The name of an animation to reveal a new fragment when prepending or appending content or opening/closing a layer",
		insertText:       `up-animation="${1|fade-in,fade-out,move-to-top,move-from-top,move-to-bottom,move-from-bottom,move-to-left,move-from-left,move-to-right,move-from-right,none|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"fade-in", "fade-out", "move-to-top", "move-from-top", "move-to-bottom", "move-from-bottom", "move-to-left", "move-from-left", "move-to-right", "move-from-right", "none"},
	},
	"up-duration": {
		detail:     "The duration of the transition or animation (in millisconds)",
		insertText: `up-duration="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-easing": {
		detail:     "The timing function that accelerates the transition or animation",
		insertText: `up-easing="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-cache": {
		detail:     "Whether to read from and write to the cache",
		insertText: `up-cache="${1|true,auto,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "auto", "false"},
	},
	"up-revalidate": {
		detail:     "Whether to reload the targeted fragment after it was rendered from a cached response",
		insertText: `up-revalidate="${1|true,auto,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "auto", "false"},
	},
	"up-expire-cache": {
		detail:     "Whether existing cache entries will be expired with this request",
		insertText: `up-expire-cache`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-evict-cache": {
		detail:     "Whether existing cache entries will be evicted with this request",
		insertText: `up-evict-cache`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-abort": {
		detail:     "Whether to abort existing requests before rendering",
		insertText: `up-abort="target"`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-abortable": {
		detail:     "Whether this request may be aborted by other requests targeting the same fragments or layer",
		insertText: `up-abortable="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-late-delay": {
		detail:     "The number of milliseconds after which this request can cause an `up:network:late` event",
		insertText: `up-late-delay="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-timeout": {
		detail:     "The number of milliseconds after which this request fails with a timeout",
		insertText: `up-timeout="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-layer": {
		detail:     "The layer in which to match and render the fragment",
		insertText: `up-layer="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-fail-layer": {
		detail:     "The layer in which render if the server responds with an error code",
		insertText: `up-fail-layer="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-peel": {
		detail:     "Whether to close overlays obstructing the updated layer when the fragment is updated",
		insertText: `up-peel="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-context": {
		detail:     "A relaxed JSON object that will be merged into the context of the current layer once the fragment is rendered",
		insertText: `up-context="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-scroll": {
		detail:     "How to scroll after the new fragment was rendered",
		insertText: `up-scroll="${1:|auto,smooth,instant|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"auto", "smooth", "instant"},
	},
	"up-fail-scroll": {
		detail:     "How to scroll after the new fragment was rendered from a failed response",
		insertText: `up-fail-scroll="${1:|auto,smooth,instant|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"auto", "smooth", "instant"},
	},
	"up-scroll-behavior": {
		detail:     "Whether to animate the scroll motion when prepending or appending content",
		insertText: `up-scroll-behavior="${1:instant}"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-reveal-snap": {
		detail:     "When to snap to the top when scrolling to an element near the top edge of the viewport's scroll buffer",
		insertText: `up-reveal-snap="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-reveal-top": {
		detail:     "When to move a revealed element to the top when scrolling to an element",
		insertText: `up-reveal-top="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-reveal-padding": {
		detail:     "How much space to leave to the closest viewport edge when scrolling to an element",
		insertText: `up-reveal-padding="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-reveal-max": {
		detail:     "How many pixel lines of high element to reveal when scrolling to an element",
		insertText: `up-reveal-max="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-save-scroll": {
		detail:     "Whether to save scroll positions before updating the fragment",
		insertText: `up-save-scroll`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-focus": {
		detail:     "What to focus after the new fragment was rendered",
		insertText: `up-focus="${1:auto}"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-fail-focus": {
		detail:     "What to focus after the new fragment was rendered from a failed response",
		insertText: `up-fail-focus="${1:auto}"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-focus-visible": {
		detail:     "Whether the focused element should have a visible focus ring",
		insertText: `up-focus-visible="${1:auto}"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-save-focus": {
		detail:     "Whether to save focus-related state before updating the fragment",
		insertText: `up-save-focus`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-confirm": {
		detail:     "A message the user needs to confirm before fragments are updated or before the layer is closed",
		insertText: `up-confirm="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-placeholder": {
		detail:     "A placeholder to show in the targeted fragment while new content is loading",
		insertText: `up-placeholder="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-preview": {
		detail:     "The name of a preview that temporarily changes the page while new content is loading",
		insertText: `up-preview="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-revalidate-preview": {
		detail:     "The name of a preview that runs while revalidating cached content",
		insertText: `up-revalidate-preview="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-disable": {
		detail:     "Disables form controls",
		insertText: `up-disable="${1:true}"`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-disable-for": {
		detail:     "Disables this element while an input field with [up-switch] has one of the given values",
		insertText: `up-disable-for="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-enable-for": {
		detail:     "Enables this element while an input field with [up-switch] has one of the given values",
		insertText: `up-enable-for="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-feedback": {
		detail:     "Whether to set feedback classes while loading content",
		insertText: `up-feedback="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-on-loaded": {
		detail:     "A JavaScript snippet that is executed when the server responds with new HTML, but before the HTML is rendered",
		insertText: `up-on-loaded="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-on-rendered": {
		detail:     "A JavaScript snippet that is executed when Unpoly has updated fragments",
		insertText: `up-on-rendered="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-on-finished": {
		detail:     "A JavaScript snippet that is executed when no further DOM changes will be caused by this render pass",
		insertText: `up-on-finished="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-on-offline": {
		detail:     "A JavaScript snippet that is executed when the fragment could not be loaded due to a disconnect or timeout",
		insertText: `up-on-offline="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-on-error": {
		detail:     "A JavaScript snippet that is run when any error is thrown during the rendering process",
		insertText: `up-on-error="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-use-keep": {
		detail:     "❗[EXPERIMENTAL] Whether `[up-keep]` elements will be preserved in the updated fragment",
		insertText: `up-use-keep="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-use-hungry": {
		detail:     "❗[EXPERIMENTAL] Whether `[up-hungry]` elements outside the updated fragment will also be updated",
		insertText: `up-use-hungry="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-instant": {
		detail:     `Follows this link on 'mousedown' instead of 'click' ("Act on press")`,
		insertText: `up-instant="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-preload": {
		detail:     "Preloads this link before the user clicks it",
		insertText: `up-preload="${1|true,false,hover,insert,reveal|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false", "hover", "insert", "reveal"},
	},
	"up-preload-delay": {
		detail:     "The number of milliseconds before the link is preloaded",
		insertText: `up-preload-delay="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-asset": {
		detail:     "Tracks an element as a frontend asset, usually JavaScripts and stylesheets",
		insertText: `up-asset="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-data": {
		detail:     "Attaches structured data to an element, to be consumed by a compiler or event handler",
		insertText: `up-data="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-autosubmit": {
		detail:     "Automatically submits a form when a field changes",
		insertText: `up-autosubmit`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-watch-event": {
		detail:     "The type of event to watch",
		insertText: `up-watch-event="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-watch-delay": {
		detail:     "The number of milliseconds to wait between an observed event and intended result",
		insertText: `up-watch-delay="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-watch-disable": {
		detail:     "Whether to disable fields",
		insertText: `up-watch-disable`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-watch-placeholder": {
		detail:     "A placeholder to show within the targeted fragment",
		insertText: `up-watch-placeholder="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-watch-preview": {
		detail:     "One or more previews that temporarily change the page",
		insertText: `up-watch-preview="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-watch-feedback": {
		detail:     "Whether to set feedback classes",
		insertText: `up-watch-feedback="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-fieldset": {
		detail:     "Marks this element as a from group, which (usually) contains a label, input and error message",
		insertText: `up-fieldset`, insertTextFormat: plainTextInsertTextFormat, isDepreciated: true,
	},
	"up-form-group": {
		detail:     "Marks this element as a form group, which (usually) contains a label, input and error message",
		insertText: `up-form-group`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-hide-for": {
		detail:     "Hides this element if an input field with `[up-switch]` has one of the given values",
		insertText: `up-hide-for="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-observe": {
		detail:     "Watches form fields and runs a callback when a value changes",
		insertText: `up-observe="$1"`, insertTextFormat: snippetInsertTextFormat, isDepreciated: true,
	},
	"up-show-for": {
		detail:     "Only shows this element if an input field with `[up-switch]` has one of the given values",
		insertText: `up-show-for="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-submit": {
		detail:     "Submits this form via JavaScript and updates a fragment with the server response",
		insertText: "up-submit", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-switch": {
		detail:     "Show or hide elements when a form field is set to a given value",
		insertText: `up-switch="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-switch-region": {
		detail:     "A selector for the region in which elements are switched.",
		insertText: `up-switch-region="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-validate": {
		detail:     "Renders a new form state when a field changes, to show validation errors or update dependent fields",
		insertText: `up-validate`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-validate-url": {
		detail:     "The URL to which to submit the validation request. By default Unpoly will use the form's [action] attribute",
		insertText: `up-validate-url="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-validate-method": {
		detail:     "The method to use for submitting the validation request. By default Unpoly will use the form's [method] attribute",
		insertText: `up-validate-method="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-validate-params": {
		detail:     "Additional Form parameters that should be sent as the request's query string or payload",
		insertText: `up-validate-params="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-validate-headers": {
		detail:     "A relaxed JSON object with additional request headers",
		insertText: `up-validate-headers="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-validate-batch": {
		detail:     "Whether to consolidate multiple validations into a single request",
		insertText: `up-validate-batch="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-watch": {
		detail:     "Watches form fields and runs a callback when a value changes",
		insertText: `up-watch="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-accept": {
		detail:     "The overlay's acceptance value as a relaxed JSON value",
		insertText: `up-accept="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-close": {
		detail:     "When this element is clicked, closes a currently open overlay",
		insertText: `up-close`, insertTextFormat: plainTextInsertTextFormat, isDepreciated: true,
	},
	"up-close-animation": {
		detail:           "The name of the closing animation",
		insertText:       `up-close-animation="${1|fade-in,fade-out,move-to-top,move-from-top,move-to-bottom,move-from-bottom,move-to-left,move-from-left,move-to-right,move-from-right,none|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"fade-in", "fade-out", "move-to-top", "move-from-top", "move-to-bottom", "move-from-bottom", "move-to-left", "move-from-left", "move-to-right", "move-from-right", "none"},
	},
	"up-close-easing": {
		detail:     "The timing function that controls the closing animation's acceleration",
		insertText: `up-close-easing="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-close-duration": {
		detail:     "The duration of the closing animation in milliseconds",
		insertText: `up-close-duration="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-dismiss": {
		detail:     "The overlay's dismissal value as a relaxed JSON value",
		insertText: `up-dismiss="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-drawer": {
		detail:     "Clicking this link will load the destination via AJAX and open the given selector in a modal drawer that slides in from the edge of the screen",
		insertText: `up-drawer="$1"`, insertTextFormat: snippetInsertTextFormat, isDepreciated: true,
	},
	"up-mode": {
		detail:     "The kind of overlay to open",
		insertText: `up-mode="${1|root,modal,drawer,popup,cover|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"root", "modal", "drawer", "popup", "cover"},
	},
	"up-size": {
		detail:     "The size of the overlay",
		insertText: `up-size="${1|small,medium,large,grow,full|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"small", "medium", "large", "grow", "full"},
	},
	"up-class": {
		detail:     "An optional HTML class for the overlay's container element",
		insertText: `up-class="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-dismissable": {
		detail:     "How the overlay may be dismissed by the user.",
		insertText: `up-dismissable="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-on-opened": {
		detail:     "A JavaScript snippet that is called when the overlay was inserted into the DOM",
		insertText: `up-on-opened="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-on-accepted": {
		detail:     "A JavaScript snippet that is called when the overlay was accepted",
		insertText: `up-on-accepted="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-on-dismissed": {
		detail:     "A JavaScript snippet that is called when the overlay was dismissed",
		insertText: `up-on-dismissed="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-accept-event": {
		detail:     "One or more space-separated event types that will cause this overlay to automatically be accepted when a matching event occurs within the overlay",
		insertText: `up-accept-event="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-dismiss-event": {
		detail:     "One or more space-separated event types that will cause this overlay to automatically be dismissed when a matching event occurs within the overlay",
		insertText: `up-dismiss-event="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-accept-location": {
		detail:     "One or more space-separated URL patterns that will cause this overlay to automatically be accepted when the overlay reaches a matching location",
		insertText: `up-accept-location="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-dismiss-location": {
		detail:     "One or more space-separated URL patterns that will cause this overlay to automatically be dismissed when the overlay reaches a matching location",
		insertText: `up-dismiss-location="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-position": {
		detail:           "The position of the popup relative to the `{ origin }` element that opened the overlay",
		insertText:       `up-position="${1|top,right,bottom,left|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"top", "right", "bottom", "left"},
	},
	"up-align": {
		detail:           "The alignment of the popup within its `{ position }`",
		insertText:       `up-align="${1|top,right,center,bottom,left|}"`,
		insertTextFormat: snippetInsertTextFormat,
		choices:          []string{"top", "right", "center", "bottom", "left"},
	},
	"up-modal": {
		detail:     "Clicking this link will load the destination via AJAX and open the given selector in a modal overlay",
		insertText: `up-modal="$1"`, insertTextFormat: snippetInsertTextFormat, isDepreciated: true,
	},
	"up-popup": {
		detail:     "Clicking this link will load the destination via AJAX and open the given selector in a popup overlay",
		insertText: `up-popup="$1"`, insertTextFormat: snippetInsertTextFormat, isDepreciated: true,
	},
	"up-etag": {
		detail:     "Sets an ETag for the fragment's underlying data",
		insertText: `up-etag="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-id": {
		detail:     "Sets an unique identifier for this element",
		insertText: `up-id="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-keep": {
		detail:     "Elements with an `[up-keep]` attribute will be persisted during fragment updates",
		insertText: "up-keep", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-on-keep": {
		detail:     "Code to run before an existing element is kept during a page update",
		insertText: `up-on-keep="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-main": {
		detail:     "Marks this element as the primary content element of your application layout",
		insertText: "up-main", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-source": {
		detail:     "The URL from which this element and its descendants were initially requested",
		insertText: `up-source="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-time": {
		detail:     "Sets the time when the fragment's underlying data was last changed",
		insertText: `up-time="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-flashes": {
		detail:     "Use an `[up-flashes]` element to show confirmations, alerts or warnings",
		insertText: "up-flashes", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-hungry": {
		detail:     "Elements with an `[up-hungry]` attribute are updated whenever the server sends a matching element, even if the element isn't targeted",
		insertText: "up-hungry", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-on-hungry": {
		detail:     "Code to run before this element is included in a fragment update",
		insertText: `up-on-hungry="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-if-layer": {
		detail:     "Only piggy-back on updates on layers that match the given layer reference",
		insertText: `up-if-layer="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-poll": {
		detail:     "Elements with an `[up-poll]` attribute are reloaded from the server periodically",
		insertText: "up-poll", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-interval": {
		detail:     "The reload interval in milliseconds",
		insertText: `up-interval="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-keep-data": {
		detail:     "Whether to preserve the polling fragment's data object through reloads",
		insertText: `up-keep-data="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-alias": {
		detail:     "A URL pattern with alternative URLs",
		insertText: `up-alias="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-nav": {
		detail:     "Marks this element as a navigational container, such as a menu or navigation bar",
		insertText: "up-nav", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-emit": {
		detail:     "The type of the event to be emitted, e.g. `my:event`",
		insertText: `up-emit="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-emit-props": {
		detail:     "The event properties, serialized as relaxed JSON",
		insertText: `up-emit-props="$1"`, insertTextFormat: snippetInsertTextFormat,
	},
	"up-anchored": {
		detail:     "Marks this element as being anchored to the right edge of the screen, typically fixed navigation bars",
		insertText: `up-anchored="right"`, insertTextFormat: plainTextInsertTextFormat,
	},
	"up-fixed": {
		detail:     "Marks this element as being fixed to the top/bootm edge of the screen using `position: fixed`",
		insertText: `up-fixed="${1|top,bottom|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"top", "bottom"},
	},
	"up-viewport": {
		detail:     `Marks this element as a scrolling container ("viewport")`,
		insertText: "up-viewport", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-back": {
		detail:     "Changes the link's destination so it points to the previous URL",
		insertText: "up-back", insertTextFormat: plainTextInsertTextFormat,
	},
	"up-meta": {
		detail:     "Configures whether this `<head>` element is updated during history changes",
		insertText: `up-meta="${1|true,false|}"`, insertTextFormat: snippetInsertTextFormat,
		choices: []string{"true", "false"},
	},
	"up-boot": {
		detail:     "Prevent Unpoly from booting automatically",
		insertText: `up-boot="manual"`, insertTextFormat: plainTextInsertTextFormat,
	},
}

//go:embed attributes/*.md
var attributesFS embed.FS

// Must be called before performing any requests.
func LoadAttributes() error {
	for label := range attributes {
		contents, err := attributesFS.ReadFile(fmt.Sprintf("attributes/%v.md", label))
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return fmt.Errorf("unable to read file contents - %v.md, %w", label, err)
		}
		attributes[label].documentation = string(contents)
	}
	return nil
}
