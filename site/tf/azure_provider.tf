# Configure the Azure Provider
provider "azurerm" {
  # whilst the `version` attribute is optional, we recommend pinning to a given version of the Provider
  version = "=1.38.0"
}

# Create a resource group
resource "azurerm_resource_group" "gifm" {
  name     = "gifm"
  location = "West US"
}

