terraform {
  backend "http" {
    address = "http://localhost:8080/api/backends/stack-a/state"
  }
}

resource "random_id" "example" {
  byte_length = 8
}

output "random_id" {
  value = random_id.example.hex
}