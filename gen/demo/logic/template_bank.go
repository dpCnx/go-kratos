package logic

//
// import (
// 	"bytes"
// 	"strings"
// 	"text/template"
// )
//
// var httpTemplate = `
// {{$svrType := .ServiceType}}
// {{$svrName := .ServiceName}}
//
// {{- range .MethodSets}}
// const Operation{{$svrType}}{{.OriginalName}} = "/{{$svrName}}/{{.OriginalName}}"
// {{- end}}
//
// type {{.ServiceType}}HTTPServer interface {
// {{- range .MethodSets}}
// 	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
// {{- end}}
// }
//
// type {{.ServiceType}} struct {
// 	server      {{.ServiceType}}HTTPServer
// 	router      gin.IRoutes
// 	middlewares []gin.HandlerFunc
// }
//
// func (r *{{.ServiceType}}) Register{{.ServiceType}}HTTPServer() {
// 	{{- range .Methods}}
// 	r.{{.Method}}("{{.Path}}", append(r.middlewares, r.{{.Name}})...)
// 	{{- end}}
// }
//
// {{range .Methods}}
// func (r *{{.ServiceType}}) {{.Name}} (c *gin.Context){
// 		var in {{.Request}}
// 		{{- if .HasBody}}
// 		if err := ctx.Bind(&in{{.Body}}); err != nil {
// 			return err
// 		}
//
// 		{{- if not (eq .Body "")}}
// 		if err := ctx.BindQuery(&in); err != nil {
// 			return err
// 		}
// 		{{- end}}
// 		{{- else}}
// 		if err := ctx.BindQuery(&in{{.Body}}); err != nil {
// 			return err
// 		}
// 		{{- end}}
// 		{{- if .HasVars}}
// 		if err := ctx.BindVars(&in); err != nil {
// 			return err
// 		}
// 		{{- end}}
//
// 		out, err := r.server.{{.Name}}(c.Request.Context(), &in)
// 		if err != nil {
// 			return r.resp.Error(c, err)
// 		}
// 		return r.resp.Success(c, out)
// }
// {{end}}
// `
//
// type serviceDesc struct {
// 	ServiceType string // Greeter
// 	ServiceName string // helloworld.Greeter
// 	Metadata    string // api/helloworld/helloworld.proto
// 	Methods     []*methodDesc
// 	MethodSets  map[string]*methodDesc
// }
//
// type methodDesc struct {
// 	// method
// 	Name         string
// 	OriginalName string // The parsed original name
// 	Num          int
// 	Request      string
// 	Reply        string
// 	// http_rule
// 	Path         string
// 	Method       string
// 	HasVars      bool
// 	HasBody      bool
// 	Body         string
// 	ResponseBody string
// }
//
// func (s *serviceDesc) execute() string {
// 	s.MethodSets = make(map[string]*methodDesc)
// 	for _, m := range s.Methods {
// 		s.MethodSets[m.Name] = m
// 	}
// 	buf := new(bytes.Buffer)
// 	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := tmpl.Execute(buf, s); err != nil {
// 		panic(err)
// 	}
// 	return strings.Trim(buf.String(), "\r\n")
// }
