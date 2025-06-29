// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the twingate package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'twingate';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The access key for API operations. You can retrieve this from the Twingate Admin Console
     * ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
     * TWINGATE_API_TOKEN environment variable.
     */
    public readonly apiToken!: pulumi.Output<string | undefined>;
    /**
     * Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
     * `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
     * environment variable.
     */
    public readonly network!: pulumi.Output<string | undefined>;
    /**
     * The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
     */
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["apiToken"] = args?.apiToken ? pulumi.secret(args.apiToken) : undefined;
            resourceInputs["cache"] = pulumi.output(args ? args.cache : undefined).apply(JSON.stringify);
            resourceInputs["defaultTags"] = pulumi.output(args ? args.defaultTags : undefined).apply(JSON.stringify);
            resourceInputs["httpMaxRetry"] = pulumi.output(args ? args.httpMaxRetry : undefined).apply(JSON.stringify);
            resourceInputs["httpTimeout"] = pulumi.output(args ? args.httpTimeout : undefined).apply(JSON.stringify);
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }

    /**
     * This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
     */
    terraformConfig(): pulumi.Output<Provider.TerraformConfigResult> {
        return pulumi.runtime.call("pulumi:providers:twingate/terraformConfig", {
            "__self__": this,
        }, this);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The access key for API operations. You can retrieve this from the Twingate Admin Console
     * ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
     * TWINGATE_API_TOKEN environment variable.
     */
    apiToken?: pulumi.Input<string>;
    /**
     * Specifies the cache settings for the provider.
     */
    cache?: pulumi.Input<inputs.ProviderCache>;
    /**
     * A default set of tags applied globally to all resources created by the provider.
     */
    defaultTags?: pulumi.Input<inputs.ProviderDefaultTags>;
    /**
     * Specifies a retry limit for the http requests made. The default value is 10. Alternatively, this can be specified using
     * the TWINGATE_HTTP_MAX_RETRY environment variable
     */
    httpMaxRetry?: pulumi.Input<number>;
    /**
     * Specifies a time limit in seconds for the http requests made. The default value is 35 seconds. Alternatively, this can
     * be specified using the TWINGATE_HTTP_TIMEOUT environment variable
     */
    httpTimeout?: pulumi.Input<number>;
    /**
     * Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
     * `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
     * environment variable.
     */
    network?: pulumi.Input<string>;
    /**
     * The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
     */
    url?: pulumi.Input<string>;
}

export namespace Provider {
    /**
     * The results of the Provider.terraformConfig method.
     */
    export interface TerraformConfigResult {
        readonly result: {[key: string]: any};
    }

}
