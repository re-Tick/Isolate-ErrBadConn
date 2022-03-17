package main

import (
	"bytes"
	"database/sql/driver"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
)

type KError struct {
	Err error
}

// Error method returns error string stored in Err field of KError.
func (e *KError) Error() string {
	return e.Err.Error()
}

const version = 1

// GobEncode encodes the Err and returns the binary data.
func (e *KError) GobEncode() ([]byte, error) {
	r := make([]byte, 0)
	r = append(r, version)

	if e.Err != nil {
		r = append(r, e.Err.Error()...)
	}
	return r, nil
}

// GobDecode decodes the b([]byte) into error struct.
func (e *KError) GobDecode(b []byte) error {
	if b[0] != version {
		return errors.New("gob decode of errors.errorString failed: unsupported version")
	}
	if len(b) == 1 {
		e.Err = nil
	} else {
		str := string(b[1:])
		e.Err = errors.New(str)
	}

	return nil
}

func main() {
	bin := []byte{10, 255, 129, 5, 1, 2, 255, 132, 0, 0, 0, 27, 255, 130, 0, 23, 1, 100, 114, 105, 118, 101, 114, 58, 32, 98, 97, 100, 32, 99, 111, 110, 110, 101, 99, 116, 105, 111, 110}
	obj := &KError{}
	dec := gob.NewDecoder(bytes.NewBuffer(bin))
	err := dec.Decode(obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj.Err)
	fmt.Println(driver.ErrBadConn)
	err = obj.Err
	if err == driver.ErrBadConn {
		fmt.Println("err and driver.ErrBadConn are Equal")
	} else {
		fmt.Println("err and driver.ErrBadConn are NOT Equal")
	}

	if err.Error() == driver.ErrBadConn.Error() {
		fmt.Println("err and driver.ErrBadConn strings are Equal")
	} else {
		fmt.Println("err and driver.ErrBadConn strings are NOT Equal")
	}

}
