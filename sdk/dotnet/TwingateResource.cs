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
    /// <summary>
    /// Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import twingate:index/twingateResource:TwingateResource resource UmVzb3VyY2U6MzQwNDQ3
    /// ```
    /// </summary>
    [TwingateResourceType("twingate:index/twingateResource:TwingateResource")]
    public partial class TwingateResource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Restrict access to certain group
        /// </summary>
        [Output("accessGroups")]
        public Output<ImmutableArray<Outputs.TwingateResourceAccessGroup>> AccessGroups { get; private set; } = null!;

        /// <summary>
        /// Restrict access to certain service account
        /// </summary>
        [Output("accessServices")]
        public Output<ImmutableArray<Outputs.TwingateResourceAccessService>> AccessServices { get; private set; } = null!;

        /// <summary>
        /// The Resource's IP/CIDR or FQDN/DNS zone
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Set a DNS alias address for the Resource. Must be a DNS-valid name string.
        /// </summary>
        [Output("alias")]
        public Output<string?> Alias { get; private set; } = null!;

        /// <summary>
        /// Set the resource as active or inactive. Default is `true`.
        /// </summary>
        [Output("isActive")]
        public Output<bool> IsActive { get; private set; } = null!;

        [Output("isAuthoritative")]
        public Output<bool> IsAuthoritative { get; private set; } = null!;

        /// <summary>
        /// Controls whether an "Open in Browser" shortcut will be shown for this Resource in the Twingate Client. Default is `false`.
        /// </summary>
        [Output("isBrowserShortcutEnabled")]
        public Output<bool> IsBrowserShortcutEnabled { get; private set; } = null!;

        /// <summary>
        /// Controls whether this Resource will be visible in the main Resource list in the Twingate Client. Default is `true`.
        /// </summary>
        [Output("isVisible")]
        public Output<bool> IsVisible { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no restriction, and all protocols and ports are allowed.
        /// </summary>
        [Output("protocols")]
        public Output<Outputs.TwingateResourceProtocols> Protocols { get; private set; } = null!;

        /// <summary>
        /// Remote Network ID where the Resource lives
        /// </summary>
        [Output("remoteNetworkId")]
        public Output<string> RemoteNetworkId { get; private set; } = null!;

        /// <summary>
        /// The ID of a `twingate.getTwingateSecurityPolicy` to set as this Resource's Security Policy. Default is `Default Policy`.
        /// </summary>
        [Output("securityPolicyId")]
        public Output<string> SecurityPolicyId { get; private set; } = null!;

        /// <summary>
        /// The `tags` attribute consists of a key-value pairs that correspond with tags to be set on the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a TwingateResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TwingateResource(string name, TwingateResourceArgs args, CustomResourceOptions? options = null)
            : base("twingate:index/twingateResource:TwingateResource", name, args ?? new TwingateResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TwingateResource(string name, Input<string> id, TwingateResourceState? state = null, CustomResourceOptions? options = null)
            : base("twingate:index/twingateResource:TwingateResource", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/Twingate/pulumi-twingate",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TwingateResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TwingateResource Get(string name, Input<string> id, TwingateResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new TwingateResource(name, id, state, options);
        }
    }

    public sealed class TwingateResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessGroups")]
        private InputList<Inputs.TwingateResourceAccessGroupArgs>? _accessGroups;

        /// <summary>
        /// Restrict access to certain group
        /// </summary>
        public InputList<Inputs.TwingateResourceAccessGroupArgs> AccessGroups
        {
            get => _accessGroups ?? (_accessGroups = new InputList<Inputs.TwingateResourceAccessGroupArgs>());
            set => _accessGroups = value;
        }

        [Input("accessServices")]
        private InputList<Inputs.TwingateResourceAccessServiceArgs>? _accessServices;

        /// <summary>
        /// Restrict access to certain service account
        /// </summary>
        public InputList<Inputs.TwingateResourceAccessServiceArgs> AccessServices
        {
            get => _accessServices ?? (_accessServices = new InputList<Inputs.TwingateResourceAccessServiceArgs>());
            set => _accessServices = value;
        }

        /// <summary>
        /// The Resource's IP/CIDR or FQDN/DNS zone
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Set a DNS alias address for the Resource. Must be a DNS-valid name string.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Set the resource as active or inactive. Default is `true`.
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        [Input("isAuthoritative")]
        public Input<bool>? IsAuthoritative { get; set; }

        /// <summary>
        /// Controls whether an "Open in Browser" shortcut will be shown for this Resource in the Twingate Client. Default is `false`.
        /// </summary>
        [Input("isBrowserShortcutEnabled")]
        public Input<bool>? IsBrowserShortcutEnabled { get; set; }

        /// <summary>
        /// Controls whether this Resource will be visible in the main Resource list in the Twingate Client. Default is `true`.
        /// </summary>
        [Input("isVisible")]
        public Input<bool>? IsVisible { get; set; }

        /// <summary>
        /// The name of the Resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no restriction, and all protocols and ports are allowed.
        /// </summary>
        [Input("protocols")]
        public Input<Inputs.TwingateResourceProtocolsArgs>? Protocols { get; set; }

        /// <summary>
        /// Remote Network ID where the Resource lives
        /// </summary>
        [Input("remoteNetworkId", required: true)]
        public Input<string> RemoteNetworkId { get; set; } = null!;

        /// <summary>
        /// The ID of a `twingate.getTwingateSecurityPolicy` to set as this Resource's Security Policy. Default is `Default Policy`.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The `tags` attribute consists of a key-value pairs that correspond with tags to be set on the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public TwingateResourceArgs()
        {
        }
        public static new TwingateResourceArgs Empty => new TwingateResourceArgs();
    }

    public sealed class TwingateResourceState : global::Pulumi.ResourceArgs
    {
        [Input("accessGroups")]
        private InputList<Inputs.TwingateResourceAccessGroupGetArgs>? _accessGroups;

        /// <summary>
        /// Restrict access to certain group
        /// </summary>
        public InputList<Inputs.TwingateResourceAccessGroupGetArgs> AccessGroups
        {
            get => _accessGroups ?? (_accessGroups = new InputList<Inputs.TwingateResourceAccessGroupGetArgs>());
            set => _accessGroups = value;
        }

        [Input("accessServices")]
        private InputList<Inputs.TwingateResourceAccessServiceGetArgs>? _accessServices;

        /// <summary>
        /// Restrict access to certain service account
        /// </summary>
        public InputList<Inputs.TwingateResourceAccessServiceGetArgs> AccessServices
        {
            get => _accessServices ?? (_accessServices = new InputList<Inputs.TwingateResourceAccessServiceGetArgs>());
            set => _accessServices = value;
        }

        /// <summary>
        /// The Resource's IP/CIDR or FQDN/DNS zone
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Set a DNS alias address for the Resource. Must be a DNS-valid name string.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Set the resource as active or inactive. Default is `true`.
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        [Input("isAuthoritative")]
        public Input<bool>? IsAuthoritative { get; set; }

        /// <summary>
        /// Controls whether an "Open in Browser" shortcut will be shown for this Resource in the Twingate Client. Default is `false`.
        /// </summary>
        [Input("isBrowserShortcutEnabled")]
        public Input<bool>? IsBrowserShortcutEnabled { get; set; }

        /// <summary>
        /// Controls whether this Resource will be visible in the main Resource list in the Twingate Client. Default is `true`.
        /// </summary>
        [Input("isVisible")]
        public Input<bool>? IsVisible { get; set; }

        /// <summary>
        /// The name of the Resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Restrict access to certain protocols and ports. By default or when this argument is not defined, there is no restriction, and all protocols and ports are allowed.
        /// </summary>
        [Input("protocols")]
        public Input<Inputs.TwingateResourceProtocolsGetArgs>? Protocols { get; set; }

        /// <summary>
        /// Remote Network ID where the Resource lives
        /// </summary>
        [Input("remoteNetworkId")]
        public Input<string>? RemoteNetworkId { get; set; }

        /// <summary>
        /// The ID of a `twingate.getTwingateSecurityPolicy` to set as this Resource's Security Policy. Default is `Default Policy`.
        /// </summary>
        [Input("securityPolicyId")]
        public Input<string>? SecurityPolicyId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The `tags` attribute consists of a key-value pairs that correspond with tags to be set on the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public TwingateResourceState()
        {
        }
        public static new TwingateResourceState Empty => new TwingateResourceState();
    }
}
