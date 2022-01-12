package aws

import "github.com/snyk/driftctl/pkg/resource"

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initAwsAmiMetaData(resourceSchemaRepository)
	initAwsCloudfrontDistributionMetaData(resourceSchemaRepository)
	initAwsDbInstanceMetaData(resourceSchemaRepository)
	initAwsDbSubnetGroupMetaData(resourceSchemaRepository)
	initAwsDefaultSecurityGroupMetaData(resourceSchemaRepository)
	initAwsDefaultSubnetMetaData(resourceSchemaRepository)
	initAwsDefaultVpcMetaData(resourceSchemaRepository)
	initAwsDefaultRouteTableMetadata(resourceSchemaRepository)
	initAwsDynamodbTableMetaData(resourceSchemaRepository)
	initAwsEbsSnapshotMetaData(resourceSchemaRepository)
	initAwsInstanceMetaData(resourceSchemaRepository)
	initAwsInternetGatewayMetaData(resourceSchemaRepository)
	initAwsEbsVolumeMetaData(resourceSchemaRepository)
	initAwsEipMetaData(resourceSchemaRepository)
	initAwsEipAssociationMetaData(resourceSchemaRepository)
	initAwsS3BucketMetaData(resourceSchemaRepository)
	initAwsS3BucketPolicyMetaData(resourceSchemaRepository)
	initAwsS3BucketInventoryMetadata(resourceSchemaRepository)
	initAwsS3BucketMetricMetadata(resourceSchemaRepository)
	initAwsS3BucketNotificationMetadata(resourceSchemaRepository)
	initAwsS3BucketAnalyticsConfigurationMetaData(resourceSchemaRepository)
	initAwsEcrRepositoryMetaData(resourceSchemaRepository)
	initAwsRouteMetaData(resourceSchemaRepository)
	initAwsRouteTableAssociationMetaData(resourceSchemaRepository)
	initAwsRoute53RecordMetaData(resourceSchemaRepository)
	initAwsRoute53ZoneMetaData(resourceSchemaRepository)
	initAwsRoute53HealthCheckMetaData(resourceSchemaRepository)
	initAwsRouteTableMetaData(resourceSchemaRepository)
	initSnsTopicSubscriptionMetaData(resourceSchemaRepository)
	initSnsTopicPolicyMetaData(resourceSchemaRepository)
	initSnsTopicMetaData(resourceSchemaRepository)
	initSqsQueueMetaData(resourceSchemaRepository)
	initAwsIAMAccessKeyMetaData(resourceSchemaRepository)
	initAwsIAMPolicyMetaData(resourceSchemaRepository)
	initAwsIAMPolicyAttachmentMetaData(resourceSchemaRepository)
	initAwsIAMRoleMetaData(resourceSchemaRepository)
	initAwsIAMRolePolicyMetaData(resourceSchemaRepository)
	initAwsIamRolePolicyAttachmentMetaData(resourceSchemaRepository)
	initAwsIamUserPolicyAttachmentMetaData(resourceSchemaRepository)
	initAwsIAMUserMetaData(resourceSchemaRepository)
	initAwsIAMUserPolicyMetaData(resourceSchemaRepository)
	initAwsKeyPairMetaData(resourceSchemaRepository)
	initAwsKmsKeyMetaData(resourceSchemaRepository)
	initAwsKmsAliasMetaData(resourceSchemaRepository)
	initAwsLambdaFunctionMetaData(resourceSchemaRepository)
	initAwsLambdaEventSourceMappingMetaData(resourceSchemaRepository)
	initNatGatewayMetaData(resourceSchemaRepository)
	initAwsNetworkACLMetaData(resourceSchemaRepository)
	initAwsNetworkACLRuleMetaData(resourceSchemaRepository)
	initAwsDefaultNetworkACLMetaData(resourceSchemaRepository)
	initAwsSubnetMetaData(resourceSchemaRepository)
	initAwsSQSQueuePolicyMetaData(resourceSchemaRepository)
	initAwsSecurityGroupRuleMetaData(resourceSchemaRepository)
	initAwsSecurityGroupMetaData(resourceSchemaRepository)
	initAwsRDSClusterMetaData(resourceSchemaRepository)
	initAwsCloudformationStackMetaData(resourceSchemaRepository)
	initAwsVpcMetaData(resourceSchemaRepository)
	initAwsAppAutoscalingTargetMetaData(resourceSchemaRepository)
	initAwsAppAutoscalingPolicyMetaData(resourceSchemaRepository)
	initAwsLaunchTemplateMetaData(resourceSchemaRepository)
	initAwsApiGatewayV2ModelMetaData(resourceSchemaRepository)
}
