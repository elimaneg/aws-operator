package cpi

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/aws-operator/service/controller/v25/controllercontext"
	"github.com/giantswarm/aws-operator/service/controller/v25/key"
	"github.com/giantswarm/aws-operator/service/controller/v25/resource/cpi/template"
)

const (
	capabilityNamesIAM = "CAPABILITY_NAMED_IAM"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cr, err := key.ToCustomObject(obj)
	if err != nil {
		return microerror.Mask(err)
	}
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	{
		r.logger.LogCtx(ctx, "level", "debug", "message", "finding the tenant cluster's control plane initializer CF stack")

		i := &cloudformation.DescribeStacksInput{
			StackName: aws.String(key.MainHostPreStackName(cr)),
		}

		_, err = cc.Client.ControlPlane.AWS.CloudFormation.DescribeStacks(i)
		if IsNotExists(err) {
			// fall through
		} else if err != nil {
			return microerror.Mask(err)
		} else {
			r.logger.LogCtx(ctx, "level", "debug", "message", "found the tenant cluster's control plane initializer CF stack already exists")
			r.logger.LogCtx(ctx, "level", "debug", "message", "canceling resource")

			return nil
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", "did not find the tenant cluster's control plane initializer CF stack")
	}

	var templateBody string
	{
		r.logger.LogCtx(ctx, "level", "debug", "message", "computing the template of the tenant cluster's control plane initializer CF stack")

		var params *template.ParamsMain
		{
			iamRoles, err := r.newIAMRolesParams(ctx, cr)
			if err != nil {
				return microerror.Mask(err)
			}

			params = &template.ParamsMain{
				IAMRoles: iamRoles,
			}
		}

		templateBody, err = template.Render(params)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", "computed the template of the tenant cluster's control plane initializer CF stack")
	}

	{
		r.logger.LogCtx(ctx, "level", "debug", "message", "requesting the creation of the tenant cluster's control plane initializer CF stack")

		i := &cloudformation.CreateStackInput{
			Capabilities: []*string{
				aws.String(capabilityNamesIAM),
			},
			EnableTerminationProtection: aws.Bool(key.EnableTerminationProtection),
			StackName:                   aws.String(key.MainHostPreStackName(cr)),
			Tags:                        r.getCloudFormationTags(cr),
			TemplateBody:                aws.String(templateBody),
		}

		_, err = cc.Client.ControlPlane.AWS.CloudFormation.CreateStack(i)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", "requested the creation of the tenant cluster's control plane initializer CF stack")
	}

	{
		r.logger.LogCtx(ctx, "level", "debug", "message", "waiting for the creation of the tenant cluster's control plane initializer CF stack")

		i := &cloudformation.DescribeStacksInput{
			StackName: aws.String(key.MainHostPreStackName(cr)),
		}

		err = cc.Client.ControlPlane.AWS.CloudFormation.WaitUntilStackCreateComplete(i)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", "waited for the creation of the tenant cluster's control plane initializer CF stack")
	}

	return nil
}

func (r *Resource) newIAMRolesParams(ctx context.Context, cr v1alpha1.AWSConfig) (*template.ParamsMainIAMRoles, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	iamRoles := &template.ParamsMainIAMRoles{
		PeerAccessRoleName: key.PeerAccessRoleName(cr),
		Tenant: template.ParamsMainIAMRolesTenant{
			AWS: template.ParamsMainIAMRolesTenantAWS{
				Account: template.ParamsMainIAMRolesTenantAWSAccount{
					ID: cc.Status.TenantCluster.AWSAccountID,
				},
			},
		},
	}

	return iamRoles, nil
}
