package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func parseArgs() string {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect usage!!")
		os.Exit(1)
	}
	var binary_name string = os.Args[1]
	return binary_name
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func pow(num int64, power int) int64 {
	var ret int64 = 1
	for i := 0; i < power; i++ {
		ret *= num
	}
	return ret
}
func byteArrToInt(arr []byte) int64 {
	sz := len(arr)
	var ret int64 = 0
	for i := 0; i < sz; i++ {
		ret += int64(arr[i]) * pow(16, i)
	}
	return ret
}

func prettyPrintHeader(header Header) {
	fmt.Println("###### ELF Header ######")
	fmt.Printf("Class: %d\n", header.Class)
	fmt.Printf("Endian: %d\n", header.Endian)
	fmt.Printf("Version: %d\n", header.Version)
	fmt.Printf("ABI: %d\n", header.OsABI)
    fmt.Printf("Type: %d\n", header.Type)
    fmt.Printf("Machine: %d\n", header.Machine)
    fmt.Printf("Entry Point: 0x%x\n", header.Entry)
    fmt.Printf("Program Header table offset: %d\n", header.Phdr_offset)
    fmt.Printf("Section Header table offset: %d\n", header.Shdr_offset)
    fmt.Printf("Flags: %d\n", header.Flags)
    fmt.Printf("Header size: %d\n", header.Hdr_sz)
    fmt.Printf("Program header entry size: %d\n", header.Phdr_entry_sz)
    fmt.Printf("Program header entries: %d\n", header.Phdr_entries)
    fmt.Printf("Section header entry size: %d\n", header.Shdr_entry_sz)
    fmt.Printf("Section header entries: %d\n", header.Shdr_entries)
    fmt.Printf("Section name entry: %d\n", header.Shstr)
    fmt.Printf("\n")
}

func prettyPrintPheader(program_header ProgramHeader) {
    fmt.Printf("Type: %d\n", program_header.Type)
    fmt.Printf("Flags: %d\n", program_header.Flags)
    fmt.Printf("Offset: %d\n", program_header.Offset)
    fmt.Printf("Virtual address: %d\n", program_header.Virt_addr)
    fmt.Printf("Physical address: %d\n", program_header.Phy_addr)
    fmt.Printf("Segment size in file: %d\n", program_header.Segment_file_sz)
    fmt.Printf("Segment size in memory: %d\n", program_header.Segment_memory_sz)
    fmt.Printf("Alignment: %d\n", program_header.Align)
    fmt.Printf("\n")
}

func prettyPrintSheader(section_header SectionHeader) {
    fmt.Printf("Name offset: %d\n", section_header.Name)
    fmt.Printf("Type: %d\n", section_header.Type)
    fmt.Printf("Flags: %d\n", section_header.Flags)
    fmt.Printf("Virtual address: %d\n", section_header.Virt_addr)
    fmt.Printf("File offset: %d\n", section_header.Offset)
    fmt.Printf("Segment size in file: %d\n", section_header.Segment_file_sz)
    fmt.Printf("Link to another section: %d\n", section_header.Link)
    fmt.Printf("Additional info: %d\n", section_header.Info)
    fmt.Printf("Alignment: %d\n", section_header.Name)
    fmt.Printf("Entry size: %d\n", section_header.Entry_sz)
    fmt.Printf("\n")
}

func main() {
	// parse the cli-arguments
	var binary_name string = parseArgs()

	//open the binary to read
	f, err := os.Open(binary_name)
	check(err)
	defer f.Close()

	// read the header into the header struct
	var header Header = Header{}

	// check if it is a valid ELF binary
	var magic [4]byte
	err = binary.Read(f, binary.LittleEndian, &magic)
	check(err)
	MAGIC := [4]byte{127, 69, 76, 70}
	
    if magic != MAGIC {
		fmt.Println("Not an ELF binary!!")
		os.Exit(1)
	}

	// populate the struct
	err = binary.Read(f, binary.LittleEndian, &header)
	check(err)

    // fmt.Println(header)
	prettyPrintHeader(header)
    
    // seek to the start of program header
    f.Seek(int64(header.Phdr_offset), 0)

    // now we parse the program header
    var program_header ProgramHeader = ProgramHeader{}
    fmt.Println("###### ELF Program Header ######")

    for i := 0; i < int(header.Phdr_entries); i++ {
        err = binary.Read(f, binary.LittleEndian, &program_header)
        check(err)
        fmt.Printf("Entry #%d\n", i + 1)
        prettyPrintPheader(program_header)
    }

    // section header table of the ELF file
    f.Seek(int64(header.Shdr_offset), 0)
    
    // parse the section header
    var section_header SectionHeader = SectionHeader{}
    fmt.Println("###### Section Header Entries ######")

    for i := 0; i < int(header.Shdr_entries); i++ {
        err = binary.Read(f, binary.LittleEndian, &section_header)
        check(err)
        fmt.Printf("Entry #%d\n", i + 1)
        prettyPrintSheader(section_header)

    }



}
