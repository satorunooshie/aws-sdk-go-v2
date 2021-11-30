// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpChangeServerLifeCycleState struct {
}

func (*validateOpChangeServerLifeCycleState) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpChangeServerLifeCycleState) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ChangeServerLifeCycleStateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpChangeServerLifeCycleStateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateReplicationConfigurationTemplate struct {
}

func (*validateOpCreateReplicationConfigurationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateReplicationConfigurationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateReplicationConfigurationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateReplicationConfigurationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteJob struct {
}

func (*validateOpDeleteJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteReplicationConfigurationTemplate struct {
}

func (*validateOpDeleteReplicationConfigurationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteReplicationConfigurationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteReplicationConfigurationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteReplicationConfigurationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSourceServer struct {
}

func (*validateOpDeleteSourceServer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSourceServer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSourceServerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSourceServerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteVcenterClient struct {
}

func (*validateOpDeleteVcenterClient) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteVcenterClient) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteVcenterClientInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteVcenterClientInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJobLogItems struct {
}

func (*validateOpDescribeJobLogItems) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJobLogItems) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJobLogItemsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJobLogItemsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJobs struct {
}

func (*validateOpDescribeJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeReplicationConfigurationTemplates struct {
}

func (*validateOpDescribeReplicationConfigurationTemplates) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeReplicationConfigurationTemplates) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeReplicationConfigurationTemplatesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeReplicationConfigurationTemplatesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSourceServers struct {
}

func (*validateOpDescribeSourceServers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSourceServers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSourceServersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSourceServersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisconnectFromService struct {
}

func (*validateOpDisconnectFromService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisconnectFromService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisconnectFromServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisconnectFromServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpFinalizeCutover struct {
}

func (*validateOpFinalizeCutover) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpFinalizeCutover) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*FinalizeCutoverInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpFinalizeCutoverInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLaunchConfiguration struct {
}

func (*validateOpGetLaunchConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLaunchConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLaunchConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLaunchConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetReplicationConfiguration struct {
}

func (*validateOpGetReplicationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetReplicationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetReplicationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetReplicationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMarkAsArchived struct {
}

func (*validateOpMarkAsArchived) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMarkAsArchived) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MarkAsArchivedInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMarkAsArchivedInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRetryDataReplication struct {
}

func (*validateOpRetryDataReplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRetryDataReplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RetryDataReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRetryDataReplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCutover struct {
}

func (*validateOpStartCutover) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCutover) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCutoverInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCutoverInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartReplication struct {
}

func (*validateOpStartReplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartReplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartReplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTest struct {
}

func (*validateOpStartTest) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTestInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTestInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTerminateTargetInstances struct {
}

func (*validateOpTerminateTargetInstances) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTerminateTargetInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TerminateTargetInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTerminateTargetInstancesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateLaunchConfiguration struct {
}

func (*validateOpUpdateLaunchConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLaunchConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLaunchConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLaunchConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateReplicationConfiguration struct {
}

func (*validateOpUpdateReplicationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateReplicationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateReplicationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateReplicationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateReplicationConfigurationTemplate struct {
}

func (*validateOpUpdateReplicationConfigurationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateReplicationConfigurationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateReplicationConfigurationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateReplicationConfigurationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSourceServerReplicationType struct {
}

func (*validateOpUpdateSourceServerReplicationType) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSourceServerReplicationType) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSourceServerReplicationTypeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSourceServerReplicationTypeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpChangeServerLifeCycleStateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpChangeServerLifeCycleState{}, middleware.After)
}

func addOpCreateReplicationConfigurationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateReplicationConfigurationTemplate{}, middleware.After)
}

func addOpDeleteJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteJob{}, middleware.After)
}

func addOpDeleteReplicationConfigurationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteReplicationConfigurationTemplate{}, middleware.After)
}

func addOpDeleteSourceServerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSourceServer{}, middleware.After)
}

func addOpDeleteVcenterClientValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteVcenterClient{}, middleware.After)
}

func addOpDescribeJobLogItemsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJobLogItems{}, middleware.After)
}

func addOpDescribeJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJobs{}, middleware.After)
}

func addOpDescribeReplicationConfigurationTemplatesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeReplicationConfigurationTemplates{}, middleware.After)
}

func addOpDescribeSourceServersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSourceServers{}, middleware.After)
}

func addOpDisconnectFromServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisconnectFromService{}, middleware.After)
}

func addOpFinalizeCutoverValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpFinalizeCutover{}, middleware.After)
}

func addOpGetLaunchConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLaunchConfiguration{}, middleware.After)
}

func addOpGetReplicationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetReplicationConfiguration{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpMarkAsArchivedValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMarkAsArchived{}, middleware.After)
}

func addOpRetryDataReplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRetryDataReplication{}, middleware.After)
}

func addOpStartCutoverValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCutover{}, middleware.After)
}

func addOpStartReplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartReplication{}, middleware.After)
}

func addOpStartTestValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTest{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpTerminateTargetInstancesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTerminateTargetInstances{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateLaunchConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLaunchConfiguration{}, middleware.After)
}

func addOpUpdateReplicationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateReplicationConfiguration{}, middleware.After)
}

func addOpUpdateReplicationConfigurationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateReplicationConfigurationTemplate{}, middleware.After)
}

func addOpUpdateSourceServerReplicationTypeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSourceServerReplicationType{}, middleware.After)
}

func validateChangeServerLifeCycleStateSourceServerLifecycle(v *types.ChangeServerLifeCycleStateSourceServerLifecycle) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeServerLifeCycleStateSourceServerLifecycle"}
	if len(v.State) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpChangeServerLifeCycleStateInput(v *ChangeServerLifeCycleStateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChangeServerLifeCycleStateInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if v.LifeCycle == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LifeCycle"))
	} else if v.LifeCycle != nil {
		if err := validateChangeServerLifeCycleStateSourceServerLifecycle(v.LifeCycle); err != nil {
			invalidParams.AddNested("LifeCycle", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateReplicationConfigurationTemplateInput(v *CreateReplicationConfigurationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateReplicationConfigurationTemplateInput"}
	if v.StagingAreaSubnetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StagingAreaSubnetId"))
	}
	if v.AssociateDefaultSecurityGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AssociateDefaultSecurityGroup"))
	}
	if v.ReplicationServersSecurityGroupsIDs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationServersSecurityGroupsIDs"))
	}
	if v.ReplicationServerInstanceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationServerInstanceType"))
	}
	if v.UseDedicatedReplicationServer == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UseDedicatedReplicationServer"))
	}
	if len(v.DefaultLargeStagingDiskType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DefaultLargeStagingDiskType"))
	}
	if len(v.EbsEncryption) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EbsEncryption"))
	}
	if len(v.DataPlaneRouting) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DataPlaneRouting"))
	}
	if v.CreatePublicIP == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CreatePublicIP"))
	}
	if v.StagingAreaTags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StagingAreaTags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteJobInput(v *DeleteJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteJobInput"}
	if v.JobID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteReplicationConfigurationTemplateInput(v *DeleteReplicationConfigurationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteReplicationConfigurationTemplateInput"}
	if v.ReplicationConfigurationTemplateID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationConfigurationTemplateID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSourceServerInput(v *DeleteSourceServerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSourceServerInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteVcenterClientInput(v *DeleteVcenterClientInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteVcenterClientInput"}
	if v.VcenterClientID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VcenterClientID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeJobLogItemsInput(v *DescribeJobLogItemsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJobLogItemsInput"}
	if v.JobID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeJobsInput(v *DescribeJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJobsInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeReplicationConfigurationTemplatesInput(v *DescribeReplicationConfigurationTemplatesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeReplicationConfigurationTemplatesInput"}
	if v.ReplicationConfigurationTemplateIDs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationConfigurationTemplateIDs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSourceServersInput(v *DescribeSourceServersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSourceServersInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisconnectFromServiceInput(v *DisconnectFromServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisconnectFromServiceInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpFinalizeCutoverInput(v *FinalizeCutoverInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FinalizeCutoverInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLaunchConfigurationInput(v *GetLaunchConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLaunchConfigurationInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetReplicationConfigurationInput(v *GetReplicationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetReplicationConfigurationInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMarkAsArchivedInput(v *MarkAsArchivedInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MarkAsArchivedInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRetryDataReplicationInput(v *RetryDataReplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RetryDataReplicationInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartCutoverInput(v *StartCutoverInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCutoverInput"}
	if v.SourceServerIDs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerIDs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartReplicationInput(v *StartReplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartReplicationInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartTestInput(v *StartTestInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTestInput"}
	if v.SourceServerIDs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerIDs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTerminateTargetInstancesInput(v *TerminateTargetInstancesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TerminateTargetInstancesInput"}
	if v.SourceServerIDs == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerIDs"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateLaunchConfigurationInput(v *UpdateLaunchConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLaunchConfigurationInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateReplicationConfigurationInput(v *UpdateReplicationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateReplicationConfigurationInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateReplicationConfigurationTemplateInput(v *UpdateReplicationConfigurationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateReplicationConfigurationTemplateInput"}
	if v.ReplicationConfigurationTemplateID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationConfigurationTemplateID"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSourceServerReplicationTypeInput(v *UpdateSourceServerReplicationTypeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSourceServerReplicationTypeInput"}
	if v.SourceServerID == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceServerID"))
	}
	if len(v.ReplicationType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ReplicationType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
