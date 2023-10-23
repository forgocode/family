package conf

import (
	"fmt"
	"os"
	"testing"

	"github.com/forgocode/family/internal/pkg/newlog"
)

func TestLoad(t *testing.T) {
	newlog.InitLogger("", os.Stdout)
	configPath = "config.yaml"
	fmt.Println(GetConfig())
}
