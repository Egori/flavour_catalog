package app

import (
	"context"
	"flag"
	"flavor/internal/config"
	apihandler "flavor/internal/handlers/api"
	htmlHandler "flavor/internal/handlers/html"
	"os"

	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		//a.GenerateProducts,
		//a.UpdateRandProd,
		a.initApp,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	envFileName := ".env"

	flagset := flag.CommandLine
	flagset.StringVar(&envFileName, "env", envFileName, "the env file which web app will use to extract its environment variables")
	flagset.Parse(os.Args[1:])

	config.Load(envFileName)
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	a.serviceProvider.InitServices()
	return nil
}

func (a *App) initApp(_ context.Context) error {

	allTemplates := template.Must(template.ParseGlob("../../internal/templates/v1/views/*.html"))

	allTemplates = template.Must(allTemplates.ParseGlob("../../internal/templates/v1/views/partials/*.html"))
	allTemplates = template.Must(allTemplates.ParseGlob("../../internal/templates/v1/views/product/*.html"))
	allTemplates = template.Must(allTemplates.ParseGlob("../../internal/templates/v1/views/prod_list/*.html"))
	allTemplates = template.Must(allTemplates.ParseGlob("../../internal/templates/v1/views/category/*.html"))

	template := &Template{
		templates: allTemplates,
	}

	app := echo.New()
	app.Renderer = template

	catalogHandler := htmlHandler.NewCatalogHandler(a.serviceProvider.catalogservice)

	app.Static("/", "../../internal/templates/v1/static")
	app.Static("/img", "../../img")

	app.GET("/main", catalogHandler.GetTest)

	// app.GET("/catalog/:page", catalogHandler.GetMain)

	// app.GET("/catalogPage/:page", catalogHandler.GetProdListAll)

	app.GET("/category/:path", catalogHandler.GetCategoryPage)

	app.GET("/product/:path", catalogHandler.GetProdPage)

	app.GET("/catalog/:path", catalogHandler.GetCatalogPageByCategoryOwn)
	//app.GET("/catalog/list/:path", catalogHandler.GetCatalogListByCategoryOwn)

	categoryApiHandler := apihandler.NewCategoryHandler(*a.serviceProvider.catalogservice)
	app.GET("/api/category/", categoryApiHandler.GetAllMain)
	app.POST("/api/category/", categoryApiHandler.Add)

	app.Logger.Fatal(app.Start(":3333"))
	return nil
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
