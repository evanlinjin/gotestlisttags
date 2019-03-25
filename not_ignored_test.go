// +build not_ignored

package gotestlisttags

import "testing"

func TestShouldNotBeIgnored(t *testing.T) {
	t.Log("I should NOT be ignored!")
}