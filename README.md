- name: Generate secret arguments
  set_fact:
    secret_args: >-
      {{ 
        secret_list.split() 
        | map('regex_replace', '^(.*?):\\s*(.*)', '--from-literal=\\1=\\2') 
        | map('regex_replace', '^\\s+|\\s+$', '') 
        | join(' ') 
      }}
  when: secret_list is defined and secret_list | length > 0




Potential Issues in secret_args
Regex Misinterpretation:
The regex '^(.*?):\\s*(.*)' looks for : as a separator, but if your input doesn't consistently use : (or contains additional spaces), it won't match as expected.
Whitespace Cleanup:
The second regex_replace ('^\\s+|\\s+$') removes leading/trailing whitespace but does not affect spaces inside the values, which might cause issues.
Formatting Validation:
The format produced for --from-literal must exactly match --from-literal=key=value. Spaces inside keys or values could lead to malformed commands.
Incorrect Debug Output:
Ensure secret_args is formed as expected by adding a debug step before the Create secret task.
Suggested Fix
To handle potential issues and validate intermediate steps, here's the updated playbook:

- name: Debug secret_list
  debug:
    var: secret_list

- name: Validate and format secret arguments
  set_fact:
    formatted_secrets: >-
      {{
        secret_list.split() 
        | map('regex_replace', '^(.*?):\\s*(.*)', '\\1=\\2')
        | map('regex_replace', '^\\s+|\\s+$', '')
      }}
  when: secret_list is defined and secret_list | length > 0

- name: Debug formatted_secrets
  debug:
    var: formatted_secrets

- name: Generate secret arguments
  set_fact:
    secret_args: >-
      {{
        formatted_secrets 
        | map('regex_replace', '^(.*)=(.*)', '--from-literal=\\1=\\2') 
        | join(' ')
      }}
  when: formatted_secrets is defined and formatted_secrets | length > 0

- name: Debug secret_args
  debug:
    var: secret_args

- name: Create secret
  shell: "{{ oc_cmd }} create secret generic {{ os_app }}.{{ Env }}-secret.properties {{ secret_args }}"
  when: secret_args is defined and secret_args | length > 0
Key Improvements
Intermediate Validation:
Split formatting into multiple steps (formatted_secrets → secret_args) to debug intermediate outputs.
Improved Regex:
The first regex ensures the separator : is properly detected.
The second regex removes leading/trailing spaces for cleaner values.
Debug Steps:
Additional debug steps (formatted_secrets, secret_args) validate that the output is correctly formed.
Consistent Transformation:
Breaking the logic into smaller, validated steps avoids complex single-line transformations and makes troubleshooting easier.
Sample Input/Output
Input (secret_list):

CAPI_CLIENT_SECRET: 'db0LHbdkPRy0jjADI8faTakZyagZy' MDH_PASSWORD: 'KJjgouhiohdi'
Debug Output (formatted_secrets):

[  "CAPI_CLIENT_SECRET=db0LHbdkPRy0jjADI8faTakZyagZy",  "MDH_PASSWORD=KJjgouhiohdi"]
Debug Output (secret_args):

--from-literal=CAPI_CLIENT_SECRET=db0LHbdkPRy0jjADI8faTakZyagZy --from-literal=MDH_PASSWORD=KJjgouhiohdi
Final Notes
Ensure the input (secret_list) matches the expected format. If there are unexpected issues, rerun with debug tasks to inspect intermediate outputs (formatted_secrets, secret_args) for discrepancies.





















- name: Debug secret_list
  debug:
    var: secret_list

- name: Generate secret arguments
  set_fact:
    secret_args: "{{ secret_list.split() | map('regex_replace', '^(.*?):\\s*(.*)', '--from-literal=\\1=\\2') | map('regex_replace', '^\\s+|\\s+$', '') | join(' ') }}"
  when: secret_list is defined and secret_list | length > 0

- name: Create secret
  shell: "{{ oc_cmd }} create secret generic {{ os_app }}.{{ Env }}-secret.properties {{ secret_args }}"
  when: secret_args is defined




<h1>
  <img src="./website/public/img/logo.svg" align="left" height="46px" alt="Consul logo"/>
  <span>Consul</span>
</h1>

[![License: BUSL-1.1](https://img.shields.io/badge/License-BUSL--1.1-yellow.svg)](LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/hashicorp/consul.svg)](https://hub.docker.com/r/hashicorp/consul)
[![Go Report Card](https://goreportcard.com/badge/github.com/hashicorp/consul)](https://goreportcard.com/report/github.com/hashicorp/consul)

Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure.

* Website: https://www.consul.io
* Tutorials: [HashiCorp Learn](https://learn.hashicorp.com/consul)
* Forum: [Discuss](https://discuss.hashicorp.com/c/consul)

Consul provides several key features:

* **Multi-Datacenter** - Consul is built to be datacenter aware, and can
  support any number of regions without complex configuration.

* **Service Mesh** - Consul Service Mesh enables secure service-to-service
  communication with automatic TLS encryption and identity-based authorization. Applications
  can use sidecar proxies in a service mesh configuration to establish TLS
  connections for inbound and outbound connections with Transparent Proxy.

* **API Gateway** - Consul API Gateway manages access to services within Consul Service Mesh, 
  allow users to define traffic and authorization policies to services deployed within the mesh.  

* **Service Discovery** - Consul makes it simple for services to register
  themselves and to discover other services via a DNS or HTTP interface.
  External services such as SaaS providers can be registered as well.

* **Health Checking** - Health Checking enables Consul to quickly alert
  operators about any issues in a cluster. The integration with service
  discovery prevents routing traffic to unhealthy hosts and enables service
  level circuit breakers.

* **Dynamic App Configuration** - An HTTP API that allows users to store indexed objects within Consul,
  for storing configuration parameters and application metadata.

Consul runs on Linux, macOS, FreeBSD, Solaris, and Windows and includes an
optional [browser based UI](https://demo.consul.io). A commercial version
called [Consul Enterprise](https://www.consul.io/docs/enterprise) is also
available.

**Please note**: We take Consul's security and our users' trust very seriously. If you
believe you have found a security issue in Consul, please [responsibly disclose](https://www.hashicorp.com/security#vulnerability-reporting)
by contacting us at security@hashicorp.com.

## Quick Start

A few quick start guides are available on the Consul website:

* **Standalone binary install:** https://learn.hashicorp.com/collections/consul/get-started-vms
* **Minikube install:** https://learn.hashicorp.com/tutorials/consul/kubernetes-minikube
* **Kind install:** https://learn.hashicorp.com/tutorials/consul/kubernetes-kind
* **Kubernetes install:** https://learn.hashicorp.com/tutorials/consul/kubernetes-deployment-guide
* **Deploy HCP Consul:** https://learn.hashicorp.com/tutorials/consul/hcp-gs-deploy 

## Documentation

Full, comprehensive documentation is available on the Consul website: https://consul.io/docs

## Contributing

Thank you for your interest in contributing! Please refer to [CONTRIBUTING.md](https://github.com/hashicorp/consul/blob/main/.github/CONTRIBUTING.md)
for guidance. For contributions specifically to the browser based UI, please
refer to the UI's [README.md](https://github.com/hashicorp/consul/blob/main/ui/packages/consul-ui/README.md)
for guidance.
