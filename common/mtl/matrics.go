package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registy *prometheus.Registry

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	Registy = prometheus.NewRegistry()
	Registy.MustRegister(collectors.NewGoCollector())
	Registy.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	r, _ := consul.NewConsulRegister(registryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registryInfo := &registry.Info{
		ServiceName: "Prometheus",
		Addr:        addr,
		Weight:      1,
		Tags: map[string]string{
			"service": serviceName,
		},
	}
	_ = r.Register(registryInfo)
	server.RegisterShutdownHook(func() {
		_ = r.Deregister(registryInfo)
	})
	http.Handle("/metrics", promhttp.HandlerFor(Registy, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
	return r, registryInfo
}
