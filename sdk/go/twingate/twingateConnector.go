// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"errors"
	"github.com/Twingate/pulumi-twingate/sdk/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Connectors provide connectivity to Remote Networks. This resource type will create the Connector in the Twingate Admin Console, but in order to successfully deploy it, you must also generate Connector tokens that authenticate the Connector with Twingate. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).
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
//			awsNetwork, err := twingate.NewTwingateRemoteNetwork(ctx, "awsNetwork", nil)
//			if err != nil {
//				return err
//			}
//			_, err = twingate.NewTwingateConnector(ctx, "awsConnector", &twingate.TwingateConnectorArgs{
//				RemoteNetworkId:      awsNetwork.ID(),
//				StatusUpdatesEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ```sh
// $ pulumi import twingate:index/twingateConnector:TwingateConnector aws_connector Q29ubmVjdG9yOjI2NzM=
// ```
type TwingateConnector struct {
	pulumi.CustomResourceState

	// Name of the Connector, if not provided one will be generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Remote Network the Connector is attached to.
	RemoteNetworkId pulumi.StringOutput `pulumi:"remoteNetworkId"`
	// Determines whether status notifications are enabled for the Connector. Default is `true`.
	StatusUpdatesEnabled pulumi.BoolOutput `pulumi:"statusUpdatesEnabled"`
}

// NewTwingateConnector registers a new resource with the given unique name, arguments, and options.
func NewTwingateConnector(ctx *pulumi.Context,
	name string, args *TwingateConnectorArgs, opts ...pulumi.ResourceOption) (*TwingateConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RemoteNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'RemoteNetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TwingateConnector
	err := ctx.RegisterResource("twingate:index/twingateConnector:TwingateConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTwingateConnector gets an existing TwingateConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTwingateConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TwingateConnectorState, opts ...pulumi.ResourceOption) (*TwingateConnector, error) {
	var resource TwingateConnector
	err := ctx.ReadResource("twingate:index/twingateConnector:TwingateConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TwingateConnector resources.
type twingateConnectorState struct {
	// Name of the Connector, if not provided one will be generated.
	Name *string `pulumi:"name"`
	// The ID of the Remote Network the Connector is attached to.
	RemoteNetworkId *string `pulumi:"remoteNetworkId"`
	// Determines whether status notifications are enabled for the Connector. Default is `true`.
	StatusUpdatesEnabled *bool `pulumi:"statusUpdatesEnabled"`
}

type TwingateConnectorState struct {
	// Name of the Connector, if not provided one will be generated.
	Name pulumi.StringPtrInput
	// The ID of the Remote Network the Connector is attached to.
	RemoteNetworkId pulumi.StringPtrInput
	// Determines whether status notifications are enabled for the Connector. Default is `true`.
	StatusUpdatesEnabled pulumi.BoolPtrInput
}

func (TwingateConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateConnectorState)(nil)).Elem()
}

type twingateConnectorArgs struct {
	// Name of the Connector, if not provided one will be generated.
	Name *string `pulumi:"name"`
	// The ID of the Remote Network the Connector is attached to.
	RemoteNetworkId string `pulumi:"remoteNetworkId"`
	// Determines whether status notifications are enabled for the Connector. Default is `true`.
	StatusUpdatesEnabled *bool `pulumi:"statusUpdatesEnabled"`
}

// The set of arguments for constructing a TwingateConnector resource.
type TwingateConnectorArgs struct {
	// Name of the Connector, if not provided one will be generated.
	Name pulumi.StringPtrInput
	// The ID of the Remote Network the Connector is attached to.
	RemoteNetworkId pulumi.StringInput
	// Determines whether status notifications are enabled for the Connector. Default is `true`.
	StatusUpdatesEnabled pulumi.BoolPtrInput
}

func (TwingateConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateConnectorArgs)(nil)).Elem()
}

type TwingateConnectorInput interface {
	pulumi.Input

	ToTwingateConnectorOutput() TwingateConnectorOutput
	ToTwingateConnectorOutputWithContext(ctx context.Context) TwingateConnectorOutput
}

func (*TwingateConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateConnector)(nil)).Elem()
}

func (i *TwingateConnector) ToTwingateConnectorOutput() TwingateConnectorOutput {
	return i.ToTwingateConnectorOutputWithContext(context.Background())
}

func (i *TwingateConnector) ToTwingateConnectorOutputWithContext(ctx context.Context) TwingateConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateConnectorOutput)
}

