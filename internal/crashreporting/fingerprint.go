package crashreporting

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func FingerPrint(crash string) string {
	var lines string
	for _, line := range strings.Split(crash, "\n") {
		// if line contains .cpp .h .c .hpp
		// then we can assume it's a stack trace
		// and we can generate a fingerprint
		// windows stack traces have a space after the file name
		if strings.Contains(line, ".cpp ") ||
			strings.Contains(line, ".h ") ||
			strings.Contains(line, ".c ") ||
			strings.Contains(line, ".hpp ") ||
			// linux stack traces have a colon after the file name
			strings.Contains(line, ".cpp:") ||
			strings.Contains(line, ".h:") ||
			strings.Contains(line, ".c:") ||
			strings.Contains(line, ".hpp:") {

			// #10 0x00005641f1abfddf in Client::Handle_Connect_OP_ClientReady (this=0x5641f3c6cbc0, app=<optimized out>) at ../zone/client_packet.cpp:982
			// if string contains " at " then we want only the characters after " at "
			if strings.Contains(line, " at ") {
				line = line[strings.Index(line, " at ")+4:]
			} else {
				lines += line
			}

			//fmt.Println("Found stack trace line", line)
		}
	}

	// generate md5 hash from lines
	// this will be the fingerprint
	hash := md5.Sum([]byte(lines))
	return fmt.Sprintf("%x", hash)
}
