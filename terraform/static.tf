# AWS S3 bucket for static hosting
resource "aws_s3_bucket" "static" {
  bucket = "mybucketstatickamil2"
  acl = "public-read"

  tags = {
    Name = "Website"
    Environment = "production"
  }

  policy = <<EOF
{
  "Version": "2008-10-17",
  "Statement": [
    {
        "Sid": "PublicReadGetObject",
        "Effect": "Allow",
        "Principal": "*",
        "Action": "s3:GetObject",
        "Resource": "arn:aws:s3:::mybucketstatickamil2/*"
    }
  ]
}
EOF

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

resource "aws_s3_bucket_object" "index" {
  bucket = "${aws_s3_bucket.static.bucket}"
  key = "index.html"
  source = "../app/static/index.html"

  etag = filemd5("../app/static/index.html")
}


resource "aws_s3_bucket_object" "error" {
  bucket = "${aws_s3_bucket.static.bucket}"
  key = "error.html"
  source = "../app/static/error.html"

  etag = filemd5("../app/static/error.html")
}