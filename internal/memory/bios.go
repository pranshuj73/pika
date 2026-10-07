// Package memory
package memory

import "github.com/pranshuj73/pika.git/internal"

type BIOS struct {
	buffer [16 * internal.KB]byte
}
