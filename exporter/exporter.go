package exporter

import (
	"context"
	"io"

	"github.com/PlakarKorp/kloset-contracts/objects"
)

type Options struct {
	MaxConcurrency uint64

	Stdout io.Writer
	Stderr io.Writer
}

type LinkType int

const (
	HARDLINK LinkType = iota
	SYMLINK
)

type Exporter interface {
	Root(ctx context.Context) (string, error)
	CreateDirectory(ctx context.Context, pathname string) error
	CreateLink(ctx context.Context, oldname string, newname string, ltype LinkType) error
	StoreFile(ctx context.Context, pathname string, fp io.Reader, size int64) error
	SetPermissions(ctx context.Context, pathname string, fileinfo *objects.FileInfo) error
	Close(ctx context.Context) error
}

type ExporterFn func(context.Context, *Options, string, map[string]string) (Exporter, error)
