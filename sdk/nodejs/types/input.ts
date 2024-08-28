// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetTwingateDNSFilteringProfileAllowedDomains {
    /**
     * A set of allowed domains.
     */
    domains?: string[];
}

export interface GetTwingateDNSFilteringProfileAllowedDomainsArgs {
    /**
     * A set of allowed domains.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetTwingateDNSFilteringProfileContentCategories {
    /**
     * Whether to block adult content.
     */
    blockAdultContent?: boolean;
    /**
     * Whether to block dating content.
     */
    blockDating?: boolean;
    /**
     * Whether to block gambling content.
     */
    blockGambling?: boolean;
    /**
     * Whether to block games.
     */
    blockGames?: boolean;
    /**
     * Whether to block piracy sites.
     */
    blockPiracy?: boolean;
    /**
     * Whether to block social media.
     */
    blockSocialMedia?: boolean;
    /**
     * Whether to block streaming content.
     */
    blockStreaming?: boolean;
    /**
     * Whether to force safe search.
     */
    enableSafesearch?: boolean;
    /**
     * Whether to force YouTube to use restricted mode.
     */
    enableYoutubeRestrictedMode?: boolean;
}

export interface GetTwingateDNSFilteringProfileContentCategoriesArgs {
    /**
     * Whether to block adult content.
     */
    blockAdultContent?: pulumi.Input<boolean>;
    /**
     * Whether to block dating content.
     */
    blockDating?: pulumi.Input<boolean>;
    /**
     * Whether to block gambling content.
     */
    blockGambling?: pulumi.Input<boolean>;
    /**
     * Whether to block games.
     */
    blockGames?: pulumi.Input<boolean>;
    /**
     * Whether to block piracy sites.
     */
    blockPiracy?: pulumi.Input<boolean>;
    /**
     * Whether to block social media.
     */
    blockSocialMedia?: pulumi.Input<boolean>;
    /**
     * Whether to block streaming content.
     */
    blockStreaming?: pulumi.Input<boolean>;
    /**
     * Whether to force safe search.
     */
    enableSafesearch?: pulumi.Input<boolean>;
    /**
     * Whether to force YouTube to use restricted mode.
     */
    enableYoutubeRestrictedMode?: pulumi.Input<boolean>;
}

export interface GetTwingateDNSFilteringProfileDeniedDomains {
    /**
     * A set of denied domains.
     */
    domains?: string[];
}

export interface GetTwingateDNSFilteringProfileDeniedDomainsArgs {
    /**
     * A set of denied domains.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetTwingateDNSFilteringProfilePrivacyCategories {
    /**
     * Whether to block ads and trackers.
     */
    blockAdsAndTrackers?: boolean;
    /**
     * Whether to block affiliate links.
     */
    blockAffiliateLinks?: boolean;
    /**
     * Whether to block disguised third party trackers.
     */
    blockDisguisedTrackers?: boolean;
}

export interface GetTwingateDNSFilteringProfilePrivacyCategoriesArgs {
    /**
     * Whether to block ads and trackers.
     */
    blockAdsAndTrackers?: pulumi.Input<boolean>;
    /**
     * Whether to block affiliate links.
     */
    blockAffiliateLinks?: pulumi.Input<boolean>;
    /**
     * Whether to block disguised third party trackers.
     */
    blockDisguisedTrackers?: pulumi.Input<boolean>;
}

export interface GetTwingateDNSFilteringProfileSecurityCategories {
    /**
     * Whether to block cryptojacking sites.
     */
    blockCryptojacking?: boolean;
    /**
     * Blocks public DNS entries from returning private IP addresses.
     */
    blockDnsRebinding?: boolean;
    /**
     * Blocks DGA domains.
     */
    blockDomainGenerationAlgorithms?: boolean;
    /**
     * Whether to block homoglyph attacks.
     */
    blockIdnHomoglyph?: boolean;
    /**
     * Blocks newly registered domains.
     */
    blockNewlyRegisteredDomains?: boolean;
    /**
     * Block parked domains.
     */
    blockParkedDomains?: boolean;
    /**
     * Blocks typosquatted domains.
     */
    blockTyposquatting?: boolean;
    /**
     * Whether to use Google Safe browsing lists to block content.
     */
    enableGoogleSafeBrowsing?: boolean;
    /**
     * Whether to filter content using threat intelligence feeds.
     */
    enableThreatIntelligenceFeeds?: boolean;
}

export interface GetTwingateDNSFilteringProfileSecurityCategoriesArgs {
    /**
     * Whether to block cryptojacking sites.
     */
    blockCryptojacking?: pulumi.Input<boolean>;
    /**
     * Blocks public DNS entries from returning private IP addresses.
     */
    blockDnsRebinding?: pulumi.Input<boolean>;
    /**
     * Blocks DGA domains.
     */
    blockDomainGenerationAlgorithms?: pulumi.Input<boolean>;
    /**
     * Whether to block homoglyph attacks.
     */
    blockIdnHomoglyph?: pulumi.Input<boolean>;
    /**
     * Blocks newly registered domains.
     */
    blockNewlyRegisteredDomains?: pulumi.Input<boolean>;
    /**
     * Block parked domains.
     */
    blockParkedDomains?: pulumi.Input<boolean>;
    /**
     * Blocks typosquatted domains.
     */
    blockTyposquatting?: pulumi.Input<boolean>;
    /**
     * Whether to use Google Safe browsing lists to block content.
     */
    enableGoogleSafeBrowsing?: pulumi.Input<boolean>;
    /**
     * Whether to filter content using threat intelligence feeds.
     */
    enableThreatIntelligenceFeeds?: pulumi.Input<boolean>;
}

