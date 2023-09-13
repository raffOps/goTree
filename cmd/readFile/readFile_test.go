package readFile

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := "../../test/read_file.txt"
	got, err := ReadFile(filename)
	wanted := [4]string{"1", "2", "3", "4"}
	if err != nil {
		t.Errorf("%v", err)
	}
	if reflect.DeepEqual(got, wanted) {
		t.Errorf("got %s,\n wanted %s", got, wanted)
	}
}

func TestReadFileError(t *testing.T) {
	filename := "foo.txt"
	_, err := ReadFile(filename)
	if err == nil {
		t.Fail()
	}
}
