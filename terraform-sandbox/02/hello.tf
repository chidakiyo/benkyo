resource "local_file" "heloworld" {
  content = "hello world!"
  filename = "hello.txt"
}