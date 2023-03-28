package database

import (
	"encoding/json"
	"fmt"
	"strings"
)

// read []*User from db format
func readData(buf []byte) (list []*User, err error) {

	if err = json.Unmarshal(buf, &list); err != nil {
		return nil, fmt.Errorf("data is corrupted (#rD): %v", err)
	}
	return list, nil
}

// convert *User to db format
func makeData(u *User) (data []byte, err error) {

	data, err = json.MarshalIndent(u, "    ", "    ")
	if err != nil {
		return nil, fmt.Errorf("data is incorrect (#mD): %v", err)
	}
	dataRows := strings.Split(string(data), "\n")

	if len(dataRows) > 6 {
		for row := 5; row < len(dataRows)-1; row++ {
			dataRows[row] = strings.TrimSpace(dataRows[row])
		}
		dataRows[4] = strings.Join(dataRows[4:len(dataRows)-1], "")
		dataRows = append(dataRows[:5], dataRows[len(dataRows)-1])
		data = []byte(strings.Join(dataRows, "\n"))
	}
	return data, nil
}

// change part of buf to data from start
func updateData(buf, data []byte, start int) []byte {

	bufStr := strings.Split(string(buf), "\n")
	startStr := start*6 - 2
	dataStr := strings.Split(string(data), "\n")
	for i := 0; i < 4; i++ {
		bufStr[startStr+i] = dataStr[i+1]
	}
	return []byte(strings.Join(bufStr, "\n"))
}
