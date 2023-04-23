package plugin_storage_123

import (
	"github.com/alist-org/alist/v3/internal/driver"
)

type Addition struct {
	Username       string `json:"username" required:"true"`
	Password       string `json:"password" required:"true"`
	RootFolderID   string `json:"root_folder_id"`
	OrderBy        string `json:"order_by" type:"select" options:"file_name,size,update_at" default:"file_name"`
	OrderDirection string `json:"order_direction" type:"select" options:"asc,desc" default:"asc"`
	StreamUpload   bool   `json:"stream_upload"`
	AccessToken    string
}

var config = driver.Config{
	Name:        "123Pan(Plugin)",
	DefaultRoot: "0",
}
