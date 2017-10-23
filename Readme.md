# edb

Collection of muscle groups, muscles, and areas of a muscle that can be
targeted.

Each muscle group is a file found in the [docs][1] directory. The file should
contain:

* any alternate names of the muscle group
* all muscles that can be exercised in the muscle group
* all areas of the muscle that can be exercised/targeted

The file format should resemble:

```
# {Muscle Group}

aka: {alternative muscle group name}, {secondary alternate muscle group name}

## {Muscle Name}

* {Muscle Target Area Name} (aka: {alternative target area name})
```

The [shoulders.md][2] is a decent example to follow.

[1]: /tree/master/docs
[2]: /tree/master/docs/shoulders.md
