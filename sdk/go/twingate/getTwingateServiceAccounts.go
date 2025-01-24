// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.
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
//			_, err := twingate.GetTwingateServiceAccounts(ctx, &twingate.GetTwingateServiceAccountsArgs{
//				Name: pulumi.StringRef("<your service account's name>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTwingateServiceAccounts(ctx *pulumi.Context, args *GetTwingateServiceAccountsArgs, opts ...pulumi.InvokeOption) (*GetTwingateServiceAccountsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTwingateServiceAccountsResult
	err := ctx.Invoke("twingate:index/getTwingateServiceAccounts:getTwingateServiceAccounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateServiceAccounts.
type GetTwingateServiceAccountsArgs struct {
	// Returns only service accounts that exactly match this name. If no options are passed it will return all service accounts. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the service account.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the service account.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the service account must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the service account.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the service account must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
}

// A collection of values returned by getTwingateServiceAccounts.
type GetTwingateServiceAccountsResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// Returns only service accounts that exactly match this name. If no options are passed it will return all service accounts. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the service account.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the service account.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the service account must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the service account.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the service account must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
	// List of Service Accounts
	ServiceAccounts []GetTwingateServiceAccountsServiceAccount `pulumi:"serviceAccounts"`
}

func GetTwingateServiceAccountsOutput(ctx *pulumi.Context, args GetTwingateServiceAccountsOutputArgs, opts ...pulumi.InvokeOption) GetTwingateServiceAccountsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTwingateServiceAccountsResultOutput, error) {
			args := v.(GetTwingateServiceAccountsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("twingate:index/getTwingateServiceAccounts:getTwingateServiceAccounts", args, GetTwingateServiceAccountsResultOutput{}, options).(GetTwingateServiceAccountsResultOutput), nil
		}).(GetTwingateServiceAccountsResultOutput)
}

// A collection of arguments for invoking getTwingateServiceAccounts.
type GetTwingateServiceAccountsOutputArgs struct {
	// Returns only service accounts that exactly match this name. If no options are passed it will return all service accounts. Only one option can be used at a time.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Match when the value exist in the name of the service account.
	NameContains pulumi.StringPtrInput `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the service account.
	NameExclude pulumi.StringPtrInput `pulumi:"nameExclude"`
	// The name of the service account must start with the value.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// The regular expression match of the name of the service account.
	NameRegexp pulumi.StringPtrInput `pulumi:"nameRegexp"`
	// The name of the service account must end with the value.
	NameSuffix pulumi.StringPtrInput `pulumi:"nameSuffix"`
}

func (GetTwingateServiceAccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateServiceAccountsArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateServiceAccounts.
type GetTwingateServiceAccountsResultOutput struct{ *pulumi.OutputState }

func (GetTwingateServiceAccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateServiceAccountsResult)(nil)).Elem()
}

func (o GetTwingateServiceAccountsResultOutput) ToGetTwingateServiceAccountsResultOutput() GetTwingateServiceAccountsResultOutput {
	return o
}

func (o GetTwingateServiceAccountsResultOutput) ToGetTwingateServiceAccountsResultOutputWithContext(ctx context.Context) GetTwingateServiceAccountsResultOutput {
	return o
}

// The ID of this resource.
func (o GetTwingateServiceAccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns only service accounts that exactly match this name. If no options are passed it will return all service accounts. Only one option can be used at a time.
func (o GetTwingateServiceAccountsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Match when the value exist in the name of the service account.
func (o GetTwingateServiceAccountsResultOutput) NameContains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) *string { return v.NameContains }).(pulumi.StringPtrOutput)
}

// Match when the exact value does not exist in the name of the service account.
func (o GetTwingateServiceAccountsResultOutput) NameExclude() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) *string { return v.NameExclude }).(pulumi.StringPtrOutput)
}

// The name of the service account must start with the value.
func (o GetTwingateServiceAccountsResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The regular expression match of the name of the service account.
func (o GetTwingateServiceAccountsResultOutput) NameRegexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) *string { return v.NameRegexp }).(pulumi.StringPtrOutput)
}

// The name of the service account must end with the value.
func (o GetTwingateServiceAccountsResultOutput) NameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) *string { return v.NameSuffix }).(pulumi.StringPtrOutput)
}

// List of Service Accounts
func (o GetTwingateServiceAccountsResultOutput) ServiceAccounts() GetTwingateServiceAccountsServiceAccountArrayOutput {
	return o.ApplyT(func(v GetTwingateServiceAccountsResult) []GetTwingateServiceAccountsServiceAccount {
		return v.ServiceAccounts
	}).(GetTwingateServiceAccountsServiceAccountArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateServiceAccountsResultOutput{})
}
