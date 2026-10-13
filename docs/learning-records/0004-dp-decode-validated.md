# Data-processing decode: step 1 complete, extraction move consolidated

User implemented `decodeDP` (all fields + rotated immediate) solo against a white-box test in `internal/cpu/arm_decoder_test.go`. All green after ~6 fix iterations. Recurring bug class across the whole session: field width/position confusion — masks swapped between fields, shift amounts taken from the wrong edge, rotate run in wrong-width containers (uint → RotateLeft8 → RotateLeft32 progression). Each fix landed once the user *derived* the field's start-bit and width out loud instead of permuting expressions.

Evidence: `go test ./internal/cpu/ -run TestDecodeDP` green, 2026-10-13.

Implications: shift-then-mask extraction is now a demonstrated skill, though fragile under time pressure — keep requiring the "say start and width before typing" ritual in future formats (load/store has 16+ fields). Rotating immediates understood including the 32-seat carousel reasoning. Next: execDP (MOV → S-bit flags → MVN/BIC), then end-to-end tests in tests/cpu per project convention. Guided-skeleton format is working; minimal-hint debugging (name failing behavior + concept, user derives fix) also working after initial over-help was corrected.
