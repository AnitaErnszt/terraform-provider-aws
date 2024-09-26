// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../../generate/tagresource/main.go -UpdateTagsFunc=updateTagsForResource
//go:generate go run ../../generate/tags/main.go -GetTag -ListTags -ListTagsOp=ListTagsOfResource -ServiceTagsSlice -UpdateTags -Wait -WaitContinuousOccurence 2 -WaitMinTimeout 1s -WaitTimeout 2m -UpdateTagsForResource -ParentNotFoundErrCode=ResourceNotFoundException
//go:generate go run ../../generate/servicepackage/main.go
//go:generate go run ../../generate/listpages/main.go -ListOps=ListBackups -InputPaginator=ExclusiveStartBackupArn -OutputPaginator=LastEvaluatedBackupArn -- list_backups_pages_gen.go
//go:generate go run ../../generate/tagstests/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package dynamodb
