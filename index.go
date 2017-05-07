package webls

import (
	"html/template"
	"os"
	"path/filepath"

	"encoding/json"

	"net/http"

	"fmt"

	"io/ioutil"

	"io"

	"mime"

	"net/url"

	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//Common errors
var (
	ErrHidden = errors.New("this directory has listing disabled")
)

//Index serves file's from root
type Index struct {
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

	config, err := i.Config(hostpath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		i.Log.Error("get config", zap.Error(err))
		fmt.Fprintf(w, "failed to retrieve config.")
		return
	}

	infos, err := ioutil.ReadDir(hostpath)
	if err != nil {
		//try to open path as file or fail
		i.handleFile(hostpath, w, r)
		return
	}

	if config.Hide {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "this directory is hidden")
		return
	}

	if !strings.HasSuffix(cleaned, "/") {
		cleaned += "/"
	}

	//help template setup links to each directory in path
	type dir struct {
		Dir      string
		Compound string
	}
	dirs := strings.Split(filepath.Dir(cleaned), "/")
	compoundDirs := make([]dir, 0, len(dirs))
	//setup root link
	compoundDirs = append(compoundDirs, dir{
		Dir:      "/",
		Compound: "/",
	})
	for i := range dirs {
		if dirs[i] == "" {
			continue
		}
		compoundDirs = append(compoundDirs, dir{Dir: dirs[i] + "/", Compound: strings.Join(dirs[:i+1], "/") + "/"})
	}

	//prepare information for view
	finfos := make([]Finfo, len(infos))
	for i, info := range infos {
		path := filepath.Join(cleaned, infos[i].Name())
		name := info.Name()
		if info.IsDir() {
			path += "/"
			name += "/"
		}
		typ, _, _ := mime.ParseMediaType(mime.TypeByExtension(filepath.Ext(name)))
		finfos[i] = Finfo{
			Path:         path,
			Name:         name,
			Type:         typ,
			LastModified: info.ModTime().Format("01/02/2006 15:04:05"),
			Size:         humanize.Bytes(uint64(info.Size())),
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := i.ListingTemplate.Execute(w, map[string]interface{}{
		"dirs":  compoundDirs,
		"files": finfos,
	}); err != nil {
		i.Log.Error("exec template", zap.Error(err))
	}
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
	if _, err := io.Copy(w, fi); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to read file")
		i.Log.Error("read file", zap.Error(err))
	}
	return
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
	indexFi, err := os.OpenFile(filepath.Dir(hostpath)+indexConfigName, os.O_RDONLY, 0640)
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
