variable "name" {
  type = string
}

variable "heroku_region" {
  type    = string
  default = "eu"
}

locals {
  recipe_app_name = "${var.name}-recipe-app"
}

terraform {
  backend "pg" {}
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "4.6.0"
    }
  }
}

provider "heroku" {
}

resource "heroku_app" "api" {
  name   = "${local.recipe_app_name}-api"
  region = var.heroku_region

  config_vars = {
    GIN_MODE = "release"
  }
}

resource "heroku_addon" "database" {
  app  = heroku_app.api.name
  plan = "heroku-postgresql:hobby-dev"
}

resource "heroku_build" "api" {
  app        = heroku_app.api.id
  buildpacks = ["https://github.com/heroku/heroku-buildpack-go.git"]

  source {
    path = "api"
  }
}

resource "heroku_formation" "api" {
  app        = heroku_app.api.id
  type       = "web"
  quantity   = 1
  size       = "free"
  depends_on = [heroku_build.api]
}