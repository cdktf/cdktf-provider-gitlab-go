package serviceemailsonpush

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceEmailsOnPushConfig struct {
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
	// ID or full-path of the project you want to activate integration on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#project ServiceEmailsOnPush#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// Emails separated by whitespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#recipients ServiceEmailsOnPush#recipients}
	Recipients *string `field:"required" json:"recipients" yaml:"recipients"`
	// Branches to send notifications for. Valid options are `all`, `default`, `protected`, `default_and_protected`. Notifications are always fired for tag pushes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#branches_to_be_notified ServiceEmailsOnPush#branches_to_be_notified}
	BranchesToBeNotified *string `field:"optional" json:"branchesToBeNotified" yaml:"branchesToBeNotified"`
	// Disable code diffs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#disable_diffs ServiceEmailsOnPush#disable_diffs}
	DisableDiffs interface{} `field:"optional" json:"disableDiffs" yaml:"disableDiffs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#id ServiceEmailsOnPush#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable notifications for push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#push_events ServiceEmailsOnPush#push_events}
	PushEvents interface{} `field:"optional" json:"pushEvents" yaml:"pushEvents"`
	// Send from committer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#send_from_committer_email ServiceEmailsOnPush#send_from_committer_email}
	SendFromCommitterEmail interface{} `field:"optional" json:"sendFromCommitterEmail" yaml:"sendFromCommitterEmail"`
	// Enable notifications for tag push events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/16.1.1/docs/resources/service_emails_on_push#tag_push_events ServiceEmailsOnPush#tag_push_events}
	TagPushEvents interface{} `field:"optional" json:"tagPushEvents" yaml:"tagPushEvents"`
}

