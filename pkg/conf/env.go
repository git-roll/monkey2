package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	EnvCoffeeTimeUpperBound = "COFFEE_TIME"
	EnvNameLength           = "NAME_LENGTH"
	EnvWriteOnceLength      = "LENGTH_WRITE_ONCE"
	EnvPercentageFileOP     = "PERCENTAGE_FILE_OPERATION"
	EnvWorktree             = "WORKTREE"
	EnvSidecarStdFile       = "SIDECAR_STD_FILE"
)

var noticeOnce = make(map[string]bool)

func notice(key string, hint string, v interface{}) {
	if !noticeOnce[key] {
		fmt.Printf(hint+`. Set environment variable "%s" to change.`+"\n", v, key)
		noticeOnce[key] = true
	}
}

func CoffeeTimeUpperBound() string {
	v := os.Getenv(EnvCoffeeTimeUpperBound)
	if len(v) == 0 {
		v = (30 * time.Second).String()
	}

	notice(EnvCoffeeTimeUpperBound, `🚁 Coffee time would be up to %s`, v)
	return v
}

func Worktree() string {
	wt := os.Getenv(EnvWorktree)
	if len(wt) == 0 {
		wt = filepath.Join(os.TempDir(), "monkey")
	}

	notice(EnvWorktree, `🚁 The workdir will be "%s"`, wt)
	return wt
}

func SidecarStdFile() string {
	std := os.Getenv(EnvSidecarStdFile)
	if len(std) == 0 {
		std = filepath.Join(os.TempDir(), "monkey.sidecar")
	}

	notice(EnvSidecarStdFile, `🚁 Stdout of the sidecar will be written to "%s"`, std)
	return std
}

func NameLength() int {
	return envInt(EnvNameLength, 8, `🚁 Length of file/dir name would be %d`)
}

func WriteOnceLengthUpperBound() int {
	return envInt(
		EnvWriteOnceLength, 2048,
		`🚁 Length of each file write would be %d`,
	)
}

func PercentageFileOP() int {
	return envInt(EnvPercentageFileOP, 70,
		`🚁 %d%% filesystem operations would be on files`,
	)
}

func envInt(key string, def int, hint string) (i int) {
	v := os.Getenv(key)
	if len(v) > 0 {
		i64, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			panic(os.Getenv(key))
		}
		i = int(i64)
	} else {
		i = def
	}

	notice(key, hint, i)
	return
}
