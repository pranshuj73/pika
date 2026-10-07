package memory

import "github.com/pranshuj73/pika.git/internal"

type IWRAM [32 * internal.KB]byte
type EWRAM [256 * internal.KB]byte
