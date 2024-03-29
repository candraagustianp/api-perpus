package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Base64Decode(file string) ([]byte, error) {
	data := file[strings.IndexByte(file, ',')+1:]
	decode64Byte, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return nil, err
	}

	return decode64Byte, nil
}

func IsBase64(s string) bool {
	s = s[strings.IndexByte(s, ',')+1:]
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func WriteFileBase64(folderName string, fileName string, base64File []byte) error {
	if _, err := os.Stat(folderName); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.Mkdir(folderName, os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", folderName, fileName))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(base64File); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func RemoveFile(folderName string, fileName string) error {
	e := os.Remove(fmt.Sprintf("%s/%s", folderName, fileName))
	if e != nil {
		return e
	}
	return nil
}

func Rename(folderName string, oldFileName string, newFileName string) error {
	e := os.Rename(fmt.Sprintf("%s/%s", folderName, oldFileName), fmt.Sprintf("%s/%s", folderName, newFileName))
	if e != nil {
		return e
	}
	return nil
}
