# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

apiToken: Optional[str]
"""
The access key for API operations. You can retrieve this from the Twingate Admin Console
([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
TWINGATE_API_TOKEN environment variable.
"""

httpMaxRetry: Optional[int]
"""
Specifies a retry limit for the http requests made. The default value is 10. Alternatively, this can be specified using
the TWINGATE_HTTP_MAX_RETRY environment variable
"""

httpTimeout: Optional[int]
"""
Specifies a time limit in seconds for the http requests made. The default value is 35 seconds. Alternatively, this can
be specified using the TWINGATE_HTTP_TIMEOUT environment variable
"""

network: Optional[str]
"""
Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
`autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
environment variable.
"""

url: Optional[str]
"""
The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
"""

