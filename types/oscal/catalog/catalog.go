// Code generated by go generate; DO NOT EDIT.
package catalog

import (
	"encoding/xml"

	"github.com/docker/oscalkit/types/oscal/nominal_catalog"

	"github.com/docker/oscalkit/types/oscal/validation_root"
)

// A collection of controls.
type Catalog struct {
	XMLName xml.Name `xml:"http://csrc.nist.gov/ns/oscal/1.0 catalog" json:"-"`
	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Provides information about the publication and availability of the containing document.
	Metadata *validation_root.Metadata `xml:"metadata,omitempty" json:"metadata,omitempty"`
	// Back matter including references and resources.
	BackMatter *validation_root.BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
	// A group of controls, or of groups of controls.
	Group []Group `xml:"groups,omitempty" json:"groups,omitempty"`
	// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
	Control []Control `xml:"controls,omitempty" json:"controls,omitempty"`
}

// A group of controls, or of groups of controls.
type Group struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

	// A title for display and navigation
	Title validation_root.Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Prop []validation_root.Prop `xml:"properties,omitempty" json:"properties,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Param []nominal_catalog.Param `xml:"parameters,omitempty" json:"parameters,omitempty"`
	// A partition or component of a control or part
	Part []nominal_catalog.Part `xml:"parts,omitempty" json:"parts,omitempty"`
	// A group of controls, or of groups of controls.
	Group []Group `xml:"groups,omitempty" json:"groups,omitempty"`
	// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
	Control []Control `xml:"controls,omitempty" json:"controls,omitempty"`
}

// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
type Control struct {

	// Unique identifier of the containing object
	Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
	// Indicating the type or classification of the containing object
	Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

	// A title for display and navigation
	Title validation_root.Title `xml:"title,omitempty" json:"title,omitempty"`
	// A value with a name, attributed to the containing control, part, or group.
	Prop []validation_root.Prop `xml:"properties,omitempty" json:"properties,omitempty"`
	// A reference to a local or remote resource
	Link []validation_root.Link `xml:"links,omitempty" json:"links,omitempty"`
	// Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
	Param []nominal_catalog.Param `xml:"parameters,omitempty" json:"parameters,omitempty"`
	// A name/value pair with optional explanatory remarks.
	Annotation []validation_root.Annotation `xml:"annotations,omitempty" json:"annotations,omitempty"`
	// A partition or component of a control or part
	Part []nominal_catalog.Part `xml:"parts,omitempty" json:"parts,omitempty"`
	// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
	Control []Control `xml:"controls,omitempty" json:"controls,omitempty"`
}
