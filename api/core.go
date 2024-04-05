//		QMY Admin
//
//		This is QMY Admin api doc
//
//		Schemes: http, https
//		Host: localhost:9100
//		BasePath: /
//		Version: 1.3.11
//		Contact: 407193275@qq.com
//		SecurityDefinitions:
//		  Token:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//	    Consumes:
//		  - application/json
//
//		Produces:
//		  - application/json
//
// swagger:meta
package main

import (
	"flag"
	"fmt"

	"github.com/qmcloud/admin-core/api/internal/config"
	"github.com/qmcloud/admin-core/api/internal/handler"
	"github.com/qmcloud/admin-core/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithCors(c.CROSConf.Address))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
