package tests

import (
	"github.com/carbonblack/cb-event-forwarder/internal/cbapi"
	"github.com/carbonblack/cb-event-forwarder/internal/messageprocessor"
)

//var configHTTPTemplate *template.Template = template.Must(template.New("testtemplateconfig").Parse(
//`{"filename": "{{.FileName}}", "service": "carbonblack", "alerts":[{{range .Events}}{{.EventText}}{{end}}]}`))
//var config conf.Configuration = conf.Configuration{UploadEmptyFiles: false, BundleSizeMax: 1024 * 1024 * 1024, BundleSendTimeout: time.Duration(30) * time.Second, CbServerURL: "https://cbtests/", HTTPPostTemplate: configHTTPTemplate, DebugStore: ".", DebugFlag: true, EventMap: make(map[string]bool)}

var cbapihandler = cbapi.CbAPIHandler{}
var jsmp *messageprocessor.JsonMessageProcessor
var protobufEventMap = map[string]bool{
	"ingress.event.process":        true,
	"ingress.event.procstart":      true,
	"ingress.event.netconn":        true,
	"ingress.event.procend":        true,
	"ingress.event.childproc":      true,
	"ingress.event.moduleload":     true,
	"ingress.event.module":         true,
	"ingress.event.filemod":        true,
	"ingress.event.regmod":         true,
	"ingress.event.tamper":         true,
	"ingress.event.crossprocopen":  true,
	"ingress.event.remotethread":   true,
	"ingress.event.processblock":   true,
	"ingress.event.emetmitigation": true,
	"binaryinfo.#":                 true,
	"binarystore.#":                true,
	"events.partition.#":           true,
}
var pbmp = messageprocessor.PbMessageProcessor{EventMap: protobufEventMap}

var jsonEventMap = map[string]bool{
	"watchlist.hit.process": true,
}

func init() {
	jsmp = messageprocessor.NewJSONProcessor(messageprocessor.Config{
		DebugFlag:   false,
		DebugStore:  "/tmp",
		CbServerURL: "https://localhost/",
		CbAPI:       &cbapihandler,
		EventMap:    jsonEventMap,
	})
}