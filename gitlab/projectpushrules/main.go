// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectpushrules

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesA",
		reflect.TypeOf((*ProjectPushRulesA)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authorEmailRegex", GoGetter: "AuthorEmailRegex"},
			_jsii_.MemberProperty{JsiiProperty: "authorEmailRegexInput", GoGetter: "AuthorEmailRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "branchNameRegex", GoGetter: "BranchNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "branchNameRegexInput", GoGetter: "BranchNameRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterCheck", GoGetter: "CommitCommitterCheck"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterCheckInput", GoGetter: "CommitCommitterCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterNameCheck", GoGetter: "CommitCommitterNameCheck"},
			_jsii_.MemberProperty{JsiiProperty: "commitCommitterNameCheckInput", GoGetter: "CommitCommitterNameCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageNegativeRegex", GoGetter: "CommitMessageNegativeRegex"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageNegativeRegexInput", GoGetter: "CommitMessageNegativeRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageRegex", GoGetter: "CommitMessageRegex"},
			_jsii_.MemberProperty{JsiiProperty: "commitMessageRegexInput", GoGetter: "CommitMessageRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "denyDeleteTag", GoGetter: "DenyDeleteTag"},
			_jsii_.MemberProperty{JsiiProperty: "denyDeleteTagInput", GoGetter: "DenyDeleteTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fileNameRegex", GoGetter: "FileNameRegex"},
			_jsii_.MemberProperty{JsiiProperty: "fileNameRegexInput", GoGetter: "FileNameRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxFileSize", GoGetter: "MaxFileSize"},
			_jsii_.MemberProperty{JsiiProperty: "maxFileSizeInput", GoGetter: "MaxFileSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "memberCheck", GoGetter: "MemberCheck"},
			_jsii_.MemberProperty{JsiiProperty: "memberCheckInput", GoGetter: "MemberCheckInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preventSecrets", GoGetter: "PreventSecrets"},
			_jsii_.MemberProperty{JsiiProperty: "preventSecretsInput", GoGetter: "PreventSecretsInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "rejectNonDcoCommits", GoGetter: "RejectNonDcoCommits"},
			_jsii_.MemberProperty{JsiiProperty: "rejectNonDcoCommitsInput", GoGetter: "RejectNonDcoCommitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rejectUnsignedCommits", GoGetter: "RejectUnsignedCommits"},
			_jsii_.MemberProperty{JsiiProperty: "rejectUnsignedCommitsInput", GoGetter: "RejectUnsignedCommitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorEmailRegex", GoMethod: "ResetAuthorEmailRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetBranchNameRegex", GoMethod: "ResetBranchNameRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitCommitterCheck", GoMethod: "ResetCommitCommitterCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitCommitterNameCheck", GoMethod: "ResetCommitCommitterNameCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitMessageNegativeRegex", GoMethod: "ResetCommitMessageNegativeRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommitMessageRegex", GoMethod: "ResetCommitMessageRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetDenyDeleteTag", GoMethod: "ResetDenyDeleteTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileNameRegex", GoMethod: "ResetFileNameRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxFileSize", GoMethod: "ResetMaxFileSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemberCheck", GoMethod: "ResetMemberCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventSecrets", GoMethod: "ResetPreventSecrets"},
			_jsii_.MemberMethod{JsiiMethod: "resetRejectNonDcoCommits", GoMethod: "ResetRejectNonDcoCommits"},
			_jsii_.MemberMethod{JsiiMethod: "resetRejectUnsignedCommits", GoMethod: "ResetRejectUnsignedCommits"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ProjectPushRulesA{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-gitlab.projectPushRules.ProjectPushRulesAConfig",
		reflect.TypeOf((*ProjectPushRulesAConfig)(nil)).Elem(),
	)
}
