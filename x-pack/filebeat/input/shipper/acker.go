package shipper

import (
	"sync/atomic"

	"github.com/elastic/elastic-agent-libs/logp"
)

type shipperAcker struct {
	persistedIndex uint64
}

func newShipperAcker() *shipperAcker {
	return &shipperAcker{persistedIndex: 0}
}

// Update the input's persistedIndex by adding total to it.
// Despite the name, "total" here means an incremental total, i.e.
// the total number of events that are being acknowledged by this callback, not the total that have been sent overall.
// The acked parameter includes only those events that were successfully sent upstream rather than dropped by processors, etc.,
// but since we don't make that distinction in persistedIndex we can probably ignore it.
func (acker *shipperAcker) Track(_ int, total int) {
	logp.L().Infof("Tracking, total is %d", total)
	atomic.AddUint64(&acker.persistedIndex, uint64(total))
}

func (acker *shipperAcker) PersistedIndex() uint64 {
	return acker.persistedIndex
}
