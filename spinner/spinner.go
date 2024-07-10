package spinner

import (
	"time"

	"github.com/briandowns/spinner"
)

func New() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 75*time.Millisecond, spinner.WithColor("yellow"))
	s.Prefix = "Searching "
	return s
}
