// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Twingate.Twingate
{
    /// <summary>
    /// The provider type for the twingate package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [TwingateResourceType("pulumi:providers:twingate")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The access key for API operations. You can retrieve this from the Twingate Admin Console
        /// ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
        /// TWINGATE_API_TOKEN environment variable.
        /// </summary>
        [Output("apiToken")]
        public Output<string?> ApiToken { get; private set; } = null!;

        /// <summary>
        /// Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
        /// `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
        /// environment variable.
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("twingate", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/Twingate/pulumi-twingate",
                AdditionalSecretOutputs =
                {
                    "apiToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }

        /// <summary>
        /// This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
        /// </summary>
        public global::Pulumi.Output<ProviderTerraformConfigResult> TerraformConfig()
            => global::Pulumi.Deployment.Instance.Call<ProviderTerraformConfigResult>("pulumi:providers:twingate/terraformConfig", CallArgs.Empty, this);
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiToken")]
        private Input<string>? _apiToken;

        /// <summary>
        /// The access key for API operations. You can retrieve this from the Twingate Admin Console
        /// ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
        /// TWINGATE_API_TOKEN environment variable.
        /// </summary>
        public Input<string>? ApiToken
        {
            get => _apiToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _apiToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies the cache settings for the provider.
        /// </summary>
        [Input("cache", json: true)]
        public Input<Inputs.ProviderCacheArgs>? Cache { get; set; }

        /// <summary>
        /// A default set of tags applied globally to all resources created by the provider.
        /// </summary>
        [Input("defaultTags", json: true)]
        public Input<Inputs.ProviderDefaultTagsArgs>? DefaultTags { get; set; }

        /// <summary>
        /// Specifies a retry limit for the http requests made. The default value is 10. Alternatively, this can be specified using
        /// the TWINGATE_HTTP_MAX_RETRY environment variable
        /// </summary>
        [Input("httpMaxRetry", json: true)]
        public Input<int>? HttpMaxRetry { get; set; }

        /// <summary>
        /// Specifies a time limit in seconds for the http requests made. The default value is 35 seconds. Alternatively, this can
        /// be specified using the TWINGATE_HTTP_TIMEOUT environment variable
        /// </summary>
        [Input("httpTimeout", json: true)]
        public Input<int>? HttpTimeout { get; set; }

        /// <summary>
        /// Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
        /// `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
        /// environment variable.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }

    /// <summary>
    /// The results of the <see cref="Provider.TerraformConfig"/> method.
    /// </summary>
    [OutputType]
    public sealed class ProviderTerraformConfigResult
    {
        public readonly ImmutableDictionary<string, object> Result;

        [OutputConstructor]
        private ProviderTerraformConfigResult(ImmutableDictionary<string, object> result)
        {
            Result = result;
        }
    }
}
