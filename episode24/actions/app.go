package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/arschles/go-in-5-minutes/episode24/models"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_episode24_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/", HomeHandler)
		app.GET("/other/{name}", OtherHandler)

		// Create a new path group. The 'apiGroup' returned by this method is like the 'app'
		// variable - you can define HTTP routes & resources against it. The difference is that
		// all of them are prefixed with /api in the path
		apiGroup := app.Group("/api")

		// Now create a v1 group - everything on apiV1Group will now get prefixed with
		// '/api/v1'. That's really nice to do API versioning
		apiV1Group := apiGroup.Group("/v1")
		apiV1Group.GET("/things", apiV1ThingsHandler)

		// Do the same thing for a v2 API
		apiV2Group := apiGroup.Group("/v2")
		apiV2Group.GET("/things", apiV2ThingsHandler)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}
