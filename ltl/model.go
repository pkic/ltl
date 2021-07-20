package ltl

import (
	"time"
)

type PKI struct {
	Name        string `json:"name" yaml:"name" validate:"required"`
	Description string `json:"description" yaml:"description"`

	// A link to the trusted list program or vendor/product when no such
	// dedicated website exists.
	Website string `json:"website" yaml:"website" validate:"required,url"`

	// Email or website address to contact privately
	Contact string `json:"contact" yaml:"contact" validate:"omitempty,email|url"`

	// Using the Common CA Database
	CCADB bool `json:"ccadb" yaml:"ccadb"`

	// Legal context
	LegalContext string `json:"legal-context" yaml:"legal-context"`

	// How is compliance audited
	Audit []Audit `json:"audit" yaml:"audit" validate:"dive"`

	Requirements []Source   `json:"requirements" yaml:"requirements" validate:"dive"`
	Discussions  []Source   `json:"discussions" yaml:"discussions" validate:"dive"`
	Issues       []Source   `json:"issues" yaml:"issues" validate:"dive"`
	TrustLists   TrustLists `json:"trust-lists" yaml:"trust-lists"`
}

// Audit specifies the audit options and or requirements,
type Audit struct {
	Name        string        `json:"name" yaml:"name" validate:"required"`
	Description string        `json:"description" yaml:"description"`
	Frequency   time.Duration `json:"frequency" yaml:"frequency"`
	Policies    []string      `json:"policies" yaml:"policies"`
}

// TrustLists under the same PKI program.
// If one list has multiple distinct programs, create a new file for each.
type TrustLists struct {
	Info   string `json:"info" yaml:"info"`
	Policy string `json:"policy" yaml:"policy"`

	Trust []Trust `json:"trust" yaml:"trust" validate:"omitempty,dive"`
}

// Trust links to a file containing the CA certificates trusted by this PKI
// for the given purpose. A file might be referenced in multiple formats, the
// preferred format (type) is PEM.
type Trust struct {
	Purposes []string `json:"purposes" yaml:"purposes"`

	// When a program has trust lists for different purposes, there can be
	// generic requirements and requirements that are specific to one purpose.
	Requirements []Source `json:"requirements" yaml:"requirements" validate:"dive"`
	Discussions  []Source `json:"discussions" yaml:"discussions" validate:"dive"`
	Issues       []Source `json:"issues" yaml:"issues" validate:"dive"`
	Audit        []Audit  `json:"audit" yaml:"audit" validate:"dive"`

	List []Source `json:"list" yaml:"list" validate:"dive"`
}

// Source is a wrapper to allow specifing the type of source file
type Source struct {
	Type string `json:"type" yaml:"type" validate:"omitempty,oneof=HTML RSS ATOM XML JSON PEM PKCS7"`
	URL  string `json:"url" yaml:"url" validate:"required,url"`
}
