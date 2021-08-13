# Recipe backend

This just a testing repository for me to familiarise myself more with backend concepts, after doing a lot of frontend work. This hopefully will include automatic deployments, separate languages, communication between service and security.

## Folder structure

* [.github/workflows](.github/workflows) Github actions defined here that test code in PRs and check that the proposed code can be deployed
* [./api](./api) A small Go API for recipes

## Deploying

The code is deployed with terraform to [Heroku](heroku.com), if code is merged to `main` it will check if changes were made to the relevant parts of the repository and then update the deployment automatically

* [main.tf](./main.tf) The terraform definition for what needs to be deployed to Heroku
* [.github/workflows/terraform-plan.yml](.github/workflows/terraform-plan.yml) The action runs on any PR to master and will determine the changes to the deployment that will be made and adds the information to the PR
* [.github/workflows/terraform-apply.yml](.github/workflows/terraform-apply.yml) This action runs when a push to `main` happens. It will apply all the necessary changes. This should generally be the same as what was created in the plan step, but they are not directly connected so they might differ.