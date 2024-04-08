package pds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	tlvBuilder := NewBuilder()
	if tlvBuilder == nil {
		t.Error("NewBuilder() returned nil")
	}
}

func TestAddTag(t *testing.T) {
	tlvBuilder := NewBuilder()
	tlvBuilder.AddTag("1", "thiago")
	tlvBuilder.AddTag("2", "1234567890")
	tlvBuilder.AddTag("3", "!@#$%^&*()")
	tlvBuilder.AddTag("10", "Hello World")
	tlvBuilder.AddTag("55", "Golang is awesome")
	tlvBuilder.AddTag("4", "Codigo")
	tlvBuilder.AddTag("43", "TR!")
	tlvBuilder.AddTag("5", "PERSYSTE")
	tlvBuilder.AddTag("6", "ESTABELECIMENTO Comercial")
	tlvBuilder.AddTag("7", "00.00.0000/0000-00")
	tlvBuilder.AddTag("8", "PAN")
	tlvBuilder.AddTag("9", "DEC")
	tlvBuilder.AddTag("11", "RESERVADO")
	tlvBuilder.AddTag("12", "DISC1001")
	tlvBuilder.AddTag("13", "DE48")
}

func TestBuild(t *testing.T) {
	tlvBuilder := NewBuilder()
	tlvBuilder.AddTag("1", "thiago")
	tlvBuilder.AddTag("2", "1234567890")
	tlvBuilder.AddTag("3", "!@#$%^&*()")
	tlvBuilder.AddTag("10", "Hello World")
	tlvBuilder.AddTag("55", "Golang is awesome")
	tlvBuilder.AddTag("4", "Codigo")
	tlvBuilder.AddTag("43", "TR!")
	tlvBuilder.AddTag("5", "PERSYSTE")
	tlvBuilder.AddTag("6", "ESTABELECIMENTO Comercial")
	tlvBuilder.AddTag("7", "00.00.0000/0000-00")
	tlvBuilder.AddTag("8", "PAN")
	tlvBuilder.AddTag("9", "DEC")
	tlvBuilder.AddTag("11", "RESERVADO")
	tlvBuilder.AddTag("12", "DISC1001")
	tlvBuilder.AddTag("13", "DE48")

	_, err := tlvBuilder.Build()

	assert.NoError(t, err)
}

func TestBuildPositional(t *testing.T) {
	tlvBuilder := NewBuilder()
	tlvBuilder.AddTag("1", "thiago")
	tlvBuilder.AddTag("2", "1234567890")
	tlvBuilder.AddTag("3", "!@#$%^&*()")
	tlvBuilder.AddTag("10", "Hello World")
	tlvBuilder.AddTag("55", "Golang is awesome")
	tlvBuilder.AddTag("4", "Codigo")
	tlvBuilder.AddTag("43", "TR!")
	tlvBuilder.AddTag("5", "PERSYSTE")
	tlvBuilder.AddTag("6", "ESTABELECIMENTO Comercial")
	tlvBuilder.AddTag("7", "00.00.0000/0000-00")
	tlvBuilder.AddTag("8", "PAN")
	tlvBuilder.AddTag("9", "DEC")
	tlvBuilder.AddTag("11", "RESERVADO")
	tlvBuilder.AddTag("12", "DISC1001")
	tlvBuilder.AddTag("13", "DE48")

	positional := tlvBuilder.BuildPositional()

	assert.NotEmpty(t, positional)
}

func TestParse(t *testing.T) {
	tlvBuilder := NewBuilder()
	tlvBuilder.AddTag("1", "thiago")
	tlvBuilder.AddTag("2", "1234567890")
	tlvBuilder.AddTag("3", "!@#$%^&*()")
	tlvBuilder.AddTag("10", "Hello World")
	tlvBuilder.AddTag("55", "Golang is awesome")
	tlvBuilder.AddTag("4", "Codigo")
	tlvBuilder.AddTag("43", "TR!")
	tlvBuilder.AddTag("5", "PERSYSTE")
	tlvBuilder.AddTag("6", "ESTABELECIMENTO Comercial")
	tlvBuilder.AddTag("7", "00.00.0000/0000-00")
	tlvBuilder.AddTag("8", "PAN")
	tlvBuilder.AddTag("9", "DEC")
	tlvBuilder.AddTag("11", "RESERVADO")
	tlvBuilder.AddTag("12", "DISC1001")
	tlvBuilder.AddTag("13", "DE48")

	data, err := tlvBuilder.Build()

	assert.NoError(t, err)

	tlvs, err := Parse(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, tlvs)
}

func TestParsePositional(t *testing.T) {
	tlvBuilder := NewBuilder()
	tlvBuilder.AddTag("1", "thiago")
	tlvBuilder.AddTag("2", "1234567890")
	tlvBuilder.AddTag("3", "!@#$%^&*()")
	tlvBuilder.AddTag("10", "Hello World")
	tlvBuilder.AddTag("55", "Golang is awesome")
	tlvBuilder.AddTag("4", "Codigo")
	tlvBuilder.AddTag("43", "TR!")
	tlvBuilder.AddTag("5", "PERSYSTE")
	tlvBuilder.AddTag("6", "ESTABELECIMENTO Comercial")
	tlvBuilder.AddTag("7", "00.00.0000/0000-00")
	tlvBuilder.AddTag("8", "PAN")
	tlvBuilder.AddTag("9", "DEC")
	tlvBuilder.AddTag("11", "RESERVADO")
	tlvBuilder.AddTag("12", "DISC1001")
	tlvBuilder.AddTag("13", "DE48")

	positional := tlvBuilder.BuildPositional()

	assert.NotEmpty(t, positional)

	tlvs, err := ParsePositional(positional)

	assert.NoError(t, err)
	assert.NotEmpty(t, tlvs)
}

func TestParsePositionalError(t *testing.T) {
	_, err := ParsePositional("1234")

	assert.Error(t, err)
}

func TestParseError(t *testing.T) {
	_, err := Parse([]byte{0x01, 0x02, 0x03, 0x04})

	assert.Error(t, err)
}

func TestParseErrorIncompleteTag(t *testing.T) {
	_, err := Parse([]byte{0x00, 0x02, 0x03, 0x04})

	assert.Error(t, err)
}

func TestParseErrorIncompleteLength(t *testing.T) {
	_, err := Parse([]byte{0x00, 0x02, 0x00, 0x04})

	assert.Error(t, err)
}
