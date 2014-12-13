package sem_test

import (
	"io/ioutil"
	"testing"

	"github.com/nejstastnejsistene/sem-almanac-go"
)

func TestUnmarshal(t *testing.T) {
	buf, err := ioutil.ReadFile("test_data/almanac.al3")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := sem.Unmarshal(buf); err != nil {
		t.Fatal(err)
	}
}
