// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about the batch policy.
type BatchPolicy struct {

	// The number of active simulation jobs create as part of the batch that can be in
	// an active state at the same time. Active states include: Pending,Preparing,
	// Running, Restarting, RunningFailed and Terminating. All other states are
	// terminal states.
	MaxConcurrency *int32

	// The amount of time, in seconds, to wait for the batch to complete. If a batch
	// times out, and there are pending requests that were failing due to an internal
	// failure (like InternalServiceError), they will be moved to the failed list and
	// the batch status will be Failed. If the pending requests were failing for any
	// other reason, the failed pending requests will be moved to the failed list and
	// the batch status will be TimedOut.
	TimeoutInSeconds *int64

	noSmithyDocumentSerde
}

// Compute information for the simulation job.
type Compute struct {

	// Compute type information for the simulation job.
	ComputeType ComputeType

	// Compute GPU unit limit for the simulation job. It is the same as the number of
	// GPUs allocated to the SimulationJob.
	GpuUnitLimit *int32

	// The simulation unit limit. Your simulation is allocated CPU and memory
	// proportional to the supplied simulation unit limit. A simulation unit is 1 vcpu
	// and 2GB of memory. You are only billed for the SU utilization you consume up to
	// the maximum value provided. The default is 15.
	SimulationUnitLimit *int32

	noSmithyDocumentSerde
}

// Compute information for the simulation job
type ComputeResponse struct {

	// Compute type response information for the simulation job.
	ComputeType ComputeType

	// Compute GPU unit limit for the simulation job. It is the same as the number of
	// GPUs allocated to the SimulationJob.
	GpuUnitLimit *int32

	// The simulation unit limit. Your simulation is allocated CPU and memory
	// proportional to the supplied simulation unit limit. A simulation unit is 1 vcpu
	// and 2GB of memory. You are only billed for the SU utilization you consume up to
	// the maximum value provided. The default is 15.
	SimulationUnitLimit *int32

	noSmithyDocumentSerde
}

// Information about a data source.
type DataSource struct {

	// The location where your files are mounted in the container image. If you've
	// specified the type of the data source as an Archive, you must provide an Amazon
	// S3 object key to your archive. The object key must point to either a .zip or
	// .tar.gz file. If you've specified the type of the data source as a Prefix, you
	// provide the Amazon S3 prefix that points to the files that you are using for
	// your data source. If you've specified the type of the data source as a File, you
	// provide the Amazon S3 path to the file that you're using as your data source.
	Destination *string

	// The name of the data source.
	Name *string

	// The S3 bucket where the data files are located.
	S3Bucket *string

	// The list of S3 keys identifying the data source files.
	S3Keys []S3KeyOutput

	// The data type for the data source that you're using for your container image or
	// simulation job. You can use this field to specify whether your data source is an
	// Archive, an Amazon S3 prefix, or a file. If you don't specify a field, the
	// default value is File.
	Type DataSourceType

	noSmithyDocumentSerde
}

// Information about a data source.
type DataSourceConfig struct {

	// The name of the data source.
	//
	// This member is required.
	Name *string

	// The S3 bucket where the data files are located.
	//
	// This member is required.
	S3Bucket *string

	// The list of S3 keys identifying the data source files.
	//
	// This member is required.
	S3Keys []string

	// The location where your files are mounted in the container image. If you've
	// specified the type of the data source as an Archive, you must provide an Amazon
	// S3 object key to your archive. The object key must point to either a .zip or
	// .tar.gz file. If you've specified the type of the data source as a Prefix, you
	// provide the Amazon S3 prefix that points to the files that you are using for
	// your data source. If you've specified the type of the data source as a File, you
	// provide the Amazon S3 path to the file that you're using as your data source.
	Destination *string

	// The data type for the data source that you're using for your container image or
	// simulation job. You can use this field to specify whether your data source is an
	// Archive, an Amazon S3 prefix, or a file. If you don't specify a field, the
	// default value is File.
	Type DataSourceType

	noSmithyDocumentSerde
}

