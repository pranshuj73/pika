# Prior knowledge baseline

User is new to Go but has already scaffolded the emulator skeleton unaided: package layout (`internal/cpu`, `internal/memory`, `internal/gba`), fixed-size arrays for registers (`[16]uint32`), typed constants, and an `InstructionState` enum via `iota`. That demonstrates basic Go syntax and package structure are in place.

Evidence: repo state 2026-10-07 (`internal/cpu/*.go`, `internal/memory/*.go`).

Implications: lessons can assume structs, packages, `iota`, and fixed arrays. Not yet demonstrated: bit manipulation, methods with pointer receivers, interfaces, testing, goroutines, profiling. Next lessons should target those through CPU work.
