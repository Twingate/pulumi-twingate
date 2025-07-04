// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Remote Network represents a single private network in Twingate that can have one or more Connectors and Resources assigned to it. You must create a Remote Network before creating Resources and Connectors that belong to it. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/remote-networks).
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
//			_, err := twingate.GetTwingateRemoteNetworks(ctx, &twingate.GetTwingateRemoteNetworksArgs{
//				Name: pulumi.StringRef("<your network's name>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTwingateRemoteNetworks(ctx *pulumi.Context, args *GetTwingateRemoteNetworksArgs, opts ...pulumi.InvokeOption) (*GetTwingateRemoteNetworksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTwingateRemoteNetworksResult
	err := ctx.Invoke("twingate:index/getTwingateRemoteNetworks:getTwingateRemoteNetworks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateRemoteNetworks.
type GetTwingateRemoteNetworksArgs struct {
	// Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the remote network.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the remote network.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the remote network must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the remote network.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the remote network must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
}

// A collection of values returned by getTwingateRemoteNetworks.
type GetTwingateRemoteNetworksResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the remote network.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the remote network.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the remote network must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the remote network.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the remote network must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
	// List of Remote Networks
	RemoteNetworks []GetTwingateRemoteNetworksRemoteNetwork `pulumi:"remoteNetworks"`
}

func GetTwingateRemoteNetworksOutput(ctx *pulumi.Context, args GetTwingateRemoteNetworksOutputArgs, opts ...pulumi.InvokeOption) GetTwingateRemoteNetworksResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTwingateRemoteNetworksResultOutput, error) {
			args := v.(GetTwingateRemoteNetworksArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("twingate:index/getTwingateRemoteNetworks:getTwingateRemoteNetworks", args, GetTwingateRemoteNetworksResultOutput{}, options).(GetTwingateRemoteNetworksResultOutput), nil
		}).(GetTwingateRemoteNetworksResultOutput)
}

// A collection of arguments for invoking getTwingateRemoteNetworks.
type GetTwingateRemoteNetworksOutputArgs struct {
	// Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Match when the value exist in the name of the remote network.
	NameContains pulumi.StringPtrInput `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the remote network.
	NameExclude pulumi.StringPtrInput `pulumi:"nameExclude"`
	// The name of the remote network must start with the value.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// The regular expression match of the name of the remote network.
	NameRegexp pulumi.StringPtrInput `pulumi:"nameRegexp"`
	// The name of the remote network must end with the value.
	NameSuffix pulumi.StringPtrInput `pulumi:"nameSuffix"`
}

func (GetTwingateRemoteNetworksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateRemoteNetworksArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateRemoteNetworks.
type GetTwingateRemoteNetworksResultOutput struct{ *pulumi.OutputState }

func (GetTwingateRemoteNetworksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateRemoteNetworksResult)(nil)).Elem()
}

func (o GetTwingateRemoteNetworksResultOutput) ToGetTwingateRemoteNetworksResultOutput() GetTwingateRemoteNetworksResultOutput {
	return o
}

func (o GetTwingateRemoteNetworksResultOutput) ToGetTwingateRemoteNetworksResultOutputWithContext(ctx context.Context) GetTwingateRemoteNetworksResultOutput {
	return o
}

// The ID of this resource.
func (o GetTwingateRemoteNetworksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
func (o GetTwingateRemoteNetworksResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Match when the value exist in the name of the remote network.
func (o GetTwingateRemoteNetworksResultOutput) NameContains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) *string { return v.NameContains }).(pulumi.StringPtrOutput)
}

// Match when the exact value does not exist in the name of the remote network.
func (o GetTwingateRemoteNetworksResultOutput) NameExclude() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) *string { return v.NameExclude }).(pulumi.StringPtrOutput)
}

// The name of the remote network must start with the value.
func (o GetTwingateRemoteNetworksResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The regular expression match of the name of the remote network.
func (o GetTwingateRemoteNetworksResultOutput) NameRegexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) *string { return v.NameRegexp }).(pulumi.StringPtrOutput)
}

// The name of the remote network must end with the value.
func (o GetTwingateRemoteNetworksResultOutput) NameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) *string { return v.NameSuffix }).(pulumi.StringPtrOutput)
}

// List of Remote Networks
func (o GetTwingateRemoteNetworksResultOutput) RemoteNetworks() GetTwingateRemoteNetworksRemoteNetworkArrayOutput {
	return o.ApplyT(func(v GetTwingateRemoteNetworksResult) []GetTwingateRemoteNetworksRemoteNetwork {
		return v.RemoteNetworks
	}).(GetTwingateRemoteNetworksRemoteNetworkArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateRemoteNetworksResultOutput{})
}
