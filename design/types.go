package design

import (
	. "goa.design/goa/v3/dsl"
)

var HelloResponse = Type("HelloResponse", func() {
	Description("Response containing a greeting message")

	Attribute("message", String, "greeting message", func() {
		Example("Hello, World!")
	})

	Required("message")

})
