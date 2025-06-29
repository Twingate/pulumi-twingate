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
    public sealed class TwingateResourceAccessGroup
    {
        /// <summary>
        /// This will set the approval model on the edge. The valid values are `AUTOMATIC` and `MANUAL`.
        /// </summary>
        public readonly string? ApprovalMode;
        /// <summary>
        /// Group ID that will have permission to access the Resource.
        /// </summary>
        public readonly string? GroupId;
        /// <summary>
        /// The ID of a `twingate.getTwingateSecurityPolicy` to use as the access policy for the group IDs in the access block.
        /// </summary>
        public readonly string? SecurityPolicyId;
        /// <summary>
        /// The usage-based auto-lock duration configured on the edge (in days).
        /// </summary>
        public readonly int? UsageBasedAutolockDurationDays;

        [OutputConstructor]
        private TwingateResourceAccessGroup(
            string? approvalMode,

            string? groupId,

            string? securityPolicyId,

            int? usageBasedAutolockDurationDays)
        {
            ApprovalMode = approvalMode;
            GroupId = groupId;
            SecurityPolicyId = securityPolicyId;
            UsageBasedAutolockDurationDays = usageBasedAutolockDurationDays;
        }
    }
}
