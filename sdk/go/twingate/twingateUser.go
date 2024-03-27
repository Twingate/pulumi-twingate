// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"errors"
	"github.com/Twingate/pulumi-twingate/sdk/v2/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Users provides different levels of write capabilities across the Twingate Admin Console. For more information, see Twingate's [documentation](https://www.twingate.com/docs/users).
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/Twingate/pulumi-twingate/sdk/v2/go/twingate"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := twingate.NewTwingateUser(ctx, "user", &twingate.TwingateUserArgs{
//				Email:      pulumi.String("sample@company.com"),
//				FirstName:  pulumi.String("Twin"),
//				LastName:   pulumi.String("Gate"),
//				Role:       pulumi.String("DEVOPS"),
//				SendInvite: pulumi.Bool(true),
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
type TwingateUser struct {
	pulumi.CustomResourceState

	// The User's email address
	Email pulumi.StringOutput `pulumi:"email"`
	// The User's first name
	FirstName pulumi.StringOutput `pulumi:"firstName"`
	// Determines whether the User is active or not. Inactive users will be not able to sign in.
	IsActive pulumi.BoolOutput `pulumi:"isActive"`
	// The User's last name
	LastName pulumi.StringOutput `pulumi:"lastName"`
	// Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
	Role pulumi.StringOutput `pulumi:"role"`
	// Determines whether to send an email invitation to the User. True by default.
	SendInvite pulumi.BoolOutput `pulumi:"sendInvite"`
	// Indicates the User's type. Either MANUAL or SYNCED.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTwingateUser registers a new resource with the given unique name, arguments, and options.
func NewTwingateUser(ctx *pulumi.Context,
	name string, args *TwingateUserArgs, opts ...pulumi.ResourceOption) (*TwingateUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TwingateUser
	err := ctx.RegisterResource("twingate:index/twingateUser:TwingateUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTwingateUser gets an existing TwingateUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTwingateUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TwingateUserState, opts ...pulumi.ResourceOption) (*TwingateUser, error) {
	var resource TwingateUser
	err := ctx.ReadResource("twingate:index/twingateUser:TwingateUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TwingateUser resources.
type twingateUserState struct {
	// The User's email address
	Email *string `pulumi:"email"`
	// The User's first name
	FirstName *string `pulumi:"firstName"`
	// Determines whether the User is active or not. Inactive users will be not able to sign in.
	IsActive *bool `pulumi:"isActive"`
	// The User's last name
	LastName *string `pulumi:"lastName"`
	// Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
	Role *string `pulumi:"role"`
	// Determines whether to send an email invitation to the User. True by default.
	SendInvite *bool `pulumi:"sendInvite"`
	// Indicates the User's type. Either MANUAL or SYNCED.
	Type *string `pulumi:"type"`
}

type TwingateUserState struct {
	// The User's email address
	Email pulumi.StringPtrInput
	// The User's first name
	FirstName pulumi.StringPtrInput
	// Determines whether the User is active or not. Inactive users will be not able to sign in.
	IsActive pulumi.BoolPtrInput
	// The User's last name
	LastName pulumi.StringPtrInput
	// Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
	Role pulumi.StringPtrInput
	// Determines whether to send an email invitation to the User. True by default.
	SendInvite pulumi.BoolPtrInput
	// Indicates the User's type. Either MANUAL or SYNCED.
	Type pulumi.StringPtrInput
}

func (TwingateUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateUserState)(nil)).Elem()
}

type twingateUserArgs struct {
	// The User's email address
	Email string `pulumi:"email"`
	// The User's first name
	FirstName *string `pulumi:"firstName"`
	// Determines whether the User is active or not. Inactive users will be not able to sign in.
	IsActive *bool `pulumi:"isActive"`
	// The User's last name
	LastName *string `pulumi:"lastName"`
	// Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
	Role *string `pulumi:"role"`
	// Determines whether to send an email invitation to the User. True by default.
	SendInvite *bool `pulumi:"sendInvite"`
}

// The set of arguments for constructing a TwingateUser resource.
type TwingateUserArgs struct {
	// The User's email address
	Email pulumi.StringInput
	// The User's first name
	FirstName pulumi.StringPtrInput
	// Determines whether the User is active or not. Inactive users will be not able to sign in.
	IsActive pulumi.BoolPtrInput
	// The User's last name
	LastName pulumi.StringPtrInput
	// Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
	Role pulumi.StringPtrInput
	// Determines whether to send an email invitation to the User. True by default.
	SendInvite pulumi.BoolPtrInput
}

func (TwingateUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*twingateUserArgs)(nil)).Elem()
}

type TwingateUserInput interface {
	pulumi.Input

	ToTwingateUserOutput() TwingateUserOutput
	ToTwingateUserOutputWithContext(ctx context.Context) TwingateUserOutput
}

func (*TwingateUser) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateUser)(nil)).Elem()
}

