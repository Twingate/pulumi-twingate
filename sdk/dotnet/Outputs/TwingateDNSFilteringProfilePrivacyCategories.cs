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
    public sealed class TwingateDNSFilteringProfilePrivacyCategories
    {
        /// <summary>
        /// Whether to block ads and trackers. Defaults to false.
        /// </summary>
        public readonly bool? BlockAdsAndTrackers;
        /// <summary>
        /// Whether to block affiliate links. Defaults to false.
        /// </summary>
        public readonly bool? BlockAffiliateLinks;
        /// <summary>
        /// Whether to block disguised third party trackers. Defaults to false.
        /// </summary>
        public readonly bool? BlockDisguisedTrackers;

        [OutputConstructor]
        private TwingateDNSFilteringProfilePrivacyCategories(
            bool? blockAdsAndTrackers,

            bool? blockAffiliateLinks,

            bool? blockDisguisedTrackers)
        {
            BlockAdsAndTrackers = blockAdsAndTrackers;
            BlockAffiliateLinks = blockAffiliateLinks;
            BlockDisguisedTrackers = blockDisguisedTrackers;
        }
    }
}
