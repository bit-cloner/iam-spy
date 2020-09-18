package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" || os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_DEFAULT_REGION") == "" {
		fmt.Println("AWS credentails not found. Set these credentials uisng aws configure or through environment variables")
		fmt.Println("Make sure these variables are set and have appropriate permission for route53 AWS_SECRET_ACCESS_KEY AWS_ACCESS_KEY_ID AWS_DEFAULT_REGION")
		fmt.Println("aborting..")
		os.Exit(1)
	}
	fmt.Println("Found AWS credentials")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	services := []string{"ServiceShortName", "a4b", "execute-api", "appstream", "athena", "chime", "clouddirectory", "cloudfront", "cloudsearch", "cloudwatch", "logs", "synthetics", "cognito-identity", "cognito-sync", "cognito-idp", "comprehend", "connect", "dlm", "dynamodb", "dax", "ec2", "autoscaling", "imagebuilder", "ec2-instance-connect", "ecr", "ecs", "eks", "elasticfilesystem", "elastic-inference", "elasticmapreduce", "elastictranscoder", "elasticache", "es", "events", "schemas", "forecast", "frauddetector", "freertos", "fsx", "gamelift", "glacier", "groundtruthlabeling", "guardduty", "inspector", "kendra", "kinesis", "kinesisanalytics", "kinesisanalytics", "firehose", "kinesisvideo", "lex", "lightsail", "machinelearning", "macie", "managedblockchain", "kafka", "mechanicalturk", "ec2messages", "mobileanalytics", "mq", "neptune-db", "personalize", "mobiletargeting", "ses", "sms-voice", "polly", "qldb", "quicksight", "rds", "rds-data", "rds-db", "redshift", "rekognition", "tag", "route53", "route53resolver", "route53domains", "s3", "sagemaker", "ses", "ssmmessages", "swf", "sdb", "sns", "sqs", "storagegateway", "sumerian", "textract", "transcribe", "translate", "workdocs", "worklink", "workmail", "workmailmessageflow", "workspaces", "wam", "application-autoscaling", "discovery", "account", "amplify", "appmesh", "appmesh-preview", "appsync", "artifact", "autoscaling-plans", "backup", "backup-storage", "batch", "aws-portal", "budgets", "acm", "acm-pca", "chatbot", "servicediscovery", "cloud9", "cloudformation", "cloudhsm", "cloudtrail", "signer", "codebuild", "codecommit", "codedeploy", "codepipeline", "codestar", "codestar-notifications", "config", "cur", "ce", "dataexchange", "dms", "deeplens", "deepracer", "devicefarm", "directconnect", "ds", "elasticbeanstalk", "mediaconnect", "mediaconvert", "medialive", "mediapackage", "mediapackage-vod", "mediastore", "mediatailor", "fms", "globalaccelerator", "glue", "groundstation", "health", "importexport", "iot", "iot1click", "iotanalytics", "iotevents", "greengrass", "iotsitewise", "iotthingsgraph", "iq", "iq-permission", "kms", "lakeformation", "lambda", "license-manager", "cassandra", "aws-marketplace", "aws-marketplace", "aws-marketplace", "aws-marketplace", "aws-marketplace-management", "aws-marketplace", "aws-marketplace", "mgh", "mobilehub", "opsworks", "opsworks-cm", "organizations", "pi", "pricing", "aws-marketplace", "ram", "resource-groups", "robomaker", "savingsplans", "secretsmanager", "securityhub", "sts", "sms", "serverlessrepo", "servicecatalog", "shield", "snowball", "sso", "sso-directory", "states", "support", "ssm", "transfer", "trustedadvisor", "waf", "waf-regional", "wellarchitected", "xray", "applicationinsights", "comprehendmedical", "compute-optimizer", "datapipeline", "dbqms", "datasync", "elasticloadbalancing", "elasticloadbalancing", "access-analyzer", "iam", "launchwizard", "apigateway", "networkmanager", "servicequotas"}
	selectedsvc := ""
	prompt := &survey.Select{
		Message: "Select one of the follwing AWS services for analysis",
		Options: services,
	}
	survey.AskOne(prompt, &selectedsvc, nil)
	fmt.Println(selectedsvc)
	svc := iam.New(sess)
	input1 := &iam.ListUsersInput{MaxItems: aws.Int64(5)}
	result1, err := svc.ListUsers(input1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var userlist []string
	for _, item := range result1.Users {
		if item == nil {
			continue
		}
		userlist = append(userlist, *item.UserName)
	}
	for _, user := range userlist {
		instring := "arn:aws:iam::372998229667:user/" + string(user)
		lpgsainput := &iam.ListPoliciesGrantingServiceAccessInput{Arn: aws.String(instring),
			ServiceNamespaces: []*string{
				aws.String(selectedsvc),
			},
		}
		result2, err := svc.ListPoliciesGrantingServiceAccess(lpgsainput)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if result2.PoliciesGrantingServiceAccess[0].Policies == nil {
			continue
		}
		fmt.Println("user: ", user, "has access to", selectedsvc, "service")
		for _, policies := range result2.PoliciesGrantingServiceAccess {
			for _, policy := range policies.Policies {
				fmt.Println("access granted by policy ", *policy.PolicyName)
			}
		}
	}
}
