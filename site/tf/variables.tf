variable "docker_repo" {
    type = string
}
variable "docker_image_tag" {
    type = string
}

variable "environment" {
    type = string
    default = "localtesting"
}
