// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagitlabartifactfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabArtifactFileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Path to the artifact file within the job's artifacts archive.
	//
	// This path is relative to the archive contents (not the local filesystem). Ensure each `gitlab_artifact_file` data source in your configuration uses a unique artifact_path to avoid ambiguity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/data-sources/artifact_file#artifact_path DataGitlabArtifactFile#artifact_path}
	ArtifactPath *string `field:"required" json:"artifactPath" yaml:"artifactPath"`
	// The name of the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/data-sources/artifact_file#job DataGitlabArtifactFile#job}
	Job *string `field:"required" json:"job" yaml:"job"`
	// The ID or URL-encoded path of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/data-sources/artifact_file#project DataGitlabArtifactFile#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The name of the branch, tag, or commit SHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/data-sources/artifact_file#ref DataGitlabArtifactFile#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// Maximum bytes to read from the artifact. Defaults to 10MB (10485760 bytes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.6.0/docs/data-sources/artifact_file#max_size_bytes DataGitlabArtifactFile#max_size_bytes}
	MaxSizeBytes *float64 `field:"optional" json:"maxSizeBytes" yaml:"maxSizeBytes"`
}

