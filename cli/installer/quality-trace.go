package installer

import (
	"bytes"
	_ "embed"
	"html/template"

	"fmt"

	cliUI "github.com/kubeshop/tracetest/cli/ui"
)

func configureDemoApp(conf configuration, ui cliUI.UI) configuration {
	conf.set("demo.enable.pokeshop", !conf.Bool("installer.only_quality-trace"))
	conf.set("demo.enable.otel", false)

	switch conf.String("installer") {
	case "docker-compose":
		conf.set("demo.endpoint.pokeshop.http", "http://demo-api:8081")
		conf.set("demo.endpoint.pokeshop.grpc", "demo-rpc:8082")
		conf.set("demo.endpoint.pokeshop.kafka", "stream:9092")
		conf.set("demo.endpoint.otel.frontend", "http://otel-frontend:8084")
		conf.set("demo.endpoint.otel.product_catalog", "otel-productcatalogservice:3550")
		conf.set("demo.endpoint.otel.cart", "otel-cartservice:7070")
		conf.set("demo.endpoint.otel.checkout", "otel-checkoutservice:5050")
	case "kubernetes":
		conf.set("demo.endpoint.pokeshop.http", "http://demo-pokemon-api.demo")
		conf.set("demo.endpoint.pokeshop.grpc", "demo-pokemon-api.demo:8082")
		conf.set("demo.endpoint.pokeshop.kafka", "stream.demo:9092")
		conf.set("demo.endpoint.otel.frontend", "http://otel-frontend.otel-demo:8084")
		conf.set("demo.endpoint.otel.product_catalog", "otel-productcatalogservice.otel-demo:3550")
		conf.set("demo.endpoint.otel.cart", "otel-cartservice.otel-demo:7070")
		conf.set("demo.endpoint.otel.checkout", "otel-checkoutservice.otel-demo:5050")
	}

	return conf
}

func configureQualitytrace(conf configuration, ui cliUI.UI) configuration {
	conf = configureBackend(conf, ui)
	conf.set("quality-trace.analytics", true)

	return conf
}

func configureBackend(conf configuration, ui cliUI.UI) configuration {
	installBackend := !conf.Bool("installer.only_quality-trace")
	conf.set("quality-trace.backend.install", installBackend)

	if !installBackend {
		conf.set("quality-trace.backend.type", "")
		return conf
	}

	// default values
	switch conf.String("installer") {
	case "docker-compose":
		conf.set("quality-trace.backend.type", "otlp")
		conf.set("quality-trace.backend.tls.insecure", true)
		conf.set("quality-trace.backend.endpoint.collector", "http://otel-collector:4317")
		conf.set("quality-trace.backend.endpoint", "quality-trace:4317")
	case "kubernetes":
		conf.set("quality-trace.backend.type", "otlp")
		conf.set("quality-trace.backend.tls.insecure", true)
		conf.set("quality-trace.backend.endpoint.collector", "http://otel-collector.quality-trace:4317")
		conf.set("quality-trace.backend.endpoint", "quality-trace:4317")

	default:
		conf.set("quality-trace.backend.type", "")
	}

	return conf
}

//go:embed templates/config.yaml.tpl
var configTemplate string

func getQualitytraceConfigFileContents(pHost, pUser, pPasswd string, ui cliUI.UI, config configuration) []byte {
	vals := map[string]string{
		"pHost":   pHost,
		"pUser":   pUser,
		"pPasswd": pPasswd,
	}

	tpl, err := template.New("page").Parse(configTemplate)
	if err != nil {
		ui.Panic(fmt.Errorf("cannot parse config template: %w", err))
	}

	out := &bytes.Buffer{}
	tpl.Execute(out, vals)

	return out.Bytes()
}

//go:embed templates/provision.yaml.tpl
var provisionTemplate string

func getQualitytraceProvisionFileContents(ui cliUI.UI, config configuration) []byte {
	vals := map[string]string{
		"installBackend":   fmt.Sprintf("%t", config.Bool("quality-trace.backend.install")),
		"backendType":      config.String("quality-trace.backend.type"),
		"backendEndpoint":  config.String("quality-trace.backend.endpoint.query"),
		"backendInsecure":  config.String("quality-trace.backend.tls.insecure"),
		"backendAddresses": config.String("quality-trace.backend.addresses"),
		"backendIndex":     config.String("quality-trace.backend.index"),
		"backendToken":     config.String("quality-trace.backend.token"),
		"backendRealm":     config.String("quality-trace.backend.realm"),

		"analyticsEnabled": fmt.Sprintf("%t", config.Bool("quality-trace.analytics")),

		"enablePokeshopDemo": fmt.Sprintf("%t", config.Bool("demo.enable.pokeshop")),
		"enableOtelDemo":     fmt.Sprintf("%t", config.Bool("demo.enable.otel")),
		"pokeshopHttp":       config.String("demo.endpoint.pokeshop.http"),
		"pokeshopGrpc":       config.String("demo.endpoint.pokeshop.grpc"),
		"pokeshopKafka":      config.String("demo.endpoint.pokeshop.kafka"),
		"otelFrontend":       config.String("demo.endpoint.otel.frontend"),
		"otelProductCatalog": config.String("demo.endpoint.otel.product_catalog"),
		"otelCart":           config.String("demo.endpoint.otel.cart"),
		"otelCheckout":       config.String("demo.endpoint.otel.checkout"),
	}

	tpl, err := template.New("page").Parse(provisionTemplate)
	if err != nil {
		ui.Panic(fmt.Errorf("cannot parse config template: %w", err))
	}

	out := &bytes.Buffer{}
	tpl.Execute(out, vals)

	return out.Bytes()
}
