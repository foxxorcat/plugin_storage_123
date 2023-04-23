package plugin_storage_123

import (
	"github.com/alist-org/alist/v3/internal/model"
	yaegi_storage "github.com/alist-org/alist/v3/internal/plugin/yaegi/adapter/storage"
)

func Plugin() (model.PluginControl, error) {
	return &model.PluginControlHelper{
		LoadFunc: func() error {
			yaegi_storage.RegisterPluginDriver(func() yaegi_storage.DriverPlugin {
				return &Pan123{}
			})
			yaegi_storage.LoadPluginStorage(config.Name)
			return nil
		},
		UnloadFunc: func() error {
			yaegi_storage.UnRegisterPluginDriver(func() yaegi_storage.DriverPlugin {
				return &Pan123{}
			})
			yaegi_storage.DropPluginStorage(config.Name)
			return nil
		},
		ConfigFunc: func() (model.PluginConfig, error) {
			return &model.PluginConfigHelper{
				UUID:        "5EB1E22C-9B3B-568E-554D-570ED3C2EA0F",
				Types:       []string{"storage"},
				Name:        "pan123(plugin)",
				Version:     "v1.0.0",
				ApiVersions: []string{"v1"},
			}, nil
		},
		ReleaseFunc: func() error {
			return nil
		},
	}, nil
}
