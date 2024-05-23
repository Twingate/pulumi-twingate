# Python Example
This example demonstrate how Twingate remote network, service account, service key and resources can be created and configured in Python.

## Pre-requisite
* Python and PIP
* Pulumi

## How to Use
* Clone the repository
* `cd /path/to/repo/examples/python`
* `virtualenv venv`
* `source venv/bin/activate`
* `./venv/bin/pip install -r requirements.txt`
* Setup Pulumi-Twingate Provider, see configuration section [here](../../README.md)
* `pulumi up`

Note: `pulumi up` should automatically download the required Python dependency and Pulumi Plugins.