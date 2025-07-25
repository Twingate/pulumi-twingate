// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Twingate.Twingate.Outputs
{

    [OutputType]
    public sealed class GetTwingateDNSFilteringProfileSecurityCategoriesResult
    {
        /// <summary>
        /// Whether to block cryptojacking sites.
        /// </summary>
        public readonly bool BlockCryptojacking;
        /// <summary>
        /// Blocks public DNS entries from returning private IP addresses.
        /// </summary>
        public readonly bool BlockDnsRebinding;
        /// <summary>
        /// Blocks DGA domains.
        /// </summary>
        public readonly bool BlockDomainGenerationAlgorithms;
        /// <summary>
        /// Whether to block homoglyph attacks.
        /// </summary>
        public readonly bool BlockIdnHomoglyph;
        /// <summary>
        /// Blocks newly registered domains.
        /// </summary>
        public readonly bool BlockNewlyRegisteredDomains;
        /// <summary>
        /// Block parked domains.
        /// </summary>
        public readonly bool BlockParkedDomains;
        /// <summary>
        /// Blocks typosquatted domains.
        /// </summary>
        public readonly bool BlockTyposquatting;
        /// <summary>
        /// Whether to use Google Safe browsing lists to block content.
        /// </summary>
        public readonly bool EnableGoogleSafeBrowsing;
        /// <summary>
        /// Whether to filter content using threat intelligence feeds.
        /// </summary>
        public readonly bool EnableThreatIntelligenceFeeds;

        [OutputConstructor]
        private GetTwingateDNSFilteringProfileSecurityCategoriesResult(
            bool blockCryptojacking,

            bool blockDnsRebinding,

            bool blockDomainGenerationAlgorithms,

            bool blockIdnHomoglyph,

            bool blockNewlyRegisteredDomains,

            bool blockParkedDomains,

            bool blockTyposquatting,

            bool enableGoogleSafeBrowsing,

            bool enableThreatIntelligenceFeeds)
        {
            BlockCryptojacking = blockCryptojacking;
            BlockDnsRebinding = blockDnsRebinding;
            BlockDomainGenerationAlgorithms = blockDomainGenerationAlgorithms;
            BlockIdnHomoglyph = blockIdnHomoglyph;
            BlockNewlyRegisteredDomains = blockNewlyRegisteredDomains;
            BlockParkedDomains = blockParkedDomains;
            BlockTyposquatting = blockTyposquatting;
            EnableGoogleSafeBrowsing = enableGoogleSafeBrowsing;
            EnableThreatIntelligenceFeeds = enableThreatIntelligenceFeeds;
        }
    }
}
