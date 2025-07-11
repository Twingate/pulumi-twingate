// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Twingate.Twingate.Inputs
{

    public sealed class GetTwingateDNSFilteringProfileSecurityCategoriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to block cryptojacking sites.
        /// </summary>
        [Input("blockCryptojacking", required: true)]
        public bool BlockCryptojacking { get; set; }

        /// <summary>
        /// Blocks public DNS entries from returning private IP addresses.
        /// </summary>
        [Input("blockDnsRebinding", required: true)]
        public bool BlockDnsRebinding { get; set; }

        /// <summary>
        /// Blocks DGA domains.
        /// </summary>
        [Input("blockDomainGenerationAlgorithms", required: true)]
        public bool BlockDomainGenerationAlgorithms { get; set; }

        /// <summary>
        /// Whether to block homoglyph attacks.
        /// </summary>
        [Input("blockIdnHomoglyph", required: true)]
        public bool BlockIdnHomoglyph { get; set; }

        /// <summary>
        /// Blocks newly registered domains.
        /// </summary>
        [Input("blockNewlyRegisteredDomains", required: true)]
        public bool BlockNewlyRegisteredDomains { get; set; }

        /// <summary>
        /// Block parked domains.
        /// </summary>
        [Input("blockParkedDomains", required: true)]
        public bool BlockParkedDomains { get; set; }

        /// <summary>
        /// Blocks typosquatted domains.
        /// </summary>
        [Input("blockTyposquatting", required: true)]
        public bool BlockTyposquatting { get; set; }

        /// <summary>
        /// Whether to use Google Safe browsing lists to block content.
        /// </summary>
        [Input("enableGoogleSafeBrowsing", required: true)]
        public bool EnableGoogleSafeBrowsing { get; set; }

        /// <summary>
        /// Whether to filter content using threat intelligence feeds.
        /// </summary>
        [Input("enableThreatIntelligenceFeeds", required: true)]
        public bool EnableThreatIntelligenceFeeds { get; set; }

        public GetTwingateDNSFilteringProfileSecurityCategoriesArgs()
        {
        }
        public static new GetTwingateDNSFilteringProfileSecurityCategoriesArgs Empty => new GetTwingateDNSFilteringProfileSecurityCategoriesArgs();
    }
}
