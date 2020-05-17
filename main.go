package main

import (
	"flag"

	"github.com/hanwen/go-fuse/fuse"
)

func main() {
	fsConfig := initializeFlags()
	flag.Parse()

	fsConfig.init()

	server, err := fuse.NewServer(fsConfig.FsConnector.RawFS(), fsConfig.FsRootPath, &fsConfig.MountOps)
	must(err)
	server.Serve()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
