package main

// Beacon : general beacon info
type Beacon struct {
	ID             string
	Name           string
	APIVersion     string
	Organization   BeaconOrganization
	Description    string
	Version        string
	WelcomeURL     string
	AlternativeURL string
	CreateDateTime string
	UpdateDateTime string
	Datasets       map[string]BeaconDataset
}

// BeaconDataset is a description of the shared dataset
type BeaconDataset struct {
	ID                string
	Name              string
	AssemblyID        string // example: GRCh38
	CreateDateTime    string // The time the dataset was created (ISO 8601 format).
	UpdateDateTime    string // The time the dataset was created (ISO 8601 format).
	Version           string
	VariantCount      uint   // Total number of variants in the dataset.
	CallCount         uint   // Total number of calls in the dataset.
	SampleCount       uint   // Total number of samples in the dataset.
	ExternalURL       string // URL to an external system providing more dataset information (RFC 3986 format).
	Info              map[string]string
	dataUseConditions DataUseConditions
}

// BeaconAlleleRequest as interpreted by the beacon
type BeaconAlleleRequest struct {
	ReferenceName           string // Reference name (chromosome). Accepting values 1-22, X, Y.
	Start                   uint64
	End                     uint64
	StartMin                uint64
	StartMax                uint64
	EndMin                  uint64
	EndMax                  uint64
	ReferenceBases          string // Reference bases for this variant (starting from start). For accepted values see the REF field in VCF 4.2 specification
	AlternateBases          string // The bases that appear instead of the reference bases. For accepted values see the ALT field in VCF 4.2 specification
	AssemblyID              string // example: GRCh38
	DatasetIds              []string
	IncludeDatasetResponses string // Valid values: ALL, HIT, MISS, NONE. If null (not specified), the default value of NONE is assumed.
}

// BeaconAlleleResponse : query response struct
type BeaconAlleleResponse struct {
	BeaconID               string
	APIVersion             string // Version of the API. If specified, the value must match APIVersion in Beacon
	Exists                 bool   // This should be non-null, unless there was an error, in which case error has to be non-null.
	AlleleRequest          BeaconAlleleRequest
	DatasetAlleleResponses []BeaconDatasetAlleleResponse
	Error                  BeaconError
}

// BeaconOrganization Organization owning the beacon.
type BeaconOrganization struct {
	ID          string // Unique identifier of the organization.
	Name        string
	Description string
	Address     string
	WelcomeURL  string // URL of the website of the organization (RFC 3986 format).
	ContactURL  string // URL with the contact for the beacon operator/maintainer, e.g. link to a contact form (RFC 3986 format) or an email (RFC 2368 format).
	LogoURL     string // URL to the logo (PNG/JPG format) of the organization (RFC 3986 format).
	Info        []KeyValuePair
}

// BeaconDatasetAlleleResponse - beacon dataset response object
type BeaconDatasetAlleleResponse struct {
	DatasetID    string
	Exists       bool // Indicator of whether the given allele was observed in the dataset. This should be non-null, unless there was an error, in which case error has to be non-null.
	Error        BeaconError
	Frequency    float64 // Frequency of this allele in the dataset. Between 0 and 1, inclusive.
	VariantCount uint    // Number of variants matching the allele request in the dataset.
	CallCount    uint    // Number of calls matching the allele request in the dataset.
	SampleCount  uint    // Number of samples matching the allele request in the dataset
	Note         string  // Additional note or description of the response.
	ExternalURL  string  // URL to an external system, such as a secured beacon or a system providing more information about a given allele (RFC 3986 format).
	Info         []KeyValuePair
}

// BeaconError : Beacon-specific error. This should be non-null in exceptional situations only, in which case exists has to be null.
type BeaconError struct {
	ErrorCode    int32
	ErrorMessage string
}

// KeyValuePair : KV pairing for additional info
type KeyValuePair struct {
	Key   string
	Value string
}

// DataUseConditions : Data use conditions ruling this dataset.
type DataUseConditions struct {
	ConsentCodeDataUse ConsentCodeDataUse
}

// ConsentCodeDataUseConditionPrimary : see below
type ConsentCodeDataUseConditionPrimary struct {
	// Primary data use categories:
	//
	// NRES: no restrictions - No restrictions on data use.
	// GRU(CC): general research use and clinical care - For health/medical/biomedical purposes and other biological research, including the study of population origins or ancestry.
	// HMB(CC): health/medical/biomedical research and clinical care - Use of the data is limited to health/medical/biomedical purposes, does not include the study of population origins or ancestry.
	// DS-XX: disease-specific research and clinical care - Use of the data must be related to [disease].
	// POA: population origins/ancestry research - Use of the data is limited to the study of population origins or ancestry.
	Code        string
	Description string
}

// ConsentCodeDataUseConditionSecondary : see below
type ConsentCodeDataUseConditionSecondary struct {
	// Secondary data use categories:
	//
	// RS-[XX]: other research-specific restrictions - Use of the data is limited to studies of [research type] (e.g., pediatric research).
	// RUO: research use only - Use of data is limited to research purposes (e.g., does not include its use in clinical care).
	// NMDS: no “general methods” research - Use of the data includes methods development research (e.g., development of software or algorithms) ONLY within the bounds of other data use limitations.
	// GSO: genetic studies only - Use of the data is limited to genetic studies only (i.e., no research using only the phenotype data).
	Code        string
	Description string
}

// ConsentCodeDataUseConditionRequirement : see below
type ConsentCodeDataUseConditionRequirement struct {
	// Data use requirements:
	//
	// NPU: not-for-profit use only - Use of the data is limited to not-for-profit organizations.
	// PUB: publication required - Requestor agrees to make results of studies using the data available to the larger scientific community.
	// COL-[XX]: collaboration required - Requestor must agree to collaboration with the primary study investigator(s).
	// RTN: return data to database/resource - Requestor must return derived/enriched data to the database/resource.
	// IRB: ethics approval required - Requestor must provide documentation of local IRB/REC approval.
	// GS-[XX]: geographical restrictions - Use of the data is limited to within [geographic region].
	// MOR-[XX]: publication moratorium/embargo - Requestor agrees not to publish results of studies until [date].
	// TS-[XX]: time limits on use - Use of data is approved for [x months].
	// US: user-specific restrictions - Use of data is limited to use by approved users.
	// PS: project-specific restrictions - Use of data is limited to use within an approved project.
	// IS: institution-specific restrictions - Use of data is limited to use within an approved institution.
	Code        *string
	Description string
}

// ConsentCodeDataUse : the categories the dta is available for use with
type ConsentCodeDataUse struct {
	PrimaryCategory     ConsentCodeDataUseConditionPrimary
	SecondaryCategories []ConsentCodeDataUseConditionSecondary
	requirements        []ConsentCodeDataUseConditionRequirement
	Version             string // Version of the data use specification., e.g., v0.1
}
