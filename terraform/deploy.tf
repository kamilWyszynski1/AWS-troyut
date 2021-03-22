resource "aws_ecr_repository" "aws_tryout" {
  name = "aws_tryout"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}

data "aws_ecr_image" "app" {
  repository_name = "${aws_ecr_repository.aws_tryout.name}"
  image_tag = "latest"
}