// Information about a deployment application configuration.
type DeploymentApplicationConfig struct {

	// The Amazon Resource Name (ARN) of the robot application.
	//
	// This member is required.
	Application *string

	// The version of the application.
	//
	// This member is required.
	ApplicationVersion *string

	// The launch configuration.
	//
	// This member is required.
	LaunchConfig *DeploymentLaunchConfig

	noSmithyDocumentSerde
}

// Information about a deployment configuration.
type DeploymentConfig struct {

	// The percentage of robots receiving the deployment at the same time.
	ConcurrentDeploymentPercentage *int32

	// The download condition file.
	DownloadConditionFile *S3Object

	// The percentage of deployments that need to fail before stopping deployment.
	FailureThresholdPercentage *int32

	// The amount of time, in seconds, to wait for deployment to a single robot to
	// complete. Choose a time between 1 minute and 7 days. The default is 5 hours.
	RobotDeploymentTimeoutInSeconds *int64

	noSmithyDocumentSerde
}

// Information about a deployment job.
type DeploymentJob struct {

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string

	// The time, in milliseconds since the epoch, when the deployment job was created.
	CreatedAt *time.Time

	// The deployment application configuration.
	DeploymentApplicationConfigs []DeploymentApplicationConfig

	// The deployment configuration.
	DeploymentConfig *DeploymentConfig

	// The deployment job failure code.
	FailureCode DeploymentJobErrorCode

	// A short description of the reason why the deployment job failed.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the fleet.
	Fleet *string

	// The status of the deployment job.
	Status DeploymentStatus

	noSmithyDocumentSerde
}

// Configuration information for a deployment launch.
type DeploymentLaunchConfig struct {

	// The launch file name.
	//
	// This member is required.
	LaunchFile *string

	// The package name.
	//
	// This member is required.
	PackageName *string

	// An array of key/value pairs specifying environment variables for the robot
	// application
	EnvironmentVariables map[string]string

	// The deployment post-launch file. This file will be executed after the launch
	// file.
	PostLaunchFile *string

	// The deployment pre-launch file. This file will be executed prior to the launch
	// file.
	PreLaunchFile *string

	noSmithyDocumentSerde
}

// The object that contains the Docker image URI for either your robot or
// simulation applications.
type Environment struct {

	// The Docker image URI for either your robot or simulation applications.
	Uri *string

	noSmithyDocumentSerde
}

// Information about a failed create simulation job request.
type FailedCreateSimulationJobRequest struct {

	// The time, in milliseconds since the epoch, when the simulation job batch failed.
	FailedAt *time.Time

	// The failure code.
	FailureCode SimulationJobErrorCode

	// The failure reason of the simulation job request.
	FailureReason *string

	// The simulation job request.
	Request *SimulationJobRequest

	noSmithyDocumentSerde
}

// Information about worlds that failed.
type FailureSummary struct {

	// The worlds that failed.
	Failures []WorldFailure

	// The total number of failures.
	TotalFailureCount int32

	noSmithyDocumentSerde
}

// Information about a filter.
type Filter struct {

	// The name of the filter.
	Name *string

	// A list of values.
	Values []string

	noSmithyDocumentSerde
}

// Information about worlds that finished.
type FinishedWorldsSummary struct {

	// Information about worlds that failed.
	FailureSummary *FailureSummary

	// The total number of finished worlds.
	FinishedCount int32

	// A list of worlds that succeeded.
	SucceededWorlds []string

	noSmithyDocumentSerde
}

// Information about a fleet.
type Fleet struct {

	// The Amazon Resource Name (ARN) of the fleet.
	Arn *string

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string

	// The status of the last fleet deployment.
	LastDeploymentStatus DeploymentStatus

	// The time of the last deployment.
	LastDeploymentTime *time.Time

	// The name of the fleet.
	Name *string

	noSmithyDocumentSerde
}

