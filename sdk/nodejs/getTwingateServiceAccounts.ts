// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateServiceAccounts({
 *     name: "<your service account's name>",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTwingateServiceAccounts(args?: GetTwingateServiceAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateServiceAccountsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("twingate:index/getTwingateServiceAccounts:getTwingateServiceAccounts", {
        "name": args.name,
        "nameContains": args.nameContains,
        "nameExclude": args.nameExclude,
        "namePrefix": args.namePrefix,
        "nameRegexp": args.nameRegexp,
        "nameSuffix": args.nameSuffix,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateServiceAccounts.
 */
export interface GetTwingateServiceAccountsArgs {
    /**
     * Name of the Service Account
     */
    name?: string;
    /**
     * Match when the value exist in the name of the service account.
     */
    nameContains?: string;
    /**
     * Match when the exact value does not exist in the name of the service account.
     */
    nameExclude?: string;
    /**
     * The name of the service account must start with the value.
     */
    namePrefix?: string;
    /**
     * The regular expression match of the name of the service account.
     */
    nameRegexp?: string;
    /**
     * The name of the service account must end with the value.
     */
    nameSuffix?: string;
}

/**
 * A collection of values returned by getTwingateServiceAccounts.
 */
export interface GetTwingateServiceAccountsResult {
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * Returns only service accounts that exactly match this name. If no options are passed it will return all service accounts. Only one option can be used at a time.
     */
    readonly name?: string;
    /**
     * Match when the value exist in the name of the service account.
     */
    readonly nameContains?: string;
    /**
     * Match when the exact value does not exist in the name of the service account.
     */
    readonly nameExclude?: string;
    /**
     * The name of the service account must start with the value.
     */
    readonly namePrefix?: string;
    /**
     * The regular expression match of the name of the service account.
     */
    readonly nameRegexp?: string;
    /**
     * The name of the service account must end with the value.
     */
    readonly nameSuffix?: string;
    /**
     * List of Service Accounts
     */
    readonly serviceAccounts: outputs.GetTwingateServiceAccountsServiceAccount[];
}
/**
 * Service Accounts offer a way to provide programmatic, centrally-controlled, and consistent access controls.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const foo = twingate.getTwingateServiceAccounts({
 *     name: "<your service account's name>",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getTwingateServiceAccountsOutput(args?: GetTwingateServiceAccountsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTwingateServiceAccountsResult> {
    return pulumi.output(args).apply((a: any) => getTwingateServiceAccounts(a, opts))
}

/**
 * A collection of arguments for invoking getTwingateServiceAccounts.
 */
export interface GetTwingateServiceAccountsOutputArgs {
    /**
     * Name of the Service Account
     */
    name?: pulumi.Input<string>;
    /**
     * Match when the value exist in the name of the service account.
     */
    nameContains?: pulumi.Input<string>;
    /**
     * Match when the exact value does not exist in the name of the service account.
     */
    nameExclude?: pulumi.Input<string>;
    /**
     * The name of the service account must start with the value.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The regular expression match of the name of the service account.
     */
    nameRegexp?: pulumi.Input<string>;
    /**
     * The name of the service account must end with the value.
     */
    nameSuffix?: pulumi.Input<string>;
}
