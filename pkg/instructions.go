package pkg

import "github.com/prorochestvo/hugeint/internal/instruction"

// Add ...
func Add(a, b uint32) (hDWord, lDWord uint32) {
	return instruction.AddV2(a, b)
}
