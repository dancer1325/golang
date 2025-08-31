package abs_test // 's package != NO test's file package				-- BUT NOT error / location == no-test file

import (
	"std/testing/examples/abs"
	"testing"
)

func TestAbs(t *testing.T) {
	got := abs.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
