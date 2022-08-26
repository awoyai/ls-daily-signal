package model

type Data struct {
	Driver          string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	DbUser          string `protobuf:"bytes,2,opt,name=db_user,json=dbUser,proto3" json:"db_user,omitempty"`
	DbPassword      string `protobuf:"bytes,3,opt,name=db_password,json=dbPassword,proto3" json:"db_password,omitempty"`
	DbUrl           string `protobuf:"bytes,4,opt,name=db_url,json=dbUrl,proto3" json:"db_url,omitempty"`
	DryRun          bool   `protobuf:"varint,5,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	MaxIdleConns    int64  `protobuf:"varint,6,opt,name=max_idle_conns,json=maxIdleConns,proto3" json:"max_idle_conns,omitempty"`
	MaxOpenConns    int64  `protobuf:"varint,7,opt,name=max_open_conns,json=maxOpenConns,proto3" json:"max_open_conns,omitempty"`
	ConnMaxLifeTime string `protobuf:"bytes,8,opt,name=conn_max_life_time,json=connMaxLifeTime,proto3" json:"conn_max_life_time,omitempty"`
	DbName          string `protobuf:"bytes,9,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}
