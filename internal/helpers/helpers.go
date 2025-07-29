package helpers

import (
	"fmt"
)

func SplitDBInput(s []map[string]string, chunkSize uint) ([][]map[string]string, error) {
	sLen := uint(len(s))

	if chunkSize >= sLen {
		return [][]map[string]string{}, fmt.Errorf("n can't be smaller or equal to slice lenght")
	}

	split := [][]map[string]string{}
	chunks := sLen / chunkSize
	reminder := sLen % chunkSize

	fmt.Printf("chunks: %v\n", chunks)
	fmt.Printf("reminder: %v\n", reminder)

	for i := 0; i < int(chunks); i++ {
		split = append(split, s[i*int(chunkSize):(i+1)*int(chunkSize)])
	}
	if reminder > 0 {
		split = append(split, s[chunks*chunkSize:])
	}

	return split, nil
}
