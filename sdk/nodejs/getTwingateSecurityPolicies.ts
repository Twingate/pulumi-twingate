// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Security Policies are defined in the Twingate Admin Console and determine user and device authentication requirements for Resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const all = twingate.getTwingateSecurityPolicies({
 *     name: "<your security policy's name>",
 * });
 * ```
 */
export function getTwingateSecurityPolicies(args?: GetTwingateSecurityPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetTwingateSecurityPoliciesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("twingate:index/getTwingateSecurityPolicies:getTwingateSecurityPolicies", {
        "name": args.name,
        "nameContains": args.nameContains,
        "nameExclude": args.nameExclude,
        "namePrefix": args.namePrefix,
        "nameRegexp": args.nameRegexp,
        "nameSuffix": args.nameSuffix,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateSecurityPolicies.
 */
export interface GetTwingateSecurityPoliciesArgs {
    /**
     * Returns only security policies that exactly match this name. If no options are passed it will return all security policies. Only one option can be used at a time.
     */
    name?: string;
    /**
     * Match when the value exist in the name of the security policy.
     */
    nameContains?: string;
    /**
     * Match when the exact value does not exist in the name of the security policy.
     */
    nameExclude?: string;
    /**
     * The name of the security policy must start with the value.
     */
    namePrefix?: string;
    /**
     * The regular expression match of the name of the security policy.
     */
    nameRegexp?: string;
    /**
     * The name of the security policy must end with the value.
     */
    nameSuffix?: string;
}

/**
 * A collection of values returned by getTwingateSecurityPolicies.
 */
export interface GetTwingateSecurityPoliciesResult {
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * Returns only security policies that exactly match this name. If no options are passed it will return all security policies. Only one option can be used at a time.
     */
    readonly name?: string;
    /**
     * Match when the value exist in the name of the security policy.
     */
    readonly nameContains?: string;
    /**
     * Match when the exact value does not exist in the name of the security policy.
     */
    readonly nameExclude?: string;
    /**
     * The name of the security policy must start with the value.
     */
    readonly namePrefix?: string;
    /**
     * The regular expression match of the name of the security policy.
     */
    readonly nameRegexp?: string;
    /**
     * The name of the security policy must end with the value.
     */
    readonly nameSuffix?: string;
    readonly securityPolicies: outputs.GetTwingateSecurityPoliciesSecurityPolicy[];
}
/**
 * Security Policies are defined in the Twingate Admin Console and determine user and device authentication requirements for Resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as twingate from "@pulumi/twingate";
 *
 * const all = twingate.getTwingateSecurityPolicies({
 *     name: "<your security policy's name>",
 * });
 * ```
 */
export function getTwingateSecurityPoliciesOutput(args?: GetTwingateSecurityPoliciesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTwingateSecurityPoliciesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("twingate:index/getTwingateSecurityPolicies:getTwingateSecurityPolicies", {
        "name": args.name,
        "nameContains": args.nameContains,
        "nameExclude": args.nameExclude,
        "namePrefix": args.namePrefix,
        "nameRegexp": args.nameRegexp,
        "nameSuffix": args.nameSuffix,
    }, opts);
}

/**
 * A collection of arguments for invoking getTwingateSecurityPolicies.
 */
export interface GetTwingateSecurityPoliciesOutputArgs {
    /**
     * Returns only security policies that exactly match this name. If no options are passed it will return all security policies. Only one option can be used at a time.
     */
    name?: pulumi.Input<string>;
    /**
     * Match when the value exist in the name of the security policy.
     */
    nameContains?: pulumi.Input<string>;
    /**
     * Match when the exact value does not exist in the name of the security policy.
     */
    nameExclude?: pulumi.Input<string>;
    /**
     * The name of the security policy must start with the value.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The regular expression match of the name of the security policy.
     */
    nameRegexp?: pulumi.Input<string>;
    /**
     * The name of the security policy must end with the value.
     */
    nameSuffix?: pulumi.Input<string>;
}
