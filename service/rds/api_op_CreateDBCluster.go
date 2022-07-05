// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Aurora DB cluster or Multi-AZ DB cluster. You can use the
// ReplicationSourceIdentifier parameter to create an Amazon Aurora DB cluster as a
// read replica of another DB cluster or Amazon RDS MySQL or PostgreSQL DB
// instance. For cross-Region replication where the DB cluster identified by
// ReplicationSourceIdentifier is encrypted, also specify the PreSignedUrl
// parameter. For more information on Amazon Aurora, see  What is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see  Multi-AZ deployments with two readable standby DB instances
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) CreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*Options)) (*CreateDBClusterOutput, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBCluster", params, optFns, c.addOperationCreateDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBClusterInput struct {

	// The DB cluster identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	// * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	// * First
	// character must be a letter.
	//
	// * Can't end with a hyphen or contain two
	// consecutive hyphens.
	//
	// Example: my-cluster1 Valid for: Aurora DB clusters and
	// Multi-AZ DB clusters
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the database engine to be used for this DB cluster. Valid Values:
	//
	// *
	// aurora (for MySQL 5.6-compatible Aurora)
	//
	// * aurora-mysql (for MySQL
	// 5.7-compatible and MySQL 8.0-compatible Aurora)
	//
	// * aurora-postgresql
	//
	// * mysql
	//
	// *
	// postgres
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	//
	// This member is required.
	Engine *string

	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the
	// Multi-AZ DB cluster. This setting is required to create a Multi-AZ DB cluster.
	// Valid for: Multi-AZ DB clusters only
	AllocatedStorage *int32

	// A value that indicates whether minor engine upgrades are applied automatically
	// to the DB cluster during the maintenance window. By default, minor engine
	// upgrades are applied automatically. Valid for: Multi-AZ DB clusters only
	AutoMinorVersionUpgrade *bool

	// A list of Availability Zones (AZs) where DB instances in the DB cluster can be
	// created. For information on Amazon Web Services Regions and Availability Zones,
	// see Choosing the Regions and Availability Zones
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	AvailabilityZones []string

	// The target backtrack window, in seconds. To disable backtracking, set this value
	// to 0. Default: 0 Constraints:
	//
	// * If specified, this value must be set to a
	// number from 0 to 259,200 (72 hours).
	//
	// Valid for: Aurora MySQL DB clusters only
	BacktrackWindow *int64

	// The number of days for which automated backups are retained. Default: 1
	// Constraints:
	//
	// * Must be a value from 1 to 35
	//
	// Valid for: Aurora DB clusters and
	// Multi-AZ DB clusters
	BackupRetentionPeriod *int32

	// A value that indicates that the DB cluster should be associated with the
	// specified CharacterSet. Valid for: Aurora DB clusters only
	CharacterSetName *string

	// A value that indicates whether to copy all tags from the DB cluster to snapshots
	// of the DB cluster. The default is not to copy them. Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	CopyTagsToSnapshot *bool

	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster,
	// for example db.m6g.xlarge. Not all DB instance classes are available in all
	// Amazon Web Services Regions, or for all database engines. For the full list of
	// DB instance classes and availability for your engine, see DB instance class
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide. This setting is required to create a Multi-AZ DB
	// cluster. Valid for: Multi-AZ DB clusters only
	DBClusterInstanceClass *string

	// The name of the DB cluster parameter group to associate with this DB cluster. If
	// you do not specify a value, then the default DB cluster parameter group for the
	// specified DB engine and version is used. Constraints:
	//
	// * If supplied, must match
	// the name of an existing DB cluster parameter group.
	//
	// Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	DBClusterParameterGroupName *string

	// A DB subnet group to associate with this DB cluster. This setting is required to
	// create a Multi-AZ DB cluster. Constraints: Must match the name of an existing
	// DBSubnetGroup. Must not be default. Example: mydbsubnetgroup Valid for: Aurora
	// DB clusters and Multi-AZ DB clusters
	DBSubnetGroupName *string

	// The name for your database of up to 64 alphanumeric characters. If you do not
	// provide a name, Amazon RDS doesn't create a database in the DB cluster you are
	// creating. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	DatabaseName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection isn't enabled. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	DeletionProtection *bool

	// The Active Directory directory ID to create the DB cluster in. For Amazon Aurora
	// DB clusters, Amazon RDS can use Kerberos authentication to authenticate users
	// that connect to the DB cluster. For more information, see Kerberos
	// authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	Domain *string

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service. Valid for: Aurora DB clusters only
	DomainIAMRoleName *string

	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// The values in the list depend on the DB engine being used. RDS for MySQL
	// Possible values are error, general, and slowquery. RDS for PostgreSQL Possible
	// values are postgresql and upgrade. Aurora MySQL Possible values are audit,
	// error, general, and slowquery. Aurora PostgreSQL Possible value is postgresql.
	// For more information about exporting CloudWatch Logs for Amazon RDS, see
	// Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon RDS User Guide. For more information about exporting CloudWatch
	// Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	EnableCloudwatchLogsExports []string

	// A value that indicates whether to enable this DB cluster to forward write
	// operations to the primary cluster of an Aurora global database (GlobalCluster).
	// By default, write operations are not allowed on Aurora DB clusters that are
	// secondary clusters in an Aurora global database. You can set this value only on
	// Aurora DB clusters that are members of an Aurora global database. With this
	// parameter enabled, a secondary cluster can forward writes to the current primary
	// cluster and the resulting changes are replicated back to this cluster. For the
	// primary DB cluster of an Aurora global database, this value is used immediately
	// if the primary is demoted by the FailoverGlobalCluster API operation, but it
	// does nothing until then. Valid for: Aurora DB clusters only
	EnableGlobalWriteForwarding *bool

	// A value that indicates whether to enable the HTTP endpoint for an Aurora
	// Serverless v1 DB cluster. By default, the HTTP endpoint is disabled. When
	// enabled, the HTTP endpoint provides a connectionless web service API for running
	// SQL queries on the Aurora Serverless v1 DB cluster. You can also query your
	// database from inside the RDS console with the query editor. For more
	// information, see Using the Data API for Aurora Serverless v1
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in
	// the Amazon Aurora User Guide. Valid for: Aurora DB clusters only
	EnableHttpEndpoint *bool

	// A value that indicates whether to enable mapping of Amazon Web Services Identity
	// and Access Management (IAM) accounts to database accounts. By default, mapping
	// isn't enabled. For more information, see  IAM Database Authentication
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide.. Valid for: Aurora DB clusters only
	EnableIAMDatabaseAuthentication *bool

	// A value that indicates whether to turn on Performance Insights for the DB
	// cluster. For more information, see  Using Amazon Performance Insights
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon RDS User Guide. Valid for: Multi-AZ DB clusters only
	EnablePerformanceInsights *bool

	// The DB engine mode of the DB cluster, either provisioned, serverless,
	// parallelquery, global, or multimaster. The parallelquery engine mode isn't
	// required for Aurora MySQL version 1.23 and higher 1.x versions, and version 2.09
	// and higher 2.x versions. The global engine mode isn't required for Aurora MySQL
	// version 1.22 and higher 1.x versions, and global engine mode isn't required for
	// any 2.x versions. The multimaster engine mode only applies for DB clusters
	// created with Aurora MySQL version 5.6.10a. For Aurora PostgreSQL, the global
	// engine mode isn't required, and both the parallelquery and the multimaster
	// engine modes currently aren't supported. Limitations and requirements apply to
	// some DB engine modes. For more information, see the following sections in the
	// Amazon Aurora User Guide:
	//
	// * Limitations of Aurora Serverless v1
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)
	//
	// *
	// Limitations of Parallel Query
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)
	//
	// *
	// Limitations of Aurora Global Databases
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)
	//
	// *
	// Limitations of Multi-Master Clusters
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-multi-master.html#aurora-multi-master-limitations)
	//
	// Valid
	// for: Aurora DB clusters only
	EngineMode *string

	// The version number of the database engine to use. To list all of the available
	// engine versions for MySQL 5.6-compatible Aurora, use the following command: aws
	// rds describe-db-engine-versions --engine aurora --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora, use the following
	// command: aws rds describe-db-engine-versions --engine aurora-mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for Aurora PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine aurora-postgresql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for MySQL, use the following command: aws rds
	// describe-db-engine-versions --engine mysql --query
	// "DBEngineVersions[].EngineVersion" To list all of the available engine versions
	// for RDS for PostgreSQL, use the following command: aws rds
	// describe-db-engine-versions --engine postgres --query
	// "DBEngineVersions[].EngineVersion" Aurora MySQL For information, see MySQL on
	// Amazon RDS Versions
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html)
	// in the Amazon Aurora User Guide. Aurora PostgreSQL For information, see Amazon
	// Aurora PostgreSQL releases and engine versions
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html)
	// in the Amazon Aurora User Guide. MySQL For information, see MySQL on Amazon RDS
	// Versions
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt)
	// in the Amazon RDS User Guide. PostgreSQL For information, see Amazon RDS for
	// PostgreSQL versions and extensions
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts)
	// in the Amazon RDS User Guide. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	EngineVersion *string

	// The global cluster ID of an Aurora cluster that becomes the primary cluster in
	// the new global database cluster. Valid for: Aurora DB clusters only
	GlobalClusterIdentifier *string

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for each DB instance in the Multi-AZ DB cluster. For
	// information about valid Iops values, see Amazon RDS Provisioned IOPS storage to
	// improve performance
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide. This setting is required to create a Multi-AZ DB
	// cluster. Constraints: Must be a multiple between .5 and 50 of the storage amount
	// for the DB cluster. Valid for: Multi-AZ DB clusters only
	Iops *int32

	// The Amazon Web Services KMS key identifier for an encrypted DB cluster. The
	// Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or
	// alias name for the KMS key. To use a KMS key in a different Amazon Web Services
	// account, specify the key ARN or alias ARN. When a KMS key isn't specified in
	// KmsKeyId:
	//
	// * If ReplicationSourceIdentifier identifies an encrypted source, then
	// Amazon RDS will use the KMS key used to encrypt the source. Otherwise, Amazon
	// RDS will use your default KMS key.
	//
	// * If the StorageEncrypted parameter is
	// enabled and ReplicationSourceIdentifier isn't specified, then Amazon RDS will
	// use your default KMS key.
	//
	// There is a default KMS key for your Amazon Web
	// Services account. Your Amazon Web Services account has a different default KMS
	// key for each Amazon Web Services Region. If you create a read replica of an
	// encrypted DB cluster in another Amazon Web Services Region, you must set
	// KmsKeyId to a KMS key identifier that is valid in the destination Amazon Web
	// Services Region. This KMS key is used to encrypt the read replica in that Amazon
	// Web Services Region. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	KmsKeyId *string

	// The password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain
	// from 8 to 41 characters. Valid for: Aurora DB clusters and Multi-AZ DB clusters
	MasterUserPassword *string

	// The name of the master user for the DB cluster. Constraints:
	//
	// * Must be 1 to 16
	// letters or numbers.
	//
	// * First character must be a letter.
	//
	// * Can't be a reserved
	// word for the chosen database engine.
	//
	// Valid for: Aurora DB clusters and Multi-AZ
	// DB clusters
	MasterUsername *string

	// The interval, in seconds, between points when Enhanced Monitoring metrics are
	// collected for the DB cluster. To turn off collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0. If MonitoringRoleArn is specified, also
	// set MonitoringInterval to a value other than 0. Valid Values: 0, 1, 5, 10, 15,
	// 30, 60 Valid for: Multi-AZ DB clusters only
	MonitoringInterval *int32

	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send
	// Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is
	// arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring
	// role, see Setting up and enabling Enhanced Monitoring
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)
	// in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than
	// 0, supply a MonitoringRoleArn value. Valid for: Multi-AZ DB clusters only
	MonitoringRoleArn *string

	// A value that indicates that the DB cluster should be associated with the
	// specified option group. DB clusters are associated with a default option group
	// that can't be modified.
	OptionGroupName *string

	// The Amazon Web Services KMS key identifier for encryption of Performance
	// Insights data. The Amazon Web Services KMS key identifier is the key ARN, key
	// ID, alias ARN, or alias name for the KMS key. If you don't specify a value for
	// PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is
	// a default KMS key for your Amazon Web Services account. Your Amazon Web Services
	// account has a different default KMS key for each Amazon Web Services Region.
	// Valid for: Multi-AZ DB clusters only
	PerformanceInsightsKMSKeyId *string

	// The amount of time, in days, to retain Performance Insights data. Valid values
	// are 7 or 731 (2 years). Valid for: Multi-AZ DB clusters only
	PerformanceInsightsRetentionPeriod *int32

	// The port number on which the instances in the DB cluster accept connections. RDS
	// for MySQL and Aurora MySQL Default: 3306 Valid values: 1150-65535 RDS for
	// PostgreSQL and Aurora PostgreSQL Default: 5432 Valid values: 1150-65535 Valid
	// for: Aurora DB clusters and Multi-AZ DB clusters
	Port *int32

	// A URL that contains a Signature Version 4 signed request for the CreateDBCluster
	// action to be called in the source Amazon Web Services Region where the DB
	// cluster is replicated from. Specify PreSignedUrl only when you are performing
	// cross-Region replication from an encrypted DB cluster. The pre-signed URL must
	// be a valid request for the CreateDBCluster API action that can be executed in
	// the source Amazon Web Services Region that contains the encrypted DB cluster to
	// be copied. The pre-signed URL request must contain the following parameter
	// values:
	//
	// * KmsKeyId - The Amazon Web Services KMS key identifier for the KMS key
	// to use to encrypt the copy of the DB cluster in the destination Amazon Web
	// Services Region. This should refer to the same KMS key for both the
	// CreateDBCluster action that is called in the destination Amazon Web Services
	// Region, and the action contained in the pre-signed URL.
	//
	// * DestinationRegion -
	// The name of the Amazon Web Services Region that Aurora read replica will be
	// created in.
	//
	// * ReplicationSourceIdentifier - The DB cluster identifier for the
	// encrypted DB cluster to be copied. This identifier must be in the Amazon
	// Resource Name (ARN) format for the source Amazon Web Services Region. For
	// example, if you are copying an encrypted DB cluster from the us-west-2 Amazon
	// Web Services Region, then your ReplicationSourceIdentifier would look like
	// Example: arn:aws:rds:us-west-2:123456789012:cluster:aurora-cluster1.
	//
	// To learn
	// how to generate a Signature Version 4 signed request, see  Authenticating
	// Requests: Using Query Parameters (Amazon Web Services Signature Version 4)
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// and  Signature Version 4 Signing Process
	// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). If you
	// are using an Amazon Web Services SDK tool or the CLI, you can specify
	// SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl
	// manually. Specifying SourceRegion autogenerates a pre-signed URL that is a valid
	// request for the operation that can be executed in the source Amazon Web Services
	// Region. Valid for: Aurora DB clusters only
	PreSignedUrl *string

	// The daily time range during which automated backups are created if automated
	// backups are enabled using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region. To view the time blocks available, see  Backup window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow)
	// in the Amazon Aurora User Guide. Constraints:
	//
	// * Must be in the format
	// hh24:mi-hh24:mi.
	//
	// * Must be in Universal Coordinated Time (UTC).
	//
	// * Must not
	// conflict with the preferred maintenance window.
	//
	// * Must be at least 30
	// minutes.
	//
	// Valid for: Aurora DB clusters and Multi-AZ DB clusters
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region, occurring on a random day of the week. To see the time
	// blocks available, see  Adjusting the Preferred DB Cluster Maintenance Window
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora)
	// in the Amazon Aurora User Guide. Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	// Constraints: Minimum 30-minute window. Valid for: Aurora DB clusters and
	// Multi-AZ DB clusters
	PreferredMaintenanceWindow *string

	// A value that indicates whether the DB cluster is publicly accessible. When the
	// DB cluster is publicly accessible, its Domain Name System (DNS) endpoint
	// resolves to the private IP address from within the DB cluster's virtual private
	// cloud (VPC). It resolves to the public IP address from outside of the DB
	// cluster's VPC. Access to the DB cluster is ultimately controlled by the security
	// group it uses. That public access isn't permitted if the security group assigned
	// to the DB cluster doesn't permit it. When the DB cluster isn't publicly
	// accessible, it is an internal DB cluster with a DNS name that resolves to a
	// private IP address. Default: The default behavior varies depending on whether
	// DBSubnetGroupName is specified. If DBSubnetGroupName isn't specified, and
	// PubliclyAccessible isn't specified, the following applies:
	//
	// * If the default VPC
	// in the target Region doesn’t have an internet gateway attached to it, the DB
	// cluster is private.
	//
	// * If the default VPC in the target Region has an internet
	// gateway attached to it, the DB cluster is public.
	//
	// If DBSubnetGroupName is
	// specified, and PubliclyAccessible isn't specified, the following applies:
	//
	// * If
	// the subnets are part of a VPC that doesn’t have an internet gateway attached to
	// it, the DB cluster is private.
	//
	// * If the subnets are part of a VPC that has an
	// internet gateway attached to it, the DB cluster is public.
	//
	// Valid for: Multi-AZ
	// DB clusters only
	PubliclyAccessible *bool

	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this
	// DB cluster is created as a read replica. Valid for: Aurora DB clusters only
	ReplicationSourceIdentifier *string

	// For DB clusters in serverless DB engine mode, the scaling properties of the DB
	// cluster. Valid for: Aurora DB clusters only
	ScalingConfiguration *types.ScalingConfiguration

	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster. For
	// more information, see Using Amazon Aurora Serverless v2
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html)
	// in the Amazon Aurora User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string

	// A value that indicates whether the DB cluster is encrypted. Valid for: Aurora DB
	// clusters and Multi-AZ DB clusters
	StorageEncrypted *bool

	// Specifies the storage type to be associated with the DB cluster. This setting is
	// required to create a Multi-AZ DB cluster. Valid values: io1 When specified, a
	// value for the Iops parameter is required. Default: io1 Valid for: Multi-AZ DB
	// clusters only
	StorageType *string

	// Tags to assign to the DB cluster. Valid for: Aurora DB clusters and Multi-AZ DB
	// clusters
	Tags []types.Tag

	// A list of EC2 VPC security groups to associate with this DB cluster. Valid for:
	// Aurora DB clusters and Multi-AZ DB clusters
	VpcSecurityGroupIds []string

	// Used by the SDK's PresignURL autofill customization to specify the region the of
	// the client's request.
	destinationRegion *string

	noSmithyDocumentSerde
}

type CreateDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters,
	// FailoverDBCluster, ModifyDBCluster, PromoteReadReplicaDBCluster,
	// RestoreDBClusterFromS3, RestoreDBClusterFromSnapshot,
	// RestoreDBClusterToPointInTime, StartDBCluster, and StopDBCluster. For a Multi-AZ
	// DB cluster, this data type is used as a response element in the operations
	// CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster,
	// ModifyDBCluster, RebootDBCluster, RestoreDBClusterFromSnapshot, and
	// RestoreDBClusterToPointInTime. For more information on Amazon Aurora DB
	// clusters, see  What is Amazon Aurora?
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see  Multi-AZ deployments with two readable standby DB instances
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateDBClusterPresignURLMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBCluster(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func copyCreateDBClusterInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCreateDBClusterPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCreateDBClusterSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCreateDBClusterPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCreateDBClusterdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCreateDBClusterPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCreateDBClusterPreSignedUrl,

			GetSourceRegion: getCreateDBClusterSourceRegion,

			CopyInput: copyCreateDBClusterInputForPresign,

			SetDestinationRegion: setCreateDBClusterdestinationRegion,

			SetPresignedURL: setCreateDBClusterPreSignedUrl,
		},
		Presigner: &presignAutoFillCreateDBClusterClient{client: NewPresignClient(New(options))},
	})
}

type presignAutoFillCreateDBClusterClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCreateDBClusterClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CreateDBClusterInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateDBClusterInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCreateDBCluster(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCreateDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBCluster",
	}
}

// PresignCreateDBCluster is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCreateDBCluster(ctx context.Context, params *CreateDBClusterInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil {
		params = &CreateDBClusterInput{}
	}
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)

	result, _, err := c.client.invokeOperation(ctx, "CreateDBCluster", params, clientOptFns,
		c.client.addOperationCreateDBClusterMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil {
		return nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}
