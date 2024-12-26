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
	readFlag      = flag.String("r", "", "Read barcode from a file (short flag)")
	readLongFlag  = flag.String("read", "", "Read barcode from a file (long flag)")
	writeFlag     = flag.String("w", "", "Text to encode as barcode (short flag)")
	writeLongFlag = flag.String("write", "", "Text to encode as barcode (long flag)")
	qrFlag        = flag.Bool("qrcode", false, "Generate a QR Code (default)")
	code128Flag   = flag.Bool("code128", false, "Generate a Code 128 barcode")
)

func main() {
	// Setup for usage
	flag.Usage = printUsage
	flag.Parse()

	// Handle the case when only a file is provided (default to read action)
	if len(flag.Args()) == 1 && *readFlag == "" && *readLongFlag == "" && *writeFlag == "" && *writeLongFlag == "" {
		readBarcode(flag.Arg(0))
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

	// If read action is provided, read the barcode from the file
	if readFile != "" {
		readBarcode(readFile)
	} else if writeText != "" {
		// If write action is provided, create barcode from text and save to file
		if len(flag.Args()) == 1 {
			filename := flag.Arg(0)
			barcodeType := detectBarcodeType()
			writeBarcode(writeText, filename, barcodeType)
		} else {
			printUsage()
			os.Exit(1)
		}
	} else if len(flag.Args()) == 1 {
		// Default to read action if only a file is provided
		readBarcode(flag.Arg(0))
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
  -w, --write <text> <file>     Write text as a barcode and save to file

Type (optional):
  -qrcode                       Generate a QR Code (default)
  -code128                      Generate a Code 128 barcode

Arguments:
  <FILE>                        The input image file to read, or the output file to save the barcode

Examples:
  ./app -r input.png            Read barcode from the input image file
  ./app -w "qr text" output.png Write a QR Code barcode to output.png
  ./app -w "qr text" -code128 output.png Write a Code 128 barcode to output.png`)
}

// Function to read barcode from an image
func readBarcode(filename string) {
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

	log.Println("Failed to decode barcode")
}

// Function to detect the barcode type, defaulting to QR Code
func detectBarcodeType() string {
	if *code128Flag {
		return "code128"
	}
	return "qrcode"
}

// Function to write barcode as an image file
func writeBarcode(text, filename, barcodeType string) {
	var img image.Image
	var err error

	switch barcodeType {
	case "qrcode":
		img, err = qrcode.NewQRCodeWriter().Encode(text, gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
	case "code128":
		img, err = oned.NewCode128Writer().Encode(text, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	default:
		log.Fatalf("Unsupported barcode type: %s", barcodeType)
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

	fmt.Printf("%s barcode saved to %s\n", barcodeType, filename)
}
