module "iam_policy" {
  source = "terraform-aws-modules/iam/aws//modules/iam-policy"

  name        = "policy-eks"
  path        = "/"
  description = "criacao cluster eks"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
          "autoscaling:*",
          "ec2:*",
          "elasticloadbalancing:*",
          "iam:*",
          "kms:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}