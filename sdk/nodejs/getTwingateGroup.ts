// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Groups are how users are authorized to access Resources. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/groups).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateGroup({
 *     id: "<your group's id>",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTwingateGroup(args: GetTwingateGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("twingate:index/getTwingateGroup:getTwingateGroup", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateGroup.
 */
export interface GetTwingateGroupArgs {
    /**
     * The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: string;
}

/**
 * A collection of values returned by getTwingateGroup.
 */
export interface GetTwingateGroupResult {
    /**
     * The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
     */
    readonly id: string;
    /**
     * Indicates if the Group is active
     */
    readonly isActive: boolean;
    /**
     * The name of the Group
     */
    readonly name: string;
    /**
     * The Security Policy assigned to the Group.
     */
    readonly securityPolicyId: string;
    /**
     * The type of the Group
     */
    readonly type: string;
}
/**
 * Groups are how users are authorized to access Resources. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/groups).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateGroup({
 *     id: "<your group's id>",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTwingateGroupOutput(args: GetTwingateGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTwingateGroupResult> {
    return pulumi.output(args).apply((a: any) => getTwingateGroup(a, opts))
}

/**
 * A collection of arguments for invoking getTwingateGroup.
 */
export interface GetTwingateGroupOutputArgs {
    /**
     * The ID of the Group. The ID for the Group can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: pulumi.Input<string>;
}
