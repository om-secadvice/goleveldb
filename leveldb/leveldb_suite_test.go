package leveldb

import (
	"testing"

	"github.com/om-secadvice/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
