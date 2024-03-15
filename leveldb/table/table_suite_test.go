package table

import (
	"testing"

	"github.com/om-secadvice/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
