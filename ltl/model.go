package ltl

import (
	"time"
)

type PKI struct {
	Name        string `yaml:"name" validate:"required"`
	Description string `yaml:"description" validate:"required"`
	Website     string `yaml:"website" validate:"required,url"`

	// Email or website address to contact privately
	Contact string `yaml:"contact" validate:"email|url"`

	// Using the Common CA Database
	CCADB bool `yaml:"ccadb"`

	// Legal context
	LegalContext string `yaml:"legal-context"`

	// How is compliance audited
	Audit []Audit `yaml:"audit" validate:"dive"`

	Requirements []Source   `yaml:"requirements" validate:"dive"`
	Discussions  []Source   `yaml:"discussions" validate:"dive"`
	Issues       []Source   `yaml:"issues" validate:"dive"`
	TrustLists   TrustLists `yaml:"trust-lists"`
}

// Audit specifies the audit options and or requirements,
type Audit struct {
	Name        string        `yaml:"name" validate:"required"`
	Description string        `yaml:"description"`
	Frequency   time.Duration `yaml:"frequency"`
	Policies    []string      `yaml:"policies"`
}

// TrustLists under the same PKI program.
// If one list has multiple distinct programs, create a new file for each.
type TrustLists struct {
	Info   string `yaml:"info"`
	Policy string `yaml:"policy"`

	Trust []Trust `yaml:"trust" validate:"dive"`
}

// Trust links to a file containing the CA certificates trusted by this PKI
// for the given purpose. A file might be referenced in multiple formats, the
// preferred format (type) is PEM.
type Trust struct {
	Purposes []string `yaml:"purposes"`

	// When a program has trust lists for different purposes, there can be
	// generic requirements and requirements that are specific to one purpose.
	Requirements []Source `yaml:"requirements" validate:"dive"`
	Discussions  []Source `yaml:"discussions" validate:"dive"`
	Issues       []Source `yaml:"issues" validate:"dive"`
	Audit        []Audit  `yaml:"audit" validate:"dive"`

	List []Source `yaml:"list" validate:"dive"`
}

// Source is a wrapper to allow specifing the type of source file
type Source struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url" validate:"required,url"`
}