// Information about a launch configuration.
type LaunchConfig struct {

	// If you've specified General as the value for your RobotSoftwareSuite, you can
	// use this field to specify a list of commands for your container image. If you've
	// specified SimulationRuntime as the value for your SimulationSoftwareSuite, you
	// can use this field to specify a list of commands for your container image.
	Command []string

	// The environment variables for the application launch.
	EnvironmentVariables map[string]string

	// The launch file name.
	LaunchFile *string

	// The package name.
	PackageName *string

	// The port forwarding configuration.
	PortForwardingConfig *PortForwardingConfig

	// Boolean indicating whether a streaming session will be configured for the
	// application. If True, AWS RoboMaker will configure a connection so you can
	// interact with your application as it is running in the simulation. You must
	// configure and launch the component. It must have a graphical user interface.
	StreamUI bool

	noSmithyDocumentSerde
}

// The logging configuration.
type LoggingConfig struct {

	// A boolean indicating whether to record all ROS topics. This API is no longer
	// supported and will throw an error if used.
	//
	// Deprecated: AWS RoboMaker is ending support for ROS software suite. For
	// additional information, see
	// https://docs.aws.amazon.com/robomaker/latest/dg/software-support-policy.html.
	RecordAllRosTopics *bool

	noSmithyDocumentSerde
}

// Describes a network interface.
type NetworkInterface struct {

	// The ID of the network interface.
	NetworkInterfaceId *string

	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress *string

	// The IPv4 public address of the network interface.
	PublicIpAddress *string

	noSmithyDocumentSerde
}

// The output location.
type OutputLocation struct {

	// The S3 bucket for output.
	S3Bucket *string

	// The S3 folder in the s3Bucket where output files will be placed.
	S3Prefix *string

	noSmithyDocumentSerde
}

// Configuration information for port forwarding.
type PortForwardingConfig struct {

	// The port mappings for the configuration.
	PortMappings []PortMapping

	noSmithyDocumentSerde
}

// An object representing a port mapping.
type PortMapping struct {

	// The port number on the application.
	//
	// This member is required.
	ApplicationPort int32

	// The port number on the simulation job instance to use as a remote connection
	// point.
	//
	// This member is required.
	JobPort int32

	// A Boolean indicating whether to enable this port mapping on public IP.
	EnableOnPublicIp bool

	noSmithyDocumentSerde
}

// Information about the progress of a deployment job.
type ProgressDetail struct {

	// The current progress status. Validating Validating the deployment.
	// DownloadingExtracting Downloading and extracting the bundle on the robot.
	// ExecutingPreLaunch Executing pre-launch script(s) if provided. Launching
	// Launching the robot application. ExecutingPostLaunch Executing post-launch
	// script(s) if provided. Finished Deployment is complete.
	CurrentProgress RobotDeploymentStep

	// Estimated amount of time in seconds remaining in the step. This currently only
	// applies to the Downloading/Extracting step of the deployment. It is empty for
	// other steps.
	EstimatedTimeRemainingSeconds *int32

	// Precentage of the step that is done. This currently only applies to the
	// Downloading/Extracting step of the deployment. It is empty for other steps.
	PercentDone *float32

	// The Amazon Resource Name (ARN) of the deployment job.
	TargetResource *string

	noSmithyDocumentSerde
}

// Information about a rendering engine.
type RenderingEngine struct {

	// The name of the rendering engine.
	Name RenderingEngineType

	// The version of the rendering engine.
	Version *string

	noSmithyDocumentSerde
}

// Information about a robot.
type Robot struct {

	// The architecture of the robot.
	Architecture Architecture

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the fleet.
	FleetArn *string

	// The Greengrass group associated with the robot.
	GreenGrassGroupId *string

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string

	// The time of the last deployment.
	LastDeploymentTime *time.Time

	// The name of the robot.
	Name *string

	// The status of the robot.
	Status RobotStatus

	noSmithyDocumentSerde
}

