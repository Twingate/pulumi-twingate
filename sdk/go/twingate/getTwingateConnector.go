// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/go/twingate/internal"
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
//	"github.com/Twingate/pulumi-twingate/sdk/go/twingate"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := twingate.LookupTwingateConnector(ctx, &twingate.LookupTwingateConnectorArgs{
//				Id: "<your connector's id>",
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
func LookupTwingateConnector(ctx *pulumi.Context, args *LookupTwingateConnectorArgs, opts ...pulumi.InvokeOption) (*LookupTwingateConnectorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTwingateConnectorResult
	err := ctx.Invoke("twingate:index/getTwingateConnector:getTwingateConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateConnector.
type LookupTwingateConnectorArgs struct {
	// The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
	Id string `pulumi:"id"`
}

// A collection of values returned by getTwingateConnector.
type LookupTwingateConnectorResult struct {
	// The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
	Id string `pulumi:"id"`
	// The name of the Connector.
	Name string `pulumi:"name"`
	// The ID of the Remote Network the Connector is attached to.
	RemoteNetworkId string `pulumi:"remoteNetworkId"`
	// Determines whether status notifications are enabled for the Connector.
	StatusUpdatesEnabled bool `pulumi:"statusUpdatesEnabled"`
}

func LookupTwingateConnectorOutput(ctx *pulumi.Context, args LookupTwingateConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupTwingateConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTwingateConnectorResult, error) {
			args := v.(LookupTwingateConnectorArgs)
			r, err := LookupTwingateConnector(ctx, &args, opts...)
			var s LookupTwingateConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTwingateConnectorResultOutput)
}

// A collection of arguments for invoking getTwingateConnector.
type LookupTwingateConnectorOutputArgs struct {
	// The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTwingateConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateConnectorArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateConnector.
type LookupTwingateConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupTwingateConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateConnectorResult)(nil)).Elem()
}

func (o LookupTwingateConnectorResultOutput) ToLookupTwingateConnectorResultOutput() LookupTwingateConnectorResultOutput {
	return o
}

func (o LookupTwingateConnectorResultOutput) ToLookupTwingateConnectorResultOutputWithContext(ctx context.Context) LookupTwingateConnectorResultOutput {
	return o
}

// The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
func (o LookupTwingateConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Connector.
func (o LookupTwingateConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Remote Network the Connector is attached to.
func (o LookupTwingateConnectorResultOutput) RemoteNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateConnectorResult) string { return v.RemoteNetworkId }).(pulumi.StringOutput)
}

// Determines whether status notifications are enabled for the Connector.
func (o LookupTwingateConnectorResultOutput) StatusUpdatesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTwingateConnectorResult) bool { return v.StatusUpdatesEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTwingateConnectorResultOutput{})
}
