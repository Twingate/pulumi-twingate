// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    public sealed class TwingateDNSFilteringProfileAllowedDomains
    {
        /// <summary>
        /// A set of allowed domains. Defaults to an empty set.
        /// </summary>
        public readonly ImmutableArray<string> Domains;
        public readonly bool? IsAuthoritative;

        [OutputConstructor]
        private TwingateDNSFilteringProfileAllowedDomains(
            ImmutableArray<string> domains,

            bool? isAuthoritative)
        {
            Domains = domains;
            IsAuthoritative = isAuthoritative;
        }
    }
}
