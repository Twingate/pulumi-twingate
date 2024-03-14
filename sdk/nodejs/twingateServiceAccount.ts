// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const githubActionsProd = new twingate.TwingateServiceAccount("githubActionsProd", {});
 * ```
 */
export class TwingateServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing TwingateServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TwingateServiceAccountState, opts?: pulumi.CustomResourceOptions): TwingateServiceAccount {
        return new TwingateServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'twingate:index/twingateServiceAccount:TwingateServiceAccount';

    /**
     * Returns true if the given object is an instance of TwingateServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TwingateServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TwingateServiceAccount.__pulumiType;
    }

    /**
     * The name of the Service Account in Twingate
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a TwingateServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TwingateServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TwingateServiceAccountArgs | TwingateServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TwingateServiceAccountState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as TwingateServiceAccountArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TwingateServiceAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TwingateServiceAccount resources.
 */
export interface TwingateServiceAccountState {
    /**
     * The name of the Service Account in Twingate
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TwingateServiceAccount resource.
 */
export interface TwingateServiceAccountArgs {
    /**
     * The name of the Service Account in Twingate
     */
    name?: pulumi.Input<string>;
}
