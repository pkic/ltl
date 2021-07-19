# Contributing to this repository 

You must read and adhere to the [code of conduct](https://pkic.org/code-of-conduct/) of the PKI Consortium.

Some good first issues can be found [here](https://github.com/pkic/ltl/contribute). Feel free to select [another issue](https://github.com/pkic/ltl/issues) within your expertise or create [a new issue](https://github.com/pkic/ltl/issues/new).

Changes can be made directly on github.com or using a plain text / YAML editor of your choice such a [Visual Studio Code](https://code.visualstudio.com/).

After creating a [Pull Request](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork) an automatic check will verify if your changes comply with the [document model](https://github.com/pkic/ltl/blob/main/ltl/model.go), you must ensure the checks succeed before your changed can be accepted.

## Add a Trust List to the list

Please create an issue or pull request, the easiest way is to copy and modify an existing file.

## What fields can I use?

The available field and their requirements can be found in the [document model](https://github.com/pkic/ltl/blob/main/ltl/model.go).

### Minimal fields
These are the bare minimal requirements to add a trust list. Please try to add as much as information possible using the supported fields.
```yaml
name: <string>

// A link to the trusted list program or vendor/product when no such
// dedicated website exists.
website: <url>
```

### Supported fields
Please check the document model for the currently supported fields, the example below might be out of date.

```yaml
name: <string>
description: <string>

// A link to the trusted list program or vendor/product when no such
// dedicated website exists.
website: <url>

// Email or website address to contact privately
// (do NOT include non-public contact information)
contact: <email|url>

// Using the Common CA Database
ccadb: <true|false>

// Legal context
legal-context: <string>

// Audit specifies the audit options and or requirements
audit:
  - name: <string>
    description: <string>
    frequency: 0s
    policies: [<string>, <string>, ...]

requirements:
  - type: <string>
    url: <url>

discussions:
  - type: <string>
    url: <url>

issues:
  - type: <string>
    url: <url>

// TrustLists under the same PKI program.
// If one list has multiple distinct programs, create a new file for each.
trust-lists:
    info: <string>
    policy: <string>
    
    // Trust links to a file containing the CA certificates trusted by this PKI
    // for the given purpose. A file might be referenced in multiple formats, the
    // preferred format (type) is PEM.
    trust:
        - purposes: [<string>, <string>, ...]
          
          // When a program has trust lists for different purposes, there can be
          // generic requirements and requirements that are specific to one purpose.
          requirements:
            - type: <oneof>
              url: <url>
              
          discussions:
            - type: <oneof>
              url: <url>
              
          issues:
            - type: <oneof>
              url: <url>
          
          // Audit specifies the audit options specific to this purpose.
          audit:
            - name: <string>
              description: <string>
              frequency: 0s
              policies: [<string>, <string>, ...]
              
          // Links to a file containing the CA certificates trusted by this PKI for the
          // given purpose. A file might be referenced in multiple formats, the preferred
          // format (type) is HTML and PEM.
          list:
            - type: <oneof>
              url: <url>
```
For the type `<oneof>` see https://github.com/pkic/ltl/blob/main/ltl/model.go#L68
