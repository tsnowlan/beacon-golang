package main

// Beacon : general beacon info
type Beacon struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	APIVersion     string                   `json:"apiVersion"`
	Organization   BeaconOrganization       `json:"organization"`
	Description    string                   `json:"description"`
	Version        string                   `json:"version"`
	WelcomeURL     string                   `json:"welcomeUrl"`
	AlternativeURL string                   `json:"alternativeUrl"`
	CreateDateTime string                   `json:"createDateTime"`
	UpdateDateTime string                   `json:"updateDateTime"`
	Datasets       map[string]BeaconDataset `json:"datasets"`
}

// BeaconDataset is a description of the shared dataset
type BeaconDataset struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	AssemblyID        string            `json:"assemblyId"`     // example: GRCh38
	CreateDateTime    string            `json:"createDateTime"` // The time the dataset was created (ISO 8601 format).
	UpdateDateTime    string            `json:"updateDateTime"` // The time the dataset was created (ISO 8601 format).
	Version           string            `json:"version"`
	VariantCount      uint              `json:"variantCount"` // Total number of variants in the dataset.
	CallCount         uint              `json:"callCoutner"`  // Total number of calls in the dataset.
	SampleCount       uint              `json:"sampleCount"`  // Total number of samples in the dataset.
	ExternalURL       string            `json:"externalUrl"`  // URL to an external system providing more dataset information (RFC 3986 format).
	Info              []KeyValuePair    `json:"info"`
	DataUseConditions DataUseConditions `json:"dataUseConditions"`
}

// BeaconAlleleRequest as interpreted by the beacon
type BeaconAlleleRequest struct {
	ReferenceName           string   `json:"referenceName"` // Reference name (chromosome). Accepting values 1-22, X, Y.
	Start                   uint64   `json:"start"`
	End                     uint64   `json:"end"`
	StartMin                uint64   `json:"startMin"`
	StartMax                uint64   `json:"startMax"`
	EndMin                  uint64   `json:"endMin"`
	EndMax                  uint64   `json:"endMax"`
	ReferenceBases          string   `json:"referenceBases"` // Reference bases for this variant (starting from start). For accepted values see the REF field in VCF 4.2 specification
	AlternateBases          string   `json:"alternateBases"` // The bases that appear instead of the reference bases. For accepted values see the ALT field in VCF 4.2 specification
	AssemblyID              string   `json:"assemblyId"`     // example: GRCh38
	DatasetIDs              []string `json:"datasetIds"`
	IncludeDatasetResponses string   `json:"includeDatasetResponses"` // Valid values: ALL, HIT, MISS, NONE. If null (not specified), the default value of NONE is assumed.
}

// BeaconAlleleResponse : query response struct
type BeaconAlleleResponse struct {
	BeaconID               string                        `json:"beaconId"`
	APIVersion             string                        `json:"apiVersion"` // Version of the API. If specified, the value must match APIVersion in Beacon
	Exists                 bool                          `json:"exists"`     // This should be non-null, unless there was an error, in which case error has to be non-null.
	AlleleRequest          BeaconAlleleRequest           `json:"alleleRequest"`
	DatasetAlleleResponses []BeaconDatasetAlleleResponse `json:"datasetAlleleResponses"`
	Error                  BeaconError                   `json:"error"`
}

// BeaconOrganization Organization owning the beacon.
type BeaconOrganization struct {
	ID          string         `json:"id"` // Unique identifier of the organization.
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Address     string         `json:"address"`
	WelcomeURL  string         `json:"welcomeUrl"` // URL of the website of the organization (RFC 3986 format).
	ContactURL  string         `json:"contactUrl"` // URL with the contact for the beacon operator/maintainer, e.g. link to a contact form (RFC 3986 format) or an email (RFC 2368 format).
	LogoURL     string         `json:"logoUrl"`    // URL to the logo (PNG/JPG format) of the organization (RFC 3986 format).
	Info        []KeyValuePair `json:"info"`
}

// BeaconDatasetAlleleResponse - beacon dataset response object
type BeaconDatasetAlleleResponse struct {
	DatasetID    string         `json:"datasetId"`
	Exists       bool           `json:"exists"` // Indicator of whether the given allele was observed in the dataset. This should be non-null, unless there was an error, in which case error has to be non-null.
	Error        BeaconError    `json:"error"`
	Frequency    float64        `json:"frequency"`    // Frequency of this allele in the dataset. Between 0 and 1, inclusive.
	VariantCount uint           `json:"variantCount"` // Number of variants matching the allele request in the dataset.
	CallCount    uint           `json:"callCount"`    // Number of calls matching the allele request in the dataset.
	SampleCount  uint           `json:"sampleCount"`  // Number of samples matching the allele request in the dataset
	Note         string         `json:"note"`         // Additional note or description of the response.
	ExternalURL  string         `json:"externalUrl"`  // URL to an external system, such as a secured beacon or a system providing more information about a given allele (RFC 3986 format).
	Info         []KeyValuePair `json:"info"`
}

// BeaconError : Beacon-specific error. This should be non-null in exceptional situations only, in which case exists has to be null.
type BeaconError struct {
	ErrorCode    int32  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// KeyValuePair : KV pairing for additional info
type KeyValuePair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// DataUseConditions : Data use conditions ruling this dataset.
type DataUseConditions struct {
	ConsentCodeDataUse ConsentCodeDataUse `json:"consentCodeDataUse"`
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
	Code        string `json:"code"`
	Description string `json:"description"`
}

// ConsentCodeDataUseConditionSecondary : see below
type ConsentCodeDataUseConditionSecondary struct {
	// Secondary data use categories:
	//
	// RS-[XX]: other research-specific restrictions - Use of the data is limited to studies of [research type] (e.g., pediatric research).
	// RUO: research use only - Use of data is limited to research purposes (e.g., does not include its use in clinical care).
	// NMDS: no “general methods” research - Use of the data includes methods development research (e.g., development of software or algorithms) ONLY within the bounds of other data use limitations.
	// GSO: genetic studies only - Use of the data is limited to genetic studies only (i.e., no research using only the phenotype data).
	Code        string `json:"code"`
	Description string `json:"description"`
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
	Code        string `json:"code"`
	Description string `json:"description"`
}

// ConsentCodeDataUse : the categories the dta is available for use with
type ConsentCodeDataUse struct {
	PrimaryCategory     ConsentCodeDataUseConditionPrimary       `json:"primaryCategory"`
	SecondaryCategories []ConsentCodeDataUseConditionSecondary   `json:"secondaryCategory"`
	Requirements        []ConsentCodeDataUseConditionRequirement `json:"requirements"`
	Version             string                                   `json:"version"` // Version of the data use specification., e.g., v0.1
}
