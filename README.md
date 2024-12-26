# QR Code and Barcode Generator

This application allows you to generate and read QR codes and Code 128 barcodes from image files. It provides a command-line interface to encode text into barcodes and decode barcodes from images.

## Usage

```sh
./app [TYPE] [ACTION] <FILE> ...
```

### Actions

- `-r, --read <file>`: Read barcode or QR code from an image file.
- `-w, --write <text> <file>`: Write text as a barcode and save to file.

### Types (optional)

- `-qrcode`: Generate a QR Code (default).
- `-code128`: Generate a Code 128 barcode.

### Arguments

- `<FILE>`: The input image file to read, or the output file to save the barcode.

### Examples

- Read barcode from the input image file:
  ```sh
  ./app -r input.png
  ```

- Write a QR Code barcode to `output.png`:
  ```sh
  ./app -w "qr text" output.png
  ```

- Write a Code 128 barcode to `output.png`:
  ```sh
  ./app -w "qr text" -code128 output.png
  ```

- Read barcode from `file.png` (default to read action):
  ```sh
  ./app file.png
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

3. Build the application:
   ```sh
   go build -o app main.go
   ```

## Dependencies

- [gozxing](https://github.com/makiuchi-d/gozxing): A Go port of ZXing library for barcode scanning and generation.

## License

This project is licensed under the MIT License.
