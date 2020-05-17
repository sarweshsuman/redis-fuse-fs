package main

import (
	"fmt"
	"path"

	"github.com/garyburd/redigo/redis"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func (fs *RedisFuseConfig) GetAttr(name string, ctx *fuse.Context) (*fuse.Attr, fuse.Status) {
	fmt.Printf("Request to GetAttr of %v\n", name)

	if name == "" {
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	}

	// ignore hidden files
	if string(name[0]) == "." {
		return nil, fuse.ENOENT
	}

	key := path.Base(name)

	conn := fs.RedisConnPool.Get()
	defer conn.Close()

	content, err1 := redis.String(conn.Do("GET", key))
	list, err2 := redis.Strings(conn.Do("KEYS", key+"/*"))

	switch {
	case err2 == nil && len(list) > 0:
		return &fuse.Attr{
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
		break
	case err1 == nil:
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644,
			Size: uint64(len(content)),
		}, fuse.OK
		break
	}

	return nil, fuse.ENOENT
}

func (fs *RedisFuseConfig) OpenDir(name string, ctx *fuse.Context) ([]fuse.DirEntry, fuse.Status) {
	fmt.Printf("Request to Open Dir %v\n", name)
	return []fuse.DirEntry{}, fuse.OK
}

func (fs *RedisFuseConfig) Open(name string, flags uint32, ctx *fuse.Context) (nodefs.File, fuse.Status) {
	fmt.Printf("Request to Open %v\n", name)
	return nil, fuse.ENOENT
}

func (fs *RedisFuseConfig) Create(name string, flags uint32, mode uint32, ctx *fuse.Context) (nodefs.File, fuse.Status) {
	fmt.Printf("Request to Create %v\n", name)
	return nil, fuse.ENOENT
}

func (fs *RedisFuseConfig) Rename(oldName string, newName string, ctx *fuse.Context) fuse.Status {
	fmt.Printf("Request to Rename %v\n", oldName)
	return fuse.ENOENT
}

func (fs *RedisFuseConfig) Unlink(name string, ctx *fuse.Context) fuse.Status {
	fmt.Printf("Request to Unlink %v\n", name)
	return fuse.ENOENT
}

func (fs *RedisFuseConfig) Rmdir(name string, ctx *fuse.Context) fuse.Status {
	fmt.Printf("Request to Rmdir %v\n", name)
	return fuse.ENOENT
}

func (fs *RedisFuseConfig) Mkdir(name string, mode uint32, ctx *fuse.Context) fuse.Status {
	fmt.Printf("Request to Mkdir %v\n", name)
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
