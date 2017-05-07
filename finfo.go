package webls

import (
	"mime"
	"os"
	"path/filepath"

	humanize "github.com/dustin/go-humanize"
)

//Finfo contains file info prettified
type Finfo struct {
	Path         string
	Name         string
	Type         string
	LastModified string
	Size         string
}

//ConvertFileInfos converts file infos into "prettier" versions useful for our index view
//cleaned should contain the clean path as accessed by a web client.
func ConvertFileInfos(cleaned string, infos []os.FileInfo) []Finfo {
	//prepare information for view
	finfos := make([]Finfo, len(infos))
	for i, info := range infos {
		path := filepath.Join(cleaned, info.Name())
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
	return finfos
}