// Application configuration information for a robot.
type RobotApplicationConfig struct {

	// The application information for the robot application.
	//
	// This member is required.
	Application *string

	// The launch configuration for the robot application.
	//
	// This member is required.
	LaunchConfig *LaunchConfig

	// The version of the robot application.
	ApplicationVersion *string

	// Information about tools configured for the robot application.
	Tools []Tool

	// The upload configurations for the robot application.
	UploadConfigurations []UploadConfiguration

	// A Boolean indicating whether to use default robot application tools. The default
	// tools are rviz, rqt, terminal and rosbag record. The default is False. This API
	// is no longer supported and will throw an error if used.
	//
	// Deprecated: AWS RoboMaker is ending support for ROS software suite. For
	// additional information, see
	// https://docs.aws.amazon.com/robomaker/latest/dg/software-support-policy.html.
	UseDefaultTools *bool

	// A Boolean indicating whether to use default upload configurations. By default,
	// .ros and .gazebo files are uploaded when the application terminates and all ROS
	// topics will be recorded. If you set this value, you must specify an
	// outputLocation. This API is no longer supported and will throw an error if used.
	//
	// Deprecated: AWS RoboMaker is ending support for ROS software suite. For
	// additional information, see
	// https://docs.aws.amazon.com/robomaker/latest/dg/software-support-policy.html.
	UseDefaultUploadConfigurations *bool

	noSmithyDocumentSerde
}

// Summary information for a robot application.
type RobotApplicationSummary struct {

	// The Amazon Resource Name (ARN) of the robot.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot application was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the robot application.
	Name *string

	// Information about a robot software suite (ROS distribution).
	RobotSoftwareSuite *RobotSoftwareSuite

	// The version of the robot application.
	Version *string

	noSmithyDocumentSerde
}

// Information about a robot deployment.
type RobotDeployment struct {

	// The robot deployment Amazon Resource Name (ARN).
	Arn *string

	// The time, in milliseconds since the epoch, when the deployment finished.
	DeploymentFinishTime *time.Time

	// The time, in milliseconds since the epoch, when the deployment was started.
	DeploymentStartTime *time.Time

	// The robot deployment failure code.
	FailureCode DeploymentJobErrorCode

	// A short description of the reason why the robot deployment failed.
	FailureReason *string

	// Information about how the deployment is progressing.
	ProgressDetail *ProgressDetail

	// The status of the robot deployment.
	Status RobotStatus

	noSmithyDocumentSerde
}

// Information about a robot software suite (ROS distribution).
type RobotSoftwareSuite struct {

	// The name of the robot software suite (ROS distribution).
	Name RobotSoftwareSuiteType

	// The version of the robot software suite (ROS distribution).
	Version RobotSoftwareSuiteVersionType

	noSmithyDocumentSerde
}

// Information about S3 keys.
type S3KeyOutput struct {

	// The etag for the object.
	Etag *string

	// The S3 key.
	S3Key *string

	noSmithyDocumentSerde
}

// Information about an S3 object.
type S3Object struct {

	// The bucket containing the object.
	//
	// This member is required.
	Bucket *string

	// The key of the object.
	//
	// This member is required.
	Key *string

	// The etag of the object.
	Etag *string

	noSmithyDocumentSerde
}

