// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources in Twingate represent any network destination address that you wish to provide private access to for users authorized via the Twingate Client application. Resources can be defined by either IP or DNS address, and all private DNS addresses will be automatically resolved with no client configuration changes. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
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
//			_, err := twingate.LookupTwingateResource(ctx, &twingate.LookupTwingateResourceArgs{
//				Id: "<your resource's id>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTwingateResource(ctx *pulumi.Context, args *LookupTwingateResourceArgs, opts ...pulumi.InvokeOption) (*LookupTwingateResourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTwingateResourceResult
	err := ctx.Invoke("twingate:index/getTwingateResource:getTwingateResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateResource.
type LookupTwingateResourceArgs struct {
	// The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
	Id string `pulumi:"id"`
	// By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
	Protocols *GetTwingateResourceProtocols `pulumi:"protocols"`
}

// A collection of values returned by getTwingateResource.
type LookupTwingateResourceResult struct {
	// The Resource's address, which may be an IP address, CIDR range, or DNS address
	Address string `pulumi:"address"`
	// The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
	Id string `pulumi:"id"`
	// The name of the Resource
	Name string `pulumi:"name"`
	// By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
	Protocols *GetTwingateResourceProtocols `pulumi:"protocols"`
	// The Remote Network ID that the Resource is associated with. Resources may only be associated with a single Remote Network.
	RemoteNetworkId string `pulumi:"remoteNetworkId"`
	// The `tags` attribute consists of a key-value pairs that correspond with tags to be set on the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupTwingateResourceOutput(ctx *pulumi.Context, args LookupTwingateResourceOutputArgs, opts ...pulumi.InvokeOption) LookupTwingateResourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTwingateResourceResultOutput, error) {
			args := v.(LookupTwingateResourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("twingate:index/getTwingateResource:getTwingateResource", args, LookupTwingateResourceResultOutput{}, options).(LookupTwingateResourceResultOutput), nil
		}).(LookupTwingateResourceResultOutput)
}

// A collection of arguments for invoking getTwingateResource.
type LookupTwingateResourceOutputArgs struct {
	// The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
	Id pulumi.StringInput `pulumi:"id"`
	// By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
	Protocols GetTwingateResourceProtocolsPtrInput `pulumi:"protocols"`
}

func (LookupTwingateResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateResourceArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateResource.
type LookupTwingateResourceResultOutput struct{ *pulumi.OutputState }

func (LookupTwingateResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateResourceResult)(nil)).Elem()
}

func (o LookupTwingateResourceResultOutput) ToLookupTwingateResourceResultOutput() LookupTwingateResourceResultOutput {
	return o
}

func (o LookupTwingateResourceResultOutput) ToLookupTwingateResourceResultOutputWithContext(ctx context.Context) LookupTwingateResourceResultOutput {
	return o
}

// The Resource's address, which may be an IP address, CIDR range, or DNS address
func (o LookupTwingateResourceResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateResourceResult) string { return v.Address }).(pulumi.StringOutput)
}

// The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
func (o LookupTwingateResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Resource
func (o LookupTwingateResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
func (o LookupTwingateResourceResultOutput) Protocols() GetTwingateResourceProtocolsPtrOutput {
	return o.ApplyT(func(v LookupTwingateResourceResult) *GetTwingateResourceProtocols { return v.Protocols }).(GetTwingateResourceProtocolsPtrOutput)
}

// The Remote Network ID that the Resource is associated with. Resources may only be associated with a single Remote Network.
func (o LookupTwingateResourceResultOutput) RemoteNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateResourceResult) string { return v.RemoteNetworkId }).(pulumi.StringOutput)
}

// The `tags` attribute consists of a key-value pairs that correspond with tags to be set on the resource.
func (o LookupTwingateResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTwingateResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTwingateResourceResultOutput{})
}
