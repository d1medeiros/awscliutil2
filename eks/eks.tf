data "aws_subnets" "sub_private" {
  filter {
    name   = "tag:Name"
    values = ["ex-vpc-private-us-east-1a", "ex-vpc-private-us-east-1b", "ex-vpc-private-us-east-1c"]
  }
}

data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}

resource "aws_kms_key" "eks" {
  description             = "eks secret"
  deletion_window_in_days = 7
  enable_key_rotation     = true
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.cluster.token
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "17.23.0"

  cluster_name              = var.cluster_name
  cluster_version           = var.cluster_version
  subnets                   = data.aws_subnets.sub_private.ids
  vpc_id                    = var.vpc
  cluster_enabled_log_types = ["api", "authenticator", "controllerManager"]
  write_kubeconfig          = true
  cluster_encryption_config = [
    {
      provider_key_arn = aws_kms_key.eks.arn
      resources        = ["secrets"]
    }
  ]
  worker_groups = [
    {
      asg_desired_capacity = 3
      asg_max_size         = 5
      instance_type        = "m5.large"
    }
  ]
}