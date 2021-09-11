# argocd-templatize

Templatize finds `./*.json` and `./*.yaml` and passes them through [text/template](https://pkg.go.dev/text/template) with the environment variables named `TPZ_*`.

`TPZ_SOME_THING` is available as `{{.SomeThing}}` in the template.

Its intended usage is as a custom plugin for argocd, to be able to use cluster specific variables in your YAML.

You can install this into [ArgoCD](https://argoproj.github.io/argo-cd/) with the following Kustomization: https://github.com/Jille/argocd-templatize/blob/master/kustomization.yaml

Now you can enable templatize in your `Application`:

```diff
 apiVersion: argoproj.io/v1alpha1
 kind: Application
 spec:
   project: default
   source:
     repoURL: ssh://git@github.com:/Jille/configuration.git
+    plugin:
+      name: templatize
+      env:
+      - name: TPZ_CLUSTER
+        value: prod
```

`TPZ_CLUSTER` will now be available as `{{.Cluster}}` in your YAML. You can use all functionality of [text/template](https://pkg.go.dev/text/template).
