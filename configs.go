package main

import (
	"flag"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	"github.com/sarweshsuman/redis-fuse-fs/redisutils"
)

type RedisFuseConfig struct {
	pathfs.FileSystem
	FsRootPath  string
	MountOps    fuse.MountOptions
	redisConfig *redisutils.RedisConfig
	NodeFsConv  *pathfs.PathNodeFs
	FsConnector *nodefs.FileSystemConnector
}

func initializeFlags() *RedisFuseConfig {
	config := &RedisFuseConfig{}
	config.redisConfig = &redisutils.RedisConfig{}
	config.redisConfig.InitializeRedisFlags()

	flag.StringVar(&config.FsRootPath, "fs", "DEFAULT_NONE", "filesystem to mount")
	config.MountOps = fuse.MountOptions{}
	flag.BoolVar(&config.MountOps.AllowOther, "allow-other", false, "allow other users to access this mount point")

	return config
}

func (fsConfig *RedisFuseConfig) init() {
	fsConfig.FileSystem = pathfs.NewDefaultFileSystem()
	fsConfig.NodeFsConv = pathfs.NewPathNodeFs(fsConfig, nil)
	fsConfig.FsConnector = nodefs.NewFileSystemConnector(fsConfig.NodeFsConv.Root(), nil)
}
