package container

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/docker/docker/pkg/idtools"
	"golang.org/x/sys/unix"
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

func TestProc(t *testing.T) {
	rootIdentity := idtools.Identity{}
	// initLayerFs := containerfs.ContainerFS{}
	initLayer := "/"

	for pth, typ := range map[string]string{
		// "/dev/pts":         "dir",
		// "/dev/shm":         "dir",
		"/proc": "dir",
		// "/sys":             "dir",
		// "/.dockerenv":      "file",
		// "/etc/resolv.conf": "file",
		// "/etc/hosts":       "file",
		// "/etc/hostname":    "file",
		// "/dev/console":     "file",
		// "/etc/mtab":        "/proc/mounts",
	} {
		parts := strings.Split(pth, "/")
		prev := "/"
		for _, p := range parts[1:] {
			prev = filepath.Join(prev, p)
			unix.Unlink(filepath.Join(initLayer, prev))
		}

		if _, err := os.Stat(filepath.Join(initLayer, pth)); err != nil {
			if os.IsNotExist(err) {
				if err := idtools.MkdirAllAndChownNew(filepath.Join(initLayer, filepath.Dir(pth)), 0755, rootIdentity); err != nil {
					return
				}
				switch typ {
				case "dir":
					if err := idtools.MkdirAllAndChownNew(filepath.Join(initLayer, pth), 0755, rootIdentity); err != nil {
						return
					}
				// case "file":
				// 	f, err := os.OpenFile(filepath.Join(initLayer, pth), os.O_CREATE, 0755)
				// 	if err != nil {
				// 		return
				// 	}
				// 	f.Chown(rootIdentity.UID, rootIdentity.GID)
				// 	f.Close()
				default:
					if err := os.Symlink(typ, filepath.Join(initLayer, pth)); err != nil {
						return
					}
				}
			} else {
				return
			}
		}
	}
}
