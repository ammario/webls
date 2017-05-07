package webls

const indexConfigName = ".index"

//DefaultIndexConfig is used if no index config is provided
var DefaultIndexConfig = &IndexConfig{}

//IndexConfig represents an index configuration or .index file
type IndexConfig struct {
	MaxFiles int    `json:"max_files"`
	Password string `json:"password"`
	Hide     bool   `json:"hide"`
}
