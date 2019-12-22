resource "azurerm_container_registry" "site" {
  name = "gifmsite"
  resource_group_name = azurerm_resource_group.gifm.name
  location = azurerm_resource_group.gifm.location
  sku = "Premium"
  admin_enabled = false
  georeplication_locations = ["West US"]
}
