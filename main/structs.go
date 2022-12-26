package main

// structure for the ELF header
type Header struct {
	Class         uint8
	Endian        uint8
	Version       uint8
	OsABI         uint8
	_             uint64
	Type          uint16
	Machine       uint16
	Version_e     uint32
	Entry         uint64
	Phdr_offset   uint64
	Shdr_offset   uint64
	Flags         uint32
	Hdr_sz        uint16
	Phdr_entry_sz uint16
	Phdr_entries  uint16
	Shdr_entry_sz uint16
	Shdr_entries  uint16
	Shstr         uint16
}

//structure for program header entries
type ProgramHeader struct {
    Type              uint32
    Flags             uint32
    Offset            uint64
    Virt_addr         uint64
    Phy_addr          uint64
    Segment_file_sz   uint64
    Segment_memory_sz uint64
    Align             uint64    
}

//structure for section header entries
type SectionHeader struct {
    Name            uint32
    Type            uint32
    Flags           uint64
    Virt_addr       uint64
    Offset          uint64
    Segment_file_sz uint64
    Link            uint32
    Info            uint32
    Align           uint64
    Entry_sz        uint64
}
