// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Groups are how users are authorized to access Resources. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/groups).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@twingate/pulumi-twingate";
 *
 * const aws = new twingate.TwingateGroup("aws", {});
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import twingate:index/twingateGroup:TwingateGroup aws R3JvdXA6MzQ4OTE=
 * ```
 */
export class TwingateGroup extends pulumi.CustomResource {
    /**
     * Get an existing TwingateGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TwingateGroupState, opts?: pulumi.CustomResourceOptions): TwingateGroup {
        return new TwingateGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'twingate:index/twingateGroup:TwingateGroup';

    /**
     * Returns true if the given object is an instance of TwingateGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TwingateGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TwingateGroup.__pulumiType;
    }

    public readonly isAuthoritative!: pulumi.Output<boolean>;
    /**
     * The name of the group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Defines which Security Policy applies to this Group. The Security Policy ID can be obtained from the `twingate.getTwingateSecurityPolicy` and `twingate.getTwingateSecurityPolicies` data sources.
     */
    public readonly securityPolicyId!: pulumi.Output<string>;
    /**
     * List of User IDs that have permission to access the Group.
     */
    public readonly userIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a TwingateGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TwingateGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TwingateGroupArgs | TwingateGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TwingateGroupState | undefined;
            resourceInputs["isAuthoritative"] = state ? state.isAuthoritative : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["securityPolicyId"] = state ? state.securityPolicyId : undefined;
            resourceInputs["userIds"] = state ? state.userIds : undefined;
        } else {
            const args = argsOrState as TwingateGroupArgs | undefined;
            resourceInputs["isAuthoritative"] = args ? args.isAuthoritative : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["securityPolicyId"] = args ? args.securityPolicyId : undefined;
            resourceInputs["userIds"] = args ? args.userIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TwingateGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TwingateGroup resources.
 */
export interface TwingateGroupState {
    isAuthoritative?: pulumi.Input<boolean>;
    /**
     * The name of the group
     */
    name?: pulumi.Input<string>;
    /**
     * Defines which Security Policy applies to this Group. The Security Policy ID can be obtained from the `twingate.getTwingateSecurityPolicy` and `twingate.getTwingateSecurityPolicies` data sources.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * List of User IDs that have permission to access the Group.
     */
    userIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a TwingateGroup resource.
 */
export interface TwingateGroupArgs {
    isAuthoritative?: pulumi.Input<boolean>;
    /**
     * The name of the group
     */
    name?: pulumi.Input<string>;
    /**
     * Defines which Security Policy applies to this Group. The Security Policy ID can be obtained from the `twingate.getTwingateSecurityPolicy` and `twingate.getTwingateSecurityPolicies` data sources.
     */
    securityPolicyId?: pulumi.Input<string>;
    /**
     * List of User IDs that have permission to access the Group.
     */
    userIds?: pulumi.Input<pulumi.Input<string>[]>;
}
