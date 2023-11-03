package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/discovery/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	srcgrpc "google.golang.org/grpc"
	"juice/app/user/api/user/v1"
	"juice/app/user/internal/conf"
	"os"
	"time"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	//服务注册，本地跑可以先注释掉
	r := discovery.New(&discovery.Config{
		Nodes:  []string{"0.0.0.0:7171"},
		Env:    "dev",
		Region: "sh1",
		Zone:   "zone1",
		Host:   "hostname",
	})
	Name = "user"
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{"color": "gray"}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	)
}

func callGRPC(conn *srcgrpc.ClientConn) {
	client := v1.NewUserBasicClient(conn)
	reply, err := client.GetUserInfo(context.Background(), &v1.GetUserInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("[grpc] SayHello %+v\n", reply)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	r := discovery.New(&discovery.Config{
		Nodes:  []string{"0.0.0.0:7171"},
		Env:    "dev",
		Region: "sh1",
		Zone:   "zone1",
		Host:   "localhost",
	})

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///video"),
		// use discovery
		grpc.WithDiscovery(r),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		callGRPC(conn)
		time.Sleep(time.Second)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
