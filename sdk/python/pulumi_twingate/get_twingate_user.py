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
    'GetTwingateUserResult',
    'AwaitableGetTwingateUserResult',
    'get_twingate_user',
    'get_twingate_user_output',
]

@pulumi.output_type
class GetTwingateUserResult:
    """
    A collection of values returned by getTwingateUser.
    """
    def __init__(__self__, email=None, first_name=None, id=None, last_name=None, role=None, type=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        pulumi.set(__self__, "first_name", first_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        pulumi.set(__self__, "last_name", last_name)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email address of the User
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> str:
        """
        The first name of the User
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the User. The ID for the User can be obtained from the Admin API or the URL string in the Admin Console.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> str:
        """
        The last name of the User
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Indicates the User's role. Either ADMIN, DEVOPS, SUPPORT, or MEMBER
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Indicates the User's type. Either MANUAL or SYNCED.
        """
        return pulumi.get(self, "type")


class AwaitableGetTwingateUserResult(GetTwingateUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTwingateUserResult(
            email=self.email,
            first_name=self.first_name,
            id=self.id,
            last_name=self.last_name,
            role=self.role,
            type=self.type)


def get_twingate_user(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTwingateUserResult:
    """
    Users in Twingate can be given access to Twingate Resources and may either be added manually or automatically synchronized with a 3rd party identity provider. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/users).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_twingate as twingate

    foo = twingate.get_twingate_user(id="<your user's id>")
    ```


    :param str id: The ID of the User. The ID for the User can be obtained from the Admin API or the URL string in the Admin Console.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('twingate:index/getTwingateUser:getTwingateUser', __args__, opts=opts, typ=GetTwingateUserResult).value

    return AwaitableGetTwingateUserResult(
        email=pulumi.get(__ret__, 'email'),
        first_name=pulumi.get(__ret__, 'first_name'),
        id=pulumi.get(__ret__, 'id'),
        last_name=pulumi.get(__ret__, 'last_name'),
        role=pulumi.get(__ret__, 'role'),
        type=pulumi.get(__ret__, 'type'))
def get_twingate_user_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTwingateUserResult]:
    """
    Users in Twingate can be given access to Twingate Resources and may either be added manually or automatically synchronized with a 3rd party identity provider. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/users).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_twingate as twingate

    foo = twingate.get_twingate_user(id="<your user's id>")
    ```


    :param str id: The ID of the User. The ID for the User can be obtained from the Admin API or the URL string in the Admin Console.
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('twingate:index/getTwingateUser:getTwingateUser', __args__, opts=opts, typ=GetTwingateUserResult)
    return __ret__.apply(lambda __response__: GetTwingateUserResult(
        email=pulumi.get(__response__, 'email'),
        first_name=pulumi.get(__response__, 'first_name'),
        id=pulumi.get(__response__, 'id'),
        last_name=pulumi.get(__response__, 'last_name'),
        role=pulumi.get(__response__, 'role'),
        type=pulumi.get(__response__, 'type')))
