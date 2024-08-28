// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Twingate.Twingate.Inputs
{

    public sealed class TwingateDNSFilteringProfileContentCategoriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to block adult content. Defaults to false.
        /// </summary>
        [Input("blockAdultContent")]
        public Input<bool>? BlockAdultContent { get; set; }

        /// <summary>
        /// Whether to block dating content. Defaults to false.
        /// </summary>
        [Input("blockDating")]
        public Input<bool>? BlockDating { get; set; }

        /// <summary>
        /// Whether to block gambling content. Defaults to false.
        /// </summary>
        [Input("blockGambling")]
        public Input<bool>? BlockGambling { get; set; }

        /// <summary>
        /// Whether to block games. Defaults to false.
        /// </summary>
        [Input("blockGames")]
        public Input<bool>? BlockGames { get; set; }

        /// <summary>
        /// Whether to block piracy sites. Defaults to false.
        /// </summary>
        [Input("blockPiracy")]
        public Input<bool>? BlockPiracy { get; set; }

        /// <summary>
        /// Whether to block social media. Defaults to false.
        /// </summary>
        [Input("blockSocialMedia")]
        public Input<bool>? BlockSocialMedia { get; set; }

        /// <summary>
        /// Whether to block streaming content. Defaults to false.
        /// </summary>
        [Input("blockStreaming")]
        public Input<bool>? BlockStreaming { get; set; }

        /// <summary>
        /// Whether to force safe search. Defaults to false.
        /// </summary>
        [Input("enableSafesearch")]
        public Input<bool>? EnableSafesearch { get; set; }

        /// <summary>
        /// Whether to force YouTube to use restricted mode. Defaults to false.
        /// </summary>
        [Input("enableYoutubeRestrictedMode")]
        public Input<bool>? EnableYoutubeRestrictedMode { get; set; }

        public TwingateDNSFilteringProfileContentCategoriesArgs()
        {
        }
        public static new TwingateDNSFilteringProfileContentCategoriesArgs Empty => new TwingateDNSFilteringProfileContentCategoriesArgs();
    }
}
