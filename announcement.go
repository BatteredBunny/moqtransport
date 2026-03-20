package moqtransport

import "github.com/BatteredBunny/moqtransport/internal/wire"

type announcement struct {
	requestID  uint64
	namespace  []string
	parameters wire.KVPList

	response chan error
}
