package main

import (
	// "bytes"
	// "crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

// Write Ahead Log (WAL) Log entry
type LogEntry struct{
	Data []byte
	Checksum uint32
}

func GenerateChecksum(data []byte) uint32{
	return crc32.ChecksumIEEE(data)
}

func WriteLog(filename string, entry LogEntry) error{
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	// Write the length of the data
	dataLen := uint32(len(entry.Data))
	lengthBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBytes, dataLen)
	_, err = file.Write(lengthBytes)
	if err != nil {
		return err
	}

	// write actual datta
	_, err = file.Write(entry.Data)

	// write checksum
	checksum := make([]byte, 4)
	binary.LittleEndian.PutUint32(checksum, entry.Checksum) // converts a uint32 checksum value into its Little-Endian byte representation
	_, err = file.Write(checksum)

	return err
}

func ReadLog(filename string) ([]LogEntry, error){
	file, err := os.Open(filename)
	if err != nil{
		return nil, err
	}
	defer file.Close()

	if err != nil{
		return nil, err
	}

	var entries []LogEntry
	for {
		// Read data len
		lengthBytes := make([]byte, 4)
		_, err := file.Read(lengthBytes)
		if err == io.EOF{
			break
		}
		if err != nil{
			return nil, err
		}

		// Read data
		dataLen := binary.LittleEndian.Uint32(lengthBytes)
		data := make([]byte, dataLen)
		_, err = file.Read(data)
		if err != nil {
			return nil, err
		}
		
		// Read checksums
		checksumBytes := make([]byte, 4)
		_, err = file.Read(checksumBytes)
		if err != nil {
			return nil, err
		}
		expectedChecksum := binary.LittleEndian.Uint32(checksumBytes)
		if GenerateChecksum(data) != expectedChecksum {
			offset, _ := file.Seek(0, io.SeekCurrent)
			return nil, fmt.Errorf("data corruption detected at byte %d", offset)
		}

		// Add valid entry to the list
		entries = append(entries, LogEntry{Data: data, Checksum: expectedChecksum})
	}

	return entries, nil
}

func main(){
	filename := "wal.log"
	data := []byte("My file contents server with chicken nugets and soup!!")
	logEntry := LogEntry{Data: data, Checksum: GenerateChecksum(data)}

	err := WriteLog(filename, logEntry)
	// err = WriteLog(filename, logEntry)
	if err != nil{
		log.Fatalf("Error writing log: %v", err)
	}

	file, err := ReadLog(filename)
	if err != nil{
		log.Fatalf("Error reading the log %v", err)
	}

	for _, entry := range(file){
		fmt.Printf("Contents of the file:\n%s \nChecksum: %d\n", string(entry.Data), entry.Checksum)
	}

}