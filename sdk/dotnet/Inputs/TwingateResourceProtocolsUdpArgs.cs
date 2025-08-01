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

    public sealed class TwingateResourceProtocolsUdpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to allow or deny all ports, or restrict protocol access within certain port ranges: Can be `RESTRICTED` (only listed ports are allowed), `ALLOW_ALL`, or `DENY_ALL`
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// List of port ranges between 1 and 65535 inclusive, in the format `100-200` for a range, or `8080` for a single port
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        public TwingateResourceProtocolsUdpArgs()
        {
        }
        public static new TwingateResourceProtocolsUdpArgs Empty => new TwingateResourceProtocolsUdpArgs();
    }
}
