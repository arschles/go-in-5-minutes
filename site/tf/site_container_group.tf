resource "azurerm_container_group" "site" {
  name = "site"
  location = azurerm_resource_group.gifm.location
  resource_group_name = azurerm_resource_group.gifm.name
  ip_address_type = "public"
  dns_name_label = "gifm-site"
  os_type = "Linux"

  container {
    name = "site"
    # image = "${var.docker_repo}:${var.docker_image_tag}"
    image = "arschles/gifm-site:latest"
    cpu = "1"
    memory = "2"

    ports {
      port = 3000
      protocol = "TCP"
    }
  }

  tags = {
    environment = var.environment
  }
}
