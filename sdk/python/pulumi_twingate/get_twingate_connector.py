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

__all__ = [
    'GetTwingateConnectorResult',
    'AwaitableGetTwingateConnectorResult',
    'get_twingate_connector',
    'get_twingate_connector_output',
]

@pulumi.output_type
class GetTwingateConnectorResult:
    """
    A collection of values returned by getTwingateConnector.
    """
    def __init__(__self__, hostname=None, id=None, name=None, private_ips=None, public_ip=None, remote_network_id=None, state=None, status_updates_enabled=None, version=None):
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_ips and not isinstance(private_ips, list):
            raise TypeError("Expected argument 'private_ips' to be a list")
        pulumi.set(__self__, "private_ips", private_ips)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if remote_network_id and not isinstance(remote_network_id, str):
            raise TypeError("Expected argument 'remote_network_id' to be a str")
        pulumi.set(__self__, "remote_network_id", remote_network_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status_updates_enabled and not isinstance(status_updates_enabled, bool):
            raise TypeError("Expected argument 'status_updates_enabled' to be a bool")
        pulumi.set(__self__, "status_updates_enabled", status_updates_enabled)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        The hostname of the machine hosting the Connector.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Connector.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateIps")
    def private_ips(self) -> Sequence[str]:
        """
        The Connector's private IP addresses.
        """
        return pulumi.get(self, "private_ips")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        The Connector's public IP address.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="remoteNetworkId")
    def remote_network_id(self) -> str:
        """
        The ID of the Remote Network the Connector is attached to.
        """
        return pulumi.get(self, "remote_network_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The Connector's state. One of `ALIVE`, `DEAD_NO_HEARTBEAT`, `DEAD_HEARTBEAT_TOO_OLD` or `DEAD_NO_RELAYS`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusUpdatesEnabled")
    def status_updates_enabled(self) -> bool:
        """
        Determines whether status notifications are enabled for the Connector.
        """
        return pulumi.get(self, "status_updates_enabled")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The Connector's version.
        """
        return pulumi.get(self, "version")


class AwaitableGetTwingateConnectorResult(GetTwingateConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTwingateConnectorResult(
            hostname=self.hostname,
            id=self.id,
            name=self.name,
            private_ips=self.private_ips,
            public_ip=self.public_ip,
            remote_network_id=self.remote_network_id,
            state=self.state,
            status_updates_enabled=self.status_updates_enabled,
            version=self.version)


def get_twingate_connector(id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTwingateConnectorResult:
    """
    Connectors provide connectivity to Remote Networks. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_twingate as twingate

    foo = twingate.get_twingate_connector(id="<your connector's id>")
    ```


    :param str id: The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('twingate:index/getTwingateConnector:getTwingateConnector', __args__, opts=opts, typ=GetTwingateConnectorResult).value

    return AwaitableGetTwingateConnectorResult(
        hostname=pulumi.get(__ret__, 'hostname'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        private_ips=pulumi.get(__ret__, 'private_ips'),
        public_ip=pulumi.get(__ret__, 'public_ip'),
        remote_network_id=pulumi.get(__ret__, 'remote_network_id'),
        state=pulumi.get(__ret__, 'state'),
        status_updates_enabled=pulumi.get(__ret__, 'status_updates_enabled'),
        version=pulumi.get(__ret__, 'version'))
def get_twingate_connector_output(id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTwingateConnectorResult]:
    """
    Connectors provide connectivity to Remote Networks. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_twingate as twingate

    foo = twingate.get_twingate_connector(id="<your connector's id>")
    ```


    :param str id: The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('twingate:index/getTwingateConnector:getTwingateConnector', __args__, opts=opts, typ=GetTwingateConnectorResult)
    return __ret__.apply(lambda __response__: GetTwingateConnectorResult(
        hostname=pulumi.get(__response__, 'hostname'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        private_ips=pulumi.get(__response__, 'private_ips'),
        public_ip=pulumi.get(__response__, 'public_ip'),
        remote_network_id=pulumi.get(__response__, 'remote_network_id'),
        state=pulumi.get(__response__, 'state'),
        status_updates_enabled=pulumi.get(__response__, 'status_updates_enabled'),
        version=pulumi.get(__response__, 'version')))
