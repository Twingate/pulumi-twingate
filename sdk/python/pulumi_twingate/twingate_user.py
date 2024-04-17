# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TwingateUserArgs', 'TwingateUser']

@pulumi.input_type
class TwingateUserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 first_name: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 send_invite: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TwingateUser resource.
        :param pulumi.Input[str] email: The User's email address
        :param pulumi.Input[str] first_name: The User's first name
        :param pulumi.Input[bool] is_active: Determines whether the User is active or not. Inactive users will be not able to sign in.
        :param pulumi.Input[str] last_name: The User's last name
        :param pulumi.Input[str] role: Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        :param pulumi.Input[bool] send_invite: Determines whether to send an email invitation to the User. True by default.
        """
        pulumi.set(__self__, "email", email)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if send_invite is not None:
            pulumi.set(__self__, "send_invite", send_invite)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        The User's email address
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        The User's first name
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether the User is active or not. Inactive users will be not able to sign in.
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        The User's last name
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="sendInvite")
    def send_invite(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether to send an email invitation to the User. True by default.
        """
        return pulumi.get(self, "send_invite")

    @send_invite.setter
    def send_invite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_invite", value)


@pulumi.input_type
class _TwingateUserState:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 send_invite: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TwingateUser resources.
        :param pulumi.Input[str] email: The User's email address
        :param pulumi.Input[str] first_name: The User's first name
        :param pulumi.Input[bool] is_active: Determines whether the User is active or not. Inactive users will be not able to sign in.
        :param pulumi.Input[str] last_name: The User's last name
        :param pulumi.Input[str] role: Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        :param pulumi.Input[bool] send_invite: Determines whether to send an email invitation to the User. True by default.
        :param pulumi.Input[str] type: Indicates the User's type. Either MANUAL or SYNCED.
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if send_invite is not None:
            pulumi.set(__self__, "send_invite", send_invite)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The User's email address
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        The User's first name
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether the User is active or not. Inactive users will be not able to sign in.
        """
        return pulumi.get(self, "is_active")

    @is_active.setter
    def is_active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_active", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        The User's last name
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="sendInvite")
    def send_invite(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether to send an email invitation to the User. True by default.
        """
        return pulumi.get(self, "send_invite")

    @send_invite.setter
    def send_invite(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "send_invite", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the User's type. Either MANUAL or SYNCED.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class TwingateUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 send_invite: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Users provides different levels of write capabilities across the Twingate Admin Console. For more information, see Twingate's [documentation](https://www.twingate.com/docs/users).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_twingate as twingate

        user = twingate.TwingateUser("user",
            email="sample@company.com",
            first_name="Twin",
            last_name="Gate",
            role="DEVOPS",
            send_invite=True)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The User's email address
        :param pulumi.Input[str] first_name: The User's first name
        :param pulumi.Input[bool] is_active: Determines whether the User is active or not. Inactive users will be not able to sign in.
        :param pulumi.Input[str] last_name: The User's last name
        :param pulumi.Input[str] role: Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        :param pulumi.Input[bool] send_invite: Determines whether to send an email invitation to the User. True by default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TwingateUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Users provides different levels of write capabilities across the Twingate Admin Console. For more information, see Twingate's [documentation](https://www.twingate.com/docs/users).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_twingate as twingate

        user = twingate.TwingateUser("user",
            email="sample@company.com",
            first_name="Twin",
            last_name="Gate",
            role="DEVOPS",
            send_invite=True)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param TwingateUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TwingateUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 is_active: Optional[pulumi.Input[bool]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 send_invite: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TwingateUserArgs.__new__(TwingateUserArgs)

            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["first_name"] = first_name
            __props__.__dict__["is_active"] = is_active
            __props__.__dict__["last_name"] = last_name
            __props__.__dict__["role"] = role
            __props__.__dict__["send_invite"] = send_invite
            __props__.__dict__["type"] = None
        super(TwingateUser, __self__).__init__(
            'twingate:index/twingateUser:TwingateUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email: Optional[pulumi.Input[str]] = None,
            first_name: Optional[pulumi.Input[str]] = None,
            is_active: Optional[pulumi.Input[bool]] = None,
            last_name: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            send_invite: Optional[pulumi.Input[bool]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'TwingateUser':
        """
        Get an existing TwingateUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: The User's email address
        :param pulumi.Input[str] first_name: The User's first name
        :param pulumi.Input[bool] is_active: Determines whether the User is active or not. Inactive users will be not able to sign in.
        :param pulumi.Input[str] last_name: The User's last name
        :param pulumi.Input[str] role: Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        :param pulumi.Input[bool] send_invite: Determines whether to send an email invitation to the User. True by default.
        :param pulumi.Input[str] type: Indicates the User's type. Either MANUAL or SYNCED.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TwingateUserState.__new__(_TwingateUserState)

        __props__.__dict__["email"] = email
        __props__.__dict__["first_name"] = first_name
        __props__.__dict__["is_active"] = is_active
        __props__.__dict__["last_name"] = last_name
        __props__.__dict__["role"] = role
        __props__.__dict__["send_invite"] = send_invite
        __props__.__dict__["type"] = type
        return TwingateUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The User's email address
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[str]:
        """
        The User's first name
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> pulumi.Output[bool]:
        """
        Determines whether the User is active or not. Inactive users will be not able to sign in.
        """
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[str]:
        """
        The User's last name
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Determines the User's role. Either ADMIN, DEVOPS, SUPPORT or MEMBER.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="sendInvite")
    def send_invite(self) -> pulumi.Output[bool]:
        """
        Determines whether to send an email invitation to the User. True by default.
        """
        return pulumi.get(self, "send_invite")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Indicates the User's type. Either MANUAL or SYNCED.
        """
        return pulumi.get(self, "type")

