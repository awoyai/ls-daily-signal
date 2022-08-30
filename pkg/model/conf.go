package model

import "encoding/json"

type Data struct {
	Driver          string `json:"driver,omitempty"`
	DbUser          string `json:"db_user,omitempty"`
	DbPassword      string `json:"db_password,omitempty"`
	DbUrl           string `json:"db_url,omitempty"`
	DryRun          bool   `json:"dry_run,omitempty"`
	MaxIdleConns    int64  `json:"max_idle_conns,omitempty"`
	MaxOpenConns    int64  `json:"max_open_conns,omitempty"`
	ConnMaxLifeTime string `json:"conn_max_life_time,omitempty"`
	DbName          string `json:"db_name,omitempty"`
}

type Conf struct {
	Data        Data        `json:"data,omitempty"`
	PlateConfig PlateConfig `json:"plate_config,omitempty"`
}
type PlateConfig struct {
	PlateName    json.RawMessage `json:"plate_name,omitempty"`
	PlateSymbols json.RawMessage `json:"plate_symbols,omitempty"`
	FilePath     string          `json:"file_path,omitempty"`
}
