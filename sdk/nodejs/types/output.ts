// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetTwingateConnectorsConnector {
    /**
     * The hostname of the machine hosting the Connector.
     */
    hostname: string;
    /**
     * The ID of the Connector.
     */
    id: string;
    /**
     * The Name of the Connector.
     */
    name: string;
    /**
     * The Connector's private IP addresses.
     */
    privateIps: string[];
    /**
     * The Connector's public IP address.
     */
    publicIp: string;
    /**
     * The ID of the Remote Network attached to the Connector.
     */
    remoteNetworkId: string;
    /**
     * The Connector's state. One of `ALIVE`, `DEAD_NO_HEARTBEAT`, `DEAD_HEARTBEAT_TOO_OLD` or `DEAD_NO_RELAYS`.
     */
    state: string;
    /**
     * Determines whether status notifications are enabled for the Connector.
     */
    statusUpdatesEnabled: boolean;
    /**
     * The Connector's version.
     */
    version: string;
}

export interface GetTwingateDNSFilteringProfileAllowedDomains {
    /**
     * A set of allowed domains.
     */
    domains: string[];
}

export interface GetTwingateDNSFilteringProfileContentCategories {
    /**
     * Whether to block adult content.
     */
    blockAdultContent: boolean;
    /**
     * Whether to block dating content.
     */
    blockDating: boolean;
    /**
     * Whether to block gambling content.
     */
    blockGambling: boolean;
    /**
     * Whether to block games.
     */
    blockGames: boolean;
    /**
     * Whether to block piracy sites.
     */
    blockPiracy: boolean;
    /**
     * Whether to block social media.
     */
    blockSocialMedia: boolean;
    /**
     * Whether to block streaming content.
     */
    blockStreaming: boolean;
    /**
     * Whether to force safe search.
     */
    enableSafesearch: boolean;
    /**
     * Whether to force YouTube to use restricted mode.
     */
    enableYoutubeRestrictedMode: boolean;
}

export interface GetTwingateDNSFilteringProfileDeniedDomains {
    /**
     * A set of denied domains.
     */
    domains: string[];
}

export interface GetTwingateDNSFilteringProfilePrivacyCategories {
    /**
     * Whether to block ads and trackers.
     */
    blockAdsAndTrackers: boolean;
    /**
     * Whether to block affiliate links.
     */
    blockAffiliateLinks: boolean;
    /**
     * Whether to block disguised third party trackers.
     */
    blockDisguisedTrackers: boolean;
}

export interface GetTwingateDNSFilteringProfileSecurityCategories {
    /**
     * Whether to block cryptojacking sites.
     */
    blockCryptojacking: boolean;
    /**
     * Blocks public DNS entries from returning private IP addresses.
     */
    blockDnsRebinding: boolean;
    /**
     * Blocks DGA domains.
     */
    blockDomainGenerationAlgorithms: boolean;
    /**
     * Whether to block homoglyph attacks.
     */
    blockIdnHomoglyph: boolean;
    /**
     * Blocks newly registered domains.
     */
    blockNewlyRegisteredDomains: boolean;
    /**
     * Block parked domains.
     */
    blockParkedDomains: boolean;
    /**
     * Blocks typosquatted domains.
     */
    blockTyposquatting: boolean;
    /**
     * Whether to use Google Safe browsing lists to block content.
     */
    enableGoogleSafeBrowsing: boolean;
    /**
     * Whether to filter content using threat intelligence feeds.
     */
    enableThreatIntelligenceFeeds: boolean;
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
    /**
     * The type of the Remote Network. Must be one of the following: REGULAR, EXIT.
     */
    type: string;
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
    /**
     * The `tags` attribute consists of a key-value pairs that correspond with tags to be set on the resource.
     */
    tags: {[key: string]: string};
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

export interface TwingateDNSFilteringProfileAllowedDomains {
    /**
     * A set of allowed domains. Defaults to an empty set.
     */
    domains: string[];
    isAuthoritative: boolean;
}

export interface TwingateDNSFilteringProfileContentCategories {
    /**
     * Whether to block adult content. Defaults to false.
     */
    blockAdultContent: boolean;
    /**
     * Whether to block dating content. Defaults to false.
     */
    blockDating: boolean;
    /**
     * Whether to block gambling content. Defaults to false.
     */
    blockGambling: boolean;
    /**
     * Whether to block games. Defaults to false.
     */
    blockGames: boolean;
    /**
     * Whether to block piracy sites. Defaults to false.
     */
    blockPiracy: boolean;
    /**
     * Whether to block social media. Defaults to false.
     */
    blockSocialMedia: boolean;
    /**
     * Whether to block streaming content. Defaults to false.
     */
    blockStreaming: boolean;
    /**
     * Whether to force safe search. Defaults to false.
     */
    enableSafesearch: boolean;
    /**
     * Whether to force YouTube to use restricted mode. Defaults to false.
     */
    enableYoutubeRestrictedMode: boolean;
}

export interface TwingateDNSFilteringProfileDeniedDomains {
    /**
     * A set of denied domains. Defaults to an empty set.
     */
    domains: string[];
    isAuthoritative: boolean;
}

export interface TwingateDNSFilteringProfilePrivacyCategories {
    /**
     * Whether to block ads and trackers. Defaults to false.
     */
    blockAdsAndTrackers: boolean;
    /**
     * Whether to block affiliate links. Defaults to false.
     */
    blockAffiliateLinks: boolean;
    /**
     * Whether to block disguised third party trackers. Defaults to false.
     */
    blockDisguisedTrackers: boolean;
}

export interface TwingateDNSFilteringProfileSecurityCategories {
    /**
     * Whether to block cryptojacking sites. Defaults to true.
     */
    blockCryptojacking: boolean;
    /**
     * Blocks public DNS entries from returning private IP addresses. Defaults to true.
     */
    blockDnsRebinding: boolean;
    /**
     * Blocks DGA domains. Defaults to true.
     */
    blockDomainGenerationAlgorithms: boolean;
    /**
     * Whether to block homoglyph attacks. Defaults to true.
     */
    blockIdnHomoglyph: boolean;
    /**
     * Blocks newly registered domains. Defaults to true.
     */
    blockNewlyRegisteredDomains: boolean;
    /**
     * Block parked domains. Defaults to true.
     */
    blockParkedDomains: boolean;
    /**
     * Blocks typosquatted domains. Defaults to true.
     */
    blockTyposquatting: boolean;
    /**
     * Whether to use Google Safe browsing lists to block content. Defaults to true.
     */
    enableGoogleSafeBrowsing: boolean;
    /**
     * Whether to filter content using threat intelligence feeds. Defaults to true.
     */
    enableThreatIntelligenceFeeds: boolean;
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
    /**
     * The usage-based auto-lock duration configured on the edge (in days).
     */
    usageBasedAutolockDurationDays: number;
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

