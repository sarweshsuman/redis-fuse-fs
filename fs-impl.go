package main

import (
	"fmt"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func (fs *RedisFuseConfig) GetAttr(name string, ctx *fuse.Context) (*fuse.Attr, fuse.Status) {
	fmt.Print("Request to GetAttr of %v\n", name)
	return nil, fuse.ENOENT
}

func (fs *RedisFuseConfig) OpenDir(name string, ctx *fuse.Context) ([]fuse.DirEntry, fuse.Status) {
	fmt.Print("Request to Open Dir %v\n", name)
	return []fuse.DirEntry{}, fuse.OK
}

func (fs *RedisFuseConfig) Open(name string, flags uint32, ctx *fuse.Context) (nodefs.File, fuse.Status) {
	fmt.Print("Request to Open %v\n", name)
	return nil, fuse.ENOENT
}

func (fs *RedisFuseConfig) Create(name string, flags uint32, mode uint32, ctx *fuse.Context) (nodefs.File, fuse.Status) {
	fmt.Print("Request to Create %v\n", name)
	return nil, fuse.ENOENT
}

func (fs *RedisFuseConfig) Rename(oldName string, newName string, ctx *fuse.Context) fuse.Status {
	fmt.Print("Request to Rename %v\n", oldName)
	return fuse.ENOENT
}

func (fs *RedisFuseConfig) Unlink(name string, ctx *fuse.Context) fuse.Status {
	fmt.Print("Request to Unlink %v\n", name)
	return fuse.ENOENT
}

func (fs *RedisFuseConfig) Rmdir(name string, ctx *fuse.Context) fuse.Status {
	fmt.Print("Request to Rmdir %v\n", name)
	return fuse.ENOENT
}

func (fs *RedisFuseConfig) Mkdir(name string, mode uint32, ctx *fuse.Context) fuse.Status {
	fmt.Print("Request to Mkdir %v\n", name)
	return fuse.OK
}

/*
func (fs *RedisFuseConfig) nameToPattern(name string) string {
  return ""
}

func (fs *RedisFuseConfig) dirsToEntries(dir string, m map[string]bool) []fuse.DirEntry {
  return []fuse.DirEntry{}
}

func (fs *RedisFuseConfig) resToEntries(dir string, list []string, m map[string]bool) []fuse.DirEntry {
  return []fuse.DirEntry{}
}
*/
