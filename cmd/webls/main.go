package main

import (
	"net/http"

	"runtime"

	"html/template"

	"github.com/alecthomas/kingpin"
	"github.com/ammario/webls"
	"go.uber.org/zap"
)

func main() {
	//flags
	var (
		brand    = kingpin.Flag("brand", "brand displayed on the bottom of all listings").Default("webls").String()
		brandURL = kingpin.Flag("brand-url", "url which brand is linked to").Default("//github.com/ammario/webls").String()
		root     = kingpin.Arg("root", "directory to serve").Required().ExistingDir()
		bindAddr = kingpin.Flag("address", "Address to bind to").Default(":80").String()
	)
	kingpin.Parse()

	//logging
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	//get listing template
	index := &webls.Index{
		Brand:           *brand,
		BrandURL:        *brandURL,
		ListingTemplate: template.Must(template.New("listing.tmpl").Funcs(webls.TemplateFuncs).Parse(string(MustAsset("listing.tmpl")))),
		Root:            *root,
		Log:             log,
	}

	//http
	go func() {
		log.Fatal("Failed to server HTTP", zap.Error(http.ListenAndServe(*bindAddr, index)))
	}()
	runtime.Gosched()

	log.Info("Started HTTP server", zap.String("bind addr", *bindAddr))
	select {}
}
