
```golang
package main

import (
        "encoding/json"
        // "k8s.io/kubernetes/pkg/api"
        "fmt"
        "k8s.io/kubernetes/pkg/apis/extensions"
        podutil "k8s.io/kubernetes/pkg/util/pod"
)

const (
        DeploymentNewStr = `{
        ...
         "spec": {
         ...
          "template": {
          ...
          "spec": {
                "containers": [
                    {
                    ...
                    "image": "img.reg.3g:15000/yce-alpha:201703211640",
                    ...
                    }
                }
            }
        }
    }`

    DeploymentOldStr = `{
        ...
         "spec": {
         ...
          "template": {
          ...
          "spec": {
                "containers": [
                    {
                    ...
                    "image": "img.reg.3g:15000/yce-alpha:201703151133",
                    ...
                    }
                }
            }
        }
    }`

func main() {

        newDp:= new(extensions.Deployment)
        json.Unmarshal([]byte(DeploymentNewStr), &newDp)

        oldDp:= new(extensions.Deployment)
        json.Unmarshal([]byte(DeploymentOldStr), &oldDp)

        podTemplateSpecHash := podutil.GetPodTemplateSpecHash(newDp.Spec.Template)
        podTemplateSpecHashOld := podutil.GetPodTemplateSpecHash(oldDp.Spec.Template)

        fmt.Printf("%d, %s\n", podTemplateSpecHash, newDp.Spec.Template.Spec.Containers[0].Image)
        fmt.Printf("%d, %s\n", podTemplateSpecHashOld, oldDp.Spec.Template.Spec.Containers[0].Image)
}
```




