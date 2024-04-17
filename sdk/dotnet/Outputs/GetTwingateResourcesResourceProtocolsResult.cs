// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Twingate.Outputs
{

    [OutputType]
    public sealed class GetTwingateResourcesResourceProtocolsResult
    {
        /// <summary>
        /// Whether to allow ICMP (ping) traffic
        /// </summary>
        public readonly bool AllowIcmp;
        public readonly Outputs.GetTwingateResourcesResourceProtocolsTcpResult Tcp;
        public readonly Outputs.GetTwingateResourcesResourceProtocolsUdpResult Udp;

        [OutputConstructor]
        private GetTwingateResourcesResourceProtocolsResult(
            bool allowIcmp,

            Outputs.GetTwingateResourcesResourceProtocolsTcpResult tcp,

            Outputs.GetTwingateResourcesResourceProtocolsUdpResult udp)
        {
            AllowIcmp = allowIcmp;
            Tcp = tcp;
            Udp = udp;
        }
    }
}