func (i *TwingateUser) ToTwingateUserOutput() TwingateUserOutput {
	return i.ToTwingateUserOutputWithContext(context.Background())
}

func (i *TwingateUser) ToTwingateUserOutputWithContext(ctx context.Context) TwingateUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateUserOutput)
}

// TwingateUserArrayInput is an input type that accepts TwingateUserArray and TwingateUserArrayOutput values.
// You can construct a concrete instance of `TwingateUserArrayInput` via:
//
//	TwingateUserArray{ TwingateUserArgs{...} }
type TwingateUserArrayInput interface {
	pulumi.Input

	ToTwingateUserArrayOutput() TwingateUserArrayOutput
	ToTwingateUserArrayOutputWithContext(context.Context) TwingateUserArrayOutput
}

type TwingateUserArray []TwingateUserInput

func (TwingateUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateUser)(nil)).Elem()
}

func (i TwingateUserArray) ToTwingateUserArrayOutput() TwingateUserArrayOutput {
	return i.ToTwingateUserArrayOutputWithContext(context.Background())
}

func (i TwingateUserArray) ToTwingateUserArrayOutputWithContext(ctx context.Context) TwingateUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateUserArrayOutput)
}

// TwingateUserMapInput is an input type that accepts TwingateUserMap and TwingateUserMapOutput values.
// You can construct a concrete instance of `TwingateUserMapInput` via:
//
//	TwingateUserMap{ "key": TwingateUserArgs{...} }
type TwingateUserMapInput interface {
	pulumi.Input

	ToTwingateUserMapOutput() TwingateUserMapOutput
	ToTwingateUserMapOutputWithContext(context.Context) TwingateUserMapOutput
}

type TwingateUserMap map[string]TwingateUserInput

func (TwingateUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateUser)(nil)).Elem()
}

func (i TwingateUserMap) ToTwingateUserMapOutput() TwingateUserMapOutput {
	return i.ToTwingateUserMapOutputWithContext(context.Background())
}

func (i TwingateUserMap) ToTwingateUserMapOutputWithContext(ctx context.Context) TwingateUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwingateUserMapOutput)
}

type TwingateUserOutput struct{ *pulumi.OutputState }

func (TwingateUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwingateUser)(nil)).Elem()
}

func (o TwingateUserOutput) ToTwingateUserOutput() TwingateUserOutput {
	return o
}

func (o TwingateUserOutput) ToTwingateUserOutputWithContext(ctx context.Context) TwingateUserOutput {
	return o
}

// The User's email address
func (o TwingateUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The User's first name
func (o TwingateUserOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.StringOutput { return v.FirstName }).(pulumi.StringOutput)
}

// Determines whether the User is active or not. Inactive users will be not able to sign in.
func (o TwingateUserOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.BoolOutput { return v.IsActive }).(pulumi.BoolOutput)
}

// The User's last name
func (o TwingateUserOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.StringOutput { return v.LastName }).(pulumi.StringOutput)
}

// Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
func (o TwingateUserOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Determines whether to send an email invitation to the User. True by default.
func (o TwingateUserOutput) SendInvite() pulumi.BoolOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.BoolOutput { return v.SendInvite }).(pulumi.BoolOutput)
}

// Indicates the User's type. Either MANUAL or SYNCED.
func (o TwingateUserOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TwingateUser) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type TwingateUserArrayOutput struct{ *pulumi.OutputState }

func (TwingateUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TwingateUser)(nil)).Elem()
}

func (o TwingateUserArrayOutput) ToTwingateUserArrayOutput() TwingateUserArrayOutput {
	return o
}

func (o TwingateUserArrayOutput) ToTwingateUserArrayOutputWithContext(ctx context.Context) TwingateUserArrayOutput {
	return o
}

func (o TwingateUserArrayOutput) Index(i pulumi.IntInput) TwingateUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TwingateUser {
		return vs[0].([]*TwingateUser)[vs[1].(int)]
	}).(TwingateUserOutput)
}

type TwingateUserMapOutput struct{ *pulumi.OutputState }

func (TwingateUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TwingateUser)(nil)).Elem()
}

func (o TwingateUserMapOutput) ToTwingateUserMapOutput() TwingateUserMapOutput {
	return o
}

func (o TwingateUserMapOutput) ToTwingateUserMapOutputWithContext(ctx context.Context) TwingateUserMapOutput {
	return o
}

func (o TwingateUserMapOutput) MapIndex(k pulumi.StringInput) TwingateUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TwingateUser {
		return vs[0].(map[string]*TwingateUser)[vs[1].(string)]
	}).(TwingateUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateUserInput)(nil)).Elem(), &TwingateUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateUserArrayInput)(nil)).Elem(), TwingateUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TwingateUserMapInput)(nil)).Elem(), TwingateUserMap{})
	pulumi.RegisterOutputType(TwingateUserOutput{})
	pulumi.RegisterOutputType(TwingateUserArrayOutput{})
	pulumi.RegisterOutputType(TwingateUserMapOutput{})
}
