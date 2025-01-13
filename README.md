# qrcode - read or write qr and bar codes

This application allows you to generate and read QR codes and Code 128 barcodes from image files. It provides a command-line interface to encode text into barcodes and decode barcodes from images.

## Usage

```
user@pc:~$ qrcode 
Usage: ./app [TYPE] [ACTION] <FILE> ...

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
  ./app -barcode -w "text"      Write a barcode (Code 128) to barcode.png
```

## Installation

1. Clone the repository:
   ```sh
   git clone <repository-url>
   ```

2. Navigate to the project directory:
   ```sh
   cd qrcode
   ```

3. Build and run application:
   ```sh
   go install
   qrcode
   ```

## Dependencies

- [gozxing](https://github.com/makiuchi-d/gozxing): A Go port of ZXing library for barcode scanning and generation.

## License

This project is licensed under the MIT License.
