// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
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
//			_, err := twingate.GetTwingateResources(ctx, &twingate.GetTwingateResourcesArgs{
//				Name: pulumi.StringRef("<your resource's name>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTwingateResources(ctx *pulumi.Context, args *GetTwingateResourcesArgs, opts ...pulumi.InvokeOption) (*GetTwingateResourcesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTwingateResourcesResult
	err := ctx.Invoke("twingate:index/getTwingateResources:getTwingateResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateResources.
type GetTwingateResourcesArgs struct {
	// Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the resource.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the resource.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the resource must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the resource.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the resource must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
	// Returns only resources that exactly match the given tags.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getTwingateResources.
type GetTwingateResourcesResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the resource.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the resource.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the resource must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the resource.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the resource must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
	// List of Resources
	Resources []GetTwingateResourcesResource `pulumi:"resources"`
	// Returns only resources that exactly match the given tags.
	Tags map[string]string `pulumi:"tags"`
}

func GetTwingateResourcesOutput(ctx *pulumi.Context, args GetTwingateResourcesOutputArgs, opts ...pulumi.InvokeOption) GetTwingateResourcesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTwingateResourcesResultOutput, error) {
			args := v.(GetTwingateResourcesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("twingate:index/getTwingateResources:getTwingateResources", args, GetTwingateResourcesResultOutput{}, options).(GetTwingateResourcesResultOutput), nil
		}).(GetTwingateResourcesResultOutput)
}

// A collection of arguments for invoking getTwingateResources.
type GetTwingateResourcesOutputArgs struct {
	// Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Match when the value exist in the name of the resource.
	NameContains pulumi.StringPtrInput `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the resource.
	NameExclude pulumi.StringPtrInput `pulumi:"nameExclude"`
	// The name of the resource must start with the value.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// The regular expression match of the name of the resource.
	NameRegexp pulumi.StringPtrInput `pulumi:"nameRegexp"`
	// The name of the resource must end with the value.
	NameSuffix pulumi.StringPtrInput `pulumi:"nameSuffix"`
	// Returns only resources that exactly match the given tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetTwingateResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateResourcesArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateResources.
type GetTwingateResourcesResultOutput struct{ *pulumi.OutputState }

func (GetTwingateResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateResourcesResult)(nil)).Elem()
}

func (o GetTwingateResourcesResultOutput) ToGetTwingateResourcesResultOutput() GetTwingateResourcesResultOutput {
	return o
}

func (o GetTwingateResourcesResultOutput) ToGetTwingateResourcesResultOutputWithContext(ctx context.Context) GetTwingateResourcesResultOutput {
	return o
}

// The ID of this resource.
func (o GetTwingateResourcesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
func (o GetTwingateResourcesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Match when the value exist in the name of the resource.
func (o GetTwingateResourcesResultOutput) NameContains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) *string { return v.NameContains }).(pulumi.StringPtrOutput)
}

// Match when the exact value does not exist in the name of the resource.
func (o GetTwingateResourcesResultOutput) NameExclude() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) *string { return v.NameExclude }).(pulumi.StringPtrOutput)
}

// The name of the resource must start with the value.
func (o GetTwingateResourcesResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The regular expression match of the name of the resource.
func (o GetTwingateResourcesResultOutput) NameRegexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) *string { return v.NameRegexp }).(pulumi.StringPtrOutput)
}

// The name of the resource must end with the value.
func (o GetTwingateResourcesResultOutput) NameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) *string { return v.NameSuffix }).(pulumi.StringPtrOutput)
}

// List of Resources
func (o GetTwingateResourcesResultOutput) Resources() GetTwingateResourcesResourceArrayOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) []GetTwingateResourcesResource { return v.Resources }).(GetTwingateResourcesResourceArrayOutput)
}

// Returns only resources that exactly match the given tags.
func (o GetTwingateResourcesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetTwingateResourcesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateResourcesResultOutput{})
}
