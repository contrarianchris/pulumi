// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import {Custom} from "./index";

/**
 * A component resource that accepts a reference to a custom resource. The input resource's `value` is used to create a child custom resource inside the component, before a reference to this child is returned.
 */
export class ComponentCustomRefInputOutput extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'component:index:ComponentCustomRefInputOutput';

    /**
     * Returns true if the given object is an instance of ComponentCustomRefInputOutput.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComponentCustomRefInputOutput {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComponentCustomRefInputOutput.__pulumiType;
    }

    public readonly inputRef!: pulumi.Output<Custom>;
    public /*out*/ readonly outputRef!: pulumi.Output<Custom>;

    /**
     * Create a ComponentCustomRefInputOutput resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentCustomRefInputOutputArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.inputRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'inputRef'");
            }
            resourceInputs["inputRef"] = args ? args.inputRef : undefined;
            resourceInputs["outputRef"] = undefined /*out*/;
        } else {
            resourceInputs["inputRef"] = undefined /*out*/;
            resourceInputs["outputRef"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComponentCustomRefInputOutput.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a ComponentCustomRefInputOutput resource.
 */
export interface ComponentCustomRefInputOutputArgs {
    inputRef: pulumi.Input<Custom>;
}