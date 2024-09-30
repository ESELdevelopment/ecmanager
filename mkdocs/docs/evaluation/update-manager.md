# Update manager

We want to build a secure and up-to-date software. Therefore, we have to update
out code regularly. To make this process as easy as possible, we want to use an
update manager.

This eval compares 2 different update managers: `Renovate` and `Dependabot`.

## Requirements

Following requirements are important for us:

- group dependencies by type (e.g. `go-Dependencies`, `python-Dependencies`)
- easy to use
- minimal overhead to maintain

## Renovate

[Renovate](https://docs.renovatebot.com/) is tool, build by Mend and can be
self-hosted or used as a service. For configuration, it uses a `renovate.json`
file in the repository.

To use Renovate, we can either host a server and deploy Renovate or
install the Renovate GitHub [App](https://github.com/apps/renovate)
(for the entire Organization).

## Dependabot

[Dependabot](https://github.com/dependabot) is a tool, build by GitHub and can
be used as a service. For configuration, it uses a `.github/dependabot.yml` file
in the repository.

To use Dependabot, we have to enable it in the
[repository settings](https://docs.github.com/de/code-security/getting-started/dependabot-quickstart-guide).

## Comparison

<!-- markdownlint-disable MD013 -->
| Feature       | Renovate                                                                                          | Dependabot                                          |
|---------------|---------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| Hosting       | :red_circle: as service, connected via GitHub App (self hosted is possible, but no option for us) | :green_circle: native in GitHub                     |
| Configuration | :green_circle: endless configuration via JSON and Environment Variables                           | :orange: limited, but simple configuration via YAML |
| Cost          | :green_circle: free                                                                               | :green_circle: free                                 |
| Security      | :red_circle: hosted by Mend (write access from an app)                                            | :green_circle: native in GitHub, no rights needed   |
| Grouping      | :green_circle: via `packageRules`                                                                 | :white_circle: no grouping                          |
<!-- markdownlint-enable MD013 -->

## Conclusion

We have a very simple setup and don't need a lot of configuration. Therefore,
Dependabot is the better choice for us, because it is native in GitHub and has
no overhead to maintain. Because of the reduced features, it hat also a
reduced complexity.

{{ decision("We use dependabot") }}
