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

// A Service Key authorizes access to all Resources assigned to a Service Account.
//
// ## Example Usage
//
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
//			githubActionsProd, err := twingate.NewTwingateServiceAccount(ctx, "githubActionsProd", nil)
//			if err != nil {
//				return err
//			}
//			_, err = twingate.NewTwingateServiceAccountKey(ctx, "githubKey", &twingate.TwingateServiceAccountKeyArgs{
//				ServiceAccountId: githubActionsProd.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TwingateServiceAccountKey struct {
	pulumi.CustomResourceState

	// The name of the Service Key
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the Service Account
	ServiceAccountId pulumi.StringOutput `pulumi:"serviceAccountId"`
	// Autogenerated Service Key token. Used to configure a Twingate Client running in headless mode.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewTwingateServiceAccountKey registers a new resource with the given unique name, arguments, and options.
func NewTwingateServiceAccountKey(ctx *pulumi.Context,
	name string, args *TwingateServiceAccountKeyArgs, opts ...pulumi.ResourceOption) (*TwingateServiceAccountKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TwingateServiceAccountKey
	err := ctx.RegisterResource("twingate:index/twingateServiceAccountKey:TwingateServiceAccountKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTwingateServiceAccountKey gets an existing TwingateServiceAccountKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTwingateServiceAccountKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TwingateServiceAccountKeyState, opts ...pulumi.ResourceOption) (*TwingateServiceAccountKey, error) {
	var resource TwingateServiceAccountKey
	err := ctx.ReadResource("twingate:index/twingateServiceAccountKey:TwingateServiceAccountKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TwingateServiceAccountKey resources.
type twingateServiceAccountKeyState struct {
	// The name of the Service Key
	Name *string `pulumi:"name"`
	// The id of the Service Account
	ServiceAccountId *string `pulumi:"serviceAccountId"`
	// Autogenerated Service Key token. Used to configure a Twingate Client running in headless mode.
	Token *string `pulumi:"token"`
}

type TwingateServiceAccountKeyState struct {
	// The name of the Service Key
	Name pulumi.StringPtrInput
	// The id of the Service Account
	ServiceAccountId pulumi.StringPtrInput
	// Autogenerated Service Key token. Used to configure a Twingate Client running in headless mode.
	Token pulumi.StringPtrInput
}

func (TwingateServiceAccountKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateServiceAccountKeyState)(nil)).Elem()
}

type twingateServiceAccountKeyArgs struct {
	// The name of the Service Key
	Name *string `pulumi:"name"`
	// The id of the Service Account
	ServiceAccountId string `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a TwingateServiceAccountKey resource.
type TwingateServiceAccountKeyArgs struct {
	// The name of the Service Key
	Name pulumi.StringPtrInput
	// The id of the Service Account
	ServiceAccountId pulumi.StringInput
}

func (TwingateServiceAccountKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateServiceAccountKeyArgs)(nil)).Elem()
}

type TwingateServiceAccountKeyInput interface {
	pulumi.Input

	ToTwingateServiceAccountKeyOutput() TwingateServiceAccountKeyOutput
	ToTwingateServiceAccountKeyOutputWithContext(ctx context.Context) TwingateServiceAccountKeyOutput
}

func (*TwingateServiceAccountKey) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateServiceAccountKey)(nil)).Elem()
}

func (i *TwingateServiceAccountKey) ToTwingateServiceAccountKeyOutput() TwingateServiceAccountKeyOutput {
	return i.ToTwingateServiceAccountKeyOutputWithContext(context.Background())
}

func (i *TwingateServiceAccountKey) ToTwingateServiceAccountKeyOutputWithContext(ctx context.Context) TwingateServiceAccountKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateServiceAccountKeyOutput)
}

// TwingateServiceAccountKeyArrayInput is an input type that accepts TwingateServiceAccountKeyArray and TwingateServiceAccountKeyArrayOutput values.
// You can construct a concrete instance of `TwingateServiceAccountKeyArrayInput` via:
//
//	TwingateServiceAccountKeyArray{ TwingateServiceAccountKeyArgs{...} }
type TwingateServiceAccountKeyArrayInput interface {
	pulumi.Input

	ToTwingateServiceAccountKeyArrayOutput() TwingateServiceAccountKeyArrayOutput
	ToTwingateServiceAccountKeyArrayOutputWithContext(context.Context) TwingateServiceAccountKeyArrayOutput
}

type TwingateServiceAccountKeyArray []TwingateServiceAccountKeyInput

func (TwingateServiceAccountKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateServiceAccountKey)(nil)).Elem()
}

func (i TwingateServiceAccountKeyArray) ToTwingateServiceAccountKeyArrayOutput() TwingateServiceAccountKeyArrayOutput {
	return i.ToTwingateServiceAccountKeyArrayOutputWithContext(context.Background())
}

func (i TwingateServiceAccountKeyArray) ToTwingateServiceAccountKeyArrayOutputWithContext(ctx context.Context) TwingateServiceAccountKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateServiceAccountKeyArrayOutput)
}

// TwingateServiceAccountKeyMapInput is an input type that accepts TwingateServiceAccountKeyMap and TwingateServiceAccountKeyMapOutput values.
// You can construct a concrete instance of `TwingateServiceAccountKeyMapInput` via:
//
//	TwingateServiceAccountKeyMap{ "key": TwingateServiceAccountKeyArgs{...} }
type TwingateServiceAccountKeyMapInput interface {
	pulumi.Input

	ToTwingateServiceAccountKeyMapOutput() TwingateServiceAccountKeyMapOutput
	ToTwingateServiceAccountKeyMapOutputWithContext(context.Context) TwingateServiceAccountKeyMapOutput
}

type TwingateServiceAccountKeyMap map[string]TwingateServiceAccountKeyInput

func (TwingateServiceAccountKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateServiceAccountKey)(nil)).Elem()
}

func (i TwingateServiceAccountKeyMap) ToTwingateServiceAccountKeyMapOutput() TwingateServiceAccountKeyMapOutput {
	return i.ToTwingateServiceAccountKeyMapOutputWithContext(context.Background())
}

func (i TwingateServiceAccountKeyMap) ToTwingateServiceAccountKeyMapOutputWithContext(ctx context.Context) TwingateServiceAccountKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateServiceAccountKeyMapOutput)
}

type TwingateServiceAccountKeyOutput struct{ *pulumi.OutputState }

func (TwingateServiceAccountKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateServiceAccountKey)(nil)).Elem()
}

func (o TwingateServiceAccountKeyOutput) ToTwingateServiceAccountKeyOutput() TwingateServiceAccountKeyOutput {
	return o
}

func (o TwingateServiceAccountKeyOutput) ToTwingateServiceAccountKeyOutputWithContext(ctx context.Context) TwingateServiceAccountKeyOutput {
	return o
}

// The name of the Service Key
func (o TwingateServiceAccountKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateServiceAccountKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the Service Account
func (o TwingateServiceAccountKeyOutput) ServiceAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateServiceAccountKey) pulumi.StringOutput { return v.ServiceAccountId }).(pulumi.StringOutput)
}

// Autogenerated Service Key token. Used to configure a Twingate Client running in headless mode.
func (o TwingateServiceAccountKeyOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateServiceAccountKey) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type TwingateServiceAccountKeyArrayOutput struct{ *pulumi.OutputState }

func (TwingateServiceAccountKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateServiceAccountKey)(nil)).Elem()
}

func (o TwingateServiceAccountKeyArrayOutput) ToTwingateServiceAccountKeyArrayOutput() TwingateServiceAccountKeyArrayOutput {
	return o
}

func (o TwingateServiceAccountKeyArrayOutput) ToTwingateServiceAccountKeyArrayOutputWithContext(ctx context.Context) TwingateServiceAccountKeyArrayOutput {
	return o
}

func (o TwingateServiceAccountKeyArrayOutput) Index(i pulumi.IntInput) TwingateServiceAccountKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TwingateServiceAccountKey {
		return vs[0].([]*TwingateServiceAccountKey)[vs[1].(int)]
	}).(TwingateServiceAccountKeyOutput)
}

type TwingateServiceAccountKeyMapOutput struct{ *pulumi.OutputState }

func (TwingateServiceAccountKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateServiceAccountKey)(nil)).Elem()
}

func (o TwingateServiceAccountKeyMapOutput) ToTwingateServiceAccountKeyMapOutput() TwingateServiceAccountKeyMapOutput {
	return o
}

func (o TwingateServiceAccountKeyMapOutput) ToTwingateServiceAccountKeyMapOutputWithContext(ctx context.Context) TwingateServiceAccountKeyMapOutput {
	return o
}

func (o TwingateServiceAccountKeyMapOutput) MapIndex(k pulumi.StringInput) TwingateServiceAccountKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TwingateServiceAccountKey {
		return vs[0].(map[string]*TwingateServiceAccountKey)[vs[1].(string)]
	}).(TwingateServiceAccountKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateServiceAccountKeyInput)(nil)).Elem(), &TwingateServiceAccountKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateServiceAccountKeyArrayInput)(nil)).Elem(), TwingateServiceAccountKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateServiceAccountKeyMapInput)(nil)).Elem(), TwingateServiceAccountKeyMap{})
	pulumi.RegisterOutputType(TwingateServiceAccountKeyOutput{})
	pulumi.RegisterOutputType(TwingateServiceAccountKeyArrayOutput{})
	pulumi.RegisterOutputType(TwingateServiceAccountKeyMapOutput{})
}
