---
name: ASCII JUnit Test Report
description:
  A simple Woodpecker CI plugin that prints out JUnit summaries in ASCII.
author: Wouter (@wgroeneveld)
tags: [testing, java, junit]
containerImage: ghcr.io/wgroeneveld/woodpecker-ascii-junit:main
containerImageUrl: https://github.com/wgroeneveld/woodpecker-ascii-junit/pkgs/container/woodpecker-ascii-junit
url: https://github.com/wgroeneveld/woodpecker-ascii-junit
---

# Woodpecker-ascii-junit

A simple Woodpecker CI plugin that prints out JUnit summaries in ASCII:

```

JUnit Test Results: 3 Test Suites Found
----------------------------------------

| Passed ✅ | Failed ❌ | Errored 🚫 | Skipped ⏭️ | Total 📈 |
_______________________________________________________________
| 7         | 5          | 0          | 1          | 13       | 

⏱️ Total time: 1.378s

❌ Failed Test Details
----------------------
  🧪 Test TestSubtests#package/subtests (⏱️0s) Failure: Failed
  🧪 Test TestSubtests/Subtest#01#package/subtests (⏱️0s) Failure: Failed
    subtests_test.go:10: error message
  🧪 Test TestFailingSubtestWithNestedSubtest#package/subtests (⏱️0s) Failure: Failed
  🧪 Test TestFailingSubtestWithNestedSubtest/Subtest#package/subtests (⏱️0s) Failure: Failed
    subtests_test.go:31: Subtest error message
  🧪 Test someTestName#org.SomeTest (⏱️1.311s) Failure: java.lang.AssertionError: this should be that.
java.lang.AssertionError: this should be that
            at com.tngtech.archunit.lang.ArchRule$Factory$SimpleArchRule.verifyNoEmptyShouldIfEnabled(ArchRule.java:201)
            at com.tngtech.archunit.lang.ArchRule$Factory$SimpleArchRule.evaluate(ArchRule.java:181)
            at com.tngtech.archunit.lang.ArchRule$Assertions.check(ArchRule.java:84)
            at com.tngtech.archunit.lang.ArchRule$Factory$SimpleArchRule.check(ArchRule.java:165)
            at com.tngtech.archunit.lang.syntax.ObjectsShouldInternal.check(ObjectsShouldInternal.java:81)
            at com.tngtech.archunit.junit.internal.ArchUnitTestDescriptor$ArchUnitRuleDescriptor.execute(ArchUnitTestDescriptor.java:168)
            at com.tngtech.archunit.junit.internal.ArchUnitTestDescriptor$ArchUnitRuleDescriptor.execute(ArchUnitTestDescriptor.java:151)
            at java.base/java.util.ArrayList.forEach(ArrayList.java:1596)
            at java.base/java.util.ArrayList.forEach(ArrayList.java:1596)
```

This includes coloured console text output from [chalk](https://github.com/vinay03/chalk).

If there are failing tests, details of those will be printed as well.

If you are using Drone-CI, consider using [drone-junit](https://github.com/rohit-gohri/drone-junit/) instead 
that has a nice Adaptive Card UI which is currently not supported by Woodpecker-CI.

This plugin reads JUnit XML files in a `path` glob pattern. If you're running JS Jest tests, add [jest-junit](https://github.com/jest-community/jest-junit) as a reporter and it'll be integrated as well.

## Configuration

See `docker-compose.yml` as an example:

- `PLUGIN_PATH` env var or `path` setting in Woodpecker
- Optional: `PLUGIN_LOG_LEVEL` env var or `log-level` (built-in Woodpecker plugin)

Here's an example how to include it in your Woodpecker workflow:

```
  - name: junit-reports
    image: ghcr.io/wgroeneveld/woodpecker-ascii-junit:main
    settings:
      log-level: debug
      path: /tmp/reports/**/*.xml
    when:
      status: [
        'success',
        'failure',
      ]
```
