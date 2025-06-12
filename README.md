# What
Is able to parse a collection an unstructured (json, xml) version of the same data type and output generated types based on the best approximation of matching type names found across the input file.

Takes considerations for partial representations found across the input and the same type being found under different parents

# Usage

```shell
go install github.com/Jordation/j2s
```

```shell
js2 -i=<source file path> -o=<output file name> -l=<languages (go, ts)>, -tn=<name of the root type. Optional>
```