// TwingateConnectorArrayInput is an input type that accepts TwingateConnectorArray and TwingateConnectorArrayOutput values.
// You can construct a concrete instance of `TwingateConnectorArrayInput` via:
//
//	TwingateConnectorArray{ TwingateConnectorArgs{...} }
type TwingateConnectorArrayInput interface {
	pulumi.Input

	ToTwingateConnectorArrayOutput() TwingateConnectorArrayOutput
	ToTwingateConnectorArrayOutputWithContext(context.Context) TwingateConnectorArrayOutput
}

type TwingateConnectorArray []TwingateConnectorInput

func (TwingateConnectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateConnector)(nil)).Elem()
}

func (i TwingateConnectorArray) ToTwingateConnectorArrayOutput() TwingateConnectorArrayOutput {
	return i.ToTwingateConnectorArrayOutputWithContext(context.Background())
}

func (i TwingateConnectorArray) ToTwingateConnectorArrayOutputWithContext(ctx context.Context) TwingateConnectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateConnectorArrayOutput)
}

// TwingateConnectorMapInput is an input type that accepts TwingateConnectorMap and TwingateConnectorMapOutput values.
// You can construct a concrete instance of `TwingateConnectorMapInput` via:
//
//	TwingateConnectorMap{ "key": TwingateConnectorArgs{...} }
type TwingateConnectorMapInput interface {
	pulumi.Input

	ToTwingateConnectorMapOutput() TwingateConnectorMapOutput
	ToTwingateConnectorMapOutputWithContext(context.Context) TwingateConnectorMapOutput
}

type TwingateConnectorMap map[string]TwingateConnectorInput

func (TwingateConnectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateConnector)(nil)).Elem()
}

func (i TwingateConnectorMap) ToTwingateConnectorMapOutput() TwingateConnectorMapOutput {
	return i.ToTwingateConnectorMapOutputWithContext(context.Background())
}

func (i TwingateConnectorMap) ToTwingateConnectorMapOutputWithContext(ctx context.Context) TwingateConnectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateConnectorMapOutput)
}

type TwingateConnectorOutput struct{ *pulumi.OutputState }

func (TwingateConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateConnector)(nil)).Elem()
}

func (o TwingateConnectorOutput) ToTwingateConnectorOutput() TwingateConnectorOutput {
	return o
}

func (o TwingateConnectorOutput) ToTwingateConnectorOutputWithContext(ctx context.Context) TwingateConnectorOutput {
	return o
}

// Name of the Connector, if not provided one will be generated.
func (o TwingateConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Remote Network the Connector is attached to.
func (o TwingateConnectorOutput) RemoteNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateConnector) pulumi.StringOutput { return v.RemoteNetworkId }).(pulumi.StringOutput)
}

// Determines whether status notifications are enabled for the Connector. Default is `true`.
func (o TwingateConnectorOutput) StatusUpdatesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *TwingateConnector) pulumi.BoolOutput { return v.StatusUpdatesEnabled }).(pulumi.BoolOutput)
}

type TwingateConnectorArrayOutput struct{ *pulumi.OutputState }

func (TwingateConnectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateConnector)(nil)).Elem()
}

func (o TwingateConnectorArrayOutput) ToTwingateConnectorArrayOutput() TwingateConnectorArrayOutput {
	return o
}

func (o TwingateConnectorArrayOutput) ToTwingateConnectorArrayOutputWithContext(ctx context.Context) TwingateConnectorArrayOutput {
	return o
}

func (o TwingateConnectorArrayOutput) Index(i pulumi.IntInput) TwingateConnectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TwingateConnector {
		return vs[0].([]*TwingateConnector)[vs[1].(int)]
	}).(TwingateConnectorOutput)
}

type TwingateConnectorMapOutput struct{ *pulumi.OutputState }

func (TwingateConnectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateConnector)(nil)).Elem()
}

func (o TwingateConnectorMapOutput) ToTwingateConnectorMapOutput() TwingateConnectorMapOutput {
	return o
}

func (o TwingateConnectorMapOutput) ToTwingateConnectorMapOutputWithContext(ctx context.Context) TwingateConnectorMapOutput {
	return o
}

func (o TwingateConnectorMapOutput) MapIndex(k pulumi.StringInput) TwingateConnectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TwingateConnector {
		return vs[0].(map[string]*TwingateConnector)[vs[1].(string)]
	}).(TwingateConnectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateConnectorInput)(nil)).Elem(), &TwingateConnector{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateConnectorArrayInput)(nil)).Elem(), TwingateConnectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateConnectorMapInput)(nil)).Elem(), TwingateConnectorMap{})
	pulumi.RegisterOutputType(TwingateConnectorOutput{})
	pulumi.RegisterOutputType(TwingateConnectorArrayOutput{})
	pulumi.RegisterOutputType(TwingateConnectorMapOutput{})
}
