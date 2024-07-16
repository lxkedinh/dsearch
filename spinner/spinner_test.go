package spinner_test

import (
	"testing"
	"time"

	"github.com/lxkedinh/dsearch/spinner"
)

func TestLoadingSpinner(t *testing.T) {
	spinner := spinner.New()
	spinner.Prefix = "Loading Spinner Creation Test "
	spinner.Start()
	time.Sleep(time.Millisecond * 500)
	spinner.Stop()
}
