package main

import (
	"os"
	"testing"
)

var (
	smallFile = "smallfile.txt"
	longFile  = "largefile.txt"
)

func init() {
	os.Remove(longFile)
	os.Remove(smallFile)
	lg := 1024 * 1024 * 1024 // 1Gb
	sm := 1024 * 512         // 512kb
	lgData := make([]byte, lg)
	smData := make([]byte, sm)
	tmpFileLg, err := os.Create(longFile)
	if err != nil {
		panic(err)
	}
	tmpFileSm, err := os.Create(smallFile)
	if err != nil {
		panic(err)
	}
	_, err = tmpFileSm.WriteAt(smData, 0)
	if err != nil {
		panic(err)
	}
	_, err = tmpFileLg.WriteAt(lgData, 0)
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadWriteSmallFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open(smallFile)
		if err != nil {
			b.Fatalf("failed opening small file")
		}
		defer f.Close()
		if err := readWrite(f); err != nil {
			b.Fatalf("failed reading small file")
		}
	}
}

func BenchmarkReadWriteLongFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open(longFile)
		if err != nil {
			b.Fatalf("failed opening long file")
		}
		defer f.Close()
		if err := readWrite(f); err != nil {
			b.Fatalf("failed reading small file")
		}
	}
}

func BenchmarkReadWriteBufferedSmallFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open(smallFile)
		if err != nil {
			b.Fatalf("failed opening small file")
		}
		defer f.Close()
		if err := readWriteBuffered(f); err != nil {
			b.Fatalf("failed reading small file")
		}
	}
}

func BenchmarkReadWriteBufferedLongFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open(longFile)
		if err != nil {
			b.Fatalf("failed opening long file")
		}
		defer f.Close()
		if err := readWriteBuffered(f); err != nil {
			b.Fatalf("failed reading small file")
		}
	}
}
