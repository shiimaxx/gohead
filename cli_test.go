package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestRun_versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("gohead -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("gohead version %s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_linesFlag_default(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	content := []byte(`aaaaa
bbbbb
ccccc
ddddd
eeeee
fffff
ggggg
hhhhh
iiiii
jjjjj
kkkkk
lllll
mmmmm
nnnnn
ooooo
`)
	tempfile, _ := ioutil.TempFile("", "temp")
	defer os.Remove(tempfile.Name())
	if _, err := tempfile.Write(content); err != nil {
		fmt.Print(err)
	}

	args := strings.Split(fmt.Sprintf("gohead %s", tempfile.Name()), " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := `aaaaa
bbbbb
ccccc
ddddd
eeeee
fffff
ggggg
hhhhh
iiiii
jjjjj
`
	if !strings.EqualFold(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRun_linesFlag_5(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	content := []byte(`aaaaa
bbbbb
ccccc
ddddd
eeeee
fffff
ggggg
hhhhh
iiiii
jjjjj
kkkkk
lllll
mmmmm
nnnnn
ooooo
`)
	tempfile, _ := ioutil.TempFile("", "temp")
	defer os.Remove(tempfile.Name())
	if _, err := tempfile.Write(content); err != nil {
		fmt.Print(err)
	}

	args := strings.Split(fmt.Sprintf("gohead -n=5 %s", tempfile.Name()), " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := `aaaaa
bbbbb
ccccc
ddddd
eeeee
`
	if !strings.EqualFold(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRun_linesFlag_15(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	content := []byte(`aaaaa
bbbbb
ccccc
ddddd
eeeee
fffff
ggggg
hhhhh
iiiii
jjjjj
kkkkk
lllll
mmmmm
nnnnn
ooooo
`)
	tempfile, _ := ioutil.TempFile("", "temp")
	defer os.Remove(tempfile.Name())
	if _, err := tempfile.Write(content); err != nil {
		fmt.Print(err)
	}

	args := strings.Split(fmt.Sprintf("gohead -n=15 %s", tempfile.Name()), " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := `aaaaa
bbbbb
ccccc
ddddd
eeeee
fffff
ggggg
hhhhh
iiiii
jjjjj
kkkkk
lllll
mmmmm
nnnnn
ooooo
`
	if !strings.EqualFold(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRun_fileNotExists(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split("gohead dummy_file", " ")

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := "dummy_file: No such file or directory"
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_noArguments(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := []string{"gohead"}

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}

	expected := "Missing filename"
	if !strings.EqualFold(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}
