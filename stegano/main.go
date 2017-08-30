package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
	Usable only with bmp (index 55)
	Other formats can be identified using image.Decode()
	Just make sure to include the following imports to be able to decode all formats
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ "golang.org/x/image/bmp"
*/
const (
	bmpIndex = 55
)

func main() {
	if len(os.Args) == 2 {
		msg, err := decryptFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println("Message:\n", msg)
	} else if len(os.Args) == 4 {
		in := os.Args[1]
		out := os.Args[2]
		text := os.Args[3]
		encryptFile(in, out, text)
	}
}

func encryptFile(in, out, text string) error {
	inFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer inFile.Close()

	inContent, err := ioutil.ReadAll(inFile)
	if err != nil {
		return err
	}

	outContent := encrypt(inContent, text)

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}

	_, err = outFile.Write(outContent)
	if err != nil {
		return err
	}

	// Do not defer Close writable files
	err = outFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func encrypt(img []byte, message string) []byte {
	out := make([]byte, len(img))
	copy(out, img)

	msg := []byte(message)
	// Use the null char to mark the end of the message (useful for decryption)
	msg = append(msg, 0)

	index := bmpIndex
	for _, char := range msg {
		// Iterate through each of the 7 bits of each char
		for mask := byte(1 << 7); mask > 0; mask >>= 1 {
			bit := char&mask == mask

			// And write each bit "inside" one image byte
			if bit {
				out[index] = img[index] | byte(0x01)
			} else {
				out[index] = img[index] & byte(0xfe)
			}
			index++
		}
	}
	return out
}

func decryptFile(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	message := decrypt(content)
	return message, nil
}

func decrypt(img []byte) string {
	var msg []byte

	index := bmpIndex

	var letter byte
	for i, b := range img[index:] {
		// Get the LSB of the image byte
		bit := b&0x01 == 0x01
		if bit {
			letter = letter | 1
		} else {
			letter = letter &^ 1
		}

		// If we are on the last bit of the current letter
		if i%8 == 7 {
			// If we reached the end of the message
			if letter == 0 {
				break
			}
			msg = append(msg, letter)
			letter = 0
		} else {
			// Let's write the next bit of the current letter
			letter <<= 1
		}
	}

	return string(msg)
}
