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

    public sealed class TwingateResourceAccessGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group ID that will have permission to access the Resource.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The ID of a `twingate.getTwingateSecurityPolicy` to use as the access policy for the group IDs in the access block.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        /// <summary>
        /// The usage-based auto-lock duration configured on the edge (in days).
        /// </summary>
        [Input("usageBasedAutolockDurationDays")]
        public Input<int>? UsageBasedAutolockDurationDays { get; set; }

        public TwingateResourceAccessGroupArgs()
        {
        }
        public static new TwingateResourceAccessGroupArgs Empty => new TwingateResourceAccessGroupArgs();
    }
}
