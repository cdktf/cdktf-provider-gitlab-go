package repositoryfile


type RepositoryFileTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/repository_file#create RepositoryFile#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/repository_file#delete RepositoryFile#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/15.10.0/docs/resources/repository_file#update RepositoryFile#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

