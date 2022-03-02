# CM Generator

CM Generator is a github action to compile a bunch of jsonnet files. It compiles files placed on in/ directory and write the result in one out/ directory.

## How to use it?

Create `.github/workflows/<workflow_name>.yml`

```yaml
on:
  pull_request:
    branches:
      - main
      - master
name: Pull request workflow
jobs:
  validate_configurations:
    name: Validate configurations
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: vtex/action-cm-generator@v0.1.1
        with:
          in: "in"
          out: "gen"
```

Or if you prefer, the docker version (17x faster)

```yaml
on:
  pull_request:
    branches:
      - main
      - master
name: Pull request workflow
jobs:
  validate_configurations:
    name: Validate configurations
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: docker://public.ecr.aws/n0a0a3c3/action-cm-generator:latest
        with:
          in: "in"
          out: "gen"
```

| Inputs | Required | Default | Description                            |
| ------ | -------- | ------- | -------------------------------------- |
| `in`   | No       | "in"    | The folder that contains jsonnet files |
| `out`  | No       | "out"   | The output folder                      |
