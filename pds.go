package pds

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Builder ----
type PDSBuilder struct {
	pdsMap map[string][]byte
}

func NewBuilder() *PDSBuilder {
	return &PDSBuilder{pdsMap: make(map[string][]byte)}
}

func (b *PDSBuilder) AddTag(tag string, value string) {
	b.pdsMap[tag] = []byte(value)
}

func (b *PDSBuilder) Build() ([]byte, error) {
	var buf bytes.Buffer

	for tag, value := range b.pdsMap {
		binary.Write(&buf, binary.BigEndian, uint16(len(tag)))
		binary.Write(&buf, binary.BigEndian, []byte(tag))
		binary.Write(&buf, binary.BigEndian, uint16(len(value)))
		binary.Write(&buf, binary.BigEndian, value)
	}

	return buf.Bytes(), nil
}

func (b *PDSBuilder) BuildPositional() string {
	var result strings.Builder

	keys := []string{}

	for k := range b.pdsMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		vi, vierr := strconv.Atoi(keys[i])
		vj, vjerr := strconv.Atoi(keys[j])

		if vierr != nil {
			vi = 0
		}

		if vjerr != nil {
			vj = 0
		}

		return vi < vj
	})

	for _, tag := range keys {
		value := b.pdsMap[tag]

		// Ensure the tag is always 3 characters long,
		// pad if necessary (assuming it's numeric)
		formattedTag := fmt.Sprintf("%04s", tag)

		// Convert value to string and calculate its length
		valueStr := string(value)
		valueLength := len(valueStr)

		// Ensure the length is always 2 characters long, pad if necessary
		formattedLength := fmt.Sprintf("%03d", valueLength)

		// Construct the positional string for this TLV block
		// and append it to the result
		result.WriteString(fmt.Sprintf("%s%s%s", formattedTag,
			formattedLength, valueStr))
	}

	return result.String()
}

// Parser ----

// This data element contains one or more PDSs in PDS-encoded format.
// PDSs are formatted using a tag-length-data encoding procedure, shown in this table.
type PDS struct {
	Tag   string
	Value []byte
}

// TLV is a type-length-value structure.
func Parse(data []byte) ([]PDS, error) {
	var tlvs []PDS

	for len(data) >= 4 {
		// Read TLV header
		tagLength := binary.BigEndian.Uint16(data[0:2])
		if len(data) < int(2+tagLength) {
			return nil, errors.New("incomplete tag data")
		}
		tag := string(data[2 : 2+tagLength]) // Convert tag to string
		data = data[2+tagLength:]

		if len(data) < 2 {
			return nil, errors.New("incomplete length data")
		}
		length := binary.BigEndian.Uint16(data[0:2])
		data = data[2:]

		if uint16(len(data)) < length {
			return nil, errors.New("incomplete value data")
		}

		// // Append to PDS slice
		tlvs = append(tlvs, PDS{
			Tag:   tag, // Use the string tag directly
			Value: data[:length],
		})

		// Move to the next PDS block
		data = data[length:]
	}

	sort.Slice(tlvs, func(i, j int) bool {
		typeI, errI := strconv.Atoi(tlvs[i].Tag)
		typeJ, errJ := strconv.Atoi(tlvs[j].Tag)

		if errI != nil {
			typeI = 0
		}

		if errJ != nil {
			typeJ = 0
		}

		return typeI < typeJ
	})

	return tlvs, nil
}

// ParsePositional parses a positional PDS string and returns a slice of PDS structures.
func ParsePositional(s string) ([]PDS, error) {
	var tlvs []PDS

	// Define the fixed lengths for the tag and length fields
	const tagLength = 4
	const lengthFieldLength = 3

	for i := 0; i < len(s); {
		if i+tagLength+lengthFieldLength > len(s) {
			return nil, errors.New("incomplete TLV block")
		}

		// Extract Tag, Length, and Value
		tag := s[i : i+tagLength]

		lengthStr := s[i+tagLength : i+tagLength+lengthFieldLength]

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid length value: %v", err)
		}

		// Check if there's enough room for the value
		if i+tagLength+lengthFieldLength+length > len(s) {
			return nil, errors.New("incomplete value field")
		}

		value := s[i+tagLength+lengthFieldLength : i+tagLength+lengthFieldLength+length]

		// Append to PDS slice
		tlvs = append(tlvs, PDS{
			Tag:   tag,
			Value: []byte(value),
		})

		// Move to the next PDS block
		i += tagLength + lengthFieldLength + length
	}

	return tlvs, nil
}
