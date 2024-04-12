// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Connectors provide connectivity to Remote Networks. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := twingate.GetTwingateConnectors(ctx, &twingate.GetTwingateConnectorsArgs{
//				Name: pulumi.StringRef("<your connector's name>"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetTwingateConnectors(ctx *pulumi.Context, args *GetTwingateConnectorsArgs, opts ...pulumi.InvokeOption) (*GetTwingateConnectorsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTwingateConnectorsResult
	err := ctx.Invoke("twingate:index/getTwingateConnectors:getTwingateConnectors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateConnectors.
type GetTwingateConnectorsArgs struct {
	// The Name of the Connector.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the connector.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the connector.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the connector must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the connector.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the connector must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
}

// A collection of values returned by getTwingateConnectors.
type GetTwingateConnectorsResult struct {
	// List of Connectors
	Connectors []GetTwingateConnectorsConnector `pulumi:"connectors"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// Returns only connectors that exactly match this name. If no options are passed it will return all connectors. Only one option can be used at a time.
	Name *string `pulumi:"name"`
	// Match when the value exist in the name of the connector.
	NameContains *string `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the connector.
	NameExclude *string `pulumi:"nameExclude"`
	// The name of the connector must start with the value.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression match of the name of the connector.
	NameRegexp *string `pulumi:"nameRegexp"`
	// The name of the connector must end with the value.
	NameSuffix *string `pulumi:"nameSuffix"`
}

func GetTwingateConnectorsOutput(ctx *pulumi.Context, args GetTwingateConnectorsOutputArgs, opts ...pulumi.InvokeOption) GetTwingateConnectorsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTwingateConnectorsResult, error) {
			args := v.(GetTwingateConnectorsArgs)
			r, err := GetTwingateConnectors(ctx, &args, opts...)
			var s GetTwingateConnectorsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTwingateConnectorsResultOutput)
}

// A collection of arguments for invoking getTwingateConnectors.
type GetTwingateConnectorsOutputArgs struct {
	// The Name of the Connector.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Match when the value exist in the name of the connector.
	NameContains pulumi.StringPtrInput `pulumi:"nameContains"`
	// Match when the exact value does not exist in the name of the connector.
	NameExclude pulumi.StringPtrInput `pulumi:"nameExclude"`
	// The name of the connector must start with the value.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// The regular expression match of the name of the connector.
	NameRegexp pulumi.StringPtrInput `pulumi:"nameRegexp"`
	// The name of the connector must end with the value.
	NameSuffix pulumi.StringPtrInput `pulumi:"nameSuffix"`
}

func (GetTwingateConnectorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateConnectorsArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateConnectors.
type GetTwingateConnectorsResultOutput struct{ *pulumi.OutputState }

func (GetTwingateConnectorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTwingateConnectorsResult)(nil)).Elem()
}

func (o GetTwingateConnectorsResultOutput) ToGetTwingateConnectorsResultOutput() GetTwingateConnectorsResultOutput {
	return o
}

func (o GetTwingateConnectorsResultOutput) ToGetTwingateConnectorsResultOutputWithContext(ctx context.Context) GetTwingateConnectorsResultOutput {
	return o
}

// List of Connectors
func (o GetTwingateConnectorsResultOutput) Connectors() GetTwingateConnectorsConnectorArrayOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) []GetTwingateConnectorsConnector { return v.Connectors }).(GetTwingateConnectorsConnectorArrayOutput)
}

// The ID of this resource.
func (o GetTwingateConnectorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns only connectors that exactly match this name. If no options are passed it will return all connectors. Only one option can be used at a time.
func (o GetTwingateConnectorsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Match when the value exist in the name of the connector.
func (o GetTwingateConnectorsResultOutput) NameContains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) *string { return v.NameContains }).(pulumi.StringPtrOutput)
}

// Match when the exact value does not exist in the name of the connector.
func (o GetTwingateConnectorsResultOutput) NameExclude() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) *string { return v.NameExclude }).(pulumi.StringPtrOutput)
}

// The name of the connector must start with the value.
func (o GetTwingateConnectorsResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The regular expression match of the name of the connector.
func (o GetTwingateConnectorsResultOutput) NameRegexp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) *string { return v.NameRegexp }).(pulumi.StringPtrOutput)
}

// The name of the connector must end with the value.
func (o GetTwingateConnectorsResultOutput) NameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTwingateConnectorsResult) *string { return v.NameSuffix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTwingateConnectorsResultOutput{})
}
