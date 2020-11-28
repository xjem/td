package telegram

import (
	"bytes"
	"testing"

	"github.com/ernado/td/bin"
	"github.com/ernado/td/internal/proto"
)

func TestDecrypt(t *testing.T) {
	// Test vector from grammers.

	c := &Client{
		rand: Zero{},
		authKey: [256]byte{
			93, 46, 125, 101, 244, 158, 194, 139, 208, 41, 168, 135, 97, 234, 39, 184, 164, 199,
			159, 18, 34, 101, 37, 68, 62, 125, 124, 89, 110, 243, 48, 53, 48, 219, 33, 7, 232, 154,
			169, 151, 199, 160, 22, 74, 182, 148, 24, 122, 222, 255, 21, 107, 214, 239, 113, 24,
			161, 150, 35, 71, 117, 60, 14, 126, 137, 160, 53, 75, 142, 195, 100, 249, 153, 126,
			113, 188, 105, 35, 251, 134, 232, 228, 52, 145, 224, 16, 96, 106, 108, 232, 69, 226,
			250, 1, 148, 9, 119, 239, 10, 163, 42, 223, 90, 151, 219, 246, 212, 40, 236, 4, 52,
			215, 23, 162, 211, 173, 25, 98, 44, 192, 88, 135, 100, 33, 19, 199, 150, 95, 251, 134,
			42, 62, 60, 203, 10, 185, 90, 221, 218, 87, 248, 146, 69, 219, 215, 107, 73, 35, 72,
			248, 233, 75, 213, 167, 192, 224, 184, 72, 8, 82, 60, 253, 30, 168, 11, 50, 254, 154,
			209, 152, 188, 46, 16, 63, 206, 183, 213, 36, 146, 236, 192, 39, 58, 40, 103, 75, 201,
			35, 238, 229, 146, 101, 171, 23, 160, 2, 223, 31, 74, 162, 197, 155, 129, 154, 94, 94,
			29, 16, 94, 193, 23, 51, 111, 92, 118, 198, 177, 135, 3, 125, 75, 66, 112, 206, 233,
			204, 33, 7, 29, 151, 233, 188, 162, 32, 198, 215, 176, 27, 153, 140, 242, 229, 205,
			185, 165, 14, 205, 161, 133, 42, 54, 230, 53, 105, 12, 142,
		},
	}
	c.authKeyID = c.authKey.ID()

	var msg proto.EncryptedMessage
	b := &bin.Buffer{Buf: []byte{
		122, 113, 131, 194, 193, 14, 79, 77, 249, 69, 250, 154, 154, 189, 53, 231, 195, 132,
		11, 97, 240, 69, 48, 79, 57, 103, 76, 25, 192, 226, 9, 120, 79, 80, 246, 34, 106, 7,
		53, 41, 214, 117, 201, 44, 191, 11, 250, 140, 153, 167, 155, 63, 57, 199, 42, 93, 154,
		2, 109, 67, 26, 183, 64, 124, 160, 78, 204, 85, 24, 125, 108, 69, 241, 120, 113, 82,
		78, 221, 144, 206, 160, 46, 215, 40, 225, 77, 124, 177, 138, 234, 42, 99, 97, 88, 240,
		148, 89, 169, 67, 119, 16, 216, 148, 199, 159, 54, 140, 78, 129, 100, 183, 100, 126,
		169, 134, 18, 174, 254, 148, 44, 93, 146, 18, 26, 203, 141, 176, 45, 204, 206, 182,
		109, 15, 135, 32, 172, 18, 160, 109, 176, 88, 43, 253, 149, 91, 227, 79, 54, 81, 24,
		227, 186, 184, 205, 8, 12, 230, 180, 91, 40, 234, 197, 109, 205, 42, 41, 55, 78,
	}}
	if err := msg.Decode(b); err != nil {
		t.Fatal(err)
	}
	plaintext, err := c.decrypt(&msg)
	if err != nil {
		t.Fatal(err)
	}
	expectedPlaintext := []byte{
		252, 130, 106, 2, 36, 139, 40, 253, 96, 242, 196, 130, 36, 67, 173, 104, 1, 240, 193,
		194, 145, 139, 48, 94, 2, 0, 0, 0, 88, 0, 0, 0, 220, 248, 241, 115, 2, 0, 0, 0, 1, 168,
		193, 194, 145, 139, 48, 94, 1, 0, 0, 0, 28, 0, 0, 0, 8, 9, 194, 158, 196, 253, 51, 173,
		145, 139, 48, 94, 24, 168, 142, 166, 7, 238, 88, 22, 252, 130, 106, 2, 36, 139, 40,
		253, 1, 204, 193, 194, 145, 139, 48, 94, 2, 0, 0, 0, 20, 0, 0, 0, 197, 115, 119, 52,
		196, 253, 51, 173, 145, 139, 48, 94, 100, 8, 48, 0, 0, 0, 0, 0, 252, 230, 103, 4, 163,
		205, 142, 233, 208, 174, 111, 171, 103, 44, 96, 192, 74, 63, 31, 212, 73, 14, 81, 246,
	}
	if !bytes.Equal(expectedPlaintext, plaintext) {
		t.Error("mismatch")
	}
}