// Information about a simulation application configuration.
type SimulationApplicationConfig struct {

	// The application information for the simulation application.
	//
	// This member is required.
	Application *string

	// The launch configuration for the simulation application.
	//
	// This member is required.
	LaunchConfig *LaunchConfig

	// The version of the simulation application.
	ApplicationVersion *string

	// Information about tools configured for the simulation application.
	Tools []Tool

	// Information about upload configurations for the simulation application.
	UploadConfigurations []UploadConfiguration

	// A Boolean indicating whether to use default simulation application tools. The
	// default tools are rviz, rqt, terminal and rosbag record. The default is False.
	// This API is no longer supported and will throw an error if used.
	//
	// Deprecated: AWS RoboMaker is ending support for ROS software suite. For
	// additional information, see
	// https://docs.aws.amazon.com/robomaker/latest/dg/software-support-policy.html.
	UseDefaultTools *bool

	// A Boolean indicating whether to use default upload configurations. By default,
	// .ros and .gazebo files are uploaded when the application terminates and all ROS
	// topics will be recorded. If you set this value, you must specify an
	// outputLocation. This API is no longer supported and will throw an error if used.
	//
	// Deprecated: AWS RoboMaker is ending support for ROS software suite. For
	// additional information, see
	// https://docs.aws.amazon.com/robomaker/latest/dg/software-support-policy.html.
	UseDefaultUploadConfigurations *bool

	// A list of world configurations.
	WorldConfigs []WorldConfig

	noSmithyDocumentSerde
}

// Summary information for a simulation application.
type SimulationApplicationSummary struct {

	// The Amazon Resource Name (ARN) of the simulation application.
	Arn *string

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// The name of the simulation application.
	Name *string

	// Information about a robot software suite (ROS distribution).
	RobotSoftwareSuite *RobotSoftwareSuite

	// Information about a simulation software suite.
	SimulationSoftwareSuite *SimulationSoftwareSuite

	// The version of the simulation application.
	Version *string

	noSmithyDocumentSerde
}

// Information about a simulation job.
type SimulationJob struct {

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string

	// A unique identifier for this SimulationJob request.
	ClientRequestToken *string

	// Compute information for the simulation job
	Compute *ComputeResponse

	// The data sources for the simulation job.
	DataSources []DataSource

	// The failure behavior the simulation job. Continue Leaves the host running for
	// its maximum timeout duration after a 4XX error code. Fail Stop the simulation
	// job and terminate the instance.
	FailureBehavior FailureBehavior

	// The failure code of the simulation job if it failed.
	FailureCode SimulationJobErrorCode

	// The reason why the simulation job failed.
	FailureReason *string

	// The IAM role that allows the simulation instance to call the AWS APIs that are
	// specified in its associated policies on your behalf. This is how credentials are
	// passed in to your simulation job.
	IamRole *string

	// The time, in milliseconds since the epoch, when the simulation job was last
	// started.
	LastStartedAt *time.Time

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time

	// The logging configuration.
	LoggingConfig *LoggingConfig

	// The maximum simulation job duration in seconds. The value must be 8 days
	// (691,200 seconds) or less.
	MaxJobDurationInSeconds int64

	// The name of the simulation job.
	Name *string

	// Information about a network interface.
	NetworkInterface *NetworkInterface

	// Location for output files generated by the simulation job.
	OutputLocation *OutputLocation

	// A list of robot applications.
	RobotApplications []RobotApplicationConfig

	// A list of simulation applications.
	SimulationApplications []SimulationApplicationConfig

	// The simulation job execution duration in milliseconds.
	SimulationTimeMillis int64

	// Status of the simulation job.
	Status SimulationJobStatus

	// A map that contains tag keys and tag values that are attached to the simulation
	// job.
	Tags map[string]string

	// VPC configuration information.
	VpcConfig *VPCConfigResponse

	noSmithyDocumentSerde
}

