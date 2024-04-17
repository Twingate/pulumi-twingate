// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetTwingateConnectorsConnector {
    /**
     * The ID of the Connector.
     */
    id: string;
    /**
     * The Name of the Connector.
     */
    name: string;
    /**
     * The ID of the Remote Network attached to the Connector.
     */
    remoteNetworkId: string;
    /**
     * Determines whether status notifications are enabled for the Connector.
     */
    statusUpdatesEnabled: boolean;
}

export interface GetTwingateGroupsGroup {
    /**
     * The ID of the Group
     */
    id: string;
    /**
     * Indicates if the Group is active
     */
    isActive: boolean;
    /**
     * The name of the Group
     */
    name: string;
    /**
     * The Security Policy assigned to the Group.
     */
    securityPolicyId: string;
    /**
     * The type of the Group
     */
    type: string;
}

export interface GetTwingateRemoteNetworksRemoteNetwork {
    /**
     * The ID of the Remote Network.
     */
    id: string;
    /**
     * The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE*CLOUD, ON*PREMISE, OTHER.
     */
    location: string;
    /**
     * The name of the Remote Network.
     */
    name?: string;
}

export interface GetTwingateResourceProtocols {
    /**
     * Whether to allow ICMP (ping) traffic
     */
    allowIcmp: boolean;
    tcp?: outputs.GetTwingateResourceProtocolsTcp;
    udp?: outputs.GetTwingateResourceProtocolsUdp;
}

export interface GetTwingateResourceProtocolsTcp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports: string[];
}

export interface GetTwingateResourceProtocolsUdp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports: string[];
}

export interface GetTwingateResourcesResource {
    /**
     * The Resource's IP/CIDR or FQDN/DNS zone
     */
    address: string;
    /**
     * The id of the Resource
     */
    id: string;
    /**
     * The name of the Resource
     */
    name: string;
    /**
     * Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no restriction, and all protocols and ports are allowed.
     */
    protocols: outputs.GetTwingateResourcesResourceProtocols;
    /**
     * Remote Network ID where the Resource lives
     */
    remoteNetworkId: string;
}

export interface GetTwingateResourcesResourceProtocols {
    /**
     * Whether to allow ICMP (ping) traffic
     */
    allowIcmp: boolean;
    tcp: outputs.GetTwingateResourcesResourceProtocolsTcp;
    udp: outputs.GetTwingateResourcesResourceProtocolsUdp;
}

export interface GetTwingateResourcesResourceProtocolsTcp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports: string[];
}

export interface GetTwingateResourcesResourceProtocolsUdp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports: string[];
}

export interface GetTwingateSecurityPoliciesSecurityPolicy {
    /**
     * Return a matching Security Policy by its ID. The ID for the Security Policy can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: string;
    /**
     * Return a Security Policy that exactly matches this name.
     */
    name: string;
}

export interface GetTwingateServiceAccountsServiceAccount {
    /**
     * ID of the Service Account resource
     */
    id: string;
    /**
     * List of twingate*service*account_key IDs that are assigned to the Service Account.
     */
    keyIds: string[];
    /**
     * Name of the Service Account
     */
    name: string;
    /**
     * List of twingate.TwingateResource IDs that the Service Account is assigned to.
     */
    resourceIds: string[];
}

export interface GetTwingateUsersUser {
    /**
     * The email address of the User
     */
    email: string;
    /**
     * The first name of the User
     */
    firstName: string;
    /**
     * The ID of the User
     */
    id: string;
    /**
     * The last name of the User
     */
    lastName: string;
    /**
     * Indicates the User's role. Either ADMIN, DEVOPS, SUPPORT, or MEMBER.
     */
    role: string;
    /**
     * Indicates the User's type. Either MANUAL or SYNCED.
     */
    type: string;
}

export interface TwingateResourceAccessGroup {
    /**
     * Group ID that will have permission to access the Resource.
     */
    groupId: string;
    /**
     * The ID of a `twingate.getTwingateSecurityPolicy` to use as the access policy for the group IDs in the access block.
     */
    securityPolicyId: string;
}

export interface TwingateResourceAccessService {
    /**
     * The ID of the service account that should have access to this Resource.
     */
    serviceAccountId: string;
}

export interface TwingateResourceProtocols {
    /**
     * Whether to allow ICMP (ping) traffic
     */
    allowIcmp: boolean;
    tcp: outputs.TwingateResourceProtocolsTcp;
    udp: outputs.TwingateResourceProtocolsUdp;
}

export interface TwingateResourceProtocolsTcp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports: string[];
}

export interface TwingateResourceProtocolsUdp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports: string[];
}

