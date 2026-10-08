# Notes

- User is new to Go; every lesson must pair an emulator concept with a Go mechanic actually exercised in their codebase.
- Repo `pika` already has skeleton: `internal/cpu` (Registers [16]uint32, CPSR uint32, Mode/ARM/THUMB enums), `internal/memory` (bus stub, IWRAM/EWRAM arrays), `internal/gba`. Module path: `github.com/pranshuj73/pika.git` (note the `.git` suffix — unusual, may bite later with tooling).
- User's stated learning loop: grills ChatGPT manually, reads Lycan + DawnGBA. Lessons should teach them to verify claims against GBATEK instead of trusting any single source.
- README shows current confusion points: "cpsr (wtf is cpsr?)", "cpu modes (again wtf?)", "barrel shifter (wtf is this now)". These are ready-made lesson topics.
- Self-imposed constraints to honor in teaching: terminal-only rendering, goroutines, performance.

- Preference: keep the repo clean — all teaching files live under docs/, never at repo root.
- Preference: ZERO assumed background on GBA/low-level/hardware concepts. Every lesson section gets an explicit ELI5 block (plain-English analogy) in addition to the technical explanation.
- Learning moment (LR-worthy): user hit Go distinct-named-type mixing error (CPSR vs uint32) in Set(). Understood the fix = same-type constants. Signals: basic type system grasp forming; typed constants + conversions idioms are live teaching topics.
- Preference: NO code in lessons. Give plain-instruction objectives and explanations only; user writes all code themselves. (stated re lesson 3)
- Preference: hints must be minimal — name the failing behavior/concept only, never hand over the fix line. User wants to derive fixes.
- Project convention: tests live in tests/<pkg>/ as external _test packages mirroring internal/ structure.
