// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Twingate.Twingate
{
    public static class GetTwingateGroup
    {
        /// <summary>
        /// Groups are how users are authorized to access Resources. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/groups).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateGroup.Invoke(new()
        ///     {
        ///         Id = "&lt;your group's id&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTwingateGroupResult> InvokeAsync(GetTwingateGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTwingateGroupResult>("twingate:index/getTwingateGroup:getTwingateGroup", args ?? new GetTwingateGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Groups are how users are authorized to access Resources. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/groups).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateGroup.Invoke(new()
        ///     {
        ///         Id = "&lt;your group's id&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTwingateGroupResult> Invoke(GetTwingateGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTwingateGroupResult>("twingate:index/getTwingateGroup:getTwingateGroup", args ?? new GetTwingateGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTwingateGroupArgs()
        {
        }
        public static new GetTwingateGroupArgs Empty => new GetTwingateGroupArgs();
    }

    public sealed class GetTwingateGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTwingateGroupInvokeArgs()
        {
        }
        public static new GetTwingateGroupInvokeArgs Empty => new GetTwingateGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetTwingateGroupResult
    {
        /// <summary>
        /// The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates if the Group is active
        /// </summary>
        public readonly bool IsActive;
        /// <summary>
        /// The name of the Group
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Security Policy assigned to the Group.
        /// </summary>
        public readonly string SecurityPolicyId;
        /// <summary>
        /// The type of the Group
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTwingateGroupResult(
            string id,

            bool isActive,

            string name,

            string securityPolicyId,

            string type)
        {
            Id = id;
            IsActive = isActive;
            Name = name;
            SecurityPolicyId = securityPolicyId;
            Type = type;
        }
    }
}
