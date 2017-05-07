package webls

import (
	"html/template"
	"os"
	"path/filepath"

	"encoding/json"

	"net/http"

	"fmt"

	"io/ioutil"

	"net/url"

	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//Common errors
var (
	ErrHidden = errors.New("this directory has listing disabled")
)

//Index serves file's from root
type Index struct {
	Brand           string
	BrandURL        string
	ListingTemplate *template.Template
	Log             *zap.Logger
	Root            string
}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cleaned, hostpath := i.securePath(r.RequestURI)

	i.Log.Info("request",
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("method", r.Method),
		zap.String("path", r.RequestURI),
	)

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "only GET is allowed to this server")
		return
	}

	infos, err := ioutil.ReadDir(hostpath)
	if err != nil {
		//try to open path as file or fail
		i.handleFile(hostpath, w, r)
		return
	}

	config, err := i.Config(hostpath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		i.Log.Error("get config", zap.Error(err))
		fmt.Fprintf(w, "failed to retrieve config.")
		return
	}

	if config.Hide {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "this directory has listing disabled")
		return
	}

	if config.MaxFiles > 0 && len(infos) > config.MaxFiles {
		infos = infos[:config.MaxFiles]
	}

	//must be directory listing
	i.handleDir(cleaned, hostpath, infos, w, r)
	return
}

func (i *Index) handleFile(hostpath string, w http.ResponseWriter, r *http.Request) {
	fi, err := os.OpenFile(hostpath, os.O_RDONLY, 0640)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "file or directory not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to open file")
		i.Log.Error("open file", zap.Error(err))
		return
	}

	info, err := fi.Stat()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to stat file")
		i.Log.Error("stat file",
			zap.String("file", fi.Name()),
			zap.Error(err),
		)
		return
	}

	http.ServeContent(w, r, fi.Name(), info.ModTime(), fi)
	return
}

func (i *Index) handleDir(cleaned string, hostpath string, infos []os.FileInfo, w http.ResponseWriter, r *http.Request) {
	//better aesthetics
	if !strings.HasSuffix(cleaned, "/") {
		cleaned += "/"
	}

	w.WriteHeader(http.StatusOK)
	if err := i.ListingTemplate.Execute(w, map[string]interface{}{
		"Dirs":     SplitPath(cleaned),
		"Files":    ConvertFileInfos(cleaned, infos),
		"Brand":    i.Brand,
		"BrandURL": i.BrandURL,
	}); err != nil {
		i.Log.Error("exec template", zap.Error(err))
	}
}

//securePath secures the provided path
func (i *Index) securePath(path string) (cleaned string, hostpath string) {
	path, err := url.QueryUnescape(path)
	if err != nil {
		return "/", i.Root
	}
	cleaned = filepath.Clean(path)
	hostpath = i.Root + cleaned
	return
}

//Config returns the index config for a hostpath
func (i *Index) Config(hostpath string) (*IndexConfig, error) {
	confName := hostpath + "/" + indexConfigName
	i.Log.Debug(".index check", zap.String("name", confName))
	indexFi, err := os.OpenFile(confName, os.O_RDONLY, 0640)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultIndexConfig, nil
		}
		return nil, errors.Wrap(err, "failed to load index file")
	}
	defer indexFi.Close()

	iconfig := &IndexConfig{}
	if err := json.NewDecoder(indexFi).Decode(iconfig); err != nil {
		return nil, errors.Wrap(err, "failed to decode config")
	}
	return iconfig, nil
}

//CompoundPath contains a fragment of a path and all of it's preceding paths
type CompoundPath struct {
	Full    string
	Segment string
}

//SplitPath splits a path into it's compound paths
func SplitPath(path string) []CompoundPath {
	dirs := strings.Split(filepath.Dir(path), "/")

	cpaths := make([]CompoundPath, 0, len(dirs))
	//setup root link
	cpaths = append(cpaths, CompoundPath{
		Full:    "/",
		Segment: "/",
	})

	for i := range dirs {
		if dirs[i] == "" {
			continue
		}
		cpaths = append(cpaths, CompoundPath{Segment: dirs[i] + "/", Full: strings.Join(dirs[:i+1], "/") + "/"})
	}
	return cpaths
}
