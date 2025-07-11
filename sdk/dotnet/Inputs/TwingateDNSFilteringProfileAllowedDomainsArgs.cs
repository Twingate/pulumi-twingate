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

    public sealed class TwingateDNSFilteringProfileAllowedDomainsArgs : global::Pulumi.ResourceArgs
    {
        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// A set of allowed domains. Defaults to an empty set.
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        [Input("isAuthoritative")]
        public Input<bool>? IsAuthoritative { get; set; }

        public TwingateDNSFilteringProfileAllowedDomainsArgs()
        {
        }
        public static new TwingateDNSFilteringProfileAllowedDomainsArgs Empty => new TwingateDNSFilteringProfileAllowedDomainsArgs();
    }
}
