package main

import (
	"os"

	"github.com/forgocode/family/internal/pkg/newlog"
)

func main() {
	newlog.InitLogger("", os.Stdout)
	newlog.Logger.Infof("im server is start!\n")
}
