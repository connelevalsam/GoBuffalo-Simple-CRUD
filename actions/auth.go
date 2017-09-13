package actions

import "github.com/gobuffalo/buffalo"

// AuthUpdate default implementation.
func AuthUpdate(c buffalo.Context) error {
	return c.Render(200, r.HTML("auth/update.html"))
}
