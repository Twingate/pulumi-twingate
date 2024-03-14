// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Twingate
{
    public static class GetTwingateResources
    {
        /// <summary>
        /// Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateResources.Invoke(new()
        ///     {
        ///         Name = "&lt;your resource's name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTwingateResourcesResult> InvokeAsync(GetTwingateResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTwingateResourcesResult>("twingate:index/getTwingateResources:getTwingateResources", args ?? new GetTwingateResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateResources.Invoke(new()
        ///     {
        ///         Name = "&lt;your resource's name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTwingateResourcesResult> Invoke(GetTwingateResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTwingateResourcesResult>("twingate:index/getTwingateResources:getTwingateResources", args ?? new GetTwingateResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("resources")]
        private List<Inputs.GetTwingateResourcesResourceArgs>? _resources;

        /// <summary>
        /// List of Resources
        /// </summary>
        public List<Inputs.GetTwingateResourcesResourceArgs> Resources
        {
            get => _resources ?? (_resources = new List<Inputs.GetTwingateResourcesResourceArgs>());
            set => _resources = value;
        }

        public GetTwingateResourcesArgs()
        {
        }
        public static new GetTwingateResourcesArgs Empty => new GetTwingateResourcesArgs();
    }

    public sealed class GetTwingateResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Resource
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("resources")]
        private InputList<Inputs.GetTwingateResourcesResourceInputArgs>? _resources;

        /// <summary>
        /// List of Resources
        /// </summary>
        public InputList<Inputs.GetTwingateResourcesResourceInputArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.GetTwingateResourcesResourceInputArgs>());
            set => _resources = value;
        }

        public GetTwingateResourcesInvokeArgs()
        {
        }
        public static new GetTwingateResourcesInvokeArgs Empty => new GetTwingateResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTwingateResourcesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of Resources
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTwingateResourcesResourceResult> Resources;

        [OutputConstructor]
        private GetTwingateResourcesResult(
            string id,

            string name,

            ImmutableArray<Outputs.GetTwingateResourcesResourceResult> resources)
        {
            Id = id;
            Name = name;
            Resources = resources;
        }
    }
}
