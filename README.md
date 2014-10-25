# GitHub Bot

Define rules for automated actions to happen on your repository based on labels being added/removed.

## Example
```json
{
    "hook-address": "https://git-control.com",
    "repos": {
        "robstrong/hook-receiver": [
            {
                "event": "pull_request",
                "action": "labeled",
                "label": "to be merged to test",
                "": {
                    "type": "merge-pull-request",
                    "params": {
                        "branch": "testing"
                    }
                }
            }
        ]
    }
}
```
When starting GitHub Bot with this configuration, it will first register a web hook for `robstrong/hook-receiver` if one doesn't exist. Then when a label named "to be merged to test" is added to a pull request, GitHub Bot will receive the web hook from GitHub and perform the action specified in the config. In this case, it will merge the pull request branch into the "testing" branch and push the changes back up to GitHub.
