# List of Trust Lists

They have many names, such as ‘trusted root list’, ‘trusted root store’, ‘trust store’, ‘approved trust list’, etc. The PKI Consortium is curating a global List of Trust Lists (a list of root, intermediate or issuing CA certificates accepted by a public, private, industry, or solution-specific PKI), one that is not limited to a specific purpose, region or size, and is open to anyone to contribute.

Each list is documented as a YAML file and hosted in this repository. This makes it easier to read for humans while retaining version control and allowing systems to process and analyze the data.

Some lists will share a common purpose or audit regime, might be extensively documented, list policies, discussion groups, etc. Others focus on a specific purpose, region or use-case and might only have some basic information.

There are many trust lists and often there is little overlap or interoperability. With this project the PKI consortium is not only building a comprehensive list of trust lists but also a place where the industry can find each other, engage, share knowledge, policies and best practices to improve security, interoperability and mutual trust.

The PKI Consortium welcomes contributions and would like to engage in related activities from other organizations or stakeholders.

# Contributing to this repository 

You must read and adhere to the [code of conduct](https://pkic.org/code-of-conduct/) of the PKI Consortium.

Some good first issues can be found [here](https://github.com/pkic/ltl/contribute). Feel free to select [another issue](https://github.com/pkic/ltl/issues) within your expertise or create [a new issue](https://github.com/pkic/ltl/issues/new).

Changes can be made directly on github.com or using a plain text / YAML editor of your choice such a [Visual Studio Code](https://code.visualstudio.com/).

After creating a [Pull Request](https://docs.github.com/en/github/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork) an automatic check will verify if your changes comply with the [document model](https://github.com/pkic/ltl/blob/main/ltl/model.go), you must ensure the checks succeed before your changed can be accepted.

## Add a Trust List to the list

Please create an issue or pull request, the easiest way is to copy and modify an existing file.

## What fields can I use?

The available field and their requirements can be found in the [document model](https://github.com/pkic/ltl/blob/main/ltl/model.go).
