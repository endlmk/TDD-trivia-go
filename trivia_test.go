package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTrivia1(t *testing.T) {
	testdata := readTestdata(t, "testdata_1673617280741338841.txt")
	result := captureStdout(t, func() { gameLoop(1673617280741338841) })
	if result != testdata {
		t.Fatalf("result was different from testdata\nresult:%v", result)
	}
}

func TestTrivia2(t *testing.T) {
	testdata := readTestdata(t, "testdata_1673618826315788903.txt")
	result := captureStdout(t, func() { gameLoop(1673618826315788903) })
	if result != testdata {
		t.Fatalf("result was different from testdata\nresult:%v", result)
	}
}
func TestTrivia3(t *testing.T) {
	testdata := readTestdata(t, "testdata_1673618894931469315.txt")
	result := captureStdout(t, func() { gameLoop(1673618894931469315) })
	if result != testdata {
		t.Fatalf("result was different from testdata\nresult:%v", result)
	}
}

func captureStdout(t *testing.T, fun func()) string {
	t.Helper()

	orgStdout := os.Stdout
	defer func() {
		os.Stdout = orgStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	fun()

	w.Close()

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}

	return strings.TrimRight(buf.String(), "\n")
}

func readTestdata(t *testing.T, path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read testdata: %v", err)
	}
	return string(b)
}
