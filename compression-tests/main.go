package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/klauspost/compress/zstd"
)

func main() {
	f, err := os.ReadFile("./ZAP-Report.json")
	if err != nil {
		fmt.Printf("failed to read file: %v", err)
	}

	// startTime = time.Now()

	zstdCompressed, err := zstdCompress(f)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("ZAP-Report")
	if err != nil {
		fmt.Printf("failed to create file: %v", err)
	}
	defer file.Close()
	_, err = file.Write(zstdCompressed)
	if err != nil {
		fmt.Printf("failed to write to file: %v", err)
	}
	f2, err := os.ReadFile("ZAP-Report")
	if err != nil {
		fmt.Printf("failed to read file: %v", err)
	}
	decompress, err := zstdDecompress(f2)
	if err != nil {
		fmt.Printf("failed to decompress file: %v", err)
	}

	// zstdCompressTime := time.Since(startTime)

	// zstdCompressionRatio := float64(len(zstdCompressed)) / float64(len(f))
	// startTime = time.Now()
	// zstdDecompressed, err := zstdDecompress(zstdCompressed)
	// if err != nil {
	// 	panic(err)
	// }
	// zstdDecompressTime := time.Since(startTime)
	isZstdContentOK := bytes.Equal(f, decompress)

	// startTime = time.Now()

	// fmt.Printf("\nZstd Compression Time: %v\n", zstdCompressTime)
	// fmt.Printf("Zstd Compression Ratio: %.2f\n", zstdCompressionRatio)
	// fmt.Printf("Zstd output file size: %v\n", len(zstdCompressed))
	fmt.Printf("Zstd Content Check: %t\n", isZstdContentOK)
	// fmt.Printf("Zstd Decompression Time: %v\n", zstdDecompressTime)

	fmt.Printf("result: %v\n", 8/3)

}

func zstdCompress(inputData []byte) ([]byte, error) {
	encoder, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.EncoderLevelFromZstd(10)))
	if err != nil {
		return nil, fmt.Errorf("failed to create zstd writer: %v", err)
	}

	var compressedData []byte
	compressedData = encoder.EncodeAll(inputData, compressedData)

	err = encoder.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close zstd writer")
	}

	return compressedData, nil
}

func zstdDecompress(compressedData []byte) ([]byte, error) {
	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create zstd reader: %v", err)
	}
	defer decoder.Close()

	var decompressed []byte
	decompressed, err = decoder.DecodeAll(compressedData, decompressed)
	if err != nil {
		return nil, fmt.Errorf("error decompressing with zstd: %v", err)
	}
	return decompressed, nil
}
