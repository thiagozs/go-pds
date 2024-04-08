package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/thiagozs/go-pds"
)

func main() {
	builder := pds.NewBuilder()
	builder.AddTag("1", "thiago")
	builder.AddTag("2", "1234567890")
	builder.AddTag("3", "!@#$%^&*()")
	builder.AddTag("10", "Hello World")
	builder.AddTag("55", "Golang is awesome")
	builder.AddTag("4", "Codigo")
	builder.AddTag("43", "TR!")
	builder.AddTag("5", "PERSYSTE")
	builder.AddTag("6", "ESTABELECIMENTO Comercial")
	builder.AddTag("7", "00.00.0000/0000-00")
	builder.AddTag("8", "PAN")
	builder.AddTag("9", "DEC")
	builder.AddTag("11", "RESERVADO")
	builder.AddTag("12", "DISC1001")
	builder.AddTag("13", "DE48")

	data, err := builder.Build()
	if err != nil {
		fmt.Println("Error building PDS:", err)
		return
	}

	hexr := hex.EncodeToString(data)

	fmt.Println("ENCODE > PDS data (hex):", hexr)

	databytes, err := hex.DecodeString(hexr)
	if err != nil {
		fmt.Println("Error decoding hex data:", err)
		return
	}

	fmt.Printf("DECODE < PDS data (bytes): %v\n", databytes)

	pdsvs, err := pds.Parse(databytes)
	if err != nil {
		fmt.Println("Error parsing PDS:", err)
		return
	}

	fmt.Println("PDS:")
	for _, pdsv := range pdsvs {
		fmt.Printf("  tag:  %s\n", pdsv.Tag)
		fmt.Printf("  value: %s\n", pdsv.Value)
	}

	positional := builder.BuildPositional()

	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Println("Positional String:", positional)

	pdsvs, err = pds.ParsePositional(positional)
	if err != nil {
		fmt.Println("Error parsing TLV positional:", err)
		return
	}

	fmt.Println("PDS:")
	for _, pdsv := range pdsvs {
		fmt.Printf("  tag:  %s\n", pdsv.Tag)
		fmt.Printf("  value: %s\n", pdsv.Value)
	}

	realword := "0002003MBK0003003MBK001500721080510023003POI0146036002901986000000000016986000000000016014800498620158031MCC4076000PD21080703 DMCNNNNNNN015906710391      0000000010391               3LA00098620N21080703210809010165001M0177002N 01910012"

	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Println("Positional String (realword):", realword)

	pdsvs, err = pds.ParsePositional(realword)
	if err != nil {
		fmt.Println("Error parsing TLV positional:", err)
		return
	}

	fmt.Println("PDS:")
	for _, pds := range pdsvs {
		fmt.Printf("  tag:  %s\n", pds.Tag)
		fmt.Printf("  value: %s\n", pds.Value)
	}

}
