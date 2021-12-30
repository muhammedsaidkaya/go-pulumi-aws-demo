
# Pulumi Login
```
pulumi login s3://{{ YOUR_BUCKET_NAME }}
```

# Create Stack
```
pulumi stack init dev
#--secrets-provider="awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-west-2"
```

# Set Configuration
```
pulumi config set aws:region us-west-2
pulumi config set --secret dbPassword fakepassword
```

# Access these configuration values in your Pulumi program:
```
package main

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)
func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        c := config.New(ctx, "")

        name := c.Require("aws:region")
        dbPassword := c.RequireSecret("dbPassword")
    }
}
```

# Provision
```
pulumi up
```