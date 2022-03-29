# 一个简单的命令行小工具 （ one simple command line tool for aws-ec2 )
> 此工具可以列出当前账号下，所有的ec2实例信息 
## 具体使用方式
### 1 配置所需的IAM认证信息  
##### 1.1 创建所需要使用的user信息，并创建相关的 （access key）
##### 1.2 创建如下的权限策略，并将此策略附加到上面的用户中
```shell
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "ec2:Describe*",
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "elasticloadbalancing:Describe*",
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "cloudwatch:ListMetrics",
                "cloudwatch:GetMetricStatistics",
                "cloudwatch:Describe*"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "autoscaling:Describe*",
            "Resource": "*"
        }
    ]
}
```
### 2 具体使用方式
```shell
wlown@debian:~/golang/aws-ec2-cli$ ./aws-ec2-cli  -key '1111111111111' -secret '222222222222'
10.80.0.124          nx-homepage-app00

10.81.200.225        nx-control-app01

10.81.0.29           nx-ldap-app01

10.81.200.15         nx-vpn-app01 

10.80.200.252        nx-nat-1a   
```