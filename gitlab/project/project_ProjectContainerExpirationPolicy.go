package project


type ProjectContainerExpirationPolicy struct {
	// The cadence of the policy. Valid values are: `1d`, `7d`, `14d`, `1month`, `3month`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#cadence Project#cadence}
	Cadence *string `field:"optional" json:"cadence" yaml:"cadence"`
	// If true, the policy is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#enabled Project#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The number of images to keep.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#keep_n Project#keep_n}
	KeepN *float64 `field:"optional" json:"keepN" yaml:"keepN"`
	// The regular expression to match image names to delete.
	//
	// **Note**: the upstream API has some inconsistencies with the `name_regex` field here. It's basically unusable at the moment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#name_regex_delete Project#name_regex_delete}
	NameRegexDelete *string `field:"optional" json:"nameRegexDelete" yaml:"nameRegexDelete"`
	// The regular expression to match image names to keep.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#name_regex_keep Project#name_regex_keep}
	NameRegexKeep *string `field:"optional" json:"nameRegexKeep" yaml:"nameRegexKeep"`
	// The number of days to keep images.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/gitlab/r/project#older_than Project#older_than}
	OlderThan *string `field:"optional" json:"olderThan" yaml:"olderThan"`
}

