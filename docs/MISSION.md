# Mission: Build a terminal-based GBA emulator in Go

## Why
The emulator is a vehicle: a project big enough to force deep, idiomatic Go — performance, concurrency (goroutines), and architecture — rather than tutorial-level familiarity. The self-imposed constraint that it must run fully in the terminal (unlike existing emulators) makes it a genuine engineering challenge, not a paint-by-numbers port.

## Success looks like
- A fully functional GBA emulator: boots BIOS, runs commercial ROMs with correct timing/accuracy (not a rudimentary implementation)
- Runs entirely in the terminal (no windowing toolkit) — rendering solved in tty
- Idiomatic, concurrent Go: goroutines/channels used where they genuinely help, with measured performance
- Can read other Go emulator codebases (Lycan, DawnGBA) fluently and critique them

## Constraints
- Reference sources: grilling ChatGPT, github.com/Aditya-1304/Lycan, github.com/akatsuki105/dawngba
- New to Go — lessons must teach Go mechanics alongside emulator concepts
- Self-directed pace

## Out of scope
- Non-GBA systems (no GB/GBC/DS)
- Writing an assembler/compiler toolchain
- Hardware-level accuracy (cycle-accurate down to dots) unless the mission demands it later
