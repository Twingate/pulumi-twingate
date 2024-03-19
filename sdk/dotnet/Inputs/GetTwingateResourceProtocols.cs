// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Twingate.Inputs
{

    public sealed class GetTwingateResourceProtocolsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to allow ICMP (ping) traffic
        /// </summary>
        [Input("allowIcmp", required: true)]
        public bool AllowIcmp { get; set; }

        [Input("tcp")]
        public Inputs.GetTwingateResourceProtocolsTcpArgs? Tcp { get; set; }

        [Input("udp")]
        public Inputs.GetTwingateResourceProtocolsUdpArgs? Udp { get; set; }

        public GetTwingateResourceProtocolsArgs()
        {
        }
        public static new GetTwingateResourceProtocolsArgs Empty => new GetTwingateResourceProtocolsArgs();
    }
}
