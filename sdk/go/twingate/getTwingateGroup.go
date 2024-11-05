// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Groups are how users are authorized to access Resources. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/groups).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := twingate.LookupTwingateGroup(ctx, &twingate.LookupTwingateGroupArgs{
//				Id: "<your group's id>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTwingateGroup(ctx *pulumi.Context, args *LookupTwingateGroupArgs, opts ...pulumi.InvokeOption) (*LookupTwingateGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTwingateGroupResult
	err := ctx.Invoke("twingate:index/getTwingateGroup:getTwingateGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateGroup.
type LookupTwingateGroupArgs struct {
	// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
	Id string `pulumi:"id"`
}

// A collection of values returned by getTwingateGroup.
type LookupTwingateGroupResult struct {
	// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
	Id string `pulumi:"id"`
	// Indicates if the Group is active
	IsActive bool `pulumi:"isActive"`
	// The name of the Group
	Name string `pulumi:"name"`
	// The Security Policy assigned to the Group.
	SecurityPolicyId string `pulumi:"securityPolicyId"`
	// The type of the Group
	Type string `pulumi:"type"`
}

func LookupTwingateGroupOutput(ctx *pulumi.Context, args LookupTwingateGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTwingateGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTwingateGroupResultOutput, error) {
			args := v.(LookupTwingateGroupArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupTwingateGroupResult
			secret, err := ctx.InvokePackageRaw("twingate:index/getTwingateGroup:getTwingateGroup", args, &rv, "", opts...)
			if err != nil {
				return LookupTwingateGroupResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupTwingateGroupResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupTwingateGroupResultOutput), nil
			}
			return output, nil
		}).(LookupTwingateGroupResultOutput)
}

// A collection of arguments for invoking getTwingateGroup.
type LookupTwingateGroupOutputArgs struct {
	// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTwingateGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateGroupArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateGroup.
type LookupTwingateGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTwingateGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateGroupResult)(nil)).Elem()
}

func (o LookupTwingateGroupResultOutput) ToLookupTwingateGroupResultOutput() LookupTwingateGroupResultOutput {
	return o
}

func (o LookupTwingateGroupResultOutput) ToLookupTwingateGroupResultOutputWithContext(ctx context.Context) LookupTwingateGroupResultOutput {
	return o
}

// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
func (o LookupTwingateGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates if the Group is active
func (o LookupTwingateGroupResultOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTwingateGroupResult) bool { return v.IsActive }).(pulumi.BoolOutput)
}

// The name of the Group
func (o LookupTwingateGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Security Policy assigned to the Group.
func (o LookupTwingateGroupResultOutput) SecurityPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateGroupResult) string { return v.SecurityPolicyId }).(pulumi.StringOutput)
}

// The type of the Group
func (o LookupTwingateGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTwingateGroupResultOutput{})
}
