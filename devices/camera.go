package devices

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dhowden/raspicam"
	"io"
	"log"
	"os"
)

type Camera struct{}

func (c *Camera) TakePicture() io.Reader {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	reader := bufio.NewReader(&buffer)
	s := raspicam.NewStill()
	s.Camera.VFlip = true
	errCh := make(chan error)
	go func() {
		for x := range errCh {
			fmt.Fprintf(os.Stderr, "%v\n", x)
		}
	}()
	log.Println("Capturing image...")
	raspicam.Capture(s, writer, errCh)
	return reader
}
