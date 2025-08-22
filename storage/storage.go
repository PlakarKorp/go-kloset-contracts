package storage

import (
	"context"
	"io"

	"github.com/PlakarKorp/go-kloset-contracts/objects"
)

type Mode uint32

const (
	ModeWrite Mode = 1 << 1
	ModeRead  Mode = 1 << 2
)

type Store interface {
	Create(ctx context.Context, config []byte) error
	Open(ctx context.Context) ([]byte, error)
	Location(ctx context.Context) (string, error)
	Mode(ctx context.Context) (Mode, error)
	Size(ctx context.Context) (int64, error) // this can be costly, call with caution

	GetStates(ctx context.Context) ([]objects.MAC, error)
	PutState(ctx context.Context, mac objects.MAC, rd io.Reader) (int64, error)
	GetState(ctx context.Context, mac objects.MAC) (io.ReadCloser, error)
	DeleteState(ctx context.Context, mac objects.MAC) error

	GetPackfiles(ctx context.Context) ([]objects.MAC, error)
	PutPackfile(ctx context.Context, mac objects.MAC, rd io.Reader) (int64, error)
	GetPackfile(ctx context.Context, mac objects.MAC) (io.ReadCloser, error)
	GetPackfileBlob(ctx context.Context, mac objects.MAC, offset uint64, length uint32) (io.ReadCloser, error)
	DeletePackfile(ctx context.Context, mac objects.MAC) error

	GetLocks(ctx context.Context) ([]objects.MAC, error)
	PutLock(ctx context.Context, lockID objects.MAC, rd io.Reader) (int64, error)
	GetLock(ctx context.Context, lockID objects.MAC) (io.ReadCloser, error)
	DeleteLock(ctx context.Context, lockID objects.MAC) error

	Close(ctx context.Context) error
}

type StoreFn func(context.Context, string, map[string]string) (Store, error)
