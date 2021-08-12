variable "name" {
  type = string
}

variable "heroku_team" {
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

  organization {
    name = var.heroku_team
  }
}

resource "heroku_addon" "database" {
  app  = heroku_app.api.name
  plan = "heroku-postgresql:hobby-basic"
}

resource "heroku_build" "api" {
  app        = heroku_app.api.id
  buildpacks = ["heroku/go"]

  source {
    path = "api"
  }
}

resource "heroku_formation" "api" {
  app        = heroku_app.api.id
  type       = "web"
  quantity   = 1
  size       = "Standard-1x"
  depends_on = [heroku_build.api]
}