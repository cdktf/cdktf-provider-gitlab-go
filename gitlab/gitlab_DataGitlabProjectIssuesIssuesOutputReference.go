// Prebuilt gitlab Provider for Terraform CDK (cdktf)
package gitlab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/jsii"

	"github.com/hashicorp/cdktf-provider-gitlab-go/gitlab/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGitlabProjectIssuesIssuesOutputReference interface {
	cdktf.ComplexObject
	AssigneeIds() *[]*float64
	AuthorId() *float64
	ClosedAt() *string
	ClosedByUserId() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Confidential() cdktf.IResolvable
	CreatedAt() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	DiscussionLocked() cdktf.IResolvable
	DiscussionToResolve() *string
	Downvotes() *float64
	DueDate() *string
	EpicId() *float64
	EpicIssueId() *float64
	ExternalId() *string
	// Experimental.
	Fqn() *string
	HumanTimeEstimate() *string
	HumanTotalTimeSpent() *string
	Iid() *float64
	InternalValue() *DataGitlabProjectIssuesIssues
	SetInternalValue(val *DataGitlabProjectIssuesIssues)
	IssueId() *float64
	IssueLinkId() *float64
	IssueType() *string
	Labels() *[]*string
	Links() cdktf.StringMap
	MergeRequestsCount() *float64
	MergeRequestToResolveDiscussionsOf() *float64
	MilestoneId() *float64
	MovedToId() *float64
	Project() *string
	References() cdktf.StringMap
	State() *string
	Subscribed() cdktf.IResolvable
	TaskCompletionStatus() DataGitlabProjectIssuesIssuesTaskCompletionStatusList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeEstimate() *float64
	Title() *string
	TotalTimeSpent() *float64
	UpdatedAt() *string
	Upvotes() *float64
	UserNotesCount() *float64
	WebUrl() *string
	Weight() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGitlabProjectIssuesIssuesOutputReference
type jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) AssigneeIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"assigneeIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) AuthorId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ClosedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"closedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ClosedByUserId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"closedByUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Confidential() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"confidential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) DiscussionLocked() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"discussionLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) DiscussionToResolve() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discussionToResolve",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Downvotes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"downvotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) DueDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dueDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) EpicId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) EpicIssueId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"epicIssueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) HumanTimeEstimate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTimeEstimate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) HumanTotalTimeSpent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTotalTimeSpent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Iid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) InternalValue() *DataGitlabProjectIssuesIssues {
	var returns *DataGitlabProjectIssuesIssues
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) IssueId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) IssueLinkId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"issueLinkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) IssueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Links() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"links",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) MergeRequestsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) MergeRequestToResolveDiscussionsOf() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mergeRequestToResolveDiscussionsOf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) MilestoneId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"milestoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) MovedToId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"movedToId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) References() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Subscribed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"subscribed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) TaskCompletionStatus() DataGitlabProjectIssuesIssuesTaskCompletionStatusList {
	var returns DataGitlabProjectIssuesIssuesTaskCompletionStatusList
	_jsii_.Get(
		j,
		"taskCompletionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) TimeEstimate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeEstimate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) TotalTimeSpent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTimeSpent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Upvotes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"upvotes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) UserNotesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userNotesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) WebUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Weight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weight",
		&returns,
	)
	return returns
}


func NewDataGitlabProjectIssuesIssuesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGitlabProjectIssuesIssuesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabProjectIssuesIssuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGitlabProjectIssuesIssuesOutputReference_Override(d DataGitlabProjectIssuesIssuesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-gitlab.DataGitlabProjectIssuesIssuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) SetInternalValue(val *DataGitlabProjectIssuesIssues) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGitlabProjectIssuesIssuesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

