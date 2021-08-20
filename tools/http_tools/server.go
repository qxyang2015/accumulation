package http_tools

import (
	"encoding/json"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/qxyang2015/accumulation/tools/error_tools"
	"github.com/qxyang2015/accumulation/tools/plugin"
	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type Config struct {
	Server struct {
		Host           string
		Log            string
		ReadTimeOut    time.Duration
		WriteTimeOut   time.Duration
		MaxConcurrency int
	}
	Modules []struct {
		Name     string //服务名称标识
		Conf     string //服务请求日志配置文件地址
		Location string //服务路由地址
	}
	Router *router.Router
}

func ServeInit(yFile string) *Config {
	maxProcs := runtime.NumCPU() //获取cpu个数
	runtime.GOMAXPROCS(maxProcs * 4)

	data, err := ioutil.ReadFile(yFile)
	error_tools.Assert(err)
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	error_tools.Assert(err)

	// Init modules
	for i := 0; i < len(conf.Modules); i++ {
		plugin.Open(conf.Modules[i].Name, conf.Modules[i].Conf)
	}
	// Dispatch locations
	conf.Router = router.New()
	for i := 0; i < len(conf.Modules); i++ {
		plugin.Dispatch(conf.Modules[i].Location, conf.Modules[i].Name, conf.Router)
	}

	return &conf

}

func ServeMain(conf *Config) {

	server := &fasthttp.Server{
		Handler:            conf.Router.Handler,
		ReadTimeout:        conf.Server.ReadTimeOut * time.Millisecond,
		WriteTimeout:       conf.Server.WriteTimeOut * time.Millisecond,
		IdleTimeout:        100 * time.Second, //服务端连接空闲超时时间,固定100s
		MaxRequestBodySize: 500 * 1024 * 1024, //最大500M

	}
	if conf.Server.MaxConcurrency > 0 {
		server.Concurrency = conf.Server.MaxConcurrency
	}
	log.Fatal(server.ListenAndServe(conf.Server.Host))
}

func HttpResponse(ctx *fasthttp.RequestCtx, res interface{}) error {
	//log4sys := logging.GetLogger()

	b, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("json.Marshal转化返回值错误, err:%s", err)
	}

	//log4sys.Debug("res json marshal complete[%s]", string(b))

	_, err = fmt.Fprint(ctx, string(b))
	return err
}

func ServeMainWithQuit(conf *Config) {

	server := &fasthttp.Server{
		//ReadBufferSize: 30 * 1024*1024,
		Handler:            conf.Router.Handler,
		ReadTimeout:        conf.Server.ReadTimeOut * time.Millisecond,
		WriteTimeout:       conf.Server.WriteTimeOut * time.Millisecond,
		MaxRequestBodySize: 500 * 1024 * 1024, //最大500M

	}
	if conf.Server.MaxConcurrency > 0 {
		server.Concurrency = conf.Server.MaxConcurrency
	}
	go server.ListenAndServe(conf.Server.Host)
	//监听退出信号
	listenSignal(server)
}

func listenSignal(server *fasthttp.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	//signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	select {
	case v := <-sigs:
		fmt.Printf("接收到退出信号[%v]\n", v)
		server.Shutdown()
		fmt.Println("服务完成安全退出")
	}
}
