package model

import (
	"github.com/phantacix/go-admin-panel/core"
	"strconv"
	"testing"
)

func init() {
	core.Init("../../../config/env/local/admin.json")
}

func TestLog_LogCreate(t *testing.T) {
	for i := 1; i <= 50; i++ {
		LogCreate(1, "su_"+strconv.Itoa(i), "aa", "bb", "cc", "127.0.0.1")
	}
}
