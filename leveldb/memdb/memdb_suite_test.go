package memdb

import (
	"testing"

	"github.com/om-secadvice/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
