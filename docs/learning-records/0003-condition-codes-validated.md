# Condition codes: implemented and exhaustively validated

User implemented `CPSR.ConditionPassed` solo (no code from teacher, per preference) against an exhaustive 256-combination truth-table test. Initial version had three bugs — a systematic off-by-one renumbering (cases started at 1), GE/LT/GT/LE consulting C instead of V, and &&/|| errors in LS/LE — all found by the test and fixed by the user across two iterations, with only the final `==` vs `!=` in LE needing a hint.

Evidence: `go test ./internal/cpu/` fully green, 2026-10-08.

Implications: bit extraction and the condition table are now solid floors. The off-by-one and value-vs-position confusions from LR-0002 recurred in a new costume but were self-corrected once a test pointed at them — user responds well to exhaustive truth-table validation. Next: lesson 3 (data-processing decode + MOV), then barrel shifter. Teaching pattern that works: user writes all code, teacher provides validation tests and feedback.
