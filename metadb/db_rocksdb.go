// +build rocksdb

package metadb

import (
	tmdb "github.com/klyed/tm-db"
	"github.com/klyed/tm-db/rocksdb"
)

func rocksDBCreator(name, dir string) (tmdb.DB, error) {
	return rocksdb.NewDB(name, dir)
}

func init() { registerDBCreator(RocksDBBackend, rocksDBCreator, true) }
