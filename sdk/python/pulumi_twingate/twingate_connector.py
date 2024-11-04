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

__all__ = ['TwingateConnectorArgs', 'TwingateConnector']

@pulumi.input_type
class TwingateConnectorArgs:
    def __init__(__self__, *,
                 remote_network_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 status_updates_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TwingateConnector resource.
        :param pulumi.Input[str] remote_network_id: The ID of the Remote Network the Connector is attached to.
        :param pulumi.Input[str] name: Name of the Connector, if not provided one will be generated.
        :param pulumi.Input[bool] status_updates_enabled: Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        pulumi.set(__self__, "remote_network_id", remote_network_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status_updates_enabled is not None:
            pulumi.set(__self__, "status_updates_enabled", status_updates_enabled)

    @property
    @pulumi.getter(name="remoteNetworkId")
    def remote_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the Remote Network the Connector is attached to.
        """
        return pulumi.get(self, "remote_network_id")

    @remote_network_id.setter
    def remote_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "remote_network_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Connector, if not provided one will be generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="statusUpdatesEnabled")
    def status_updates_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        return pulumi.get(self, "status_updates_enabled")

    @status_updates_enabled.setter
    def status_updates_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status_updates_enabled", value)


@pulumi.input_type
class _TwingateConnectorState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_network_id: Optional[pulumi.Input[str]] = None,
                 status_updates_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering TwingateConnector resources.
        :param pulumi.Input[str] name: Name of the Connector, if not provided one will be generated.
        :param pulumi.Input[str] remote_network_id: The ID of the Remote Network the Connector is attached to.
        :param pulumi.Input[bool] status_updates_enabled: Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if remote_network_id is not None:
            pulumi.set(__self__, "remote_network_id", remote_network_id)
        if status_updates_enabled is not None:
            pulumi.set(__self__, "status_updates_enabled", status_updates_enabled)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Connector, if not provided one will be generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="remoteNetworkId")
    def remote_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Remote Network the Connector is attached to.
        """
        return pulumi.get(self, "remote_network_id")

    @remote_network_id.setter
    def remote_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_network_id", value)

    @property
    @pulumi.getter(name="statusUpdatesEnabled")
    def status_updates_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        return pulumi.get(self, "status_updates_enabled")

    @status_updates_enabled.setter
    def status_updates_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status_updates_enabled", value)


class TwingateConnector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_network_id: Optional[pulumi.Input[str]] = None,
                 status_updates_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Connectors provide connectivity to Remote Networks. This resource type will create the Connector in the Twingate Admin Console, but in order to successfully deploy it, you must also generate Connector tokens that authenticate the Connector with Twingate. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_twingate as twingate

        aws_network = twingate.TwingateRemoteNetwork("awsNetwork")
        aws_connector = twingate.TwingateConnector("awsConnector",
            remote_network_id=aws_network.id,
            status_updates_enabled=True)
        ```

        ## Import

        ```sh
        $ pulumi import twingate:index/twingateConnector:TwingateConnector aws_connector Q29ubmVjdG9yOjI2NzM=
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the Connector, if not provided one will be generated.
        :param pulumi.Input[str] remote_network_id: The ID of the Remote Network the Connector is attached to.
        :param pulumi.Input[bool] status_updates_enabled: Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TwingateConnectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Connectors provide connectivity to Remote Networks. This resource type will create the Connector in the Twingate Admin Console, but in order to successfully deploy it, you must also generate Connector tokens that authenticate the Connector with Twingate. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_twingate as twingate

        aws_network = twingate.TwingateRemoteNetwork("awsNetwork")
        aws_connector = twingate.TwingateConnector("awsConnector",
            remote_network_id=aws_network.id,
            status_updates_enabled=True)
        ```

        ## Import

        ```sh
        $ pulumi import twingate:index/twingateConnector:TwingateConnector aws_connector Q29ubmVjdG9yOjI2NzM=
        ```

        :param str resource_name: The name of the resource.
        :param TwingateConnectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TwingateConnectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 remote_network_id: Optional[pulumi.Input[str]] = None,
                 status_updates_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TwingateConnectorArgs.__new__(TwingateConnectorArgs)

            __props__.__dict__["name"] = name
            if remote_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'remote_network_id'")
            __props__.__dict__["remote_network_id"] = remote_network_id
            __props__.__dict__["status_updates_enabled"] = status_updates_enabled
        super(TwingateConnector, __self__).__init__(
            'twingate:index/twingateConnector:TwingateConnector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            remote_network_id: Optional[pulumi.Input[str]] = None,
            status_updates_enabled: Optional[pulumi.Input[bool]] = None) -> 'TwingateConnector':
        """
        Get an existing TwingateConnector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the Connector, if not provided one will be generated.
        :param pulumi.Input[str] remote_network_id: The ID of the Remote Network the Connector is attached to.
        :param pulumi.Input[bool] status_updates_enabled: Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TwingateConnectorState.__new__(_TwingateConnectorState)

        __props__.__dict__["name"] = name
        __props__.__dict__["remote_network_id"] = remote_network_id
        __props__.__dict__["status_updates_enabled"] = status_updates_enabled
        return TwingateConnector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Connector, if not provided one will be generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="remoteNetworkId")
    def remote_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the Remote Network the Connector is attached to.
        """
        return pulumi.get(self, "remote_network_id")

    @property
    @pulumi.getter(name="statusUpdatesEnabled")
    def status_updates_enabled(self) -> pulumi.Output[bool]:
        """
        Determines whether status notifications are enabled for the Connector. Default is `true`.
        """
        return pulumi.get(self, "status_updates_enabled")

