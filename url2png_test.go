package url2png

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newClient() Client {
	return Client{
		Key:    os.Getenv("TEST_KEY"),
		Secret: os.Getenv("TEST_SECRET"),
	}
}

func TestScreenshot(t *testing.T) {
	t.Parallel()

	a := assert.New(t)

	c := newClient()

	r, err := c.Screenshot("http://www.fknsrs.biz/", nil)
	a.NoError(err)
	if !a.NotNil(r) {
		return
	}

	n, err := ioutil.ReadAll(r)
	a.NoError(err)
	a.NotEqual(0, n)
}

func TestBadURL(t *testing.T) {
	t.Parallel()

	a := assert.New(t)

	c := newClient()

	r, err := c.Screenshot("zzzzzzzz", nil)
	a.Error(err)
	a.Nil(r)
}