// Information about a simulation job batch.
type SimulationJobBatchSummary struct {

	// The Amazon Resource Name (ARN) of the batch.
	Arn *string

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// created.
	CreatedAt *time.Time

	// The number of created simulation job requests.
	CreatedRequestCount int32

	// The number of failed simulation job requests.
	FailedRequestCount int32

	// The time, in milliseconds since the epoch, when the simulation job batch was
	// last updated.
	LastUpdatedAt *time.Time

	// The number of pending simulation job requests.
	PendingRequestCount int32

	// The status of the simulation job batch. Pending The simulation job batch request
	// is pending. InProgress The simulation job batch is in progress. Failed The
	// simulation job batch failed. One or more simulation job requests could not be
	// completed due to an internal failure (like InternalServiceError). See
	// failureCode and failureReason for more information. Completed The simulation
	// batch job completed. A batch is complete when (1) there are no pending
	// simulation job requests in the batch and none of the failed simulation job
	// requests are due to InternalServiceError and (2) when all created simulation
	// jobs have reached a terminal state (for example, Completed or Failed). Canceled
	// The simulation batch job was cancelled. Canceling The simulation batch job is
	// being cancelled. Completing The simulation batch job is completing. TimingOut
	// The simulation job batch is timing out. If a batch timing out, and there are
	// pending requests that were failing due to an internal failure (like
	// InternalServiceError), the batch status will be Failed. If there are no such
	// failing request, the batch status will be TimedOut. TimedOut The simulation
	// batch job timed out.
	Status SimulationJobBatchStatus

	noSmithyDocumentSerde
}

// Information about a simulation job request.
type SimulationJobRequest struct {

	// The maximum simulation job duration in seconds. The value must be 8 days
	// (691,200 seconds) or less.
	//
	// This member is required.
	MaxJobDurationInSeconds int64

	// Compute information for the simulation job
	Compute *Compute

	// Specify data sources to mount read-only files from S3 into your simulation.
	// These files are available under /opt/robomaker/datasources/data_source_name.
	// There is a limit of 100 files and a combined size of 25GB for all
	// DataSourceConfig objects.
	DataSources []DataSourceConfig

	// The failure behavior the simulation job. Continue Leaves the host running for
	// its maximum timeout duration after a 4XX error code. Fail Stop the simulation
	// job and terminate the instance.
	FailureBehavior FailureBehavior

	// The IAM role name that allows the simulation instance to call the AWS APIs that
	// are specified in its associated policies on your behalf. This is how credentials
	// are passed in to your simulation job.
	IamRole *string

	// The logging configuration.
	LoggingConfig *LoggingConfig

	// The output location.
	OutputLocation *OutputLocation

	// The robot applications to use in the simulation job.
	RobotApplications []RobotApplicationConfig

	// The simulation applications to use in the simulation job.
	SimulationApplications []SimulationApplicationConfig

	// A map that contains tag keys and tag values that are attached to the simulation
	// job request.
	Tags map[string]string

	// A Boolean indicating whether to use default applications in the simulation job.
	// Default applications include Gazebo, rqt, rviz and terminal access.
	UseDefaultApplications *bool

	// If your simulation job accesses resources in a VPC, you provide this parameter
	// identifying the list of security group IDs and subnet IDs. These must belong to
	// the same VPC. You must provide at least one security group and two subnet IDs.
	VpcConfig *VPCConfig

	noSmithyDocumentSerde
}

// Summary information for a simulation job.
type SimulationJobSummary struct {

	// The Amazon Resource Name (ARN) of the simulation job.
	Arn *string

	// The compute type for the simulation job summary.
	ComputeType ComputeType

	// The names of the data sources.
	DataSourceNames []string

	// The time, in milliseconds since the epoch, when the simulation job was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the simulation job.
	Name *string

	// A list of simulation job robot application names.
	RobotApplicationNames []string

	// A list of simulation job simulation application names.
	SimulationApplicationNames []string

	// The status of the simulation job.
	Status SimulationJobStatus

	noSmithyDocumentSerde
}

// Information about a simulation software suite.
type SimulationSoftwareSuite struct {

	// The name of the simulation software suite.
	Name SimulationSoftwareSuiteType

	// The version of the simulation software suite.
	Version *string

	noSmithyDocumentSerde
}

// Information about a source.
type Source struct {

	// The taget processor architecture for the application.
	Architecture Architecture

	// A hash of the object specified by s3Bucket and s3Key.
	Etag *string

	// The s3 bucket name.
	S3Bucket *string

	// The s3 object key.
	S3Key *string

	noSmithyDocumentSerde
}

