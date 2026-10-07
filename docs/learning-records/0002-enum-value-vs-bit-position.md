# Enum value vs. bit position: first conflation bug

User implemented Has/Set/Clear/Mode correctly with proper CPSR-typed constants, but wrote `IsThumb()` as `c.Has(CPSR(THUMB))` — converting the `InstructionState` enum (THUMB=1) into a CPSR mask, which tests bit 0 (a mode bit) instead of the T bit at bit 5. Corrected to `c.Has(flagT)`.

Evidence: `internal/cpu/registers.go` review, 2026-10-07.

Implications: user grasps methods, receivers, and typed constants; the remaining gap is the distinction between *abstract Go enums* and *hardware bit positions*. Expect the same confusion with Mode enum vs. mode field values (0x10–0x1F) and later with opcode field encodings. Lesson 2 (condition codes + data-processing format) should drill "decode the bitfield, then map to a Go value" explicitly.
