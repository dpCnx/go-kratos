package swagger

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/api/metadata"
)

//go:embed dist
var swaggerUI embed.FS

func Router(router *gin.Engine, path string) {

	opts := &options{
		// Compatible with default UseJSONNamesForFields is true
		generatorOptions: []generator.Option{generator.UseJSONNamesForFields(true)},
	}

	service := New(nil, opts.generatorOptions...)

	r := router.Group(path)

	r.GET("/swagger/services", func(c *gin.Context) {
		services, err := service.ListServices(c.Request.Context(), &metadata.ListServicesRequest{})
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(200, services)
	})

	r.GET("/swagger/service/:name", func(c *gin.Context) {
		var in metadata.GetServiceDescRequest
		in.Name = c.Param("name")

		content, err := service.GetServiceOpenAPI(c.Request.Context(), &in, false)
		if err != nil {
			c.String(400, err.Error())
			return
		}
		c.String(http.StatusOK, content)

	})

	r.StaticFS("/_swagger/static", http.FS(swaggerUI))

	templ := template.Must(template.New("").ParseFS(swaggerUI, "dist/*.html"))
	router.SetHTMLTemplate(templ)

	r.GET("/_swagger", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
