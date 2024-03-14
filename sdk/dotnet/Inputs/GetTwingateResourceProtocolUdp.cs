// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Twingate.Inputs
{

    public sealed class GetTwingateResourceProtocolUdpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
        /// </summary>
        [Input("policy", required: true)]
        public string Policy { get; set; } = null!;

        [Input("ports", required: true)]
        private List<string>? _ports;

        /// <summary>
        /// List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
        /// </summary>
        public List<string> Ports
        {
            get => _ports ?? (_ports = new List<string>());
            set => _ports = value;
        }

        public GetTwingateResourceProtocolUdpArgs()
        {
        }
        public static new GetTwingateResourceProtocolUdpArgs Empty => new GetTwingateResourceProtocolUdpArgs();
    }
}
