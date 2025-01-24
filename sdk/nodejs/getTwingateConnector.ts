// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Connectors provide connectivity to Remote Networks. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateConnector({
 *     id: "<your connector's id>",
 * });
 * ```
 */
export function getTwingateConnector(args: GetTwingateConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("twingate:index/getTwingateConnector:getTwingateConnector", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateConnector.
 */
export interface GetTwingateConnectorArgs {
    /**
     * The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: string;
}

/**
 * A collection of values returned by getTwingateConnector.
 */
export interface GetTwingateConnectorResult {
    /**
     * The hostname of the machine hosting the Connector.
     */
    readonly hostname: string;
    /**
     * The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
     */
    readonly id: string;
    /**
     * The name of the Connector.
     */
    readonly name: string;
    /**
     * The Connector's private IP addresses.
     */
    readonly privateIps: string[];
    /**
     * The Connector's public IP address.
     */
    readonly publicIp: string;
    /**
     * The ID of the Remote Network the Connector is attached to.
     */
    readonly remoteNetworkId: string;
    /**
     * The Connector's state. One of `ALIVE`, `DEAD_NO_HEARTBEAT`, `DEAD_HEARTBEAT_TOO_OLD` or `DEAD_NO_RELAYS`.
     */
    readonly state: string;
    /**
     * Determines whether status notifications are enabled for the Connector.
     */
    readonly statusUpdatesEnabled: boolean;
    /**
     * The Connector's version.
     */
    readonly version: string;
}
/**
 * Connectors provide connectivity to Remote Networks. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/understanding-access-nodes).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateConnector({
 *     id: "<your connector's id>",
 * });
 * ```
 */
export function getTwingateConnectorOutput(args: GetTwingateConnectorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTwingateConnectorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("twingate:index/getTwingateConnector:getTwingateConnector", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateConnector.
 */
export interface GetTwingateConnectorOutputArgs {
    /**
     * The ID of the Connector. The ID for the Connector can be obtained from the Admin API or the URL string in the Admin Console.
     */
    id: pulumi.Input<string>;
}
