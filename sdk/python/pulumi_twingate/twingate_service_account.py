# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['TwingateServiceAccountArgs', 'TwingateServiceAccount']

@pulumi.input_type
class TwingateServiceAccountArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TwingateServiceAccount resource.
        :param pulumi.Input[str] name: The name of the Service Account in Twingate
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Service Account in Twingate
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TwingateServiceAccountState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TwingateServiceAccount resources.
        :param pulumi.Input[str] name: The name of the Service Account in Twingate
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Service Account in Twingate
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class TwingateServiceAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_twingate as twingate

        github_actions_prod = twingate.TwingateServiceAccount("githubActionsProd")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Service Account in Twingate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TwingateServiceAccountArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_twingate as twingate

        github_actions_prod = twingate.TwingateServiceAccount("githubActionsProd")
        ```

        :param str resource_name: The name of the resource.
        :param TwingateServiceAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TwingateServiceAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TwingateServiceAccountArgs.__new__(TwingateServiceAccountArgs)

            __props__.__dict__["name"] = name
        super(TwingateServiceAccount, __self__).__init__(
            'twingate:index/twingateServiceAccount:TwingateServiceAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'TwingateServiceAccount':
        """
        Get an existing TwingateServiceAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Service Account in Twingate
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TwingateServiceAccountState.__new__(_TwingateServiceAccountState)

        __props__.__dict__["name"] = name
        return TwingateServiceAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Service Account in Twingate
        """
        return pulumi.get(self, "name")

