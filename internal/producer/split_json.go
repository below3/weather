package producer

import "bytes"

func splitForJson() func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte("},")
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token with the }
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0 : i+1], nil
		}

		// If we're at EOF, we have a final line. Return it.
		// But without the closing sign ].
		if atEOF {
			return dataLen, data[0 : dataLen-2], nil
		}

		// Request more data.
		return 0, nil, nil
	}
}
