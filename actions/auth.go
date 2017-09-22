package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/connelevalsam/BuffaloProjects/simple-crud/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/pkg/errors"
	"github.com/markbates/pop"
	"database/sql"
	"strings"
	"github.com/markbates/validate"
)

// AuthUpdate default implementation.
func AdminLoginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin-login.html"))
}

//
func AdminHandler(c buffalo.Context) error {
	admin := &models.Admin{}
	// c.Bind will populate the struct with the form values
	if err := c.Bind(admin); err != nil {
		return errors.WithStack(err)
	}
	tx := c.Value("tx").(*pop.Connection)

	// find an admin with the username
	err := tx.Where("username = ?", strings.ToLower(admin.Username)).First(admin)

	// helper function to handle bad attempts
	bad := func() error {
		c.Set("admin", admin)
		verrs := validate.NewErrors()
		verrs.Add("username", "invalid username/password")
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("admin-login.html"))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the password from the db
	err = tx.Where("password = ?", admin.Password).First(admin)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied password address.
			return bad()
		}
	}
	c.Session().Set("current_user_id", admin.ID)
	c.Flash().Add("success", "Welcome Back!")

	return c.Redirect(302, "admin/lecturer")
}

func UserLogin(c buffalo.Context) error {
	return c.Render(200, r.HTML("login.html"))
}

//check admin login details
func AuthHandler(c buffalo.Context) error {
	// auth request
	username := c.Request().FormValue("uname")
	pass := c.Request().FormValue("pword")
	var u models.Admin
	err := models.DB.Where("username = ?", username).First(&u)

	if err != nil {

		r.HTMLLayout = "admin-login.html"
		//c.Set("HasLoginError", true)
		return c.Render(401, r.HTML("admin-login.html"))
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {

		r.HTMLLayout = "admin-login.html"
		//c.Set("HasLoginError", true)
		return c.Render(401, r.HTML("admin-login.html"))
	}

	// if OK, set username in session
	c.Session().Set("username", c.Request().FormValue("username"))
	c.Session().Save()

	r.HTMLLayout = "main.html"
	//c.Set("user", u)
	//c.Set("adminpage", makeAdminPage("Dashboard", "get stuff done!", "Index"))
	return c.Render(200, r.HTML("admin/lecturer.html"))
}