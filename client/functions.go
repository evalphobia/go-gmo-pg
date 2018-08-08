package client

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

// GetRandomOrderID returns randomized 27 length string.
func GetRandomOrderID() string {
	return randomAlphaNumeric(27)
}

// RandomAlphaNumeric returns random alpha nums with specified length.
func randomAlphaNumeric(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}

// ValidateCheckString validates CheckString to avoid falsification.
func ValidateCheckString(checkString string, params ...string) bool {
	b := make([]byte, 0, 4096)
	for _, param := range params {
		b = append(b, param...)
	}
	return getMD5Hash(b) == checkString
}

func getMD5Hash(textByte []byte) string {
	return fmt.Sprintf("%x", (md5.Sum(textByte)))
}

var defaultSjisEncoder = japanese.ShiftJIS.NewEncoder()

// ConvertUtf8ToSjis converts utf-8 string to Shift_JIS string.
func ConvertUtf8ToSjis(utf8 string) (string, error) {
	if utf8 == "" {
		return "", nil
	}

	strReader := strings.NewReader(utf8)
	sjisReader := transform.NewReader(strReader, defaultSjisEncoder)
	byt, err := ioutil.ReadAll(sjisReader)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

// ConvertToFullWidth converts half-width character to full-width.
func ConvertToFullWidth(str string) string {
	return width.Widen.String(str)
}