export interface GetTwingateResourceProtocols {
    /**
     * Whether to allow ICMP (ping) traffic
     */
    allowIcmp?: boolean;
    tcp?: inputs.GetTwingateResourceProtocolsTcp;
    udp?: inputs.GetTwingateResourceProtocolsUdp;
}

export interface GetTwingateResourceProtocolsArgs {
    /**
     * Whether to allow ICMP (ping) traffic
     */
    allowIcmp?: pulumi.Input<boolean>;
    tcp?: pulumi.Input<inputs.GetTwingateResourceProtocolsTcpArgs>;
    udp?: pulumi.Input<inputs.GetTwingateResourceProtocolsUdpArgs>;
}

export interface GetTwingateResourceProtocolsTcp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy?: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports?: string[];
}

export interface GetTwingateResourceProtocolsTcpArgs {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy?: pulumi.Input<string>;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetTwingateResourceProtocolsUdp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy?: string;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports?: string[];
}

export interface GetTwingateResourceProtocolsUdpArgs {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy?: pulumi.Input<string>;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface TwingateDNSFilteringProfileAllowedDomains {
    /**
     * A set of allowed domains. Defaults to an empty set.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    isAuthoritative?: pulumi.Input<boolean>;
}

export interface TwingateDNSFilteringProfileContentCategories {
    /**
     * Whether to block adult content. Defaults to false.
     */
    blockAdultContent?: pulumi.Input<boolean>;
    /**
     * Whether to block dating content. Defaults to false.
     */
    blockDating?: pulumi.Input<boolean>;
    /**
     * Whether to block gambling content. Defaults to false.
     */
    blockGambling?: pulumi.Input<boolean>;
    /**
     * Whether to block games. Defaults to false.
     */
    blockGames?: pulumi.Input<boolean>;
    /**
     * Whether to block piracy sites. Defaults to false.
     */
    blockPiracy?: pulumi.Input<boolean>;
    /**
     * Whether to block social media. Defaults to false.
     */
    blockSocialMedia?: pulumi.Input<boolean>;
    /**
     * Whether to block streaming content. Defaults to false.
     */
    blockStreaming?: pulumi.Input<boolean>;
    /**
     * Whether to force safe search. Defaults to false.
     */
    enableSafesearch?: pulumi.Input<boolean>;
    /**
     * Whether to force YouTube to use restricted mode. Defaults to false.
     */
    enableYoutubeRestrictedMode?: pulumi.Input<boolean>;
}

export interface TwingateDNSFilteringProfileDeniedDomains {
    /**
     * A set of denied domains. Defaults to an empty set.
     */
    domains?: pulumi.Input<pulumi.Input<string>[]>;
    isAuthoritative?: pulumi.Input<boolean>;
}

export interface TwingateDNSFilteringProfilePrivacyCategories {
    /**
     * Whether to block ads and trackers. Defaults to false.
     */
    blockAdsAndTrackers?: pulumi.Input<boolean>;
    /**
     * Whether to block affiliate links. Defaults to false.
     */
    blockAffiliateLinks?: pulumi.Input<boolean>;
    /**
     * Whether to block disguised third party trackers. Defaults to false.
     */
    blockDisguisedTrackers?: pulumi.Input<boolean>;
}

export interface TwingateDNSFilteringProfileSecurityCategories {
    /**
     * Whether to block cryptojacking sites. Defaults to true.
     */
    blockCryptojacking?: pulumi.Input<boolean>;
    /**
     * Blocks public DNS entries from returning private IP addresses. Defaults to true.
     */
    blockDnsRebinding?: pulumi.Input<boolean>;
    /**
     * Blocks DGA domains. Defaults to true.
     */
    blockDomainGenerationAlgorithms?: pulumi.Input<boolean>;
    /**
     * Whether to block homoglyph attacks. Defaults to true.
     */
    blockIdnHomoglyph?: pulumi.Input<boolean>;
    /**
     * Blocks newly registered domains. Defaults to true.
     */
    blockNewlyRegisteredDomains?: pulumi.Input<boolean>;
    /**
     * Block parked domains. Defaults to true.
     */
    blockParkedDomains?: pulumi.Input<boolean>;
    /**
     * Blocks typosquatted domains. Defaults to true.
     */
    blockTyposquatting?: pulumi.Input<boolean>;
    /**
     * Whether to use Google Safe browsing lists to block content. Defaults to true.
     */
    enableGoogleSafeBrowsing?: pulumi.Input<boolean>;
    /**
     * Whether to filter content using threat intelligence feeds. Defaults to true.
     */
    enableThreatIntelligenceFeeds?: pulumi.Input<boolean>;
}

export interface TwingateResourceAccessGroup {
    /**
     * Group ID that will have permission to access the Resource.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The ID of a `twingate.getTwingateSecurityPolicy` to use as the access policy for the group IDs in the access block.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * The usage-based auto-lock duration configured on the edge (in days).
     */
    usageBasedAutolockDurationDays?: pulumi.Input<number>;
}

export interface TwingateResourceAccessService {
    /**
     * The ID of the service account that should have access to this Resource.
     */
    serviceAccountId?: pulumi.Input<string>;
}

export interface TwingateResourceProtocols {
    /**
     * Whether to allow ICMP (ping) traffic
     */
    allowIcmp?: pulumi.Input<boolean>;
    tcp?: pulumi.Input<inputs.TwingateResourceProtocolsTcp>;
    udp?: pulumi.Input<inputs.TwingateResourceProtocolsUdp>;
}

export interface TwingateResourceProtocolsTcp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy?: pulumi.Input<string>;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface TwingateResourceProtocolsUdp {
    /**
     * Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
     */
    policy?: pulumi.Input<string>;
    /**
     * List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
     */
    ports?: pulumi.Input<pulumi.Input<string>[]>;
}
