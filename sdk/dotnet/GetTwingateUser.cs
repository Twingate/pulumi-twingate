// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Twingate.Twingate
{
    public static class GetTwingateUser
    {
        /// <summary>
        /// Users in Twingate can be given access to Twingate Resources and may either be added manually or automatically synchronized with a 3rd party identity provider. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/users).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateUser.Invoke(new()
        ///     {
        ///         Id = "&lt;your user's id&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTwingateUserResult> InvokeAsync(GetTwingateUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTwingateUserResult>("twingate:index/getTwingateUser:getTwingateUser", args ?? new GetTwingateUserArgs(), options.WithDefaults());

        /// <summary>
        /// Users in Twingate can be given access to Twingate Resources and may either be added manually or automatically synchronized with a 3rd party identity provider. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/users).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateUser.Invoke(new()
        ///     {
        ///         Id = "&lt;your user's id&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTwingateUserResult> Invoke(GetTwingateUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTwingateUserResult>("twingate:index/getTwingateUser:getTwingateUser", args ?? new GetTwingateUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Users in Twingate can be given access to Twingate Resources and may either be added manually or automatically synchronized with a 3rd party identity provider. For more information, see Twingate's [documentation](https://docs.twingate.com/docs/users).
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Twingate = Pulumi.Twingate;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Twingate.GetTwingateUser.Invoke(new()
        ///     {
        ///         Id = "&lt;your user's id&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTwingateUserResult> Invoke(GetTwingateUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTwingateUserResult>("twingate:index/getTwingateUser:getTwingateUser", args ?? new GetTwingateUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTwingateUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the User. The ID for the User can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTwingateUserArgs()
        {
        }
        public static new GetTwingateUserArgs Empty => new GetTwingateUserArgs();
    }

    public sealed class GetTwingateUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the User. The ID for the User can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTwingateUserInvokeArgs()
        {
        }
        public static new GetTwingateUserInvokeArgs Empty => new GetTwingateUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetTwingateUserResult
    {
        /// <summary>
        /// The email address of the User
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The first name of the User
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// The ID of the User. The ID for the User can be obtained from the Admin API or the URL string in the Admin Console.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The last name of the User
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// Indicates the User's role. Either ADMIN, DEVOPS, SUPPORT, or MEMBER
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Indicates the User's type. Either MANUAL or SYNCED.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTwingateUserResult(
            string email,

            string firstName,

            string id,

            string lastName,

            string role,

            string type)
        {
            Email = email;
            FirstName = firstName;
            Id = id;
            LastName = lastName;
            Role = role;
            Type = type;
        }
    }
}
