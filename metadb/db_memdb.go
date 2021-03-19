// +build memdb

package metadb

import (
	tmdb "github.com/klyed/tm-db"
	"github.com/klyed/tm-db/memdb"
)

func memdbDBCreator(name, dir string) (tmdb.DB, error) {
	return memdb.NewDB(), nil
}

func init() { registerDBCreator(MemDBBackend, memdbDBCreator, false) }
