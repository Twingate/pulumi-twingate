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
from . import outputs

__all__ = [
    'GetTwingateRemoteNetworksResult',
    'AwaitableGetTwingateRemoteNetworksResult',
    'get_twingate_remote_networks',
    'get_twingate_remote_networks_output',
]

@pulumi.output_type
class GetTwingateRemoteNetworksResult:
    """
    A collection of values returned by getTwingateRemoteNetworks.
    """
    def __init__(__self__, id=None, name=None, name_contains=None, name_exclude=None, name_prefix=None, name_regexp=None, name_suffix=None, remote_networks=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_contains and not isinstance(name_contains, str):
            raise TypeError("Expected argument 'name_contains' to be a str")
        pulumi.set(__self__, "name_contains", name_contains)
        if name_exclude and not isinstance(name_exclude, str):
            raise TypeError("Expected argument 'name_exclude' to be a str")
        pulumi.set(__self__, "name_exclude", name_exclude)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)
        if name_regexp and not isinstance(name_regexp, str):
            raise TypeError("Expected argument 'name_regexp' to be a str")
        pulumi.set(__self__, "name_regexp", name_regexp)
        if name_suffix and not isinstance(name_suffix, str):
            raise TypeError("Expected argument 'name_suffix' to be a str")
        pulumi.set(__self__, "name_suffix", name_suffix)
        if remote_networks and not isinstance(remote_networks, list):
            raise TypeError("Expected argument 'remote_networks' to be a list")
        pulumi.set(__self__, "remote_networks", remote_networks)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameContains")
    def name_contains(self) -> Optional[str]:
        """
        Match when the value exist in the name of the remote network.
        """
        return pulumi.get(self, "name_contains")

    @property
    @pulumi.getter(name="nameExclude")
    def name_exclude(self) -> Optional[str]:
        """
        Match when the exact value does not exist in the name of the remote network.
        """
        return pulumi.get(self, "name_exclude")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[str]:
        """
        The name of the remote network must start with the value.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="nameRegexp")
    def name_regexp(self) -> Optional[str]:
        """
        The regular expression match of the name of the remote network.
        """
        return pulumi.get(self, "name_regexp")

    @property
    @pulumi.getter(name="nameSuffix")
    def name_suffix(self) -> Optional[str]:
        """
        The name of the remote network must end with the value.
        """
        return pulumi.get(self, "name_suffix")

    @property
    @pulumi.getter(name="remoteNetworks")
    def remote_networks(self) -> Sequence['outputs.GetTwingateRemoteNetworksRemoteNetworkResult']:
        """
        List of Remote Networks
        """
        return pulumi.get(self, "remote_networks")


class AwaitableGetTwingateRemoteNetworksResult(GetTwingateRemoteNetworksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTwingateRemoteNetworksResult(
            id=self.id,
            name=self.name,
            name_contains=self.name_contains,
            name_exclude=self.name_exclude,
            name_prefix=self.name_prefix,
            name_regexp=self.name_regexp,
            name_suffix=self.name_suffix,
            remote_networks=self.remote_networks)


def get_twingate_remote_networks(name: Optional[str] = None,
                                 name_contains: Optional[str] = None,
                                 name_exclude: Optional[str] = None,
                                 name_prefix: Optional[str] = None,
                                 name_regexp: Optional[str] = None,
                                 name_suffix: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTwingateRemoteNetworksResult:
    """
    A Remote Network represents a single private network in Twingate that can have one or more Connectors and Resources assigned to it. You must create a Remote Network before creating Resources and Connectors that belong to it. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/remote-networks).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_twingate as twingate

    all = twingate.get_twingate_remote_networks(name="<your network's name>")
    ```


    :param str name: Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
    :param str name_contains: Match when the value exist in the name of the remote network.
    :param str name_exclude: Match when the exact value does not exist in the name of the remote network.
    :param str name_prefix: The name of the remote network must start with the value.
    :param str name_regexp: The regular expression match of the name of the remote network.
    :param str name_suffix: The name of the remote network must end with the value.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['nameContains'] = name_contains
    __args__['nameExclude'] = name_exclude
    __args__['namePrefix'] = name_prefix
    __args__['nameRegexp'] = name_regexp
    __args__['nameSuffix'] = name_suffix
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('twingate:index/getTwingateRemoteNetworks:getTwingateRemoteNetworks', __args__, opts=opts, typ=GetTwingateRemoteNetworksResult).value

    return AwaitableGetTwingateRemoteNetworksResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_contains=pulumi.get(__ret__, 'name_contains'),
        name_exclude=pulumi.get(__ret__, 'name_exclude'),
        name_prefix=pulumi.get(__ret__, 'name_prefix'),
        name_regexp=pulumi.get(__ret__, 'name_regexp'),
        name_suffix=pulumi.get(__ret__, 'name_suffix'),
        remote_networks=pulumi.get(__ret__, 'remote_networks'))
def get_twingate_remote_networks_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                        name_contains: Optional[pulumi.Input[Optional[str]]] = None,
                                        name_exclude: Optional[pulumi.Input[Optional[str]]] = None,
                                        name_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                                        name_regexp: Optional[pulumi.Input[Optional[str]]] = None,
                                        name_suffix: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTwingateRemoteNetworksResult]:
    """
    A Remote Network represents a single private network in Twingate that can have one or more Connectors and Resources assigned to it. You must create a Remote Network before creating Resources and Connectors that belong to it. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/remote-networks).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_twingate as twingate

    all = twingate.get_twingate_remote_networks(name="<your network's name>")
    ```


    :param str name: Returns only remote networks that exactly match this name. If no options are passed it will return all remote networks. Only one option can be used at a time.
    :param str name_contains: Match when the value exist in the name of the remote network.
    :param str name_exclude: Match when the exact value does not exist in the name of the remote network.
    :param str name_prefix: The name of the remote network must start with the value.
    :param str name_regexp: The regular expression match of the name of the remote network.
    :param str name_suffix: The name of the remote network must end with the value.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['nameContains'] = name_contains
    __args__['nameExclude'] = name_exclude
    __args__['namePrefix'] = name_prefix
    __args__['nameRegexp'] = name_regexp
    __args__['nameSuffix'] = name_suffix
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('twingate:index/getTwingateRemoteNetworks:getTwingateRemoteNetworks', __args__, opts=opts, typ=GetTwingateRemoteNetworksResult)
    return __ret__.apply(lambda __response__: GetTwingateRemoteNetworksResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        name_contains=pulumi.get(__response__, 'name_contains'),
        name_exclude=pulumi.get(__response__, 'name_exclude'),
        name_prefix=pulumi.get(__response__, 'name_prefix'),
        name_regexp=pulumi.get(__response__, 'name_regexp'),
        name_suffix=pulumi.get(__response__, 'name_suffix'),
        remote_networks=pulumi.get(__response__, 'remote_networks')))
