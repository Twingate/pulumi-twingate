# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TwingateDNSFilteringProfileArgs', 'TwingateDNSFilteringProfile']

@pulumi.input_type
class TwingateDNSFilteringProfileArgs:
    def __init__(__self__, *,
                 priority: pulumi.Input[float],
                 allowed_domains: Optional[pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs']] = None,
                 content_categories: Optional[pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs']] = None,
                 denied_domains: Optional[pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs']] = None,
                 fallback_method: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 privacy_categories: Optional[pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs']] = None,
                 security_categories: Optional[pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs']] = None):
        """
        The set of arguments for constructing a TwingateDNSFilteringProfile resource.
        :param pulumi.Input[float] priority: A floating point number representing the profile's priority.
        :param pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs'] allowed_domains: A block with the following attributes.
        :param pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs'] content_categories: A block with the following attributes.
        :param pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs'] denied_domains: A block with the following attributes.
        :param pulumi.Input[str] fallback_method: The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        :param pulumi.Input[str] name: The DNS filtering profile's name.
        :param pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs'] privacy_categories: A block with the following attributes.
        :param pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs'] security_categories: A block with the following attributes.
        """
        pulumi.set(__self__, "priority", priority)
        if allowed_domains is not None:
            pulumi.set(__self__, "allowed_domains", allowed_domains)
        if content_categories is not None:
            pulumi.set(__self__, "content_categories", content_categories)
        if denied_domains is not None:
            pulumi.set(__self__, "denied_domains", denied_domains)
        if fallback_method is not None:
            pulumi.set(__self__, "fallback_method", fallback_method)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if privacy_categories is not None:
            pulumi.set(__self__, "privacy_categories", privacy_categories)
        if security_categories is not None:
            pulumi.set(__self__, "security_categories", security_categories)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[float]:
        """
        A floating point number representing the profile's priority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[float]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="allowedDomains")
    def allowed_domains(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "allowed_domains")

    @allowed_domains.setter
    def allowed_domains(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs']]):
        pulumi.set(self, "allowed_domains", value)

    @property
    @pulumi.getter(name="contentCategories")
    def content_categories(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "content_categories")

    @content_categories.setter
    def content_categories(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs']]):
        pulumi.set(self, "content_categories", value)

    @property
    @pulumi.getter(name="deniedDomains")
    def denied_domains(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "denied_domains")

    @denied_domains.setter
    def denied_domains(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs']]):
        pulumi.set(self, "denied_domains", value)

    @property
    @pulumi.getter(name="fallbackMethod")
    def fallback_method(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        """
        return pulumi.get(self, "fallback_method")

    @fallback_method.setter
    def fallback_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fallback_method", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS filtering profile's name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privacyCategories")
    def privacy_categories(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "privacy_categories")

    @privacy_categories.setter
    def privacy_categories(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs']]):
        pulumi.set(self, "privacy_categories", value)

    @property
    @pulumi.getter(name="securityCategories")
    def security_categories(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "security_categories")

    @security_categories.setter
    def security_categories(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs']]):
        pulumi.set(self, "security_categories", value)


@pulumi.input_type
class _TwingateDNSFilteringProfileState:
    def __init__(__self__, *,
                 allowed_domains: Optional[pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs']] = None,
                 content_categories: Optional[pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs']] = None,
                 denied_domains: Optional[pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs']] = None,
                 fallback_method: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 privacy_categories: Optional[pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs']] = None,
                 security_categories: Optional[pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs']] = None):
        """
        Input properties used for looking up and filtering TwingateDNSFilteringProfile resources.
        :param pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs'] allowed_domains: A block with the following attributes.
        :param pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs'] content_categories: A block with the following attributes.
        :param pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs'] denied_domains: A block with the following attributes.
        :param pulumi.Input[str] fallback_method: The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        :param pulumi.Input[str] name: The DNS filtering profile's name.
        :param pulumi.Input[float] priority: A floating point number representing the profile's priority.
        :param pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs'] privacy_categories: A block with the following attributes.
        :param pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs'] security_categories: A block with the following attributes.
        """
        if allowed_domains is not None:
            pulumi.set(__self__, "allowed_domains", allowed_domains)
        if content_categories is not None:
            pulumi.set(__self__, "content_categories", content_categories)
        if denied_domains is not None:
            pulumi.set(__self__, "denied_domains", denied_domains)
        if fallback_method is not None:
            pulumi.set(__self__, "fallback_method", fallback_method)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if privacy_categories is not None:
            pulumi.set(__self__, "privacy_categories", privacy_categories)
        if security_categories is not None:
            pulumi.set(__self__, "security_categories", security_categories)

    @property
    @pulumi.getter(name="allowedDomains")
    def allowed_domains(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "allowed_domains")

    @allowed_domains.setter
    def allowed_domains(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileAllowedDomainsArgs']]):
        pulumi.set(self, "allowed_domains", value)

    @property
    @pulumi.getter(name="contentCategories")
    def content_categories(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "content_categories")

    @content_categories.setter
    def content_categories(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileContentCategoriesArgs']]):
        pulumi.set(self, "content_categories", value)

    @property
    @pulumi.getter(name="deniedDomains")
    def denied_domains(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "denied_domains")

    @denied_domains.setter
    def denied_domains(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileDeniedDomainsArgs']]):
        pulumi.set(self, "denied_domains", value)

    @property
    @pulumi.getter(name="fallbackMethod")
    def fallback_method(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        """
        return pulumi.get(self, "fallback_method")

    @fallback_method.setter
    def fallback_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fallback_method", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The DNS filtering profile's name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[float]]:
        """
        A floating point number representing the profile's priority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="privacyCategories")
    def privacy_categories(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "privacy_categories")

    @privacy_categories.setter
    def privacy_categories(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfilePrivacyCategoriesArgs']]):
        pulumi.set(self, "privacy_categories", value)

    @property
    @pulumi.getter(name="securityCategories")
    def security_categories(self) -> Optional[pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "security_categories")

    @security_categories.setter
    def security_categories(self, value: Optional[pulumi.Input['TwingateDNSFilteringProfileSecurityCategoriesArgs']]):
        pulumi.set(self, "security_categories", value)


class TwingateDNSFilteringProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_domains: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileAllowedDomainsArgs', 'TwingateDNSFilteringProfileAllowedDomainsArgsDict']]] = None,
                 content_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileContentCategoriesArgs', 'TwingateDNSFilteringProfileContentCategoriesArgsDict']]] = None,
                 denied_domains: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileDeniedDomainsArgs', 'TwingateDNSFilteringProfileDeniedDomainsArgsDict']]] = None,
                 fallback_method: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 privacy_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfilePrivacyCategoriesArgs', 'TwingateDNSFilteringProfilePrivacyCategoriesArgsDict']]] = None,
                 security_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileSecurityCategoriesArgs', 'TwingateDNSFilteringProfileSecurityCategoriesArgsDict']]] = None,
                 __props__=None):
        """
        DNS filtering gives you the ability to control what websites your users can access. DNS filtering is only available on certain plans. For more information, see Twingate's [documentation](https://www.twingate.com/docs/dns-filtering). DNS filtering must be enabled for this resources to work. If DNS filtering isn't enabled, the provider will throw an error.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileAllowedDomainsArgs', 'TwingateDNSFilteringProfileAllowedDomainsArgsDict']] allowed_domains: A block with the following attributes.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileContentCategoriesArgs', 'TwingateDNSFilteringProfileContentCategoriesArgsDict']] content_categories: A block with the following attributes.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileDeniedDomainsArgs', 'TwingateDNSFilteringProfileDeniedDomainsArgsDict']] denied_domains: A block with the following attributes.
        :param pulumi.Input[str] fallback_method: The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        :param pulumi.Input[str] name: The DNS filtering profile's name.
        :param pulumi.Input[float] priority: A floating point number representing the profile's priority.
        :param pulumi.Input[Union['TwingateDNSFilteringProfilePrivacyCategoriesArgs', 'TwingateDNSFilteringProfilePrivacyCategoriesArgsDict']] privacy_categories: A block with the following attributes.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileSecurityCategoriesArgs', 'TwingateDNSFilteringProfileSecurityCategoriesArgsDict']] security_categories: A block with the following attributes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TwingateDNSFilteringProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        DNS filtering gives you the ability to control what websites your users can access. DNS filtering is only available on certain plans. For more information, see Twingate's [documentation](https://www.twingate.com/docs/dns-filtering). DNS filtering must be enabled for this resources to work. If DNS filtering isn't enabled, the provider will throw an error.

        :param str resource_name: The name of the resource.
        :param TwingateDNSFilteringProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TwingateDNSFilteringProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_domains: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileAllowedDomainsArgs', 'TwingateDNSFilteringProfileAllowedDomainsArgsDict']]] = None,
                 content_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileContentCategoriesArgs', 'TwingateDNSFilteringProfileContentCategoriesArgsDict']]] = None,
                 denied_domains: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileDeniedDomainsArgs', 'TwingateDNSFilteringProfileDeniedDomainsArgsDict']]] = None,
                 fallback_method: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 privacy_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfilePrivacyCategoriesArgs', 'TwingateDNSFilteringProfilePrivacyCategoriesArgsDict']]] = None,
                 security_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileSecurityCategoriesArgs', 'TwingateDNSFilteringProfileSecurityCategoriesArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TwingateDNSFilteringProfileArgs.__new__(TwingateDNSFilteringProfileArgs)

            __props__.__dict__["allowed_domains"] = allowed_domains
            __props__.__dict__["content_categories"] = content_categories
            __props__.__dict__["denied_domains"] = denied_domains
            __props__.__dict__["fallback_method"] = fallback_method
            __props__.__dict__["groups"] = groups
            __props__.__dict__["name"] = name
            if priority is None and not opts.urn:
                raise TypeError("Missing required property 'priority'")
            __props__.__dict__["priority"] = priority
            __props__.__dict__["privacy_categories"] = privacy_categories
            __props__.__dict__["security_categories"] = security_categories
        super(TwingateDNSFilteringProfile, __self__).__init__(
            'twingate:index/twingateDNSFilteringProfile:TwingateDNSFilteringProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_domains: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileAllowedDomainsArgs', 'TwingateDNSFilteringProfileAllowedDomainsArgsDict']]] = None,
            content_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileContentCategoriesArgs', 'TwingateDNSFilteringProfileContentCategoriesArgsDict']]] = None,
            denied_domains: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileDeniedDomainsArgs', 'TwingateDNSFilteringProfileDeniedDomainsArgsDict']]] = None,
            fallback_method: Optional[pulumi.Input[str]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            privacy_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfilePrivacyCategoriesArgs', 'TwingateDNSFilteringProfilePrivacyCategoriesArgsDict']]] = None,
            security_categories: Optional[pulumi.Input[Union['TwingateDNSFilteringProfileSecurityCategoriesArgs', 'TwingateDNSFilteringProfileSecurityCategoriesArgsDict']]] = None) -> 'TwingateDNSFilteringProfile':
        """
        Get an existing TwingateDNSFilteringProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileAllowedDomainsArgs', 'TwingateDNSFilteringProfileAllowedDomainsArgsDict']] allowed_domains: A block with the following attributes.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileContentCategoriesArgs', 'TwingateDNSFilteringProfileContentCategoriesArgsDict']] content_categories: A block with the following attributes.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileDeniedDomainsArgs', 'TwingateDNSFilteringProfileDeniedDomainsArgsDict']] denied_domains: A block with the following attributes.
        :param pulumi.Input[str] fallback_method: The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        :param pulumi.Input[str] name: The DNS filtering profile's name.
        :param pulumi.Input[float] priority: A floating point number representing the profile's priority.
        :param pulumi.Input[Union['TwingateDNSFilteringProfilePrivacyCategoriesArgs', 'TwingateDNSFilteringProfilePrivacyCategoriesArgsDict']] privacy_categories: A block with the following attributes.
        :param pulumi.Input[Union['TwingateDNSFilteringProfileSecurityCategoriesArgs', 'TwingateDNSFilteringProfileSecurityCategoriesArgsDict']] security_categories: A block with the following attributes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TwingateDNSFilteringProfileState.__new__(_TwingateDNSFilteringProfileState)

        __props__.__dict__["allowed_domains"] = allowed_domains
        __props__.__dict__["content_categories"] = content_categories
        __props__.__dict__["denied_domains"] = denied_domains
        __props__.__dict__["fallback_method"] = fallback_method
        __props__.__dict__["groups"] = groups
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["privacy_categories"] = privacy_categories
        __props__.__dict__["security_categories"] = security_categories
        return TwingateDNSFilteringProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedDomains")
    def allowed_domains(self) -> pulumi.Output[Optional['outputs.TwingateDNSFilteringProfileAllowedDomains']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "allowed_domains")

    @property
    @pulumi.getter(name="contentCategories")
    def content_categories(self) -> pulumi.Output[Optional['outputs.TwingateDNSFilteringProfileContentCategories']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "content_categories")

    @property
    @pulumi.getter(name="deniedDomains")
    def denied_domains(self) -> pulumi.Output[Optional['outputs.TwingateDNSFilteringProfileDeniedDomains']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "denied_domains")

    @property
    @pulumi.getter(name="fallbackMethod")
    def fallback_method(self) -> pulumi.Output[str]:
        """
        The DNS filtering profile's fallback method. One of "AUTO" or "STRICT". Defaults to "STRICT".
        """
        return pulumi.get(self, "fallback_method")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of group IDs that have this as their DNS filtering profile. Defaults to an empty set.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The DNS filtering profile's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[float]:
        """
        A floating point number representing the profile's priority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="privacyCategories")
    def privacy_categories(self) -> pulumi.Output[Optional['outputs.TwingateDNSFilteringProfilePrivacyCategories']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "privacy_categories")

    @property
    @pulumi.getter(name="securityCategories")
    def security_categories(self) -> pulumi.Output[Optional['outputs.TwingateDNSFilteringProfileSecurityCategories']]:
        """
        A block with the following attributes.
        """
        return pulumi.get(self, "security_categories")

