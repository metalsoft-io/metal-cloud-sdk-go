// DO NOT EDIT: this file is automatically generated by docgen
package metalcloud

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var (
	OperatingSystemDoc        encoder.Doc
	NetworkOperatingSystemDoc encoder.Doc
)

func init() {
	OperatingSystemDoc.Type = "OperatingSystem"
	OperatingSystemDoc.Comments[encoder.LineComment] = "OperatingSystem describes an OS"
	OperatingSystemDoc.Description = "OperatingSystem describes an OS"
	OperatingSystemDoc.Fields = make([]encoder.Doc, 3)
	OperatingSystemDoc.Fields[0].Name = "type"
	OperatingSystemDoc.Fields[0].Type = "string"
	OperatingSystemDoc.Fields[0].Note = ""
	OperatingSystemDoc.Fields[0].Description = "description: operating system type\nexamples:\n	- Ubuntu\n - CentOS\n"
	OperatingSystemDoc.Fields[0].Comments[encoder.LineComment] = "description: operating system type"
	OperatingSystemDoc.Fields[1].Name = "version"
	OperatingSystemDoc.Fields[1].Type = "string"
	OperatingSystemDoc.Fields[1].Note = ""
	OperatingSystemDoc.Fields[1].Description = "operating system type"
	OperatingSystemDoc.Fields[1].Comments[encoder.LineComment] = "operating system type"
	OperatingSystemDoc.Fields[2].Name = "architecture"
	OperatingSystemDoc.Fields[2].Type = "string"
	OperatingSystemDoc.Fields[2].Note = ""
	OperatingSystemDoc.Fields[2].Description = "architecture of the operating system"
	OperatingSystemDoc.Fields[2].Comments[encoder.LineComment] = "architecture of the operating system"
	OperatingSystemDoc.Fields[2].Values = []string{
		"none",
		"unknown",
		"x86",
		"x86_64",
	}

	NetworkOperatingSystemDoc.Type = "NetworkOperatingSystem"
	NetworkOperatingSystemDoc.Comments[encoder.LineComment] = ""
	NetworkOperatingSystemDoc.Description = ""
	NetworkOperatingSystemDoc.Fields = make([]encoder.Doc, 7)
	NetworkOperatingSystemDoc.Fields[0].Name = "architecture"
	NetworkOperatingSystemDoc.Fields[0].Type = "string"
	NetworkOperatingSystemDoc.Fields[0].Note = ""
	NetworkOperatingSystemDoc.Fields[0].Description = "architecture of the operating system"
	NetworkOperatingSystemDoc.Fields[0].Comments[encoder.LineComment] = "architecture of the operating system"
	NetworkOperatingSystemDoc.Fields[0].Values = []string{
		"arm",
		"powerpc",
		"x86_64",
	}
	NetworkOperatingSystemDoc.Fields[1].Name = "datacenter_name"
	NetworkOperatingSystemDoc.Fields[1].Type = "string"
	NetworkOperatingSystemDoc.Fields[1].Note = ""
	NetworkOperatingSystemDoc.Fields[1].Description = "The datacenter for this operating system."
	NetworkOperatingSystemDoc.Fields[1].Comments[encoder.LineComment] = "The datacenter for this operating system."
	NetworkOperatingSystemDoc.Fields[2].Name = "machine"
	NetworkOperatingSystemDoc.Fields[2].Type = "string"
	NetworkOperatingSystemDoc.Fields[2].Note = ""
	NetworkOperatingSystemDoc.Fields[2].Description = "A switch model number to be matched against an ONIE or POAP request during ztp."
	NetworkOperatingSystemDoc.Fields[2].Comments[encoder.LineComment] = "A switch model number to be matched against an ONIE or POAP request during ztp."
	NetworkOperatingSystemDoc.Fields[3].Name = "switchDriver"
	NetworkOperatingSystemDoc.Fields[3].Type = "string"
	NetworkOperatingSystemDoc.Fields[3].Note = ""
	NetworkOperatingSystemDoc.Fields[3].Description = "The switch driver to use."
	NetworkOperatingSystemDoc.Fields[3].Comments[encoder.LineComment] = "The switch driver to use."
	NetworkOperatingSystemDoc.Fields[4].Name = "switchRole"
	NetworkOperatingSystemDoc.Fields[4].Type = "string"
	NetworkOperatingSystemDoc.Fields[4].Note = ""
	NetworkOperatingSystemDoc.Fields[4].Description = "The role of a switch"
	NetworkOperatingSystemDoc.Fields[4].Comments[encoder.LineComment] = "The role of a switch"
	NetworkOperatingSystemDoc.Fields[4].Values = []string{
		"leaf",
		"spine",
	}
	NetworkOperatingSystemDoc.Fields[5].Name = "vendor"
	NetworkOperatingSystemDoc.Fields[5].Type = "string"
	NetworkOperatingSystemDoc.Fields[5].Note = ""
	NetworkOperatingSystemDoc.Fields[5].Description = "description: The vendor of the switch\nexamples:\n- dell\n- juniper\n"
	NetworkOperatingSystemDoc.Fields[5].Comments[encoder.LineComment] = "description: The vendor of the switch"
	NetworkOperatingSystemDoc.Fields[6].Name = "version"
	NetworkOperatingSystemDoc.Fields[6].Type = "string"
	NetworkOperatingSystemDoc.Fields[6].Note = ""
	NetworkOperatingSystemDoc.Fields[6].Description = "The version of the operating system"
	NetworkOperatingSystemDoc.Fields[6].Comments[encoder.LineComment] = "The version of the operating system"
}

func (OperatingSystem) Doc() *encoder.Doc {
	return &OperatingSystemDoc
}

func (NetworkOperatingSystem) Doc() *encoder.Doc {
	return &NetworkOperatingSystemDoc
}

// GetOperatingSystemDoc returns documentation for the file ./.
func GetOperatingSystemDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "OperatingSystem",
		Description: "",
		Structs: []*encoder.Doc{
			&OperatingSystemDoc,
			&NetworkOperatingSystemDoc,
		},
	}
}