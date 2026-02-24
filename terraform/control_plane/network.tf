resource "twc_vpc" "project-vpc" {
  name = "Porject VPC"
  description = "VPC for project"
  subnet_v4 = "192.168.0.0/24"
  location = "ru-1"
}
