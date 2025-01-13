package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/makiuchi-d/gozxing/qrcode"
)

var (
	readFlag      = flag.String("r", "", "Read code from a file (short flag)")
	readLongFlag  = flag.String("read", "", "Read code from a file (long flag)")
	writeFlag     = flag.String("w", "", "Text to encode as code (short flag)")
	writeLongFlag = flag.String("write", "", "Text to encode as code (long flag)")
	barFlag       = flag.Bool("barcode", false, "Generate a Code 128 barcode")
)

func main() {
	// Setup for usage
	flag.Usage = printUsage
	flag.Parse()

	// Handle the case when only a file is provided (default to read action)
	if len(flag.Args()) == 1 && *readFlag == "" && *readLongFlag == "" && *writeFlag == "" && *writeLongFlag == "" {
		readCode(flag.Arg(0))
		return
	}

	// Determine which action is requested (read or write)
	readFile := *readFlag
	if *readLongFlag != "" {
		readFile = *readLongFlag
	}

	writeText := *writeFlag
	if *writeLongFlag != "" {
		writeText = *writeLongFlag
	}

	// If read action is provided, read the code from the file
	if readFile != "" {
		readCode(readFile)
	} else if writeText != "" {
		// If write action is provided, create code from text and save to file
		if len(flag.Args()) == 1 {
			filename := flag.Arg(0)
			codeType := detectBarcodeType()
			writeBarcode(writeText, filename, codeType)
		} else {
			filename := detectBarcodeType() + ".png"
			codeType := detectBarcodeType()
			writeBarcode(writeText, filename, codeType)
		}
	} else if len(flag.Args()) == 1 {
		// Default to read action if only a file is provided
		readCode(flag.Arg(0))
	} else {
		printUsage()
		os.Exit(1)
	}
}

// Function to display usage
func printUsage() {
	fmt.Println(`Usage: ./app [TYPE] [ACTION] <FILE> ...

Action:
  -r, --read <file>             Read barcode or QR code from an image file
  -w, --write <text> <file>     Write "text" as QR code or barcode and save to file (.png)

Type (optional):
  -qrcode                       Generate a QR Code (default)
  -barcode                      Generate a Code 128 barcode

Arguments:
  <FILE>                        Input file to read from, or output file write to

Examples:
  ./app -r input.png            Read code from the input file
  ./app -w "text" output.png    Write a QR code to output.png
  ./app -barcode -w "text"      Write a barcode (Code 128) to barcode.png`)
}

// Function to read code from an image
func readCode(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		log.Fatal(err)
	}

	if result, err := qrcode.NewQRCodeReader().Decode(bmp, nil); err == nil {
		fmt.Println(result.String())
		return
	}
	if result, err := oned.NewCode128Reader().Decode(bmp, nil); err == nil {
		fmt.Println(result.String())
		return
	}

	log.Printf("Failed to decode '%s'\n", filename)
}

// Function to detect the code type, defaulting to QR Code
func detectBarcodeType() string {
	if *barFlag {
		return "barcode"
	}
	return "qrcode"
}

// Function to write code as an image file
func writeBarcode(text, filename, codeType string) {
	var img image.Image
	var err error

	switch codeType {
	case "qrcode":
		img, err = qrcode.NewQRCodeWriter().Encode(text, gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
	case "barcode":
		img, err = oned.NewCode128Writer().Encode(text, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	default:
		log.Fatalf("Unsupported barcode type: %s", codeType)
	}
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s saved to '%s'\n", codeType, filename)
}
