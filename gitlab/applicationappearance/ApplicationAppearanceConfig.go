// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationappearance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationAppearanceConfig struct {
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
	// Markdown text shown on the sign-in and sign-up page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#description ApplicationAppearance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Add header and footer to all outgoing emails if enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#email_header_and_footer_enabled ApplicationAppearance#email_header_and_footer_enabled}
	EmailHeaderAndFooterEnabled interface{} `field:"optional" json:"emailHeaderAndFooterEnabled" yaml:"emailHeaderAndFooterEnabled"`
	// Message in the system footer bar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#footer_message ApplicationAppearance#footer_message}
	FooterMessage *string `field:"optional" json:"footerMessage" yaml:"footerMessage"`
	// Message in the system header bar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#header_message ApplicationAppearance#header_message}
	HeaderMessage *string `field:"optional" json:"headerMessage" yaml:"headerMessage"`
	// Set to true if the appearance settings should not be reset to their pre-terraform defaults on destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#keep_settings_on_destroy ApplicationAppearance#keep_settings_on_destroy}
	KeepSettingsOnDestroy interface{} `field:"optional" json:"keepSettingsOnDestroy" yaml:"keepSettingsOnDestroy"`
	// Markdown text shown on the group or project member page for users with permission to change members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#member_guidelines ApplicationAppearance#member_guidelines}
	MemberGuidelines *string `field:"optional" json:"memberGuidelines" yaml:"memberGuidelines"`
	// Background color for the system header or footer bar, in CSS hex notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#message_background_color ApplicationAppearance#message_background_color}
	MessageBackgroundColor *string `field:"optional" json:"messageBackgroundColor" yaml:"messageBackgroundColor"`
	// Font color for the system header or footer bar, in CSS hex notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#message_font_color ApplicationAppearance#message_font_color}
	MessageFontColor *string `field:"optional" json:"messageFontColor" yaml:"messageFontColor"`
	// Markdown text shown on the new project page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#new_project_guidelines ApplicationAppearance#new_project_guidelines}
	NewProjectGuidelines *string `field:"optional" json:"newProjectGuidelines" yaml:"newProjectGuidelines"`
	// Markdown text shown on the profile page below the Public Avatar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#profile_image_guidelines ApplicationAppearance#profile_image_guidelines}
	ProfileImageGuidelines *string `field:"optional" json:"profileImageGuidelines" yaml:"profileImageGuidelines"`
	// An explanation of what the Progressive Web App does. Used for the attribute `description` in `manifest.json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#pwa_description ApplicationAppearance#pwa_description}
	PwaDescription *string `field:"optional" json:"pwaDescription" yaml:"pwaDescription"`
	// Full name of the Progressive Web App. Used for the attribute `name` in `manifest.json`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#pwa_name ApplicationAppearance#pwa_name}
	PwaName *string `field:"optional" json:"pwaName" yaml:"pwaName"`
	// Short name for Progressive Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#pwa_short_name ApplicationAppearance#pwa_short_name}
	PwaShortName *string `field:"optional" json:"pwaShortName" yaml:"pwaShortName"`
	// Application title on the sign-in and sign-up page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/gitlabhq/gitlab/18.2.0/docs/resources/application_appearance#title ApplicationAppearance#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

