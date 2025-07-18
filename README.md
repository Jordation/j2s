# What
Is able to parse a collection an unstructured (json, xml) version of the same data type and output generated types based on the best approximation of matching type names found across the input file.

Takes considerations for partial representations found across the input and the same type being found under different parents

# Usage

```shell
go install github.com/Jordation/j2s@latest
```

```
$ j2s
      -i string
            the input json or xml file to build the type from
      -l string
            the output (l)anguages you want to generate types for (default "go")
      -o string
            the output file for the generated types (default "output.go")
      -tn string
            what the (t)ype(n)ame for the top level datatype should be called (default "Unnamed") 
```
