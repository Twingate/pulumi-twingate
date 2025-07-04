// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * A Remote Network represents a single private network in Twingate that can have one or more Connectors and Resources assigned to it. You must create a Remote Network before creating Resources and Connectors that belong to it. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/remote-networks).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@twingate/pulumi-twingate";
 *
 * const awsNetwork = new twingate.TwingateRemoteNetwork("awsNetwork", {});
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import twingate:index/twingateRemoteNetwork:TwingateRemoteNetwork network UmVtb3RlTmV0d29zaipgMKIkNg==
 * ```
 */
export class TwingateRemoteNetwork extends pulumi.CustomResource {
    /**
     * Get an existing TwingateRemoteNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TwingateRemoteNetworkState, opts?: pulumi.CustomResourceOptions): TwingateRemoteNetwork {
        return new TwingateRemoteNetwork(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'twingate:index/twingateRemoteNetwork:TwingateRemoteNetwork';

    /**
     * Returns true if the given object is an instance of TwingateRemoteNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TwingateRemoteNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TwingateRemoteNetwork.__pulumiType;
    }

    /**
     * The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE*CLOUD, ON*PREMISE, OTHER.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Remote Network
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the Remote Network. Must be one of the following: REGULAR, EXIT. Defaults to REGULAR.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a TwingateRemoteNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TwingateRemoteNetworkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TwingateRemoteNetworkArgs | TwingateRemoteNetworkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TwingateRemoteNetworkState | undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TwingateRemoteNetworkArgs | undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TwingateRemoteNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TwingateRemoteNetwork resources.
 */
export interface TwingateRemoteNetworkState {
    /**
     * The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE*CLOUD, ON*PREMISE, OTHER.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Remote Network
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the Remote Network. Must be one of the following: REGULAR, EXIT. Defaults to REGULAR.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TwingateRemoteNetwork resource.
 */
export interface TwingateRemoteNetworkArgs {
    /**
     * The location of the Remote Network. Must be one of the following: AWS, AZURE, GOOGLE*CLOUD, ON*PREMISE, OTHER.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Remote Network
     */
    name?: pulumi.Input<string>;
    /**
     * The type of the Remote Network. Must be one of the following: REGULAR, EXIT. Defaults to REGULAR.
     */
    type?: pulumi.Input<string>;
}
