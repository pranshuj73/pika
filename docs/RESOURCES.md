# GBA Emulation & Go Resources

## Knowledge (GBA hardware)

- [GBATEK — GBA Hardware Manual by Martin Korth](https://problemkaputt.de/gbatek.htm)
  The single most authoritative free reference for GBA internals: memory map, ARM7TDMI, PPU timing, sound, DMA. Use for: every hardware question, before any LLM. This is ground truth.
- [Cowbite GBA Hardware Spec by Jeff Frohwein](https://www.cs.rit.edu/~tjh8300/CowBite/CowBiteSpec.htm)
  Older but very readable alternative to GBATEK. Use for: a second opinion when GBATEK's phrasing is dense.
- [ARM7TDMI Technical Reference Manual (ARM)](https://developer.arm.com/documentation/ddi0210)
  Official ARM docs for the CPU core. Use for: instruction semantics edge cases, exceptions, modes.
- [GBA Wikipedia: Memory Map](https://en.wikipedia.org/wiki/Game_Boy_Advance#Memory) — quick orientation only; verify against GBATEK.

## Knowledge (Go)

- [A Tour of Go](https://go.dev/tour/) — syntax baseline. Use for: rapid pattern lookup.
- [Effective Go](https://go.dev/doc/effective_go) — idioms. Use for: "is this how a Go programmer writes it?"
- [Go Wiki: CodeReviewComments](https://go.dev/wiki/CodeReviewComments) — the style bar. Use for: before committing anything.
- [The Go Memory Model](https://go.dev/ref/mem) — REQUIRED reading before touching goroutines for the emulator's clock/scheduler.

## Knowledge (reference emulators)

- [DawnGBA (akatsuki105)](https://github.com/akatsuki105/dawngba) — user's reference; clean modern Go GBA core. Use for: comparing architecture decisions after forming your own.
- [Lycan (Aditya-1304)](https://github.com/Aditya-1304/Lycan) — user's reference. Use for: a second Go implementation to diff against.

## Wisdom (Communities)

- [r/EmuDev](https://reddit.com/r/emudev)
  Emulator development subreddit; GBA questions common, wiki full of further docs. Use for: architecture sanity checks, "why is my PPU wrong".
- [Emulation Development Discord](https://discord.gg/EmuDev)
  The active real-time community behind r/EmuDev; many GBA emulator authors present. Use for: live debugging help, accuracy questions.

## Gaps
- No trusted terminal-rendering resource yet (tcell/termenv, frame-rate in tty). Needed for the terminal constraint — to research before the PPU phase.
- No performance-profiling resource yet (pprof, escape analysis) — needed for the performant-Go goal.