// Information about a source configuration.
type SourceConfig struct {

	// The target processor architecture for the application.
	Architecture Architecture

	// The Amazon S3 bucket name.
	S3Bucket *string

	// The s3 object key.
	S3Key *string

	noSmithyDocumentSerde
}

// Information about a template location.
type TemplateLocation struct {

	// The Amazon S3 bucket name.
	//
	// This member is required.
	S3Bucket *string

	// The list of S3 keys identifying the data source files.
	//
	// This member is required.
	S3Key *string

	noSmithyDocumentSerde
}

// Summary information for a template.
type TemplateSummary struct {

	// The Amazon Resource Name (ARN) of the template.
	Arn *string

	// The time, in milliseconds since the epoch, when the template was created.
	CreatedAt *time.Time

	// The time, in milliseconds since the epoch, when the template was last updated.
	LastUpdatedAt *time.Time

	// The name of the template.
	Name *string

	// The version of the template that you're using.
	Version *string

	noSmithyDocumentSerde
}

// Information about a tool. Tools are used in a simulation job.
type Tool struct {

	// Command-line arguments for the tool. It must include the tool executable name.
	//
	// This member is required.
	Command *string

	// The name of the tool.
	//
	// This member is required.
	Name *string

	// Exit behavior determines what happens when your tool quits running. RESTART will
	// cause your tool to be restarted. FAIL will cause your job to exit. The default
	// is RESTART.
	ExitBehavior ExitBehavior

	// Boolean indicating whether logs will be recorded in CloudWatch for the tool. The
	// default is False.
	StreamOutputToCloudWatch *bool

	// Boolean indicating whether a streaming session will be configured for the tool.
	// If True, AWS RoboMaker will configure a connection so you can interact with the
	// tool as it is running in the simulation. It must have a graphical user
	// interface. The default is False.
	StreamUI *bool

	noSmithyDocumentSerde
}

// Provides upload configuration information. Files are uploaded from the
// simulation job to a location you specify.
type UploadConfiguration struct {

	// A prefix that specifies where files will be uploaded in Amazon S3. It is
	// appended to the simulation output location to determine the final path. For
	// example, if your simulation output location is s3://my-bucket and your upload
	// configuration name is robot-test, your files will be uploaded to
	// s3://my-bucket///robot-test.
	//
	// This member is required.
	Name *string

	// Specifies the path of the file(s) to upload. Standard Unix glob matching rules
	// are accepted, with the addition of ** as a super asterisk. For example,
	// specifying /var/log/**.log causes all .log files in the /var/log directory tree
	// to be collected. For more examples, see Glob Library
	// (https://github.com/gobwas/glob).
	//
	// This member is required.
	Path *string

	// Specifies when to upload the files: UPLOAD_ON_TERMINATE Matching files are
	// uploaded once the simulation enters the TERMINATING state. Matching files are
	// not uploaded until all of your code (including tools) have stopped. If there is
	// a problem uploading a file, the upload is retried. If problems persist, no
	// further upload attempts will be made. UPLOAD_ROLLING_AUTO_REMOVE Matching files
	// are uploaded as they are created. They are deleted after they are uploaded. The
	// specified path is checked every 5 seconds. A final check is made when all of
	// your code (including tools) have stopped.
	//
	// This member is required.
	UploadBehavior UploadBehavior

	noSmithyDocumentSerde
}

// If your simulation job accesses resources in a VPC, you provide this parameter
// identifying the list of security group IDs and subnet IDs. These must belong to
// the same VPC. You must provide at least one security group and two subnet IDs.
type VPCConfig struct {

	// A list of one or more subnet IDs in your VPC.
	//
	// This member is required.
	Subnets []string

	// A boolean indicating whether to assign a public IP address.
	AssignPublicIp bool

	// A list of one or more security groups IDs in your VPC.
	SecurityGroups []string

	noSmithyDocumentSerde
}

