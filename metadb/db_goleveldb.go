// +build goleveldb

package metadb

import (
	tmdb "github.com/klyed/tm-db"
	"github.com/klyed/tm-db/goleveldb"
)

func golevelDBCreator(name, dir string) (tmdb.DB, error) {
	return goleveldb.NewDB(name, dir)
}

func init() { registerDBCreator(GoLevelDBBackend, golevelDBCreator, true) }
