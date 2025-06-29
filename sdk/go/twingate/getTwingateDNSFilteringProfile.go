// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package twingate

import (
	"context"
	"reflect"

	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// DNS filtering gives you the ability to control what websites your users can access. DNS filtering is only available on certain plans. For more information, see Twingate's [documentation](https://www.twingate.com/docs/dns-filtering). DNS filtering must be enabled for this data source to work. If DNS filtering isn't enabled, the provider will throw an error.
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
//			_, err := twingate.LookupTwingateDNSFilteringProfile(ctx, &twingate.LookupTwingateDNSFilteringProfileArgs{
//				Id: "<your dns profile's id>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTwingateDNSFilteringProfile(ctx *pulumi.Context, args *LookupTwingateDNSFilteringProfileArgs, opts ...pulumi.InvokeOption) (*LookupTwingateDNSFilteringProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTwingateDNSFilteringProfileResult
	err := ctx.Invoke("twingate:index/getTwingateDNSFilteringProfile:getTwingateDNSFilteringProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTwingateDNSFilteringProfile.
type LookupTwingateDNSFilteringProfileArgs struct {
	// A block with the following attributes.
	AllowedDomains *GetTwingateDNSFilteringProfileAllowedDomains `pulumi:"allowedDomains"`
	// A block with the following attributes.
	ContentCategories *GetTwingateDNSFilteringProfileContentCategories `pulumi:"contentCategories"`
	// A block with the following attributes.
	DeniedDomains *GetTwingateDNSFilteringProfileDeniedDomains `pulumi:"deniedDomains"`
	// The DNS filtering profile's ID.
	Id string `pulumi:"id"`
	// A block with the following attributes.
	PrivacyCategories *GetTwingateDNSFilteringProfilePrivacyCategories `pulumi:"privacyCategories"`
	// A block with the following attributes.
	SecurityCategories *GetTwingateDNSFilteringProfileSecurityCategories `pulumi:"securityCategories"`
}

// A collection of values returned by getTwingateDNSFilteringProfile.
type LookupTwingateDNSFilteringProfileResult struct {
	// A block with the following attributes.
	AllowedDomains *GetTwingateDNSFilteringProfileAllowedDomains `pulumi:"allowedDomains"`
	// A block with the following attributes.
	ContentCategories *GetTwingateDNSFilteringProfileContentCategories `pulumi:"contentCategories"`
	// A block with the following attributes.
	DeniedDomains *GetTwingateDNSFilteringProfileDeniedDomains `pulumi:"deniedDomains"`
	// The DNS filtering profile's fallback method. One of AUTOMATIC or STRICT.
	FallbackMethod string `pulumi:"fallbackMethod"`
	// A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
	Groups []string `pulumi:"groups"`
	// The DNS filtering profile's ID.
	Id string `pulumi:"id"`
	// The DNS filtering profile's name.
	Name string `pulumi:"name"`
	// A floating point number representing the profile's priority.
	Priority float64 `pulumi:"priority"`
	// A block with the following attributes.
	PrivacyCategories *GetTwingateDNSFilteringProfilePrivacyCategories `pulumi:"privacyCategories"`
	// A block with the following attributes.
	SecurityCategories *GetTwingateDNSFilteringProfileSecurityCategories `pulumi:"securityCategories"`
}

func LookupTwingateDNSFilteringProfileOutput(ctx *pulumi.Context, args LookupTwingateDNSFilteringProfileOutputArgs, opts ...pulumi.InvokeOption) LookupTwingateDNSFilteringProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTwingateDNSFilteringProfileResultOutput, error) {
			args := v.(LookupTwingateDNSFilteringProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("twingate:index/getTwingateDNSFilteringProfile:getTwingateDNSFilteringProfile", args, LookupTwingateDNSFilteringProfileResultOutput{}, options).(LookupTwingateDNSFilteringProfileResultOutput), nil
		}).(LookupTwingateDNSFilteringProfileResultOutput)
}