// VPC configuration associated with your simulation job.
type VPCConfigResponse struct {

	// A boolean indicating if a public IP was assigned.
	AssignPublicIp bool

	// A list of security group IDs associated with the simulation job.
	SecurityGroups []string

	// A list of subnet IDs associated with the simulation job.
	Subnets []string

	// The VPC ID associated with your simulation job.
	VpcId *string

	noSmithyDocumentSerde
}

// Configuration information for a world.
type WorldConfig struct {

	// The world generated by Simulation WorldForge.
	World *string

	noSmithyDocumentSerde
}

// The number of worlds that will be created. You can configure the number of
// unique floorplans and the number of unique interiors for each floor plan. For
// example, if you want 1 world with 20 unique interiors, you set floorplanCount =
// 1 and interiorCountPerFloorplan = 20. This will result in 20 worlds
// (floorplanCount * interiorCountPerFloorplan). If you set floorplanCount = 4 and
// interiorCountPerFloorplan = 5, there will be 20 worlds with 5 unique floor
// plans.
type WorldCount struct {

	// The number of unique floorplans.
	FloorplanCount *int32

	// The number of unique interiors per floorplan.
	InteriorCountPerFloorplan *int32

	noSmithyDocumentSerde
}

// Information about a world export job.
type WorldExportJobSummary struct {

	// The Amazon Resource Name (ARN) of the world export job.
	Arn *string

	// The time, in milliseconds since the epoch, when the world export job was
	// created.
	CreatedAt *time.Time

	// The output location.
	OutputLocation *OutputLocation

	// The status of the world export job. Pending The world export job request is
	// pending. Running The world export job is running. Completed The world export job
	// completed. Failed The world export job failed. See failureCode for more
	// information. Canceled The world export job was cancelled. Canceling The world
	// export job is being cancelled.
	Status WorldExportJobStatus

	// A list of worlds.
	Worlds []string

	noSmithyDocumentSerde
}

// Information about a failed world.
type WorldFailure struct {

	// The failure code of the world export job if it failed: InternalServiceError
	// Internal service error. LimitExceeded The requested resource exceeds the maximum
	// number allowed, or the number of concurrent stream requests exceeds the maximum
	// number allowed. ResourceNotFound The specified resource could not be found.
	// RequestThrottled The request was throttled. InvalidInput An input parameter in
	// the request is not valid.
	FailureCode WorldGenerationJobErrorCode

	// The number of failed worlds.
	FailureCount int32

	// The sample reason why the world failed. World errors are aggregated. A sample is
	// used as the sampleFailureReason.
	SampleFailureReason *string

	noSmithyDocumentSerde
}

// Information about a world generator job.
type WorldGenerationJobSummary struct {

	// The Amazon Resource Name (ARN) of the world generator job.
	Arn *string

	// The time, in milliseconds since the epoch, when the world generator job was
	// created.
	CreatedAt *time.Time

	// The number of worlds that failed.
	FailedWorldCount int32

	// The status of the world generator job: Pending The world generator job request
	// is pending. Running The world generator job is running. Completed The world
	// generator job completed. Failed The world generator job failed. See failureCode
	// for more information. PartialFailed Some worlds did not generate. Canceled The
	// world generator job was cancelled. Canceling The world generator job is being
	// cancelled.
	Status WorldGenerationJobStatus

	// The number of worlds that were generated.
	SucceededWorldCount int32

	// The Amazon Resource Name (arn) of the world template.
	Template *string

	// Information about the world count.
	WorldCount *WorldCount

	noSmithyDocumentSerde
}

// Information about a world.
type WorldSummary struct {

	// The Amazon Resource Name (ARN) of the world.
	Arn *string

	// The time, in milliseconds since the epoch, when the world was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (arn) of the world generation job.
	GenerationJob *string

	// The Amazon Resource Name (arn) of the world template.
	Template *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
