// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Twingate
{
    public static class GetTwingateSecurityPolicy
    {
        /// <summary>
        /// Security Policies are defined in the Twingate Admin Console and determine user and device authentication requirements for Resources.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateSecurityPolicy.Invoke(new()
        ///     {
        ///         Name = "&lt;your security policy name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetTwingateSecurityPolicyResult> InvokeAsync(GetTwingateSecurityPolicyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTwingateSecurityPolicyResult>("twingate:index/getTwingateSecurityPolicy:getTwingateSecurityPolicy", args ?? new GetTwingateSecurityPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Security Policies are defined in the Twingate Admin Console and determine user and device authentication requirements for Resources.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateSecurityPolicy.Invoke(new()
        ///     {
        ///         Name = "&lt;your security policy name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetTwingateSecurityPolicyResult> Invoke(GetTwingateSecurityPolicyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTwingateSecurityPolicyResult>("twingate:index/getTwingateSecurityPolicy:getTwingateSecurityPolicy", args ?? new GetTwingateSecurityPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateSecurityPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return a Security Policy by its ID. The ID for the Security Policy can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Return a Security Policy that exactly matches this name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetTwingateSecurityPolicyArgs()
        {
        }
        public static new GetTwingateSecurityPolicyArgs Empty => new GetTwingateSecurityPolicyArgs();
    }

    public sealed class GetTwingateSecurityPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Return a Security Policy by its ID. The ID for the Security Policy can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Return a Security Policy that exactly matches this name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetTwingateSecurityPolicyInvokeArgs()
        {
        }
        public static new GetTwingateSecurityPolicyInvokeArgs Empty => new GetTwingateSecurityPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetTwingateSecurityPolicyResult
    {
        /// <summary>
        /// Return a Security Policy by its ID. The ID for the Security Policy can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Return a Security Policy that exactly matches this name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetTwingateSecurityPolicyResult(
            string? id,

            string? name)
        {
            Id = id;
            Name = name;
        }
    }
}