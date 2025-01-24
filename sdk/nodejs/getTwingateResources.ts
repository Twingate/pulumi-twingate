// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateResources({
 *     name: "<your resource's name>",
 * });
 * ```
 */
export function getTwingateResources(args?: GetTwingateResourcesArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateResourcesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("twingate:index/getTwingateResources:getTwingateResources", {
        "name": args.name,
        "nameContains": args.nameContains,
        "nameExclude": args.nameExclude,
        "namePrefix": args.namePrefix,
        "nameRegexp": args.nameRegexp,
        "nameSuffix": args.nameSuffix,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateResources.
 */
export interface GetTwingateResourcesArgs {
    /**
     * Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
     */
    name?: string;
    /**
     * Match when the value exist in the name of the resource.
     */
    nameContains?: string;
    /**
     * Match when the exact value does not exist in the name of the resource.
     */
    nameExclude?: string;
    /**
     * The name of the resource must start with the value.
     */
    namePrefix?: string;
    /**
     * The regular expression match of the name of the resource.
     */
    nameRegexp?: string;
    /**
     * The name of the resource must end with the value.
     */
    nameSuffix?: string;
}

/**
 * A collection of values returned by getTwingateResources.
 */
export interface GetTwingateResourcesResult {
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
     */
    readonly name?: string;
    /**
     * Match when the value exist in the name of the resource.
     */
    readonly nameContains?: string;
    /**
     * Match when the exact value does not exist in the name of the resource.
     */
    readonly nameExclude?: string;
    /**
     * The name of the resource must start with the value.
     */
    readonly namePrefix?: string;
    /**
     * The regular expression match of the name of the resource.
     */
    readonly nameRegexp?: string;
    /**
     * The name of the resource must end with the value.
     */
    readonly nameSuffix?: string;
    /**
     * List of Resources
     */
    readonly resources: outputs.GetTwingateResourcesResource[];
}
/**
 * Resources in Twingate represent servers on the private network that clients can connect to. Resources can be defined by IP, CIDR range, FQDN, or DNS zone. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateResources({
 *     name: "<your resource's name>",
 * });
 * ```
 */
export function getTwingateResourcesOutput(args?: GetTwingateResourcesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTwingateResourcesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("twingate:index/getTwingateResources:getTwingateResources", {
        "name": args.name,
        "nameContains": args.nameContains,
        "nameExclude": args.nameExclude,
        "namePrefix": args.namePrefix,
        "nameRegexp": args.nameRegexp,
        "nameSuffix": args.nameSuffix,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateResources.
 */
export interface GetTwingateResourcesOutputArgs {
    /**
     * Returns only resources that exactly match this name. If no options are passed it will return all resources. Only one option can be used at a time.
     */
    name?: pulumi.Input<string>;
    /**
     * Match when the value exist in the name of the resource.
     */
    nameContains?: pulumi.Input<string>;
    /**
     * Match when the exact value does not exist in the name of the resource.
     */
    nameExclude?: pulumi.Input<string>;
    /**
     * The name of the resource must start with the value.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The regular expression match of the name of the resource.
     */
    nameRegexp?: pulumi.Input<string>;
    /**
     * The name of the resource must end with the value.
     */
    nameSuffix?: pulumi.Input<string>;
}
