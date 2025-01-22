// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Resources in Twingate represent any network destination address that you wish to provide private access to for users authorized via the Twingate Client application. Resources can be defined by either IP or DNS address, and all private DNS addresses will be automatically resolved with no client configuration changes. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateResource({
 *     id: "<your resource's id>",
 * });
 * ```
 */
export function getTwingateResource(args: GetTwingateResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateResourceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("twingate:index/getTwingateResource:getTwingateResource", {
        "id": args.id,
        "protocols": args.protocols,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateResource.
 */
export interface GetTwingateResourceArgs {
    /**
     * The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: string;
    /**
     * By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
     */
    protocols?: inputs.GetTwingateResourceProtocols;
}

/**
 * A collection of values returned by getTwingateResource.
 */
export interface GetTwingateResourceResult {
    /**
     * The Resource's address, which may be an IP address, CIDR range, or DNS address
     */
    readonly address: string;
    /**
     * The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
     */
    readonly id: string;
    /**
     * The name of the Resource
     */
    readonly name: string;
    /**
     * By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
     */
    readonly protocols?: outputs.GetTwingateResourceProtocols;
    /**
     * The Remote Network ID that the Resource is associated with. Resources may only be associated with a single Remote Network.
     */
    readonly remoteNetworkId: string;
}
/**
 * Resources in Twingate represent any network destination address that you wish to provide private access to for users authorized via the Twingate Client application. Resources can be defined by either IP or DNS address, and all private DNS addresses will be automatically resolved with no client configuration changes. For more information, see the Twingate [documentation](https://docs.twingate.com/docs/resources-and-access-nodes).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateResource({
 *     id: "<your resource's id>",
 * });
 * ```
 */
export function getTwingateResourceOutput(args: GetTwingateResourceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTwingateResourceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("twingate:index/getTwingateResource:getTwingateResource", {
        "id": args.id,
        "protocols": args.protocols,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateResource.
 */
export interface GetTwingateResourceOutputArgs {
    /**
     * The ID of the Resource. The ID for the Resource can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: pulumi.Input<string>;
    /**
     * By default (when this argument is not defined) no restriction is applied, and all protocols and ports are allowed.
     */
    protocols?: pulumi.Input<inputs.GetTwingateResourceProtocolsArgs>;
}
