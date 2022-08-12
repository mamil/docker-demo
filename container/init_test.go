package container

import (
	"os"
	"os/exec"
	"testing"
)

func TestFindCgroupMountpoint(t *testing.T) {
	cmd := "ls"
	path, err := exec.LookPath(cmd)
	if err != nil {
		t.Logf("LookPath err:%v", err)
	} else {
		t.Logf("LookPath path:%v", path)
	}

	env := os.Environ()
	t.Logf("env:%v", env)
}