// A collection of arguments for invoking getTwingateDNSFilteringProfile.
type LookupTwingateDNSFilteringProfileOutputArgs struct {
	// A block with the following attributes.
	AllowedDomains GetTwingateDNSFilteringProfileAllowedDomainsPtrInput `pulumi:"allowedDomains"`
	// A block with the following attributes.
	ContentCategories GetTwingateDNSFilteringProfileContentCategoriesPtrInput `pulumi:"contentCategories"`
	// A block with the following attributes.
	DeniedDomains GetTwingateDNSFilteringProfileDeniedDomainsPtrInput `pulumi:"deniedDomains"`
	// The DNS filtering profile's ID.
	Id pulumi.StringInput `pulumi:"id"`
	// A block with the following attributes.
	PrivacyCategories GetTwingateDNSFilteringProfilePrivacyCategoriesPtrInput `pulumi:"privacyCategories"`
	// A block with the following attributes.
	SecurityCategories GetTwingateDNSFilteringProfileSecurityCategoriesPtrInput `pulumi:"securityCategories"`
}

func (LookupTwingateDNSFilteringProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateDNSFilteringProfileArgs)(nil)).Elem()
}

// A collection of values returned by getTwingateDNSFilteringProfile.
type LookupTwingateDNSFilteringProfileResultOutput struct{ *pulumi.OutputState }

func (LookupTwingateDNSFilteringProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTwingateDNSFilteringProfileResult)(nil)).Elem()
}

func (o LookupTwingateDNSFilteringProfileResultOutput) ToLookupTwingateDNSFilteringProfileResultOutput() LookupTwingateDNSFilteringProfileResultOutput {
	return o
}

func (o LookupTwingateDNSFilteringProfileResultOutput) ToLookupTwingateDNSFilteringProfileResultOutputWithContext(ctx context.Context) LookupTwingateDNSFilteringProfileResultOutput {
	return o
}

// A block with the following attributes.
func (o LookupTwingateDNSFilteringProfileResultOutput) AllowedDomains() GetTwingateDNSFilteringProfileAllowedDomainsPtrOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) *GetTwingateDNSFilteringProfileAllowedDomains {
		return v.AllowedDomains
	}).(GetTwingateDNSFilteringProfileAllowedDomainsPtrOutput)
}

// A block with the following attributes.
func (o LookupTwingateDNSFilteringProfileResultOutput) ContentCategories() GetTwingateDNSFilteringProfileContentCategoriesPtrOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) *GetTwingateDNSFilteringProfileContentCategories {
		return v.ContentCategories
	}).(GetTwingateDNSFilteringProfileContentCategoriesPtrOutput)
}

// A block with the following attributes.
func (o LookupTwingateDNSFilteringProfileResultOutput) DeniedDomains() GetTwingateDNSFilteringProfileDeniedDomainsPtrOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) *GetTwingateDNSFilteringProfileDeniedDomains {
		return v.DeniedDomains
	}).(GetTwingateDNSFilteringProfileDeniedDomainsPtrOutput)
}

// The DNS filtering profile's fallback method. One of AUTOMATIC or STRICT.
func (o LookupTwingateDNSFilteringProfileResultOutput) FallbackMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) string { return v.FallbackMethod }).(pulumi.StringOutput)
}

// A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
func (o LookupTwingateDNSFilteringProfileResultOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

// The DNS filtering profile's ID.
func (o LookupTwingateDNSFilteringProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The DNS filtering profile's name.
func (o LookupTwingateDNSFilteringProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// A floating point number representing the profile's priority.
func (o LookupTwingateDNSFilteringProfileResultOutput) Priority() pulumi.Float64Output {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) float64 { return v.Priority }).(pulumi.Float64Output)
}

// A block with the following attributes.
func (o LookupTwingateDNSFilteringProfileResultOutput) PrivacyCategories() GetTwingateDNSFilteringProfilePrivacyCategoriesPtrOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) *GetTwingateDNSFilteringProfilePrivacyCategories {
		return v.PrivacyCategories
	}).(GetTwingateDNSFilteringProfilePrivacyCategoriesPtrOutput)
}

// A block with the following attributes.
func (o LookupTwingateDNSFilteringProfileResultOutput) SecurityCategories() GetTwingateDNSFilteringProfileSecurityCategoriesPtrOutput {
	return o.ApplyT(func(v LookupTwingateDNSFilteringProfileResult) *GetTwingateDNSFilteringProfileSecurityCategories {
		return v.SecurityCategories
	}).(GetTwingateDNSFilteringProfileSecurityCategoriesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTwingateDNSFilteringProfileResultOutput{})
}
